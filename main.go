package main

import (
	"user-server/routers"
)

func main() {
	r := routers.SetupRouter()
	r.Run("localhost:8080")
}
