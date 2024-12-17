package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/MrTeacheer/ecom/service/user"
	"github.com/gorilla/mux"
)


type APIServer struct{
	addr string
	db *sql.DB
}

func NewAPIServer(addr string,db *sql.DB) *APIServer{
	return &APIServer{
		addr: addr,
		db: db,
	}
}


func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("api/v1").Subrouter()

	user_route := user.NewHandler()
	user_route.RegisterRoutes(subrouter)


	log.Println("listening",s.addr)
	return http.ListenAndServe(s.addr,router)
}