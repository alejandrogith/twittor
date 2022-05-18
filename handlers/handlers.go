package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/alejandrogith/middlew"
	"github.com/alejandrogith/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Manejadores() {
	router := mux.NewRouter()
	PORT := os.Getenv("PORT")

	router.HandleFunc("registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")

	if PORT == "" {
		PORT = "8080"
	}

	HANDLER := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, HANDLER))

}
