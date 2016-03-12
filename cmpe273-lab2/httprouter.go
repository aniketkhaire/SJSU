package main
import (
    "fmt"
    "github.com/julienschmidt/httprouter"
    "net/http"
    "encoding/json"
)

type Request struct{
    Name string `json:"name"`
}
type Response struct{
    Greeting string `json:"greeting"`
}

func myGet(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
        fmt.Fprintf(w, "Hello, %s!", p.ByName("name"))
    }

func myPost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {  
    req := Request{}
    res := Response{}

    json.NewDecoder(r.Body).Decode(&req)
    res.Greeting = "Hello, "+req.Name+"!"
    data, _ := json.Marshal(res)

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(201)
    fmt.Fprintf(w, "Hello, %s!", data)
}

func main() {
    mux := httprouter.New()
    mux.GET("/hello/:name", myGet)
    mux.POST("/hello", myPost)
    server := http.Server{
            Addr:        "0.0.0.0:8080",
            Handler: mux,
    }
    server.ListenAndServe()
}

