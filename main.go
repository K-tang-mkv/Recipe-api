// swagger:meta

package main

import (
	"github.com/recipes-api/routers"
)

func main() {

	router := routers.InitRouter()

	router.Run()
}
