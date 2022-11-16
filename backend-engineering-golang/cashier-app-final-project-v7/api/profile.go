package api

import (
	"a21hc3NpZ25tZW50/model"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

var availableImageExt = map[string]struct{}{
	".jpg":  {},
	".jpeg": {},
	".png":  {},
}

func (api *API) ImgProfileView(w http.ResponseWriter, r *http.Request) {
	username := r.Context().Value("username").(string)

	for ext := range availableImageExt {
		fileName := fmt.Sprintf("%s-avatar%s", username, ext)
		profilePicNamePath := path.Join("assets", "images", fileName)

		if _, err := os.Stat(profilePicNamePath); err == nil {
			w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
			http.ServeFile(w, r, profilePicNamePath)
			return
		}
	}

	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	defaultAvatar := path.Join("assets", "images", "img-avatar-default.png")
	http.ServeFile(w, r, defaultAvatar)
}

func (api *API) ImgProfileUpdate(w http.ResponseWriter, r *http.Request) {
	username := r.Context().Value("username").(string)

	err := r.ParseMultipartForm(2 << 20) // 2 mb
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	profilePicName := fmt.Sprintf("%s-avatar", username)

	file, header, err := r.FormFile("file-avatar")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}
	defer file.Close()

	extension := path.Ext(header.Filename)
	fileName := fmt.Sprintf("%s%s", profilePicName, extension)
	profilePicNamePath := path.Join("assets", "images", fileName)

	f, err := os.OpenFile(profilePicNamePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer f.Close()

	_, err = io.Copy(f, file)

	// for refresh page after upload image
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")

	api.dashboardView(w, r)
}
