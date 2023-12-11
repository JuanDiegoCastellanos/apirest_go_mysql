package main

import (
	"log"
	"net/http"

	"github.com/JuanDiegoCastellanos/apirest_go_mysql/commons"
	"github.com/JuanDiegoCastellanos/apirest_go_mysql/routes"
	"github.com/gorilla/mux"
)

func main() {
	commons.Migrate()

	router := mux.NewRouter()

	routes.SetPersonRoutes(router)

	server := http.Server{
		Addr:    ":9000",
		Handler: router,
	}

	log.Println("Servidor Ejecutandose sobre puerto 9000")

	log.Println(server.ListenAndServe())
}
