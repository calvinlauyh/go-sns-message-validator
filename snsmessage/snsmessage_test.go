package snsmessage

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/yuhlau/go-sns-message-validator/snsvalidator"
)

// func shouldBeInnerTypeNil(actual interface{}, expected ...interface{}) string {
//     if <some-important-condition-is-met(actual, expected)> {
//         return ""   // empty string means the assertion passed
//     } else {
//         return "<some descriptive message detailing why the assertion failed...>"
//     }
// }

// Message Template
var NotificationMessage = SNSMessage{
	Type:             "Notification",
	MessageId:        "165545c9-2a5c-472c-8df2-7ff2be2b3b1b",
	Token:            "2336412f37fb687f5d51e6e241d09c805a5a",
	TopicArn:         "arn:aws:sns:us-west-2:123456789012:MyTopic",
	Message:          "Test notification",
	Subject:          "Test subject",
	Timestamp:        "2012-04-26T20:45:04.751Z",
	SignatureVersion: "1",
	Signature:        "EXAMPLEpH+DcEwjAPg8O9mY8dReBSwksfg2S=",
	SigningCertURL:   "https://localhost/cert.pem",
	UnsubscribeURL:   "https://localhost/unsubscribe",
}
var SubscriptionMessage = SNSMessage{
	Type:             "SubscriptionConfirmation",
	MessageId:        "165545c9-2a5c-472c-8df2-7ff2be2b3b1b",
	Token:            "2336412f37fb687f5d51e6e241d09c805a5a",
	TopicArn:         "arn:aws:sns:us-west-2:123456789012:MyTopic",
	Message:          "You have chosen to subscribe to the topic",
	SubscribeURL:     "https://localhost/subscribe",
	Timestamp:        "2012-04-26T20:45:04.751Z",
	SignatureVersion: "1",
	Signature:        "EXAMPLEpH+DcEwjAPg8O9mY8dReBSwksfg2S=",
	SigningCertURL:   "https://localhost/cert.pem",
}
var UnsubscriptionMessage = SNSMessage{
	Type:             "UnubscriptionConfirmation",
	MessageId:        "165545c9-2a5c-472c-8df2-7ff2be2b3b1b",
	Token:            "2336412f37fb687f5d51e6e241d09c805a5a",
	TopicArn:         "arn:aws:sns:us-west-2:123456789012:MyTopic",
	Message:          "You have chosen to unsubscribe to the topic",
	SubscribeURL:     "https://localhost/subscribe",
	Timestamp:        "2012-04-26T20:45:04.751Z",
	SignatureVersion: "1",
	Signature:        "EXAMPLEpH+DcEwjAPg8O9mY8dReBSwksfg2S=",
	SigningCertURL:   "https://localhost/cert.pem",
}

func TestNewFromJSON(t *testing.T) {
	Convey("Given a valid JSON-encoded SNS message", t, func() {
		encoded := []byte(`{
  "Type": "SubscriptionConfirmation",
  "MessageId": "165545c9-2a5c-472c-8df2-7ff2be2b3b1b",
  "Token": "2336412f37fb687f5d51e6e241d09c805a5a",
  "TopicArn": "arn:aws:sns:us-west-2:123456789012:MyTopic",
  "Message": "You have chosen to subscribe to the topic",
  "SubscribeURL": "https://localhost/subscribe",
  "Timestamp": "2012-04-26T20:45:04.751Z",
  "SignatureVersion": "1",
  "Signature": "EXAMPLEpH+DcEwjAPg8O9mY8dReBSwksfg2S=",
  "SigningCertURL": "https://localhost/cert.pem"
}`)
		Convey("It should succeeded and return SNSMessage representation of the message", func() {
			expected := SubscriptionMessage
			message, err := NewFromJSON(encoded)

			So(*message, ShouldResemble, expected)
			So(err, ShouldBeNil)
		})

		// Details: https://golang.org/doc/faq#nil_error
		Convey("When the placeholder of the second return value is error-typed", func() {
			var actual error
			Convey("It should return an interface value nil", func() {
				_, actual = NewFromJSON(encoded)

				So(actual == nil, ShouldBeTrue)
			})
		})
	})

	Convey("Given a malformed JSON message", t, func() {
		encoded := []byte(`{"Type": "SubscriptionConfirmation`)

		Convey("It should return a SNSError of malformed JSON", func() {
			message, err := NewFromJSON(encoded)

			So(message, ShouldBeNil)
			So(err.Type(), ShouldEqual, ErrMalformedJSON)
		})
	})

	Convey("Given a JSON-encoded Lambda SNS message with \"SubscribeUrl\" and \"SigningCertUrl\"", t, func() {
		encoded := []byte(`{
  "Type": "SubscriptionConfirmation",
  "MessageId": "165545c9-2a5c-472c-8df2-7ff2be2b3b1b",
  "Token": "2336412f37fb687f5d51e6e241d09c805a5a",
  "TopicArn": "arn:aws:sns:us-west-2:123456789012:MyTopic",
  "Message": "You have chosen to subscribe to the topic",
  "SubscribeUrl": "https://localhost/subscribe",
  "Timestamp": "2012-04-26T20:45:04.751Z",
  "SignatureVersion": "1",
  "Signature": "EXAMPLEpH+DcEwjAPg8O9mY8dReBSwksfg2S=",
  "SigningCertUrl": "https://localhost/cert.pem"
}`)

		Convey(`It should return SNSMessage with "*URL" fields be the "*Url" fields in the orignla JSON`, func() {
			message, _ := NewFromJSON(encoded)

			So(message.SubscribeURL, ShouldEqual, "https://localhost/subscribe")
			So(message.SigningCertURL, ShouldEqual, "https://localhost/cert.pem")
		})
	})
}

