package main

import (
	"app.translate/models"
	"fmt"
)

func main() {
	db := models.GetDB()
	fmt.Println("translate", db)
}
