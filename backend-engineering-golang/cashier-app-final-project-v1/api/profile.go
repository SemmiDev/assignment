package api

import (
	"a21hc3NpZ25tZW50/model"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path"
	"text/template"
)

var availablePhotoExtension = map[string]struct{}{
	".jpg":  {},
	".jpeg": {},
	".png":  {},
}

func (api *API) ImgProfileView(w http.ResponseWriter, r *http.Request) {
	username := r.Context().Value("username").(string)

	for ext, _ := range availablePhotoExtension {
		fileName := username + "-photo-profile" + ext
		profilePicNamePath := path.Join("assets", "images", fileName)

		if _, err := os.Stat(profilePicNamePath); err == nil {
			w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
			http.ServeFile(w, r, profilePicNamePath)
			return
		}
		fileName = ""
	}

	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	fallback := path.Join("assets", "images", "img-avatar-default.png")
	http.ServeFile(w, r, fallback)
	return
}

func (api *API) ImgProfileUpdate(w http.ResponseWriter, r *http.Request) {
	username := r.Context().Value("username").(string)
	profilePicNamePath := username + "-photo-profile"

	err := r.ParseMultipartForm(2 << 20) // 2 mb
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	file, header, err := r.FormFile("file-avatar")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}
	defer file.Close()

	extension := path.Ext(header.Filename)
	fileName := profilePicNamePath + extension
	profilePicNamePath = path.Join("assets", "images", fileName)

	f, err := os.OpenFile(profilePicNamePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer f.Close()

	_, err = io.Copy(f, file)

	filepath := path.Join("views", "dashboard.html")
	tmpl, err := template.ParseFiles(filepath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}

	cart, err := api.cartsRepo.ReadCart()
	cart.Name = username
	listProducts, err := api.products.ReadProducts()
	data := model.Dashboard{
		Product: listProducts,
		Cart:    cart,
	}

	// for refresh page after upload image
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")

	err = tmpl.Execute(w, data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
	}
}
