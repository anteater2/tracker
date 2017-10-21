package models

// Client information
type Client struct {
	Addr   string
	Name   string
	secret int64
	id     int64
}

type ClientList []Client

var clients = make(ClientList, 0, 0)

// AddClient adds a client
func AddClient(c Client) {
	clients = append(clients, c)
}

// RemoveClient removes a client
func RemoveClient(c Client) {
	for i, c := range clients {
		if clients[i] == c {
			clients = append(clients[:i], clients[i+1:]...)
		}
	}
}

// GetClients returns a slice of all the clients
func GetClients() ClientList {
	return clients
}

// NewClient creates a new client
func NewClient(name string, addr string) Client {
	return Client{
		Name: name,
		Addr: addr,
	}
}

func (cs ClientList) String() string {
	s := ""
	for _, c := range cs {
		s += c.Name + ":" + c.Addr + " "
	}
	return s
}

// Secret returns the client's secret
func (c Client) Secret() int64 {
	return c.secret
}

// ID returns the client's ID
func (c Client) ID() int64 {
	return c.id
}
