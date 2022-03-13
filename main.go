package main

import (
	"fmt"
	"hmv-rest-api/routes"
	"os"
)

func main() {
	fmt.Println("Iniciando API HMV")
	os.Setenv("GO_CRIPYT", "45f84bddefa6c5212b60223ceaf64e61")
	routes.HandleRequest()
}
