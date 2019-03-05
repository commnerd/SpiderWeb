package node

type Api struct {
    ip string           `json:"ip"`
    domain string       `json:"domain"`
    basePath string     `json:"base_path"`
    hostPort string     `json:"host_port"`
}

var routes = map[string]func {
	"/": Welcome
	"/register": Register
	"/ports/next": NextPort
}

func (this *Api) Listen() {
	r := mux.NewRouter()

	for path, handler := range routes {
		r.HandleFunc(path, handler)
	}

    log.Fatal(http.ListenAndServe("127.0.0.1:"+this.hostPort, r))
}