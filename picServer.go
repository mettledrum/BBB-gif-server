package main

import (
	"net/http"
	"os/exec"
	"log"
)

func runCamera() {
	cmd := exec.Command("sh", "-c", "./getGif.sh")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("waiting for camera")
	err = cmd.Wait()
	if err != nil {
		log.Printf("gif fetch finished with error: %v", err)
	} else {
		log.Println("gif fetch finished")
	}
}

func getGif(endpoint string, filename string) {
    http.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
    	runCamera()
        http.ServeFile(w, r, filename)
    })
}

func main() {
    getGif("/pic", "./output_folder/animate.gif")
    http.ListenAndServe(":80", nil)
}
