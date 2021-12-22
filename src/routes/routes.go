package routes

import (
	"net/http"

	"github.com/ayaanqui/go-rest-server/src/utils"
)

func CreateRoutes(app *AppBase, server *http.ServeMux) {
	server.Handle("/", utils.HandleGet(
		http.HandlerFunc(app.Home),
	))
	server.HandleFunc("/posts", app.Posts)
}