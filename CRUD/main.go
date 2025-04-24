package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

func conexionBD() (conexion *sql.DB) {
	Driver := "mysql"
	Usuario := "root"
	Contraseña := ""
	Nombre := "sistema"

	conexion, err := sql.Open(Driver, Usuario+":"+Contraseña+"@tcp(127.0.0.1)/"+Nombre)

	if err != nil {
		panic(err.Error())
	}
	return conexion
}

var plantillas = template.Must(template.ParseGlob("plantillas/*"))

func main() {
	http.HandleFunc("/", Inicio)
	http.HandleFunc("/crear", Crear)

	log.Println("Servidor corriendo")
	fmt.Println("Servidor corriendo")
	http.ListenAndServe(":8080", nil)
}

func Inicio(w http.ResponseWriter, r *http.Request) {

	conexionEstablecida := conexionBD()
	insertarRegistro, err := conexionEstablecida.Prepare("INSERT INTO empleados(nombre,correo) VALUES ('oscar', 'correo@gmail.com')")

	if err != nil {
		panic(err.Error())
	}
	insertarRegistro.Exec()

	// fmt.Fprintf(w, "Hola Mundo")
	plantillas.ExecuteTemplate(w, "inicio", nil)
}

func Crear(w http.ResponseWriter, r *http.Request) {
	plantillas.ExecuteTemplate(w, "crear", nil)
}
