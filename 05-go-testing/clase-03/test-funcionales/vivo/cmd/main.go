package main

import (
	"github.com/NachofraML/bootcamp-go-w8/05-go-testing/clase-03/test-funcionales/vivo/cmd/server"
	"github.com/NachofraML/bootcamp-go-w8/05-go-testing/clase-03/test-funcionales/vivo/prey"
	"github.com/NachofraML/bootcamp-go-w8/05-go-testing/clase-03/test-funcionales/vivo/shark"
	"github.com/NachofraML/bootcamp-go-w8/05-go-testing/clase-03/test-funcionales/vivo/simulator"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	sim := simulator.NewCatchSimulator(35.4)

	whiteShark := shark.CreateWhiteShark(sim)
	tuna := prey.CreateTuna()

	handler := server.NewHandler(whiteShark, tuna)

	srv := server.NewServer(handler, r)
	srv.MapRoutes()

	if err := srv.Run(); err != nil {
		panic(err)
	}
}
