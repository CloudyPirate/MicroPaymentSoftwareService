package main

import (
	"github.com/CloudyPirate/MPSS/internal/router"
)

func main() {

	r := router.SetupRouter()

	r.Run(":8080") // Using Gin

}
