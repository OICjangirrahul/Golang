package route

import (
	"golang/controller"
	"net/http"
)


func SetupRoutes() *http.ServeMux {
	app := http.NewServeMux()
	app.HandleFunc("GET  /user",  controller.GetUser)
	app.HandleFunc("GET  /admin",  controller.GetAdmin)

	return app
}

