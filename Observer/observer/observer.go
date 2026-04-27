package observer

import "fmt"

type Event struct {
	Type    string
	Payload any
}

type Publisher interface {
	// Register allows a subscriber to register/subscribe itself to listen for an event.
	Register(Subscriber)
	// DeRegister allows a subscriber to de-register/un-subscribe itself from listening for an event.
	DeRegister(Subscriber)
	// Notify publishes new events to observer.
	Notify(Event)
}

// Subscriber is anything that can receive an event.
type Subscriber interface {
	OnNotify(Event)
}

type NewsAgency struct {
	name string
	// contains the list of subscriber
	Subscribers map[Subscriber]struct{}
}

func (n *NewsAgency) Register(s Subscriber) {
	fmt.Printf("Registering the subscriber: ")
	n.Subscribers[s] = struct{}{}
}

func (n *NewsAgency) DeRegister(s Subscriber) {
	fmt.Printf("De-registering the subscriber")
	delete(n.Subscribers, s)
}

func (n *NewsAgency) Notify(e Event) {
	fmt.Printf("Notifing the subscriber from:")
	for s := range n.Subscribers {
		s.OnNotify(e)
	}
}

// Concrete observers
type TVChannel struct{ Name string }
type NewsWebsite struct{ URL string }

func (t *TVChannel) OnNotify(e Event) {
	fmt.Printf("%s: News! %s\n", t.Name, e)
}

func (n *NewsWebsite) OnNotify(e Event) {
	fmt.Printf("%s: Breaking news! %s\n", n.URL, e)
}
