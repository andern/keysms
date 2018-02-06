package keysms

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"net/url"
	"time"
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
	Message string

	// The recipients of the SMS
	Recipients []string

	// The number the SMS appears to be sent from
	Sender string

	// Time when the SMS should be sent
	Time time.Time
}

type smsParams struct {
	Message   string   `json:"message"`
	Receivers []string `json:"receivers"`
	Sender    string   `json:"sender,omitempty"`
	// YYYY-MM-DD
	Date string `json:"date,omitempty"`
	// HH:mm
	Time string `json:"time,omitempty"`
}

func Auth(user, key string) {
	username = user
	apiKey = key
}

func Send(params SMSParams) (SMSResponse, error) {
	payload, err := json.Marshal(payloadParams(params))
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

func SendSMS(message string, recipients ...string) (SMSResponse, error) {
	return Send(SMSParams{
		Message:    message,
		Recipients: recipients,
	})
}

func payloadParams(p SMSParams) (res smsParams) {
	res.Message = p.Message
	res.Receivers = p.Recipients
	res.Sender = p.Sender
	if !p.Time.IsZero() {
		res.Date = p.Time.Format("2006-01-02")
		res.Time = p.Time.Format("15:04")
	}
	return
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
	return http.Post(uri, contentType, data)
}
