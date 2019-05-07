package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/academy19/myml/src/api/controllers/myml"
)
import "github.com/mercadolibre/academy19/myml/src/api/controllers/ping"

//puedo usar tags, para simplificar el uso del nombre del package, cuando use
// los metodos de gin, como en el main. por ej:  import goG "github.com/gin-gonic/gin"


//definiciiones de variabels, no se usan :
//no se pone tipo de datos, por ser tipado estatico.
const (
	port=":8080"
)


//definiciones de variables
var (
	router = gin.Default()
)



//despues de estas declaraciones, siguen las estructuras

func main() {

	router.GET( "/ping", ping.Ping)
	// "/:id" indica qeu sera un parametro
	router.GET("/users/:userId", myml.GetUserId)
	router.Run(port)

}
