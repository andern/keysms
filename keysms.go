package keysms

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"net/url"
)

const (
	messageURI = "http://app.keysms.no/messages"
	infoURI    = "http://app.keysms.no/auth/current.json"
)

var (
	username string
	apiKey   string
)

type SMSParams struct {
	// Text content of the SMS to send
	Message string `json:"message"`

	// The receivers of the SMS
	Receivers []string `json:"receivers"`

	// The number the SMS appears to be sent from
	Sender string `json:"sender,omitempty"`

	// Date when the SMS should be sent (YYYY-MM-DD)
	Date string `json:"date,omitempty"`

	// Time when the SMS should be sent (HH:mm)
	Time string `json:"time,omitempty"`
}

func Auth(user, key string) {
	username = user
	apiKey = key
}

func Send(params SMSParams) (SMSResponse, error) {
	payload, err := json.Marshal(params)
	if err != nil {
		return SMSResponse{}, err
	}

	resp, err := call(messageURI, string(payload))
	if err != nil {
		return SMSResponse{}, err
	}

	var smsres SMSResponse
	err = json.NewDecoder(resp.Body).Decode(&smsres)
	return smsres, err
}

func SendSMS(message string, receivers ...string) (SMSResponse, error) {
	return Send(SMSParams{
		Message:   message,
		Receivers: receivers,
	})
}

func sign(payload string) string {
	hash := md5.Sum([]byte(payload + apiKey))
	return hex.EncodeToString(hash[:])
}

func call(uri string, payload string) (*http.Response, error) {
	param := "payload=" + payload +
		"&signature=" + sign(payload) +
		"&username=" + username

	param = url.PathEscape(param)
	contentType := "application/x-www-form-urlencoded"
	data := bytes.NewBuffer([]byte(param))
	return http.Post(uri, "application/x-www-form-urlencoded", bytes.NewBuffer([]byte(param)))
}
