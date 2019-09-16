package main

import (
	"fmt"
	"net/http"
    
    "github.com/kaznishi/realize_sample/src/date"
)

func main() {
	http.HandleFunc("/", todayHandler)
	http.ListenAndServe(":8080", nil)
}

func todayHandler(w http.ResponseWriter, r *http.Request) {
	today := date.GetTodayCivil()
	fmt.Fprintf(w, "Today is "+today)
}
