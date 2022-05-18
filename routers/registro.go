package routers

import (
	"encoding/json"
	"net/http"

	"github.com/alejandrogith/bd"
	"github.com/alejandrogith/models"
)

//Registro es la funcion para crear en la BD el registro de usuario
func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos recibidos"+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email de usuario es requerido", 400)
	}

	if len(t.Password) < 6 {
		http.Error(w, "Debe especificar la contraseña", 400)
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado == true {
		http.Error(w, "Ya existe un usuario registrado con ese email", 400)
	}

	_, status, err := bd.InsertoRegistro(t)
	if encontrado == true {
		http.Error(w, "Ocurrio un error al intentar el registro"+err.Error(), 400)
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el registro de usuario", 400)
	}

	w.WriteHeader(http.StatusCreated)

}
