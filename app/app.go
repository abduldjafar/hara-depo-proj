package app

import (
	"github.com/gorilla/handlers"
	"log"
	"net/http"
)

func (app *App) Run(host string) {
	headersOk := handlers.AllowedHeaders([]string{"Content-Type", "Access-Control-Allow-Headers", "Authorization", "X-Requested-With", "Access-Control-Allow-Origin", "x-access-token"})
	corsObj := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	log.Fatal(http.ListenAndServe(host, handlers.CORS(corsObj, headersOk, methodsOk)(app.Router)))

}
