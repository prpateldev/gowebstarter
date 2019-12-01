package pingcontroller

import "net/http"
import "fmt"

// GET /ping
func Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("pong")))
}
