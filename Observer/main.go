package main

import (
	"design-patterns-Golang/Observer/observer"
	"time"
)

func main() {

	// Initialize a new Notifier.
	n := observer.NewsAgency{
		Subscribers: make(map[observer.Subscriber]struct{}),
	}

	// Register a couple of Observers/Subscriber.
	n.Register(&observer.TVChannel{Name: "CNN"})
	n.Register(&observer.NewsWebsite{URL: "https://www.news.com"})

	ticker := time.NewTicker(1 * time.Second)
	timer := time.NewTimer(10 * time.Second)

	for {
		select {
		case <-ticker.C:
			n.Notify(observer.Event{Type: "breaking", Payload: "something random"})
		case <-timer.C:
			n.DeRegister(&observer.TVChannel{Name: "CNN"})
			return
		}
	}
}
