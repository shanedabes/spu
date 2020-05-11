package auth

import (
	"fmt"
	"net/http"
)

func NewCallbackHandler(state string, ch chan<- string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if st := r.FormValue("state"); st != state {
			http.NotFound(w, r)
			ch <- "error"

			return
		}

		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, "Login Completed!")
		ch <- r.FormValue("code")
	}
}
