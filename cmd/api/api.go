package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/peterest/go-basic-ecom/service/user"
)

type ApiServer struct {
	addr string
	db   *sql.DB
}

func NewApiServer(addr string, db *sql.DB) *ApiServer {
	return &ApiServer{
		addr: addr,
		db:   db,
	}
}

func (server *ApiServer) Run() error {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/v1").Subrouter()

	userRepository := user.NewRepository(server.db)
	user.NewHandler(userRepository).RegisterRoutes(subRouter)

	log.Println("Starting server on", server.addr)

	return http.ListenAndServe(server.addr, router)
}
