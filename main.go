package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	http.HandleFunc("/", ExecHandler)
	port := os.Getenv("PORT")
	fmt.Println("Listening on port: " + port)
	http.ListenAndServe(":"+port, nil)
}

func ExecHandler(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("./ag", "-l", r.URL.Path[1:], "data")
	var out bytes.Buffer
	var oerr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &oerr

	fmt.Println("ag ", r.URL.Path[1:], "data")
	err := cmd.Run()
	fmt.Println(err)

	out.WriteTo(w)
}
