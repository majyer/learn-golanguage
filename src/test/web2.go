package main
 
import (
    "fmt"
    "log"
    "net/http"
)
 
type MyMux struct {}
 
func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path == "/" {
        sayhelloName(w, r)
        return
    }
 
    http.NotFound(w, r)
    return
}
 
 
func sayhelloName(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    fmt.Println("path:", r.URL.Path)
    fmt.Fprintf(w, "hello go again")
}
 
func main() {
    mux := &MyMux{}
    err := http.ListenAndServe(":9090", mux)
    if err != nil {
        log.Fatal("ListenAndServer: ", err)
    }
}
