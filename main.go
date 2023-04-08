package main

import (
	"github.com/ainmtsn1999/orm_jwt_auth/database"
	"github.com/ainmtsn1999/orm_jwt_auth/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8000")
}
