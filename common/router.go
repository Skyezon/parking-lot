package common

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"gitlab.mapan.io/playground/parking-lot-golang/handler"
)

func Router(){
    r:= chi.NewRouter()

    r.Get("/",handler.HelloWorldHandler)

    http.ListenAndServe(":8080",r)
}
