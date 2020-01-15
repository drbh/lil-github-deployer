package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

func executeShell(name string, shell_script string) {
	will_execute := "bash " + shell_script
	cmd := exec.Command("tmux", "new", "-d", "-s", name, will_execute)
	_, erro := cmd.CombinedOutput()
	if erro != nil {
		log.Fatalf("cmd.Run() failed with %s\n", erro)
	}
	fmt.Println("Started server in tmux session")
}

func killSession() {
	// this kills all open tmux session
	cmd := exec.Command("tmux", "send-keys", "C-c")
	out, erro := cmd.CombinedOutput()
	if erro != nil {

		fmt.Printf("%s\n", erro)
		// log.Fatalf("cmd.Run() failed with %s\n", erro)
		fmt.Println("No session was running")
		return
	}
	fmt.Println(string(out))
	fmt.Println("Killed tmux session")
}

func pullLatest(dir string) {
	cmd := exec.Command("git", "pull")
	cmd.Dir = dir
	_, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Println("Pull latest git repo")
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message

	name := os.Args[1]
	dir := os.Args[2]
	shell_script := os.Args[3]

	pullLatest(dir)
	killSession()

	fmt.Println("Sleep before reinit")
	time.Sleep(time.Duration(1) * time.Second)

	executeShell(name, shell_script)

	w.Write([]byte(string("200\n")))
}

func main() {
	fmt.Println("Starting deployment server")
	http.HandleFunc("/", sayHello)
	if err := http.ListenAndServe(":9723", nil); err != nil {
		panic(err)
	}
}
