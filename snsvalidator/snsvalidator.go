// Package snsvalidator is the core of SNS message validation. Its sole purpose
// is to validate a SNS message matches its signature.
package snsvalidator

import (
	"bytes"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"

	"github.com/yuhlau/go-sns-message-validator/snserrors"
)

const (
	TypeNotification             = "Notification"
	TypeSubscriptionConfirmation = "SubscriptionConfirmation"
	TypeUnsubscribeConfirmation  = "UnsubscribeConfirmation"
)

const (
	ErrMissingKey         = "MissingKey"
	ErrInvalidType        = "InvalidType"
	ErrInvalidCert        = "InvalidCert"
	ErrIncorrectSignature = "IncorrectSignature"
)

// List of AWS Signing Certificate URL trustable hosts
// sns.<region>.amazonaws.com		(AWS)
// sns.us-gov-west-1.amazonaws.com	(AWS GovCloud)
// sns.cn-north-1.amazonaws.com.cn 	(AWS China)
const defaultHostPattern = `^sns\.[a-zA-Z0-9\-]{3,}\.amazonaws\.com(\.cn)?$`

var defaultHostPatternRegexp = regexp.MustCompile(defaultHostPattern)

// Required keys of different structures are referenced from
// http://docs.aws.amazon.com/sns/latest/dg/json-formats.html

// Required keys of a SNS Message
var requiredKeys = []string{
	"Type",
	"MessageId",
	"TopicArn",
	"Message",
	"Timestamp",
	"Signature",
	"SignatureVersion",
	"SigningCertURL",
}

// required keys for SubscriptionConfirmation and UnsubscribeConfirmation
var requiredSubscriptionKeys = []string{
	"SubscribeURL",
	"Token",
}

// List of valid message types
var validMessageTypes = []string{
	TypeSubscriptionConfirmation,
	TypeUnsubscribeConfirmation,
	TypeNotification,
}

// List of Subscription message types
var subscriptionMessageTypes = []string{
	TypeSubscriptionConfirmation,
	TypeUnsubscribeConfirmation,
}

// The order of signable keys must be same when bulding the signable string
// Signable keys for SubscriptionConfirmation and UnsubscribeConfirmation
var signableKeysForSubscription = []string{
	"Message",
	"MessageId",
	// "Subject" should not appear in Subscription. Keep here just to be safe
	"Subject",
	"SubscribeURL",
	"Timestamp",
	"Token",
	"TopicArn",
	"Type",
}

// Signable keys for Notification
var signableKeysForNotification = []string{
	"Message",
	"MessageId",
	"Subject", // if included in the message
	"SubscribeURL",
	"Timestamp",
	"TopicArn",
	"Type",
}

type SNSValidator struct {
	Version    int
	MessageMap map[string]string
}

// NewV1 returns a new version 1 SNSValiator with the specified map as the
// message map.
func NewV1(messageMap map[string]string) *SNSValidator {
	return &SNSValidator{
		Version:    1,
		MessageMap: messageMap,
	}
}

// ValidateMessage validates the underlying SNS message by following the AWS
// specification.
// If the type is invalid, it returns SNSError of type ErrInvalidType
// If one or more of the required keys are missing, it returns SNSError of type
// ErrMissingKey
// If the certificate cannot be retrieved, it returns SNSError of
// ErrInvalidCert
// If the signature is incorrect, it returns SNSError of ErrIncorrectSignature
func (validator *SNSValidator) ValidateMessage() error {
	if err := validator.validateMessageStructure(); err != nil {
		return err
	}

	if err := validator.verifySignature(); err != nil {
		return err
	}

	return nil
}

// has returns boolean on whether the underlying SNS message has the key
// specified.
// Since json.Unmsrahsl() will leave missing key with a zero value - empty
// string. The method will also check for non-empty string.
func (validator *SNSValidator) has(key string) bool {
	if value, exists := validator.MessageMap[key]; exists && value != "" {
		return true
	}
	return false
}

// hasKeys returns boolean on whether the underlying SNS message map has all
// the keys specified in the slice.
// If one or more of the keys are missing. It also returns the name of the
// missing key.
func (validator *SNSValidator) hasKeys(keylist []string) (bool, string) {
	for _, key := range keylist {
		if !validator.has(key) {
			return false, key
		}
	}
	return true, ""
}

// isTypes returns boolean on whether the underlying SNS message is one of the
// specified types.
func (validator *SNSValidator) isTypes(typelist []string) bool {
	messageType := validator.MessageMap["Type"]
	for _, t := range typelist {
		if messageType == t {
			return true
		}
	}
	return false
}

// validateRequiredKeys validates the underlying SNS message contains all the
// required keys of a valid SNS message.
// If one or more of the required keys are missing, it returns an SNSError of
// type ErrMissingKey.
func (validator *SNSValidator) validateRequiredKeys() error {
	if valid, missingKey := validator.hasKeys(requiredKeys); !valid {
		return snserrors.New(
			ErrMissingKey,
			fmt.Sprintf("\"%s\" is required in SNS message", missingKey),
		)
	}
	return nil
}

