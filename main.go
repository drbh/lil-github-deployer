package main

import (
	// "fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message

	cmd := exec.Command("git", "pull")
	cmd.Dir = os.Args[1]
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	w.Write([]byte(string(out)))
}

func main() {
	http.HandleFunc("/", sayHello)
	if err := http.ListenAndServe(":9723", nil); err != nil {
		panic(err)
	}
}
