package Project1

import (
	"fmt"
	"net/http"
)

func indexView(w http.ResponseWriter,r *http.Request) {
	w.Write([]byte("page1"))
}
func main()  {
	http.HandleFunc('/',indexView)
	http.ListenAndServe(":8080",nil)
	fmt.println("run")
}