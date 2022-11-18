package helpers

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/rudsonalves/bookings/internal/config"
)

var app *config.AppConfig

// NewHerlpers sets up app config for helpers
func NewHelpers(a *config.AppConfig) {
	app = a
}

func ClientError(write http.ResponseWriter, status int) {
	app.InfoLog.Println("Clent error with status of", status)
	http.Error(write, http.StatusText(status), status)
}

func ServerError(write http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Println(trace)
	SIError := http.StatusInternalServerError
	http.Error(write, http.StatusText(SIError), SIError)
}
