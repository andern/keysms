# Import

`import "github.com/andern/keysms"`

# Send SMS

Sending an SMS to one or multiple recipients is very simple:

`keysms.Auth("username", "apiKey")`
`keysms.SendSMS("text message here", "98765432", "12345678")`

To change sender or to send an SMS in the future you should use `SMSParam`:
`
msg := keysms.SMSParams{
    Message:   "text message here",
    Receivers: []string{"98765432", "12345678"},
    Sender:    "99999999",
    Date:      "2017-11-27",
    Time:      "22:00",
}
keysms.Send(msg)
`
