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
	http.HandleFunc("/insertar", Insertar)

	http.HandleFunc("/borrar", Borrar)
	http.HandleFunc("/editar", Editar)
	http.HandleFunc("/actualizar", Actualizar)

	log.Println("Servidor corriendo")
	fmt.Println("Servidor corriendo")
	http.ListenAndServe(":8080", nil)
}

func Actualizar(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		id := r.FormValue("id")
		nombre := r.FormValue("nombre")
		correo := r.FormValue("correo")

		conexionEstablecida := conexionBD()
		modificarRegistro, err := conexionEstablecida.Prepare("UPDATE empleados SET nombre=?, correo=? WHERE id=?")

		if err != nil {
			panic(err.Error())
		}
		modificarRegistro.Exec(nombre, correo, id)

		http.Redirect(w, r, "/", 301)
	}
}

func Editar(w http.ResponseWriter, r *http.Request) {
	idEmpleado := r.URL.Query().Get("id")

	conexionEstablecida := conexionBD()

	registro, err := conexionEstablecida.Query("SELECT * FROM empleados WHERE id=?", idEmpleado)

	empleado := Empleado{}

	for registro.Next() {
		var id int
		var nombre, correo string
		err = registro.Scan(&id, &nombre, &correo)
		if err != nil {
			panic(err.Error())
		}

		empleado.Id = id
		empleado.Nombre = nombre
		empleado.Correo = correo
	}
	fmt.Println(empleado)
	plantillas.ExecuteTemplate(w, "editar", empleado)

}

func Borrar(w http.ResponseWriter, r *http.Request) {
	idEmpleado := r.URL.Query().Get("id")

	conexionEstablecida := conexionBD()

	borrarRegistro, err := conexionEstablecida.Prepare("DELETE FROM empleados WHERE id=?")

	if err != nil {
		panic(err.Error())
	}
	borrarRegistro.Exec(idEmpleado)
	http.Redirect(w, r, "/", 301)
}

type Empleado struct {
	Id     int
	Nombre string
	Correo string
}

func Inicio(w http.ResponseWriter, r *http.Request) {

	conexionEstablecida := conexionBD()
	registros, err := conexionEstablecida.Query("SELECT * FROM empleados")

	if err != nil {
		panic(err.Error())
	}

	empleado := Empleado{}
	arregloEmpleado := []Empleado{}

	for registros.Next() {
		var id int
		var nombre, correo string
		err = registros.Scan(&id, &nombre, &correo)
		if err != nil {
			panic(err.Error())
		}

		empleado.Id = id
		empleado.Nombre = nombre
		empleado.Correo = correo

		arregloEmpleado = append(arregloEmpleado, empleado)
	}

	fmt.Println(arregloEmpleado)

	// fmt.Fprintf(w, "Hola Mundo")
	plantillas.ExecuteTemplate(w, "inicio", arregloEmpleado)
}

func Crear(w http.ResponseWriter, r *http.Request) {
	plantillas.ExecuteTemplate(w, "crear", nil)
}

func Insertar(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nombre := r.FormValue("nombre")
		correo := r.FormValue("correo")

		conexionEstablecida := conexionBD()
		insertarRegistro, err := conexionEstablecida.Prepare("INSERT INTO empleados(nombre,correo) VALUES (?,?)")

		if err != nil {
			panic(err.Error())
		}
		insertarRegistro.Exec(nombre, correo)

		http.Redirect(w, r, "/", 301)
	}
}
