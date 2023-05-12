package main

import (
	"fmt"
	"io"
	"net/http"
	"poc-google-calendar/events"
)

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/events", getEvents)

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Printf("ListenAndServe error: %v\n", err)
	}
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	code := r.FormValue("code")
	if code != "" {
		io.WriteString(w, "Your code is: \n"+code+"\n")
	}
}

func getEvents(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /events request\n")
	evs := events.FetchEvents()
	if len(evs.Items) == 0 {
		io.WriteString(w, "No upcoming events found.")
	} else {
		io.WriteString(w, "Upcoming events:\n")
		for _, item := range evs.Items {
			date := item.Start.DateTime
			if date == "" {
				date = item.Start.Date
			}
			io.WriteString(w, date+" - "+item.Summary+"\n")
		}
	}
}
