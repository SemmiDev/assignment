package api

import (
	"a21hc3NpZ25tZW50/model"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func (api *API) AddCart(w http.ResponseWriter, r *http.Request) {
	username := r.Context().Value("username").(string)

	r.ParseForm()

	product := r.Form.Get("product")
	if product == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Request Product Not Found"})
		return
	}

	var list []model.Product
	totalPrice := 0.0
	for _, formList := range r.Form {
		for _, v := range formList {
			item := strings.Split(v, ",")
			p, _ := strconv.ParseFloat(item[2], 64)
			q, _ := strconv.ParseFloat(item[3], 64)
			total := p * q
			totalPrice += total
			list = append(list, model.Product{
				Id:       item[0],
				Name:     item[1],
				Price:    item[2],
				Quantity: item[3],
				Total:    total,
			})
		}
	}

	cart := model.Cart{
		Name:       username,
		Cart:       list,
		TotalPrice: totalPrice,
	}

	_, err := api.cartsRepo.CartUserExist(cart.Name)
	if err != nil {
		api.cartsRepo.AddCart(cart)
	} else {
		api.cartsRepo.UpdateCart(cart)
	}
	api.dashboardView(w, r)

}
