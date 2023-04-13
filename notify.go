package notify

import (
	"golang.org/x/sync/errgroup"
)

func New(clients ...IClient) *Notify {
	n := new(Notify)
	for _, client := range clients {
		n.Extend(client.Name(), client)
	}
	return n
}

type Notify struct {
	names   []string
	clients map[string]IClient
}

func (n *Notify) Extend(name string, client IClient) {
	if n.clients == nil {
		n.clients = make(map[string]IClient)
	}
	n.names = append(n.names, name)
	n.clients[name] = client
}

func (n *Notify) Names(names ...string) *Notify {
	n.names = names
	return n
}

func (n *Notify) Sender(message IMessage) (map[string]Result, error) {
	var g errgroup.Group
	results := make(map[string]Result)
	for _, name := range n.names {
		if client, ok := n.clients[name]; ok {
			g.Go(func() error {
				result, err := client.Send(message)
				if err == nil {
					results[name] = result
				}
				return err
			})
		}
	}
	if err := g.Wait(); err != nil {
		return results, err
	}
	return results, nil
}
