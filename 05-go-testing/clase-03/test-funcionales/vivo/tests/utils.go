package tests

import (
	"bytes"
	"github.com/NachofraML/bootcamp-go-w8/05-go-testing/clase-03/test-funcionales/vivo/cmd/server"
	"github.com/NachofraML/bootcamp-go-w8/05-go-testing/clase-03/test-funcionales/vivo/prey"
	"github.com/NachofraML/bootcamp-go-w8/05-go-testing/clase-03/test-funcionales/vivo/shark"
	"github.com/NachofraML/bootcamp-go-w8/05-go-testing/clase-03/test-funcionales/vivo/simulator"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
)

func createServer() *gin.Engine {
	r := gin.Default()
	sim := simulator.NewCatchSimulator(35.4)

	whiteShark := shark.CreateWhiteShark(sim)
	tuna := prey.CreateTuna()

	handler := server.NewHandler(whiteShark, tuna)

	g := r.Group("/v1")

	g.PUT("/shark", handler.ConfigureShark())
	g.PUT("/prey", handler.ConfigurePrey())
	g.POST("/simulate", handler.SimulateHunt())

	return r
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))

	req.Header.Add("Content-Type", "application/json")
	return req, httptest.NewRecorder()
}
