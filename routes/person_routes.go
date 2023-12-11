package routes

import (
	"github.com/JuanDiegoCastellanos/apirest_go_mysql/controllers"
	"github.com/gorilla/mux"
)

func SetPersonRoutes(router *mux.Router) {
	subRoute := router.PathPrefix("/person/api").Subrouter()
	subRoute.HandleFunc("/all", controllers.GetAll).Methods("GET")
	subRoute.HandleFunc("/save", controllers.Save).Methods("POST")
	subRoute.HandleFunc("/delete/{id}", controllers.Delete).Methods("POST")
	subRoute.HandleFunc("/find/{id}", controllers.GetById).Methods("GET")
}
