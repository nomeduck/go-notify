package notify

func New(clients ...IClient) *Notify {
	n := new(Notify).clone()
	for _, client := range clients {
		n.Extend(client.Name(), client)
	}
	return n
}

// Notify 通知实体
type Notify struct {
	names   []string
	clients map[string]IClient
}

// 初始化参数
func (n Notify) clone() *Notify {
	n.clients = make(map[string]IClient)
	return &n
}

// Extend 注册自定义网关
func (n *Notify) Extend(name string, client IClient) {
	n.names = append(n.names, name)
	n.clients[name] = client
}

// Names 使用别名发送
func (n *Notify) Names(names ...string) *Notify {
	n.names = names
	return n
}

// Send 发送
func (n *Notify) Send(message IMessage) map[string]IResult {
	results := make(map[string]IResult)
	for _, name := range n.names {
		if client, ok := n.clients[name]; ok {
			message = message.I()
			results[name] = client.Send(message)
		}
	}
	return results
}
