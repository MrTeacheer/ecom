package products

import (
	"net/http"

	"github.com/MrTeacheer/ecom/types"
	"github.com/MrTeacheer/ecom/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	s types.ProductStore
}

func NewHandler(s types.ProductStore) *Handler {
	return &Handler{s: s}
}

func (h *Handler) RegisterRouter(route *mux.Router) {
	route.HandleFunc("/product", h.ProductsAPI).Methods("GET", "POST")
}

func (h *Handler) ProductsAPI(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var payload types.ProductAdd
		if err := utils.ParseJSON(r, &payload); err != nil {
			utils.WriteError(w, http.StatusBadRequest, err)
			return
		}
		if err := utils.Validate.Struct(payload); err != nil {
			erros := err.(validator.ValidationErrors)
			utils.WriteError(w, http.StatusBadRequest, erros)
			return
		}
		if err := h.s.CreateProduct(payload); err != nil {
			utils.WriteError(w, http.StatusBadRequest, err)
			return
		}

		utils.WriteJSON(w, http.StatusCreated, map[string]string{"200": "product was created"})
		return

	case "GET":
		ps, err := h.s.GetProducts()
		if err != nil {
			utils.WriteError(w, http.StatusBadRequest, err)
		}
		utils.WriteAPI(w, http.StatusAccepted, ps)
	}
}
