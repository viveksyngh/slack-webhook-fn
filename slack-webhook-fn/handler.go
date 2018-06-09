package function

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//Payload structure for slack API
type Payload struct {
	Text string `json:"text"`
}

// Handle a serverless request
func Handle(req []byte) string {
	webhook, err := ioutil.ReadFile("/var/openfaas/secrets/webhook-url")
	webhookURL := string(webhook)
	if err != nil {
		return fmt.Sprintf("Unable to read secret file")
	}

	payload := Payload{Text: string(req)}
	reqPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Sprintf("Unable to marshal")
	}
	reader := bytes.NewReader(reqPayload)

	client := http.Client{}
	request, _ := http.NewRequest(http.MethodPost, webhookURL, reader)
	_, err = client.Do(request)
	if err != nil {
		return fmt.Sprintf("Unable to send message")
	}

	return fmt.Sprintf("Message sent successfully")
}
