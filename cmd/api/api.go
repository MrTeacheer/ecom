package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/MrTeacheer/ecom/service/orders"
	"github.com/MrTeacheer/ecom/service/products"
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

// func greeting(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "HELLO TEACHER!!!!!")
// }

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	user_store := user.NewStore(s.db)
	user_route := user.NewHandler(user_store)
	user_route.RegisterRoutes(subrouter)

	product_store := products.NewStore(s.db)
	product_route := products.NewHandler(product_store)
	product_route.RegisterRouter(subrouter)

	order_store := orders.NewStore(s.db)
	order_route := orders.NewHandler(order_store)
	order_route.RegisterRouter(subrouter)

	log.Println("listening", s.addr)
	return http.ListenAndServe(s.addr, router)
}
