// Package snsmessage provides utility to prepare a SNS message to be
// structural and ready for validation.
package snsmessage

import (
	"encoding/json"

	"github.com/yuhlau/go-sns-message-validator/snserrors"
	"github.com/yuhlau/go-sns-message-validator/snsvalidator"
)

const (
	ErrMalformedJSON = "MalformedJSON"
)

// SNSMessage structure
// This structure can serve both Lambda and server end point SNS Message.
// SNS Message delivered to Lambda function will have its "*URL" fields end
// with camelcase "Url" instead of all uppdercase "URL". Go json.Unmarshal()
// will first try an exact match of struct field name of its tag, then accepts
// a case-insensitive match.
type SNSMessage struct {
	Type             string `json:"Type"`
	MessageId        string `json:"MessageId"`
	Token            string `json:"Token"`
	TopicArn         string `json:"TopicArn"`
	Message          string `json:"Message"`
	Subject          string `json:"Subject"`
	SubscribeURL     string `json:"SubscribeURL"`
	Timestamp        string `json:"Timestamp"`
	SignatureVersion string `json:"SignatureVersion"`
	Signature        string `json:"Signature"`
	SigningCertURL   string `json:"SigningCertURL"`
	UnsubscribeURL   string `json:"UnsubscribeURL"`
}

// Create a SNSMessage from JSON-encoded SNS message
func NewFromJSON(encoded []byte) (*SNSMessage, *snserrors.SNSError) {
	message := &SNSMessage{}
	err := json.Unmarshal(encoded, message)
	if err != nil {
		return nil, snserrors.New(ErrMalformedJSON, err.Error())
	}

	return message, nil
}

// Transform the SNSMessage structure to map
func (message *SNSMessage) toMap() map[string]string {
	return map[string]string{
		"Type":             message.Type,
		"MessageId":        message.MessageId,
		"Token":            message.Token,
		"TopicArn":         message.TopicArn,
		"Message":          message.Message,
		"Timestamp":        message.Timestamp,
		"Subject":          message.Subject,
		"SubscribeURL":     message.SubscribeURL,
		"SignatureVersion": message.SignatureVersion,
		"Signature":        message.Signature,
		"SigningCertURL":   message.SigningCertURL,
		"UnsubscribeURL":   message.UnsubscribeURL,
	}
}

// Get the SNSValidator of the SNSMessage
func (message *SNSMessage) GetValidator() *snsvalidator.SNSValidator {
	return snsvalidator.NewV1(message.toMap())
}
