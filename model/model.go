package model

type HostNameIPStatus struct {
	IP     string
	Status bool
}

type GetHostNameResponse struct {
	Result []string `json:"result"`
	Status string   `json:"status"`
	Error  error    `json:"error"`
}