// validateMessageType validates the underlying SNS message is one of the valid
// message types.
// If the message is none of the valid types, it returns an SNSError of type
// ErrInvalidType.
func (validator *SNSValidator) validateMessageType() error {
	if !validator.isTypes(validMessageTypes) {
		return snserrors.New(
			ErrInvalidType,
			fmt.Sprintf("Invalid message type \"%s\"", validator.MessageMap["Type"]),
		)
	}
	return nil
}

// validateSubscriptionKeys validates the underlying SNS message contains all
// the required keys of a Subscription message.
// If one or more of the required keys are missing, it returns an SNSError of
// type ErrMissingKey.
func (validator *SNSValidator) validateSubscriptionKeys() error {
	if valid, missingKey := validator.hasKeys(requiredSubscriptionKeys); !valid {
		return snserrors.New(
			ErrMissingKey,
			fmt.Sprintf("\"%s\" is required in Subscription message", missingKey),
		)
	}
	return nil
}

// validateMessageStructure validates the underlying SNS message has a valid
// SNS message structure. It validates it has a valid message type and contains
// all the required keys.
// If the type is invalid, it returns SNSError of type ErrInvalidType
// If one or more of the required keys are missing, it returns SNSError of type
// ErrMissingKey
func (validator *SNSValidator) validateMessageStructure() error {
	if err := validator.validateRequiredKeys(); err != nil {
		return err
	}

	if err := validator.validateMessageType(); err != nil {
		return err
	}

	// SubscriptionConfirmation or UnsubscriptionConfirmation message
	if validator.isTypes(subscriptionMessageTypes) {
		if err := validator.validateSubscriptionKeys(); err != nil {
			return err
		}
	}

	return nil
}

// buildSignableString returns signable string of the underlying SNS message.
// The signable string is essential to verify the signature.
func (validator *SNSValidator) buildSignableString() []byte {
	// 1000 bytes is a heuristic length for a signable strings
	signableString := bytes.NewBuffer(make([]byte, 0, 1000))

	var signableKeys []string
	if validator.isTypes(subscriptionMessageTypes) {
		signableKeys = signableKeysForSubscription
	} else {
		signableKeys = signableKeysForNotification
	}
	for _, key := range signableKeys {
		// Some keys like "Subject" are included only if it is present
		if validator.has(key) {
			signableString.WriteString(key)
			signableString.WriteString("\n")
			signableString.WriteString(validator.MessageMap[key])
			signableString.WriteString("\n")
		}
	}
	return signableString.Bytes()
}

// getCertificate tries to fetch the Signing Certificate that is used to sign
// the underlying SNS message, and returns the certifcate in slice of bytes.
// If the HTTP request fails, it also returns a SNSError of type ErrInvalidCert
// describing the error
func (validator *SNSValidator) getCertificate() ([]byte, error) {
	res, err := http.Get(validator.MessageMap["SigningCertURL"])
	if err != nil {
		return nil, snserrors.New(ErrInvalidCert, err.Error())
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, snserrors.New(ErrInvalidCert, "Could not retrive the certificate")
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, snserrors.New(ErrInvalidCert, err.Error())
	}

	return body, nil
}

// verifySignature verifieds the underlying SNS message signature is correct.
// If the certificate cannot be retrieved, it returns SNSError of type
// ErrInvalidCert
// If the signature is incorrect, it returns SNSError of type
// ErrIncorrectSignature
func (validator *SNSValidator) verifySignature() error {
	// Verify the SigningCertURL is trustworthy
	parsedUrl, err := url.Parse(validator.MessageMap["SigningCertURL"])
	if err != nil {
		return snserrors.New(ErrInvalidCert, err.Error())
	}

	if parsedUrl.Scheme != "https" {
		return snserrors.New(ErrInvalidCert, "The certificate URL is using insecure HTTP scheme")
	}

	if match := defaultHostPatternRegexp.MatchString(
		parsedUrl.Hostname(),
	); !match {
		return snserrors.New(ErrInvalidCert, "The certificate URL belongs to an untrusted host")
	}

	// Obtain the signing certificate
	certData, snserr := validator.getCertificate()
	if snserr != nil {
		return snserr
	}

	block, _ := pem.Decode(certData)
	if block == nil {
		return snserrors.New(ErrInvalidCert, "Could not decode the certificate")
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return snserrors.New(ErrInvalidCert, err.Error())
	}

	// base64 decode the signature given
	decodedSignature, err := base64.StdEncoding.DecodeString(validator.MessageMap["Signature"])
	if err != nil {
		return snserrors.New(ErrIncorrectSignature, "Could not base64 decode the signature")
	}

	// check for the validitly of signature
	if err := cert.CheckSignature(
		x509.SHA1WithRSA, validator.buildSignableString(), decodedSignature,
	); err != nil {
		return snserrors.New(ErrIncorrectSignature, fmt.Sprintf("Incorrect signature: %v", err))
	}
	return nil
}
