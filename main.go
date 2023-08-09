package main

import (
	"jzblog/server/routers"
)

func main() {
	router := routers.NewRouter()
	router.Run(":8080")
}
