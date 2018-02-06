# Import

`import "github.com/andern/keysms"`

# Send SMS

Sending an SMS to one or multiple recipients is very simple:

```golang
keysms.Auth("username", "apiKey")
keysms.SendSMS("text message here", "98765432", "12345678")
```

To change sender or to send an SMS in the future you should use `SMSParam`:

```golang
msg := keysms.SMSParams{
    Message:   "text message here",
    Receivers: []string{"98765432", "12345678"},
    Sender:    "99999999",
    Time:      time.Now().Add(2 * time.Hour),
}
keysms.Send(msg)
```
