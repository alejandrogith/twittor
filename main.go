package main

import (
	"log"

	"github.com/alejandrogith/bd"
	"github.com/alejandrogith/handlers"
)

func main() {

	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}

	handlers.Manejadores()
}
