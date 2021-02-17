package initialization

import (
	"net/http"
)

var Routes = map[string]func(w http.ResponseWriter, r *http.Request){
	"/token":          TokenHandler,
	"/reverse_tunnel": ReverseTunnelHandler,
}
