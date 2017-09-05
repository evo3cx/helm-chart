package main

import (
	"log"
	"net/http"

	"github.com/evo3cx/helm-chart/handler"
	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	handler.Index(r)
	handler.Healthz(r)

	log.Fatal("run in  :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
