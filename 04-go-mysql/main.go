package main

import (
	"dbmysql/db"
	"dbmysql/models"
	"fmt"
)

func main() {

	//Verificar conexion
	db.Connect()
	db.Ping()
	fmt.Println(db.ExistsTable("users"))

	//Clear tabla
	db.CreateTable(models.UserSchema, "users")

	user := models.ListUsers()
	fmt.Println(user)

	//Cerrar la conexion
	db.Close()

}
