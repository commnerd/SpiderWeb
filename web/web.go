// The web package provides all necessary endpoints for messaging between nodes
package web

// The web service held open for node communications
type Server interface{
    Start()
}

// The request structure passed to a server
type Request interface{
    GetVerb() string
    GetReferer() string
    GetBody() string
}

// The response structure returned from a server
type Response interface{
    SetFormat(string)
    GetFormat() string
    SetBody(string)
    GetBody() string
}