func TestToMapMethod(t *testing.T) {
	Convey("Given a SNSMessage structure", t, func() {
		message := SNSMessage{
			Type:             "SubscriptionConfirmation",
			MessageId:        "165545c9-2a5c-472c-8df2-7ff2be2b3b1b",
			Token:            "2336412f37fb687f5d51e6e241d09c805a5a",
			TopicArn:         "arn:aws:sns:us-west-2:123456789012:MyTopic",
			Message:          "You have chosen to subscribe to the topic",
			SubscribeURL:     "https://localhost/subscribe",
			Timestamp:        "2012-04-26T20:45:04.751Z",
			SignatureVersion: "1",
			Signature:        "EXAMPLEpH+DcEwjAPg8O9mY8dReBSwksfg2S=",
			SigningCertURL:   "https://localhost/cert.pem",
		}

		Convey("It should return map reprsentation of the SNSMessage", func() {
			expected := map[string]string{
				"Type":             "SubscriptionConfirmation",
				"MessageId":        "165545c9-2a5c-472c-8df2-7ff2be2b3b1b",
				"Token":            "2336412f37fb687f5d51e6e241d09c805a5a",
				"TopicArn":         "arn:aws:sns:us-west-2:123456789012:MyTopic",
				"Message":          "You have chosen to subscribe to the topic",
				"Subject":          "",
				"SubscribeURL":     "https://localhost/subscribe",
				"Timestamp":        "2012-04-26T20:45:04.751Z",
				"SignatureVersion": "1",
				"Signature":        "EXAMPLEpH+DcEwjAPg8O9mY8dReBSwksfg2S=",
				"SigningCertURL":   "https://localhost/cert.pem",
				"UnsubscribeURL":   "",
			}
			actual := message.toMap()

			So(actual, ShouldResemble, expected)
		})
	})
}

func TestGetValidatorMethod(t *testing.T) {
	Convey("Given a version 1 SNSMessage", t, func() {
		message := SNSMessage{
			Type:             "Notification",
			MessageId:        "165545c9-2a5c-472c-8df2-7ff2be2b3b1b",
			Token:            "2336412f37fb687f5d51e6e241d09c805a5a",
			TopicArn:         "arn:aws:sns:us-west-2:123456789012:MyTopic",
			Message:          "Test notification",
			Subject:          "Test subject",
			Timestamp:        "2012-04-26T20:45:04.751Z",
			SignatureVersion: "1",
			Signature:        "EXAMPLEpH+DcEwjAPg8O9mY8dReBSwksfg2S=",
			SigningCertURL:   "https://localhost/cert.pem",
			UnsubscribeURL:   "https://localhost/unsubscribe",
		}

		Convey("It should return SNSValidator", func() {
			actual := message.GetValidator()

			So(actual, ShouldHaveSameTypeAs, &snsvalidator.SNSValidator{})
			Convey("Returned SNSValidator should be version 1", func() {
				So(actual.Version, ShouldEqual, 1)
			})
			Convey("Returned SNSValidator should have a MessageMap of the original message", func() {
				So(actual.MessageMap, ShouldResemble, map[string]string{
					"Type":             "Notification",
					"MessageId":        "165545c9-2a5c-472c-8df2-7ff2be2b3b1b",
					"Token":            "2336412f37fb687f5d51e6e241d09c805a5a",
					"TopicArn":         "arn:aws:sns:us-west-2:123456789012:MyTopic",
					"Message":          "Test notification",
					"Subject":          "Test subject",
					"SubscribeURL":     "",
					"Timestamp":        "2012-04-26T20:45:04.751Z",
					"SignatureVersion": "1",
					"Signature":        "EXAMPLEpH+DcEwjAPg8O9mY8dReBSwksfg2S=",
					"SigningCertURL":   "https://localhost/cert.pem",
					"UnsubscribeURL":   "https://localhost/unsubscribe",
				})
			})
		})
	})
}
