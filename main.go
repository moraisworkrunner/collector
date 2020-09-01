package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	work_messages "github.com/moraisworkrunner/work-messages"
)

func handler(w http.ResponseWriter, r *http.Request) {
	workResponse := work_messages.SvcWorkResponse{}
	target := os.Getenv("INCOMING_QUEUE")
	if target == "" {
		target = "queue"
	}
	location := os.Getenv("INCOMING_LOCATION")
	if location == "" {
		location = "nowhere"
	}
	serviceURL := os.Getenv("SERVICE_URL")
	if serviceURL == "" {
		serviceURL = "http://:8080"
	}
	// Parse the body to send to the queue
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		workResponse.Error = &work_messages.Error{
			Message: fmt.Sprintf("Failed to read request"),
		}
		fmt.Printf("Failed to read work request: %v\n", r.Body)
		w.WriteHeader(http.StatusBadRequest)
		// Do not retry
		return
	}
	if _, err := createTask("moraisworkrunner", location, target, serviceURL, string(body)); err != nil {
		fmt.Printf("Failed to create task: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}

func main() {
	fmt.Print("starting server...\n")
	http.HandleFunc("/", handler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	fmt.Printf("listening on port %s\n", port)
	fmt.Print(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
	os.Exit(0)
}
