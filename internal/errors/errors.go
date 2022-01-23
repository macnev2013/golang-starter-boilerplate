package errors

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/intrigues/golang-starter-boilerplate/core"
)

func ServerError(w http.ResponseWriter, err error, logger core.Clog) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	logger.Errorf(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
