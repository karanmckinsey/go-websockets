package handlers


func NewHub() *Hub {
	return &Hub{
		clients			: 	make(map[*Client]bool),
		register		: 	make(chan *Client),
		unregister		: 	make(chan *Client),
	}
}

func (hub *Hub) Run() {
	for {
		select {
		case client := <- hub.register:
			UserRegisterEventHandler(hub, client)
		case client := <- hub.unregister:
			UserUnregesterEventHandler(hub, client)
		}
	}
}