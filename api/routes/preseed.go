package routes

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Renderer interface {
	GetForHostname(hostname string) ([]byte, error)
}

func GetPreseedHandler(renderer Renderer) (string, http.HandlerFunc) {
	return "/preseed/{hostname}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Got request to %s\n", r.URL.Path)
		vars := mux.Vars(r)
		hostname, ok := vars["hostname"]
		if !ok {
			fmt.Fprintln(os.Stderr, "Missing hostname from path, bad request")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		preseed, err := renderer.GetForHostname(hostname)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		w.Write(preseed)
	}
}
