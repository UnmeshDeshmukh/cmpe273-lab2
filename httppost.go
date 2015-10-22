package main

import (  
    // Standard library packages
	"encoding/json"
    "fmt"
    "net/http"
    // Third party packages
    "github.com/julienschmidt/httprouter"
    
)
type Request struct {
	   Name   string `json:"name"`
}
type Response struct
{
    Greeting string `json:"greeting"`
}

func postt(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
   // fmt.Fprintf(rw, "Hello, %s!\n", p.ByName("name"))
    u:= Request{}
	re:= Response{}   
    // Populate the user data
    json.NewDecoder(r.Body).Decode(&u)
    re.Greeting = "Hello," + u.Name + "!"
    // Marshal provided interface into JSON structure
    uj, _ := json.Marshal(re)

    // Write content-type, statuscode, payload
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(201)
    fmt.Fprintf(w, "%s!", uj)
    

    
}


func main() {  
    // Instantiate a new router
    r := httprouter.New()

//r.GET("/user/:id", gett)
r.POST("/hello", postt)
//r.GET("/user/:id",gett)
	 server := http.Server{
            Addr:        "0.0.0.0:8080",
            Handler: r,
    }
    server.ListenAndServe()

 //   http.ListenAndServe("localhost:8080", r)
}