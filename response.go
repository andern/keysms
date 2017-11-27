package keysms

type SMSResponse struct {
	OK       bool    `json:"ok"`
	Message  Message `json:"message"`
	Quantity int32   `json:"quantity"`
	Cost     float64 `json:"cost"`
	SMSPrice float64 `json:"smsPrice"`
}

type Message struct {
	ID        string     `json:"_id"`
	Sent      bool       `json:"sent"`
	Updated   string     `json:"updated"`
	Receivers []Receiver `json:"receivers"`
	Parts     Parts      `json:"parts"`
	Created   string     `json:"created"`
	Message   string     `json:"message"`
	Sender    string     `json:"sender"`
	Tags      bool       `json:"tags"`
	Groups    []string   `json:"groups"`
	Future    bool       `json:"future"`
	Status    Status     `json:"status"`
}

type Receiver struct {
	Number         string `json:"number"`
	Prefix         string `json:"prefix"`
	Country        string `json:"country"`
	DeliveryStatus int32  `json:"deliverystatus"`
	NextGWSyncTime string `json:"nextgwsynctime"`
}

type Parts struct {
	Total int32    `json:"total"`
	Parts []string `json:"parts"`
}

type Status struct {
	Value     string   `json:"value"`
	Text      string   `json:"text"`
	Aggregate []string `json:"aggregate"`
	Timed     int32    `json:"timed"`
}
