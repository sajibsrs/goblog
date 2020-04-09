// This file handles default routes of the application
// and serves templates accordingly

package handler

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Say hello to go!")
}