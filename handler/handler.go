package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//Healthz handler will make kubernetes container healthy
func Healthz(r *httprouter.Router) {
	fn := func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	}

	r.GET("/healthz", fn)
}

func Index(r *httprouter.Router) {
	fn := func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(`
     Source code  untuk https://medium.com/@agipreza/helm-sebuah-package-manager-untuk-kubernetes-6b47f5e88c59`))
	}

	r.GET("/", fn)
}
