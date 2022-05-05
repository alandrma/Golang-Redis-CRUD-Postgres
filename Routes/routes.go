package routes

import (
	"Golang_CRUD_Postgresql/Controller"
	"github.com/gorilla/mux"
)

func Routers() *mux.Router {

	routers := mux.NewRouter()
	routers.HandleFunc("/api/buku/redis", controller.AmbilRedis).Methods("GET", "OPTIONS")
	routers.HandleFunc("/api/buku", controller.AmbilSemuaBuku).Methods("GET", "OPTIONS")
	routers.HandleFunc("/api/buku/{id}", controller.AmbilBuku).Methods("GET", "OPTIONS")
	routers.HandleFunc("/api/buku", controller.TmbhBuku).Methods("POST", "OPTIONS")
	routers.HandleFunc("/api/buku/{id}", controller.UpdateBuku).Methods("PUT", "OPTIONS")
	routers.HandleFunc("/api/buku/{id}", controller.HapusBuku).Methods("DELETE", "OPTIONS")

	return routers
}

//test
