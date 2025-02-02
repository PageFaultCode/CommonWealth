// package CommonWealth are all of the assets that various microservices might need to share
package CommonWealth

import (
	"github.com/nats-io/nats.go"
)

type NatsRegistration struct {
	subscriptions []*nats.Subscription
}

func (reg *NatsRegistration) subscribeVerb(name, verb string, nc *nats.Conn, callback nats.MsgHandler) error {
	sub, suberr := nc.Subscribe(name+"/"+verb, callback)

	if suberr != nil {
		return suberr
	}
	reg.subscriptions = append(reg.subscriptions, sub)
	return nil
}

// ConfigureNATsSystem creates a nats connection and registers callbacks
func (reg *NatsRegistration) ConfigureNATsSystem() (*nats.Conn, error) {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		return nil, err
	}
	println("CommonWealth : Connected to nats")

	return nc, nil
}

// Subscribe allows users to subscribe to nats topics
func (reg *NatsRegistration) Subscribe(nc *nats.Conn, topic NatsTopics, callback nats.MsgHandler) error {
	reg.subscriptions = make([]*nats.Subscription, 0)

	// Register user requests
	if topic.Always {
		sub, suberr := nc.Subscribe(topic.Name, callback)

		if suberr != nil {
			return suberr
		}
		reg.subscriptions = append(reg.subscriptions, sub)
	} else {
		if topic.Create {
			suberr := reg.subscribeVerb(topic.Name, Create, nc, callback)
			if suberr != nil {
				return suberr
			}
		}
		if topic.Read {
			suberr := reg.subscribeVerb(topic.Name, Read, nc, callback)
			if suberr != nil {
				return suberr
			}
		}
		if topic.Update {
			suberr := reg.subscribeVerb(topic.Name, Update, nc, callback)
			if suberr != nil {
				return suberr
			}
		}
		if topic.Delete {
			suberr := reg.subscribeVerb(topic.Name, Delete, nc, callback)
			if suberr != nil {
				return suberr
			}
		}
	}

	return nil
}
