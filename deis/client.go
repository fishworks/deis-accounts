package deis

// Account represents a user's session with a controller.
type Account struct {
	Controller string `json:"controller"`
	Token      string `json:"token"`
	Username   string `json:"username"`
}

// Client is a mock-up of the relevant parts of ~/.deis/client.json
type Client struct {
	Controller string     `json:"controller"`
	Token      string     `json:"token"`
	Username   string     `json:"username"`
	Accounts   []*Account `json:"accounts"`
}
