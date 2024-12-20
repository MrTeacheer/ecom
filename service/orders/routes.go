package orders

import (
	"net/http"

	"github.com/MrTeacheer/ecom/types"
	"github.com/MrTeacheer/ecom/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)




type Handler struct{
	s types.OrderStore
}


func NewHandler(s types.OrderStore) *Handler{
	return &Handler{s: s}
}


func (h *Handler) RegisterRouter (router *mux.Router){
	router.HandleFunc("/orders",h.OrdersAPI).Methods("POST","GET")

}


func (h *Handler) OrdersAPI (w http.ResponseWriter,r *http.Request){
	switch r.Method{
	case "POST":
		var payload types.OrdersAdd
		if err:=utils.ParseJSON(r,&payload); err!= nil{
			utils.WriteError(w,http.StatusBadRequest,err)
			return
		}
		if err:=utils.Validate.Struct(payload);err!=nil{
			errors := err.(validator.ValidationErrors)
			utils.WriteError(w,http.StatusBadRequest,errors)
			return
		}
		if err := h.s.CreateOrder(payload); err!= nil{
			utils.WriteError(w,http.StatusBadRequest,err)
			return
		}
		utils.WriteAPI(w, http.StatusCreated, map[string]string{"ok": "product was created"})
		return

	case "GET":
		ods,err:=h.s.GetOrders()
		if err != nil{
			utils.WriteError(w,http.StatusBadRequest,err)
			return
		}
		utils.WriteAPI(w,http.StatusAccepted,ods)
	}
}

