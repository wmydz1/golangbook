package main
import (
    "net/http"
    "gopkg.in/unrolled/render.v1"
)

type Action func(w http.ResponseWriter, r *http.Request) error
type AppController struct {}
func (c *AppController) Action(a Action) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if err := a(w, r); err !=nil {
            http.Error(w, err.Error(), 500)
        }
    })
}

type MyController struct {
    AppController
    *render.Render
}
func (c *MyController) Index(w http.ResponseWriter, r *http.Request) error {
    c.JSON(w, http.StatusOK, map[string]string{"Name": "Samchen"})
    return nil
}
func main() {
    c := &MyController{AppController{}, render.New(render.Options{})}
    http.ListenAndServe(":12000", c.Action(c.Index))
}