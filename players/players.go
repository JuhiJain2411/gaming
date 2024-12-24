package players

import (
	"fmt"
	"net/http"
)

func PlayerServer(response http.ResponseWriter, r *http.Request) {
	fmt.Fprint(response, "20")
}
