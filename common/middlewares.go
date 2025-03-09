package common

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
)

type ErrorOut struct {
	Detail string `json:"detail"`
}

func CustomRecoverer(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			r := recover()
			if r != nil {
				switch r := r.(type) {
				case ClientError:
					w.WriteHeader(400)

					errOut := ErrorOut(r)
					dataOut, err := json.Marshal(errOut)
					if err != nil {
						panic("Custom recoverer: unexpected error.")
					}
					w.Write(dataOut)
				case UnrecoverableError:
					// log stack to stdout, uses chi default.
					// todo: try to use a proper logger instead.
					middleware.PrintPrettyStack(r)

					w.WriteHeader(500)

					errorOut := ErrorOut(r)
					dataOut, err := json.Marshal(errorOut)
					if err != nil {
						panic("Custom recoverer: unexpected error.")
					}
					w.Write(dataOut)
				default:
					// log stack to stdout, uses chi default.
					// todo: try to use a proper logger instead.
					middleware.PrintPrettyStack(r)

					w.WriteHeader(500)
					w.Write([]byte("internal server error"))
				}
			}
		}()

		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
