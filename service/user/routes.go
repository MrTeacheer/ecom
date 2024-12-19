package user

import (
	"fmt"
	"net/http"

	"github.com/MrTeacheer/ecom/config"
	"github.com/MrTeacheer/ecom/service/auth"
	"github.com/MrTeacheer/ecom/types"
	"github.com/MrTeacheer/ecom/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.HandleLogin).Methods("POST")
	router.HandleFunc("/register", h.HandleRegister).Methods("POST")
	router.HandleFunc("/users", h.GetUsers).Methods("GET")

}

func (h *Handler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	var payload types.LoginUserPayload
	if err:= utils.ParseJSON(r,&payload); err != nil{
		utils.WriteError(w,http.StatusBadRequest,err)
		return
	}

	u,err := h.store.GetUserByEmail(payload.Email)
	if err != nil{
		utils.WriteError(w,http.StatusBadRequest,err)
		return
	}

	if !auth.ComparePasswords(u.Password,[]byte(payload.Password)){
		utils.WriteError(w,http.StatusBadRequest,fmt.Errorf("pssword is incorrect"))
		return
	}
	secret := config.Envs.JWTsecret
	token, err := auth.CreateJWT([]byte(secret),u.ID)
	if err!= nil{
		utils.WriteError(w,http.StatusBadRequest,err)
		return
	}

	utils.WriteAPI(w,http.StatusOK,map[string]string{"token":token})

}

func (h *Handler) HandleRegister(w http.ResponseWriter, r *http.Request) {
	var payload types.RegisterUserPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, errors)
		return
	}

	_, err := h.store.GetUserByEmail(payload.Email)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with this email already exists"))
		return
	}
	hased_pass, err := auth.HashPassword(payload.Password)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	err = h.store.CreateUser(types.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  hased_pass,
	})
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)

}

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {

	users, err := h.store.GetAllUsers()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}
	utils.WriteAPI(w, http.StatusAccepted, users)

}
