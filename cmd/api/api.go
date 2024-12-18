package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/MrTeacheer/ecom/service/user"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func greeting(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HELLO TEACHER!!!!!")
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	user_store := user.NewStore(s.db)
	user_route := user.NewHandler(user_store)
	user_route.RegisterRoutes(subrouter)
	subrouter.HandleFunc("/greeting", greeting)

	log.Println("listening", s.addr)
	return http.ListenAndServe(s.addr, router)
}
