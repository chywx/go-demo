package main

import (
	"fmt"
	"net/http"
)

func main()  {
	fmt.Print(123)

	http.Handle("/", http.FileServer(http.Dir(".")))

	http.ListenAndServe(":8080", nil)
	
}
