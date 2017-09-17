package snsvalidator

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/h2non/gock.v1"

	"github.com/yuhlau/go-sns-message-validator/snserrors"
)

// Message Template
func newNotificationMessageValidator() SNSValidator {
	return SNSValidator{
		MessageMap: map[string]string{
			"Type":             "Notification",
			"MessageId":        "165545c9-2a5c-472c-8df2-7ff2be2b3b1b",
			"Token":            "",
			"TopicArn":         "arn:aws:sns:us-west-2:123456789012:MyTopic",
			"Message":          "Test notification",
			"SubscribeURL":     "",
			"Subject":          "Test subject",
			"Timestamp":        "2012-04-26T20:45:04.751Z",
			"SignatureVersion": "1",
			"Signature":        "EXAMPLEpH+DcEwjAPg8O9mY8dReBSwksfg2S=",
			"SigningCertURL":   "https://localhost/cert.pem",
			"UnsubscribeURL":   "https://localhost/unsubscribe",
		},
	}
}

func newSubscriptionMessageValidator() SNSValidator {
	return SNSValidator{
		Version: 1,
		MessageMap: map[string]string{
			"Type":             "SubscriptionConfirmation",
			"MessageId":        "165545c9-2a5c-472c-8df2-7ff2be2b3b1b",
			"Token":            "2336412f37fb687f5d51e6e241d09c805a5a",
			"TopicArn":         "arn:aws:sns:us-west-2:123456789012:MyTopic",
			"Message":          "You have chosen to subscribe to the topic",
			"SubscribeURL":     "https://localhost/subscribe",
			"Subject":          "",
			"Timestamp":        "2012-04-26T20:45:04.751Z",
			"SignatureVersion": "1",
			"Signature":        "EXAMPLEpH+DcEwjAPg8O9mY8dReBSwksfg2S=",
			"SigningCertURL":   "https://localhost/cert.pem",
			"UnsubscribeURL":   "",
		},
	}
}

func newUnsubscriptionMessageValidator() SNSValidator {
	return SNSValidator{
		Version: 1,
		MessageMap: map[string]string{
			"Type":             "UnubscriptionConfirmation",
			"MessageId":        "165545c9-2a5c-472c-8df2-7ff2be2b3b1b",
			"Token":            "2336412f37fb687f5d51e6e241d09c805a5a",
			"TopicArn":         "arn:aws:sns:us-west-2:123456789012:MyTopic",
			"Message":          "You have chosen to unsubscribe to the topic",
			"SubscribeURL":     "https://localhost/subscribe",
			"Subject":          "",
			"Timestamp":        "2012-04-26T20:45:04.751Z",
			"SignatureVersion": "1",
			"Signature":        "EXAMPLEpH+DcEwjAPg8O9mY8dReBSwksfg2S=",
			"SigningCertURL":   "https://localhost/cert.pem",
			"UnsubscribeURL":   "",
		},
	}
}

// Unit test
func TestNewV1(t *testing.T) {
	Convey("Given a SNS message in map", t, func() {
		messageMap := map[string]string{
			"Type":             "Notification",
			"MessageId":        "165545c9-2a5c-472c-8df2-7ff2be2b3b1b",
			"Token":            "",
			"TopicArn":         "arn:aws:sns:us-west-2:123456789012:MyTopic",
			"Message":          "Test notification",
			"SubscribeURL":     "",
			"Subject":          "Test subject",
			"Timestamp":        "2012-04-26T20:45:04.751Z",
			"SignatureVersion": "1",
			"Signature":        "EXAMPLEpH+DcEwjAPg8O9mY8dReBSwksfg2S=",
			"SigningCertURL":   "https://localhost/cert.pem",
			"UnsubscribeURL":   "https://localhost/unsubscribe",
		}

		Convey("It should return version 1 SNS Validator of the message", func() {
			actual := NewV1(messageMap)

			So(actual.Version, ShouldEqual, 1)
			So(actual.MessageMap, ShouldResemble, messageMap)
		})
	})
}

func TestValidateMessageMethod(t *testing.T) {
	// Mock HTTP request-response
	certData := `-----BEGIN CERTIFICATE-----
MIIC8zCCAlwCCQCLHrKJpPLt9TANBgkqhkiG9w0BAQsFADCBvDELMAkGA1UEBhMC
SEsxEjAQBgNVBAgMCUhvbmctS29uZzESMBAGA1UEBwwJSG9uZy1Lb25nMSEwHwYD
VQQKDBhnby1zbnMtbWVzc2FnZS12YWxpZGF0b3IxITAfBgNVBAsMGGdvLXNucy1t
ZXNzYWdlLXZhbGlkYXRvcjEhMB8GA1UEAwwYZ28tc25zLW1lc3NhZ2UtdmFsaWRh
dG9yMRwwGgYJKoZIhvcNAQkBFg1pYW1AeXVobGF1Lm1lMCAXDTE3MDkxNzE3MDA1
MloYDzIwNjcwOTA1MTcwMDUyWjCBvDELMAkGA1UEBhMCSEsxEjAQBgNVBAgMCUhv
bmctS29uZzESMBAGA1UEBwwJSG9uZy1Lb25nMSEwHwYDVQQKDBhnby1zbnMtbWVz
c2FnZS12YWxpZGF0b3IxITAfBgNVBAsMGGdvLXNucy1tZXNzYWdlLXZhbGlkYXRv
cjEhMB8GA1UEAwwYZ28tc25zLW1lc3NhZ2UtdmFsaWRhdG9yMRwwGgYJKoZIhvcN
AQkBFg1pYW1AeXVobGF1Lm1lMIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCu
rgm/5MlF24ofyJNkNG/sjabX5d7i3ZQkJ6M8f+N9f1bF9ZRyGucODK8rFtHj/5Gc
oDNCCduT/sqA0moDx0b1ChxO25srzNYRUe3cNHehpgEWtIzQJzrHpYUrqannmCgy
JqcNJSLvN0Ex7WO6pMgx8xKXyDI2+Z9JhMHLvCAvMwIDAQABMA0GCSqGSIb3DQEB
CwUAA4GBABUyTJZtvHmuOOSUZzaqE8HdwSzRMIGdLCQYZunBIb403Clf15f/+hpv
vobi+xG4NkTmVX5kxRqwFb2C9OMtNaivC+nZKMo9WcNOQ9TqRSlIEJLrqP5dgrxn
kvCIAouFRHuLo4r9wvF3nUxtWjqfFa6TUfB+xtEalTn3LgKg9mzJ
-----END CERTIFICATE-----
`

	Convey("Given SNSValidator of message without \"Message\" key", t, func() {
		validator := SNSValidator{
			Version: 1,
			MessageMap: map[string]string{
				"Type":             "SubscriptionConfirmation",
				"MessageId":        "165545c9-2a5c-472c-8df2-7ff2be2b3b1b",
				"Token":            "2336412f37fb687f5d51e6e241d09c805a5a",
				"TopicArn":         "arn:aws:sns:us-west-2:123456789012:MyTopic",
				"Message":          "You have chosen to subscribe to the topic",
				"SubscribeURL":     "https://localhost/subscribe",
				"Subject":          "",
				"Timestamp":        "2012-04-26T20:45:04.751Z",
				"SignatureVersion": "1",
				"Signature":        "EXAMPLEpH+DcEwjAPg8O9mY8dReBSwksfg2S=",
				"SigningCertURL":   "https://localhost/cert.pem",
				"UnsubscribeURL":   "",
			},
		}
		validator.MessageMap["Message"] = ""

		Convey("It should return a SNSError", func() {
			actual := validator.ValidateMessage()

			So(actual, ShouldNotBeNil)
			So(actual, ShouldHaveSameTypeAs, snserrors.New("Type", "Message"))
			Convey("Returned SNSError should be of type ErrMissingKey", func() {
				So(actual.Type(), ShouldEqual, ErrMissingKey)
			})
			Convey("Returned SNSError message should be about the missing \"Message\" key", func() {
				So(actual.Error(), ShouldEqual, "\"Message\" is required in SNS message")
			})
		})
	})

	Convey(`Given SNSValidator of message with un-parsable "SigningCertURL"`, t, func() {
		validator := SNSValidator{
			MessageMap: map[string]string{
				"Type":             "Notification",
				"MessageId":        "165545c9-2a5c-472c-8df2-7ff2be2b3b1b",
				"Token":            "token-is-unnecessary-in-notification-message",
				"TopicArn":         "arn:aws:sns:us-west-2:123456789012:MyTopic",
				"Message":          "Test notification",
				"SubscribeURL":     "",
				"Subject":          "Test subject",
				"Timestamp":        "2012-04-26T20:45:04.751Z",
				"SignatureVersion": "1",
				"Signature":        "EXAMPLEpH+DcEwjAPg8O9mY8dReBSwksfg2S=",
				"SigningCertURL":   "!@#$%^&*",
				"UnsubscribeURL":   "https://localhost/unsubscribe",
			},
		}

		Convey("It should return a SNSError", func() {
			actual := validator.ValidateMessage()

			So(actual, ShouldNotBeNil)
			So(actual, ShouldHaveSameTypeAs, snserrors.New("Type", "Message"))

			Convey("Returned SNSError should be of type ErrInvalidCert", func() {
				So(actual.Type(), ShouldEqual, ErrInvalidCert)
			})
		})
	})

	Convey("Given SNSValidator of message with valid signature", t, func() {
		gock.New("https://sns.ap-northeast-1.amazonaws.com").
			Get("cert.pem").
			Reply(200).
			BodyString(certData)

		validator := SNSValidator{
			MessageMap: map[string]string{
				"Type":             "Notification",
				"MessageId":        "165545c9-2a5c-472c-8df2-7ff2be2b3b1b",
				"Token":            "token-is-unnecessary-in-notification-message",
				"TopicArn":         "arn:aws:sns:us-west-2:123456789012:MyTopic",
				"Message":          "Test notification",
				"SubscribeURL":     "",
				"Subject":          "Test subject",
				"Timestamp":        "2012-04-26T20:45:04.751Z",
				"SignatureVersion": "1",
				"Signature":        "ol5x/KiU+7dWKRuyD6Y1EntwXo+orXlVgQbq4JDy5uh/+EBBz/mfWQ0X0LXyyxkXXCykDakEz1F0h9y9xV9UitLlYA/tEMzI7WU9ob9d9L8YTCZVaHZUtCu4S0p0eCFzT69q+ijPuH9N1znuZOzDogsJIf8E9/8owtRmi6M50Co=",
				"SigningCertURL":   "https://sns.ap-northeast-1.amazonaws.com/cert.pem",
				"UnsubscribeURL":   "https://localhost/unsubscribe",
			},
		}

		Convey("It should return nil", func() {
			actual := validator.ValidateMessage()

			So(actual, ShouldBeNil)
		})
	})
}

func TestHasMethod(t *testing.T) {
	Convey(`Given SNSValidator of a message with empty "Message"`, t, func() {
		validator := SNSValidator{
			MessageMap: map[string]string{
				"Type":             "Notification",
				"MessageId":        "165545c9-2a5c-472c-8df2-7ff2be2b3b1b",
				"Token":            "",
				"TopicArn":         "arn:aws:sns:us-west-2:123456789012:MyTopic",
				"Message":          "Test notification",
				"SubscribeURL":     "",
				"Subject":          "Test subject",
				"Timestamp":        "2012-04-26T20:45:04.751Z",
				"SignatureVersion": "1",
				"Signature":        "EXAMPLEpH+DcEwjAPg8O9mY8dReBSwksfg2S=",
				"SigningCertURL":   "https://localhost/cert.pem",
				"UnsubscribeURL":   "https://localhost/unsubscribe",
			},
		}
		validator.MessageMap["Message"] = ""

		Convey(`It should return false when checking presence of "Message"`, func() {
			actual := validator.has("Message")

			So(actual, ShouldBeFalse)
		})
	})

	Convey(`Given SNSValidator of a message with non-empty "Message"`, t, func() {
		validator := SNSValidator{
			MessageMap: map[string]string{
				"Type":             "Notification",
				"MessageId":        "165545c9-2a5c-472c-8df2-7ff2be2b3b1b",
				"Token":            "",
				"TopicArn":         "arn:aws:sns:us-west-2:123456789012:MyTopic",
				"Message":          "Test notification",
				"SubscribeURL":     "",
				"Subject":          "Test subject",
				"Timestamp":        "2012-04-26T20:45:04.751Z",
				"SignatureVersion": "1",
				"Signature":        "EXAMPLEpH+DcEwjAPg8O9mY8dReBSwksfg2S=",
				"SigningCertURL":   "https://localhost/cert.pem",
				"UnsubscribeURL":   "https://localhost/unsubscribe",
			},
		}

		Convey(`It should return true when checking with "Message"`, func() {
			actual := validator.has("Message")

			So(actual, ShouldBeTrue)
		})
	})
}

func TestHasKeysMethod(t *testing.T) {
	Convey(`Given SNSValidator of message with key "Type" and "Message"`, t, func() {
		validator := SNSValidator{
			MessageMap: map[string]string{
				"Type":    "Notification",
				"Message": "Test notification",
			},
		}

		Convey(`When check for existence of key "Type" and "Subject"`, func() {
			valid, missingKey := validator.hasKeys([]string{"Type", "Subject"})

			Convey(`It should return false and "Subject" as missing key`, func() {
				So(valid, ShouldBeFalse)
				So(missingKey, ShouldEqual, "Subject")
			})
		})

		Convey(`When check for existence of key "Type" and "Message"`, func() {
			valid, missingKey := validator.hasKeys([]string{"Type", "Message"})

			Convey(`It should return true and empty missing`, func() {
				So(valid, ShouldBeTrue)
				So(missingKey, ShouldEqual, "")
			})
		})
	})
}

func TestIsTypesMethod(t *testing.T) {
	Convey(`Given SNSValidtor of Notificaiton message`, t, func() {
		validator := SNSValidator{
			MessageMap: map[string]string{
				"Type":             "Notification",
				"MessageId":        "165545c9-2a5c-472c-8df2-7ff2be2b3b1b",
				"Token":            "",
				"TopicArn":         "arn:aws:sns:us-west-2:123456789012:MyTopic",
				"Message":          "Test notification",
				"SubscribeURL":     "",
				"Subject":          "Test subject",
				"Timestamp":        "2012-04-26T20:45:04.751Z",
				"SignatureVersion": "1",
				"Signature":        "EXAMPLEpH+DcEwjAPg8O9mY8dReBSwksfg2S=",
				"SigningCertURL":   "https://localhost/cert.pem",
				"UnsubscribeURL":   "https://localhost/unsubscribe",
			},
		}

		Convey(`It should return true when check if it is SubscriptionConfirmation message`, func() {
			actual := validator.isTypes([]string{TypeSubscriptionConfirmation})

			So(actual, ShouldBeFalse)
		})

		Convey(`It should return true when check if it is SubscriptionConfirmation or Notification message`, func() {
			actual := validator.isTypes([]string{TypeSubscriptionConfirmation, TypeNotification})

			So(actual, ShouldBeTrue)
		})
	})
}

func TestValidateRequiredKeysMethod(t *testing.T) {
	Convey("Given SNSValidator of message without \"Message\" key", t, func() {
		validator := SNSValidator{
			Version: 1,
			MessageMap: map[string]string{
				"Type":             "SubscriptionConfirmation",
				"MessageId":        "165545c9-2a5c-472c-8df2-7ff2be2b3b1b",
				"Token":            "2336412f37fb687f5d51e6e241d09c805a5a",
				"TopicArn":         "arn:aws:sns:us-west-2:123456789012:MyTopic",
				"Message":          "You have chosen to subscribe to the topic",
				"SubscribeURL":     "https://localhost/subscribe",
				"Subject":          "",
				"Timestamp":        "2012-04-26T20:45:04.751Z",
				"SignatureVersion": "1",
				"Signature":        "EXAMPLEpH+DcEwjAPg8O9mY8dReBSwksfg2S=",
				"SigningCertURL":   "https://localhost/cert.pem",
				"UnsubscribeURL":   "",
			},
		}
		validator.MessageMap["Message"] = ""

		Convey("It should return a SNSError", func() {
			actual := validator.validateRequiredKeys()

			So(actual, ShouldNotBeNil)
			So(actual, ShouldHaveSameTypeAs, snserrors.New("Type", "Message"))
			Convey("Returned SNSError should be of type ErrMissingKey", func() {
				So(actual.Type(), ShouldEqual, ErrMissingKey)
			})
			Convey("Returned SNSError message should be about the missing \"Message\" key", func() {
				So(actual.Error(), ShouldEqual, "\"Message\" is required in SNS message")
			})
		})
	})

	Convey("Given a SNSValidator of message with all the required keys", t, func() {
		validator := SNSValidator{
			MessageMap: map[string]string{
				"Type":             "Notification",
				"MessageId":        "165545c9-2a5c-472c-8df2-7ff2be2b3b1b",
				"Token":            "",
				"TopicArn":         "arn:aws:sns:us-west-2:123456789012:MyTopic",
				"Message":          "Test notification",
				"SubscribeURL":     "",
				"Subject":          "Test subject",
				"Timestamp":        "2012-04-26T20:45:04.751Z",
				"SignatureVersion": "1",
				"Signature":        "EXAMPLEpH+DcEwjAPg8O9mY8dReBSwksfg2S=",
				"SigningCertURL":   "https://localhost/cert.pem",
				"UnsubscribeURL":   "https://localhost/unsubscribe",
			},
		}

		Convey("It should return nil", func() {
			actual := validator.validateRequiredKeys()
			So(actual, ShouldBeNil)
		})
	})
}

func TestValidateMessageTypeMethod(t *testing.T) {
	Convey("Given a SNSValidator of message with invalid type", t, func() {
		validator := SNSValidator{
			MessageMap: map[string]string{
				"Type":             "Notification",
				"MessageId":        "165545c9-2a5c-472c-8df2-7ff2be2b3b1b",
				"Token":            "",
				"TopicArn":         "arn:aws:sns:us-west-2:123456789012:MyTopic",
				"Message":          "Test notification",
				"SubscribeURL":     "",
				"Subject":          "Test subject",
				"Timestamp":        "2012-04-26T20:45:04.751Z",
				"SignatureVersion": "1",
				"Signature":        "EXAMPLEpH+DcEwjAPg8O9mY8dReBSwksfg2S=",
				"SigningCertURL":   "https://localhost/cert.pem",
				"UnsubscribeURL":   "https://localhost/unsubscribe",
			},
		}
		validator.MessageMap["Type"] = "InvalidType"

		Convey("It should return a SNSError", func() {
			actual := validator.validateMessageType()

			So(actual, ShouldNotBeNil)
			So(actual, ShouldHaveSameTypeAs, snserrors.New("Type", "Message"))
			Convey("Returned SNSError should be of type ErrInvalidType", func() {
				So(actual.Type(), ShouldEqual, ErrInvalidType)
			})
			Convey("Returned SNSError message should be about the invalid type", func() {
				So(actual.Error(), ShouldEqual, `Invalid message type "InvalidType"`)
			})
		})
	})

	Convey("Given a SNSValidator of SubscriptionConfirmation message", t, func() {
		validator := SNSValidator{
			Version: 1,
			MessageMap: map[string]string{
				"Type":             "SubscriptionConfirmation",
				"MessageId":        "165545c9-2a5c-472c-8df2-7ff2be2b3b1b",
				"Token":            "2336412f37fb687f5d51e6e241d09c805a5a",
				"TopicArn":         "arn:aws:sns:us-west-2:123456789012:MyTopic",
				"Message":          "You have chosen to subscribe to the topic",
				"SubscribeURL":     "https://localhost/subscribe",
				"Subject":          "",
				"Timestamp":        "2012-04-26T20:45:04.751Z",
				"SignatureVersion": "1",
				"Signature":        "EXAMPLEpH+DcEwjAPg8O9mY8dReBSwksfg2S=",
				"SigningCertURL":   "https://localhost/cert.pem",
				"UnsubscribeURL":   "",
			},
		}

		Convey("It should return nil", func() {
			actual := validator.validateMessageType()

			So(actual, ShouldBeNil)
		})
	})
}

func TestValidateSubscriptionKeysMethod(t *testing.T) {
	Convey(`Given a SNSValidator of message without "SubscribeURL" key`, t, func() {
		validator := SNSValidator{
			MessageMap: map[string]string{
				"Type":             "Notification",
				"MessageId":        "165545c9-2a5c-472c-8df2-7ff2be2b3b1b",
				"Token":            "",
				"TopicArn":         "arn:aws:sns:us-west-2:123456789012:MyTopic",
				"Message":          "Test notification",
				"SubscribeURL":     "",
				"Subject":          "Test subject",
				"Timestamp":        "2012-04-26T20:45:04.751Z",
				"SignatureVersion": "1",
				"Signature":        "EXAMPLEpH+DcEwjAPg8O9mY8dReBSwksfg2S=",
				"SigningCertURL":   "https://localhost/cert.pem",
				"UnsubscribeURL":   "https://localhost/unsubscribe",
			},
		}
		validator.MessageMap["SubscribeURL"] = ""

		Convey("It should return a SNSError", func() {
			actual := validator.validateSubscriptionKeys()

			So(actual, ShouldNotBeNil)
			So(actual, ShouldHaveSameTypeAs, snserrors.New("Type", "Message"))
			Convey("Returned SNSError should be of type ErrMissingKey", func() {
				So(actual.Type(), ShouldEqual, ErrMissingKey)
			})
			Convey(`Returned SNSError message should be about the missing "SubscribeURL" key`, func() {
				So(actual.Error(), ShouldEqual, `"SubscribeURL" is required in Subscription message`)
			})
		})
	})

	Convey(`Given a SNSValidator of message without "Token" key`, t, func() {
		validator := SNSValidator{
			Version: 1,
			MessageMap: map[string]string{
				"Type":             "SubscriptionConfirmation",
				"MessageId":        "165545c9-2a5c-472c-8df2-7ff2be2b3b1b",
				"Token":            "2336412f37fb687f5d51e6e241d09c805a5a",
				"TopicArn":         "arn:aws:sns:us-west-2:123456789012:MyTopic",
				"Message":          "You have chosen to subscribe to the topic",
				"SubscribeURL":     "https://localhost/subscribe",
				"Subject":          "",
				"Timestamp":        "2012-04-26T20:45:04.751Z",
				"SignatureVersion": "1",
				"Signature":        "EXAMPLEpH+DcEwjAPg8O9mY8dReBSwksfg2S=",
				"SigningCertURL":   "https://localhost/cert.pem",
				"UnsubscribeURL":   "",
			},
		}
		validator.MessageMap["Token"] = ""

		Convey("It should return a SNSError", func() {
			actual := validator.validateSubscriptionKeys()

			So(actual, ShouldNotBeNil)
			So(actual, ShouldHaveSameTypeAs, snserrors.New("Type", "Message"))
			Convey("Returned SNSError should be of type ErrMissingKey", func() {
				So(actual.Type(), ShouldEqual, ErrMissingKey)
			})
			Convey(`Returned SNSError message should be about the missing "Token" key`, func() {
				So(actual.Error(), ShouldEqual, `"Token" is required in Subscription message`)
			})
		})
	})

	Convey("Given a SNSValidator of valid SubscriptionMessage", t, func() {
		validator := SNSValidator{
			Version: 1,
			MessageMap: map[string]string{
				"Type":             "SubscriptionConfirmation",
				"MessageId":        "165545c9-2a5c-472c-8df2-7ff2be2b3b1b",
				"Token":            "2336412f37fb687f5d51e6e241d09c805a5a",
				"TopicArn":         "arn:aws:sns:us-west-2:123456789012:MyTopic",
				"Message":          "You have chosen to subscribe to the topic",
				"SubscribeURL":     "https://localhost/subscribe",
				"Subject":          "",
				"Timestamp":        "2012-04-26T20:45:04.751Z",
				"SignatureVersion": "1",
				"Signature":        "EXAMPLEpH+DcEwjAPg8O9mY8dReBSwksfg2S=",
				"SigningCertURL":   "https://localhost/cert.pem",
				"UnsubscribeURL":   "",
			},
		}

		Convey("It should return nil", func() {
			actual := validator.validateSubscriptionKeys()

			So(actual, ShouldBeNil)
		})
	})
}

func TestValidateMessageStructureMethod(t *testing.T) {
	Convey(`Given SNSValidator of message without "Message" key`, t, func() {
		validator := SNSValidator{
			MessageMap: map[string]string{
				"Type":             "Notification",
				"MessageId":        "165545c9-2a5c-472c-8df2-7ff2be2b3b1b",
				"Token":            "",
				"TopicArn":         "arn:aws:sns:us-west-2:123456789012:MyTopic",
				"Message":          "Test notification",
				"SubscribeURL":     "",
				"Subject":          "Test subject",
				"Timestamp":        "2012-04-26T20:45:04.751Z",
				"SignatureVersion": "1",
				"Signature":        "EXAMPLEpH+DcEwjAPg8O9mY8dReBSwksfg2S=",
				"SigningCertURL":   "https://localhost/cert.pem",
				"UnsubscribeURL":   "https://localhost/unsubscribe",
			},
		}
		validator.MessageMap["Message"] = ""

		Convey("It should return a SNSError", func() {
			actual := validator.validateMessageStructure()

			So(actual, ShouldNotBeNil)
			So(actual, ShouldHaveSameTypeAs, snserrors.New("Type", "Message"))
			Convey("Returned SNSError should be of type ErrMissingKey", func() {
				So(actual.Type(), ShouldEqual, ErrMissingKey)
			})
			Convey(`Returned SNSError message should be about the missing "Message" key`, func() {
				So(actual.Error(), ShouldEqual, `"Message" is required in SNS message`)
			})
		})
	})

	Convey("Given a SNSValidator of message with invalid type", t, func() {
		validator := SNSValidator{
			MessageMap: map[string]string{
				"Type":             "Notification",
				"MessageId":        "165545c9-2a5c-472c-8df2-7ff2be2b3b1b",
				"Token":            "",
				"TopicArn":         "arn:aws:sns:us-west-2:123456789012:MyTopic",
				"Message":          "Test notification",
				"SubscribeURL":     "",
				"Subject":          "Test subject",
				"Timestamp":        "2012-04-26T20:45:04.751Z",
				"SignatureVersion": "1",
				"Signature":        "EXAMPLEpH+DcEwjAPg8O9mY8dReBSwksfg2S=",
				"SigningCertURL":   "https://localhost/cert.pem",
				"UnsubscribeURL":   "https://localhost/unsubscribe",
			},
		}
		validator.MessageMap["Type"] = "InvalidType"

		Convey("It should return a SNSError", func() {
			actual := validator.validateMessageStructure()

			So(actual, ShouldNotBeNil)
			So(actual, ShouldHaveSameTypeAs, snserrors.New("Type", "Message"))
			Convey("Returned SNSError should be of type ErrInvalidType", func() {
				So(actual.Type(), ShouldEqual, ErrInvalidType)
			})
			Convey("Returned SNSError message should be about the invalide type", func() {
				So(actual.Error(), ShouldEqual, `Invalid message type "InvalidType"`)
			})
		})
	})

	Convey(`Given a SNSValiator of SubscriptionMessage without "SubscribeURL" key`, t, func() {
		validator := SNSValidator{
			Version: 1,
			MessageMap: map[string]string{
				"Type":             "SubscriptionConfirmation",
				"MessageId":        "165545c9-2a5c-472c-8df2-7ff2be2b3b1b",
				"Token":            "2336412f37fb687f5d51e6e241d09c805a5a",
				"TopicArn":         "arn:aws:sns:us-west-2:123456789012:MyTopic",
				"Message":          "You have chosen to subscribe to the topic",
				"SubscribeURL":     "https://localhost/subscribe",
				"Subject":          "",
				"Timestamp":        "2012-04-26T20:45:04.751Z",
				"SignatureVersion": "1",
				"Signature":        "EXAMPLEpH+DcEwjAPg8O9mY8dReBSwksfg2S=",
				"SigningCertURL":   "https://localhost/cert.pem",
				"UnsubscribeURL":   "",
			},
		}
		validator.MessageMap["SubscribeURL"] = ""

		Convey("It should return a SNSError", func() {
			actual := validator.validateMessageStructure()

			So(actual, ShouldNotBeNil)
			So(actual, ShouldHaveSameTypeAs, snserrors.New("Type", "Message"))
			Convey("Returned SNSError should be of type ErrMissingKey", func() {
				So(actual.Type(), ShouldEqual, ErrMissingKey)
			})
			Convey(`Returned SNSError message should be about the missing "SubscribeURL" key`, func() {
				So(actual.Error(), ShouldEqual, `"SubscribeURL" is required in Subscription message`)
			})
		})
	})

	Convey("Given a valid SubscriptionMessage", t, func() {
		validator := SNSValidator{
			Version: 1,
			MessageMap: map[string]string{
				"Type":             "SubscriptionConfirmation",
				"MessageId":        "165545c9-2a5c-472c-8df2-7ff2be2b3b1b",
				"Token":            "2336412f37fb687f5d51e6e241d09c805a5a",
				"TopicArn":         "arn:aws:sns:us-west-2:123456789012:MyTopic",
				"Message":          "You have chosen to subscribe to the topic",
				"SubscribeURL":     "https://localhost/subscribe",
				"Subject":          "",
				"Timestamp":        "2012-04-26T20:45:04.751Z",
				"SignatureVersion": "1",
				"Signature":        "EXAMPLEpH+DcEwjAPg8O9mY8dReBSwksfg2S=",
				"SigningCertURL":   "https://localhost/cert.pem",
				"UnsubscribeURL":   "",
			},
		}

		Convey("It should return nil", func() {
			actual := validator.validateMessageStructure()

			So(actual, ShouldBeNil)
		})
	})
}

func TestBuildSingalbleString(t *testing.T) {
	Convey("Given a SNSValidator of Notification message", t, func() {
		validator := SNSValidator{
			MessageMap: map[string]string{
				"Type":             "Notification",
				"MessageId":        "165545c9-2a5c-472c-8df2-7ff2be2b3b1b",
				"Token":            "",
				"TopicArn":         "arn:aws:sns:us-west-2:123456789012:MyTopic",
				"Message":          "Test notification",
				"SubscribeURL":     "",
				"Subject":          "Test subject",
				"Timestamp":        "2012-04-26T20:45:04.751Z",
				"SignatureVersion": "1",
				"Signature":        "EXAMPLEpH+DcEwjAPg8O9mY8dReBSwksfg2S=",
				"SigningCertURL":   "https://localhost/cert.pem",
				"UnsubscribeURL":   "https://localhost/unsubscribe",
			},
		}

		Convey("It should build and return a signable string with only the signable subscription keys", func() {
			expected := []byte(`Message
Test notification
MessageId
165545c9-2a5c-472c-8df2-7ff2be2b3b1b
Subject
Test subject
Timestamp
2012-04-26T20:45:04.751Z
TopicArn
arn:aws:sns:us-west-2:123456789012:MyTopic
Type
Notification
`)
			actual := validator.buildSignableString()

			So(actual, ShouldResemble, expected)
		})
	})

	Convey("Given a SNSValidator of SubscriptConfirmation message", t, func() {
		validator := SNSValidator{
			Version: 1,
			MessageMap: map[string]string{
				"Type":             "SubscriptionConfirmation",
				"MessageId":        "165545c9-2a5c-472c-8df2-7ff2be2b3b1b",
				"Token":            "2336412f37fb687f5d51e6e241d09c805a5a",
				"TopicArn":         "arn:aws:sns:us-west-2:123456789012:MyTopic",
				"Message":          "You have chosen to subscribe to the topic",
				"SubscribeURL":     "https://localhost/subscribe",
				"Subject":          "",
				"Timestamp":        "2012-04-26T20:45:04.751Z",
				"SignatureVersion": "1",
				"Signature":        "EXAMPLEpH+DcEwjAPg8O9mY8dReBSwksfg2S=",
				"SigningCertURL":   "https://localhost/cert.pem",
				"UnsubscribeURL":   "",
			},
		}

		Convey("It should build and return a signable string with only the signable subscription keys", func() {
			expected := []byte(`Message
You have chosen to subscribe to the topic
MessageId
165545c9-2a5c-472c-8df2-7ff2be2b3b1b
SubscribeURL
https://localhost/subscribe
Timestamp
2012-04-26T20:45:04.751Z
Token
2336412f37fb687f5d51e6e241d09c805a5a
TopicArn
arn:aws:sns:us-west-2:123456789012:MyTopic
Type
SubscriptionConfirmation
`)
			actual := validator.buildSignableString()

			So(actual, ShouldResemble, expected)
		})
	})

	Convey(`Given a SNSValidator of Notificatoin message with unncecessary "Token" key`, t, func() {
		validator := SNSValidator{
			MessageMap: map[string]string{
				"Type":             "Notification",
				"MessageId":        "165545c9-2a5c-472c-8df2-7ff2be2b3b1b",
				"Token":            "token-is-unnecessary-in-notification-message",
				"TopicArn":         "arn:aws:sns:us-west-2:123456789012:MyTopic",
				"Message":          "Test notification",
				"SubscribeURL":     "",
				"Subject":          "Test subject",
				"Timestamp":        "2012-04-26T20:45:04.751Z",
				"SignatureVersion": "1",
				"Signature":        "EXAMPLEpH+DcEwjAPg8O9mY8dReBSwksfg2S=",
				"SigningCertURL":   "https://localhost/cert.pem",
				"UnsubscribeURL":   "https://localhost/unsubscribe",
			},
		}

		Convey("It should build and return a signable string without the Token key", func() {
			expected := []byte(`Message
Test notification
MessageId
165545c9-2a5c-472c-8df2-7ff2be2b3b1b
Subject
Test subject
Timestamp
2012-04-26T20:45:04.751Z
TopicArn
arn:aws:sns:us-west-2:123456789012:MyTopic
Type
Notification
`)
			actual := validator.buildSignableString()

			So(actual, ShouldResemble, expected)
		})
	})
}

func TestGetCertifiateMethod(t *testing.T) {
	certData := `-----BEGIN CERTIFICATE-----
MIIC8zCCAlwCCQCLHrKJpPLt9TANBgkqhkiG9w0BAQsFADCBvDELMAkGA1UEBhMC
SEsxEjAQBgNVBAgMCUhvbmctS29uZzESMBAGA1UEBwwJSG9uZy1Lb25nMSEwHwYD
VQQKDBhnby1zbnMtbWVzc2FnZS12YWxpZGF0b3IxITAfBgNVBAsMGGdvLXNucy1t
ZXNzYWdlLXZhbGlkYXRvcjEhMB8GA1UEAwwYZ28tc25zLW1lc3NhZ2UtdmFsaWRh
dG9yMRwwGgYJKoZIhvcNAQkBFg1pYW1AeXVobGF1Lm1lMCAXDTE3MDkxNzE3MDA1
MloYDzIwNjcwOTA1MTcwMDUyWjCBvDELMAkGA1UEBhMCSEsxEjAQBgNVBAgMCUhv
bmctS29uZzESMBAGA1UEBwwJSG9uZy1Lb25nMSEwHwYDVQQKDBhnby1zbnMtbWVz
c2FnZS12YWxpZGF0b3IxITAfBgNVBAsMGGdvLXNucy1tZXNzYWdlLXZhbGlkYXRv
cjEhMB8GA1UEAwwYZ28tc25zLW1lc3NhZ2UtdmFsaWRhdG9yMRwwGgYJKoZIhvcN
AQkBFg1pYW1AeXVobGF1Lm1lMIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCu
rgm/5MlF24ofyJNkNG/sjabX5d7i3ZQkJ6M8f+N9f1bF9ZRyGucODK8rFtHj/5Gc
oDNCCduT/sqA0moDx0b1ChxO25srzNYRUe3cNHehpgEWtIzQJzrHpYUrqannmCgy
JqcNJSLvN0Ex7WO6pMgx8xKXyDI2+Z9JhMHLvCAvMwIDAQABMA0GCSqGSIb3DQEB
CwUAA4GBABUyTJZtvHmuOOSUZzaqE8HdwSzRMIGdLCQYZunBIb403Clf15f/+hpv
vobi+xG4NkTmVX5kxRqwFb2C9OMtNaivC+nZKMo9WcNOQ9TqRSlIEJLrqP5dgrxn
kvCIAouFRHuLo4r9wvF3nUxtWjqfFa6TUfB+xtEalTn3LgKg9mzJ
-----END CERTIFICATE-----
`

	Convey(`Given SNSValidator of mesage with "SigningCertURL" heading to 404 Not Found`, t, func() {
		gock.New("https://sns.ap-northeast-1.amazonaws.com").
			Get("notfound.pem").
			Reply(404)

		validator := SNSValidator{
			MessageMap: map[string]string{
				"Type":             "Notification",
				"MessageId":        "165545c9-2a5c-472c-8df2-7ff2be2b3b1b",
				"Token":            "token-is-unnecessary-in-notification-message",
				"TopicArn":         "arn:aws:sns:us-west-2:123456789012:MyTopic",
				"Message":          "Test notification",
				"SubscribeURL":     "",
				"Subject":          "Test subject",
				"Timestamp":        "2012-04-26T20:45:04.751Z",
				"SignatureVersion": "1",
				"Signature":        "EXAMPLEpH+DcEwjAPg8O9mY8dReBSwksfg2S=",
				"SigningCertURL":   "https://sns.ap-northeast-1.amazonaws.com/notfound.pem",
				"UnsubscribeURL":   "https://localhost/unsubscribe",
			},
		}

		Convey("It should return a SNSError", func() {
			actualData, actualErr := validator.getCertificate()

			So(actualData, ShouldBeNil)
			So(actualErr, ShouldHaveSameTypeAs, snserrors.New("Type", "Message"))

			Convey("Returned SNSError should be of type ErrInvalidCert", func() {
				So(actualErr.Type(), ShouldEqual, ErrInvalidCert)
			})
			Convey("Returned SNSError should be about certificate could not be retrieved", func() {
				So(actualErr.Error(), ShouldEqual, "Could not retrive the certificate")
			})
		})
	})

	Convey(`Given SNSValidator of mesage with unreachable "SigningCertURL"`, t, func() {
		validator := SNSValidator{
			MessageMap: map[string]string{
				"Type":             "Notification",
				"MessageId":        "165545c9-2a5c-472c-8df2-7ff2be2b3b1b",
				"Token":            "token-is-unnecessary-in-notification-message",
				"TopicArn":         "arn:aws:sns:us-west-2:123456789012:MyTopic",
				"Message":          "Test notification",
				"SubscribeURL":     "",
				"Subject":          "Test subject",
				"Timestamp":        "2012-04-26T20:45:04.751Z",
				"SignatureVersion": "1",
				"Signature":        "EXAMPLEpH+DcEwjAPg8O9mY8dReBSwksfg2S=",
				"SigningCertURL":   "https://localhost/cert.pem",
				"UnsubscribeURL":   "https://localhost/unsubscribe",
			},
		}

		Convey("It should return a SNSError", func() {
			actualData, actualErr := validator.getCertificate()

			So(actualData, ShouldBeNil)
			So(actualErr, ShouldHaveSameTypeAs, snserrors.New("Type", "Message"))

			Convey("Returned SNSError should be of type ErrInvalidCert", func() {
				So(actualErr.Type(), ShouldEqual, ErrInvalidCert)
			})
		})
	})

	Convey(`Given SNSValidator of mesage with valid "SigningCertURL"`, t, func() {
		gock.New("https://sns.ap-northeast-1.amazonaws.com").
			Get("cert.pem").
			Reply(200).
			BodyString(certData)

		validator := SNSValidator{
			MessageMap: map[string]string{
				"Type":             "Notification",
				"MessageId":        "165545c9-2a5c-472c-8df2-7ff2be2b3b1b",
				"Token":            "token-is-unnecessary-in-notification-message",
				"TopicArn":         "arn:aws:sns:us-west-2:123456789012:MyTopic",
				"Message":          "Test notification",
				"SubscribeURL":     "",
				"Subject":          "Test subject",
				"Timestamp":        "2012-04-26T20:45:04.751Z",
				"SignatureVersion": "1",
				"Signature":        "EXAMPLEpH+DcEwjAPg8O9mY8dReBSwksfg2S=",
				"SigningCertURL":   "https://sns.ap-northeast-1.amazonaws.com/cert.pem",
				"UnsubscribeURL":   "https://localhost/unsubscribe",
			},
		}

		Convey("It should fetch the certificate and return certifitcate as slice of byte", func() {
			actualData, actualErr := validator.getCertificate()

			So(actualData, ShouldResemble, []byte(certData))
			So(actualErr, ShouldBeNil)
		})
	})
}

func TestVerifySignature(t *testing.T) {
	// Mock HTTP request-response
	certData := `-----BEGIN CERTIFICATE-----
MIIC8zCCAlwCCQCLHrKJpPLt9TANBgkqhkiG9w0BAQsFADCBvDELMAkGA1UEBhMC
SEsxEjAQBgNVBAgMCUhvbmctS29uZzESMBAGA1UEBwwJSG9uZy1Lb25nMSEwHwYD
VQQKDBhnby1zbnMtbWVzc2FnZS12YWxpZGF0b3IxITAfBgNVBAsMGGdvLXNucy1t
ZXNzYWdlLXZhbGlkYXRvcjEhMB8GA1UEAwwYZ28tc25zLW1lc3NhZ2UtdmFsaWRh
dG9yMRwwGgYJKoZIhvcNAQkBFg1pYW1AeXVobGF1Lm1lMCAXDTE3MDkxNzE3MDA1
MloYDzIwNjcwOTA1MTcwMDUyWjCBvDELMAkGA1UEBhMCSEsxEjAQBgNVBAgMCUhv
bmctS29uZzESMBAGA1UEBwwJSG9uZy1Lb25nMSEwHwYDVQQKDBhnby1zbnMtbWVz
c2FnZS12YWxpZGF0b3IxITAfBgNVBAsMGGdvLXNucy1tZXNzYWdlLXZhbGlkYXRv
cjEhMB8GA1UEAwwYZ28tc25zLW1lc3NhZ2UtdmFsaWRhdG9yMRwwGgYJKoZIhvcN
AQkBFg1pYW1AeXVobGF1Lm1lMIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCu
rgm/5MlF24ofyJNkNG/sjabX5d7i3ZQkJ6M8f+N9f1bF9ZRyGucODK8rFtHj/5Gc
oDNCCduT/sqA0moDx0b1ChxO25srzNYRUe3cNHehpgEWtIzQJzrHpYUrqannmCgy
JqcNJSLvN0Ex7WO6pMgx8xKXyDI2+Z9JhMHLvCAvMwIDAQABMA0GCSqGSIb3DQEB
CwUAA4GBABUyTJZtvHmuOOSUZzaqE8HdwSzRMIGdLCQYZunBIb403Clf15f/+hpv
vobi+xG4NkTmVX5kxRqwFb2C9OMtNaivC+nZKMo9WcNOQ9TqRSlIEJLrqP5dgrxn
kvCIAouFRHuLo4r9wvF3nUxtWjqfFa6TUfB+xtEalTn3LgKg9mzJ
-----END CERTIFICATE-----
`

	// A passphase protected RSA private key cannot be decoed correctly
	protectedCertData := `-----BEGIN RSA PRIVATE KEY-----
Proc-Type: 4,ENCRYPTED
DEK-Info: DES-EDE3-CBC,8AFA72215161D548

HXO8Es6+W8cXJYi4nDKE/2ojwRLVRqIAlFxevFoXJVMcIfhcmDEBR+Oy8n3Z2QzL
RKQatJZ5Pn7h61xSisM3B96EJ5UIfF6c1TsMtfSzmd3MHS6rs2ktVwj04iUN8JJZ
ccrMWNbHj0Exlr+IKZKw4/c5ScKiub2zFnD7lD0sGbZxrWR8doCriWb2QCUcMla7
t35ql3N526qGsts2DQZ2qmIDUP4GMEusB1FK1YSrCQ9qIL5Cn4CuKqZFdyuSIkIa
OpAWqNh1Y6JGbMjsX53lqgM+BJ31fgGlMdev7zLetiltfEfF+wURX+HmbZZHt4UP
y1BXv63L5WwiafNj0IVCWM1wXYe8ePxGI04mNndZrkx4SA+BIEKznSgIHYo8N5CT
Brdq7UOx3M1NlLEt71LiU82oa3VF7f9fuHpdbtAPCo6YqRH2ez/ya/m/yWIm9Z7l
LFP8L+E8lpyt/VRrpA2wZwOPJKr2EAWNcBnKceY+7KLujs4HV6DiEb9JQ7o7Ayof
bI/0zc9CL+O/K98vAdhXLvRQJ8DUdsVc5wiet2t/RjN8XlBLHUX8kDg0NyHyBuVp
MrWobRNuVblYlx6MOJRQMdbfN7+EU8r/VRMv13JDSmQJsOxlHQQNhhi46cO1S1DM
bUSBS7Mf+PGP1OSMPtF1f4+OdMBi19Px38LPO9FBBumFn492Q3W1Wg/7VGNyVxrS
YKBNdTuDP0D0B952ZWP8LpUi5072UIGlr3QdqZ7BLCwJyxaaMjkSkEaPMhcq/mD6
CL1EiIfbFtoy2EXXADdQr/wz2dOAiufhQXf2lkzAKS9QQq0J0L54Ng==
-----END RSA PRIVATE KEY-----
`

	invalidCertData := `-----BEGIN CERTIFICATE-----
MIIC8zCCAlwCCQCLHrKJpPLt9TANBgkqhkiG9w0BAQsFADCBvDELMAkGA1UEBhMC
SEsxEjAQBgNVBAgMCUhvbmctS29uZzESMBAGA1UEBwwJSG9uZy1Lb25nMSEwHwYD
VQQKDBhnby1zbnMtbWVzc2FnZS12YWxpZGF0b3IxITAfBgNVBAsMGGdvLXNucy1t
ZXNzYWdlLXZhbGlkYXRvcjEhMB8GA1UEAwwYZ28tc25zLW1lc3NhZ2UtdmFsaWRh
dG9yMRwwGgYJKoZIhvcNAQkBFg1pYW1AeXVobGF1Lm1lMCAXDTE3MDkxNzE3MDA1
MloYDzIwNjcwOTA1MTcwMDUyWjCBvDELMAkGA1UEBhMCSEsxEjAQBgNVBAgMCUhv
-----END CERTIFICATE-----
`

	Convey(`Given SNSValidator of message with un-parsable "SigningCertURL"`, t, func() {
		validator := SNSValidator{
			MessageMap: map[string]string{
				"Type":             "Notification",
				"MessageId":        "165545c9-2a5c-472c-8df2-7ff2be2b3b1b",
				"Token":            "token-is-unnecessary-in-notification-message",
				"TopicArn":         "arn:aws:sns:us-west-2:123456789012:MyTopic",
				"Message":          "Test notification",
				"SubscribeURL":     "",
				"Subject":          "Test subject",
				"Timestamp":        "2012-04-26T20:45:04.751Z",
				"SignatureVersion": "1",
				"Signature":        "EXAMPLEpH+DcEwjAPg8O9mY8dReBSwksfg2S=",
				"SigningCertURL":   "!@#$%^&*",
				"UnsubscribeURL":   "https://localhost/unsubscribe",
			},
		}

		Convey("It should return a SNSError", func() {
			actual := validator.verifySignature()

			So(actual, ShouldNotBeNil)
			So(actual, ShouldHaveSameTypeAs, snserrors.New("Type", "Message"))

			Convey("Returned SNSError should be of type ErrInvalidCert", func() {
				So(actual.Type(), ShouldEqual, ErrInvalidCert)
			})
		})
	})

	Convey(`Given SNSValidator of mesage with "SigningCertURL" using insecure HTTP`, t, func() {
		validator := SNSValidator{
			MessageMap: map[string]string{
				"Type":             "Notification",
				"MessageId":        "165545c9-2a5c-472c-8df2-7ff2be2b3b1b",
				"Token":            "token-is-unnecessary-in-notification-message",
				"TopicArn":         "arn:aws:sns:us-west-2:123456789012:MyTopic",
				"Message":          "Test notification",
				"SubscribeURL":     "",
				"Subject":          "Test subject",
				"Timestamp":        "2012-04-26T20:45:04.751Z",
				"SignatureVersion": "1",
				"Signature":        "EXAMPLEpH+DcEwjAPg8O9mY8dReBSwksfg2S=",
				"SigningCertURL":   "http://sns.ap-northeast-1.amazonaws.com/cert.pem",
				"UnsubscribeURL":   "https://localhost/unsubscribe",
			},
		}

		Convey("It should return a SNSError", func() {
			actual := validator.verifySignature()

			So(actual, ShouldNotBeNil)
			So(actual, ShouldHaveSameTypeAs, snserrors.New("Type", "Message"))

			Convey("Returned SNSError should be of type ErrInvalidCert", func() {
				So(actual.Type(), ShouldEqual, ErrInvalidCert)
			})
			Convey("Returned SNSError should be about insecure HTTP", func() {
				So(actual.Error(), ShouldEqual, "The certificate URL is using insecure HTTP scheme")
			})
		})
	})

	Convey(`Given SNSValidator of message with untrusted "SigningCertURL"`, t, func() {
		validator := SNSValidator{
			MessageMap: map[string]string{
				"Type":             "Notification",
				"MessageId":        "165545c9-2a5c-472c-8df2-7ff2be2b3b1b",
				"Token":            "token-is-unnecessary-in-notification-message",
				"TopicArn":         "arn:aws:sns:us-west-2:123456789012:MyTopic",
				"Message":          "Test notification",
				"SubscribeURL":     "",
				"Subject":          "Test subject",
				"Timestamp":        "2012-04-26T20:45:04.751Z",
				"SignatureVersion": "1",
				"Signature":        "EXAMPLEpH+DcEwjAPg8O9mY8dReBSwksfg2S=",
				"SigningCertURL":   "https://localhost/cert.pem",
				"UnsubscribeURL":   "https://localhost/unsubscribe",
			},
		}

		Convey("It should return a SNSError", func() {
			actual := validator.verifySignature()

			So(actual, ShouldNotBeNil)
			So(actual, ShouldHaveSameTypeAs, snserrors.New("Type", "Message"))

			Convey("Returned SNSError should be of type ErrInvalidCert", func() {
				So(actual.Type(), ShouldEqual, ErrInvalidCert)
			})
			Convey("Returned SNSError should be reatled to untrusted host", func() {
				So(actual.Error(), ShouldEqual, "The certificate URL belongs to an untrusted host")
			})
		})
	})

	Convey(`Given SNSValidator of mesage with "SigningCertURL" heading to 404 Not Found`, t, func() {
		gock.New("https://sns.ap-northeast-1.amazonaws.com").
			Get("notfound.pem").
			Reply(404)

		validator := SNSValidator{
			MessageMap: map[string]string{
				"Type":             "Notification",
				"MessageId":        "165545c9-2a5c-472c-8df2-7ff2be2b3b1b",
				"Token":            "token-is-unnecessary-in-notification-message",
				"TopicArn":         "arn:aws:sns:us-west-2:123456789012:MyTopic",
				"Message":          "Test notification",
				"SubscribeURL":     "",
				"Subject":          "Test subject",
				"Timestamp":        "2012-04-26T20:45:04.751Z",
				"SignatureVersion": "1",
				"Signature":        "EXAMPLEpH+DcEwjAPg8O9mY8dReBSwksfg2S=",
				"SigningCertURL":   "https://sns.ap-northeast-1.amazonaws.com/notfound.pem",
				"UnsubscribeURL":   "https://localhost/unsubscribe",
			},
		}

		Convey("It should return a SNSError", func() {
			actual := validator.verifySignature()

			So(actual, ShouldNotBeNil)
			So(actual, ShouldHaveSameTypeAs, snserrors.New("Type", "Message"))

			Convey("Returned SNSError should be of type ErrInvalidCert", func() {
				So(actual.Type(), ShouldEqual, ErrInvalidCert)
			})
			Convey("Returned SNSError should be about certificate could not be retrieved", func() {
				So(actual.Error(), ShouldEqual, "Could not retrive the certificate")
			})
		})
	})

	Convey(`Given SNSValidator of message with "SigningCertURL" heading to empty certificate`, t, func() {
		gock.New("https://sns.ap-northeast-1.amazonaws.com").
			Get("empty.pem").
			Reply(200).
			BodyString("")

		validator := SNSValidator{
			MessageMap: map[string]string{
				"Type":             "Notification",
				"MessageId":        "165545c9-2a5c-472c-8df2-7ff2be2b3b1b",
				"Token":            "token-is-unnecessary-in-notification-message",
				"TopicArn":         "arn:aws:sns:us-west-2:123456789012:MyTopic",
				"Message":          "Test notification",
				"SubscribeURL":     "",
				"Subject":          "Test subject",
				"Timestamp":        "2012-04-26T20:45:04.751Z",
				"SignatureVersion": "1",
				"Signature":        "EXAMPLEpH+DcEwjAPg8O9mY8dReBSwksfg2S=",
				"SigningCertURL":   "https://sns.ap-northeast-1.amazonaws.com/empty.pem",
				"UnsubscribeURL":   "https://localhost/unsubscribe",
			},
		}

		Convey("It should return a SNSError", func() {
			actual := validator.verifySignature()

			So(actual, ShouldNotBeNil)
			So(actual, ShouldHaveSameTypeAs, snserrors.New("Type", "Message"))

			Convey("Returned SNSError should be of type ErrInvalidCert", func() {
				So(actual.Type(), ShouldEqual, ErrInvalidCert)
			})
		})
	})

	Convey(`Given SNSValidator of message with "SigningCertURL" heading to un-decodable certificate`, t, func() {
		gock.New("https://sns.ap-northeast-1.amazonaws.com").
			Get("protected.pem").
			Reply(200).
			BodyString(protectedCertData)

		validator := SNSValidator{
			MessageMap: map[string]string{
				"Type":             "Notification",
				"MessageId":        "165545c9-2a5c-472c-8df2-7ff2be2b3b1b",
				"Token":            "token-is-unnecessary-in-notification-message",
				"TopicArn":         "arn:aws:sns:us-west-2:123456789012:MyTopic",
				"Message":          "Test notification",
				"SubscribeURL":     "",
				"Subject":          "Test subject",
				"Timestamp":        "2012-04-26T20:45:04.751Z",
				"SignatureVersion": "1",
				"Signature":        "EXAMPLEpH+DcEwjAPg8O9mY8dReBSwksfg2S=",
				"SigningCertURL":   "https://sns.ap-northeast-1.amazonaws.com/protected.pem",
				"UnsubscribeURL":   "https://localhost/unsubscribe",
			},
		}

		Convey("It should return a SNSError", func() {
			actual := validator.verifySignature()

			So(actual, ShouldNotBeNil)
			So(actual, ShouldHaveSameTypeAs, snserrors.New("Type", "Message"))

			Convey("Returned SNSError should be of type ErrInvalidCert", func() {
				So(actual.Type(), ShouldEqual, ErrInvalidCert)
			})
		})
	})

	Convey(`Given SNSValidator of mesage with "SigningCertURL" heading to invalid certificate`, t, func() {
		gock.New("https://sns.ap-northeast-1.amazonaws.com").
			Get("invalid.pem").
			Reply(200).
			BodyString(invalidCertData)

		validator := SNSValidator{
			MessageMap: map[string]string{
				"Type":             "Notification",
				"MessageId":        "165545c9-2a5c-472c-8df2-7ff2be2b3b1b",
				"Token":            "token-is-unnecessary-in-notification-message",
				"TopicArn":         "arn:aws:sns:us-west-2:123456789012:MyTopic",
				"Message":          "Test notification",
				"SubscribeURL":     "",
				"Subject":          "Test subject",
				"Timestamp":        "2012-04-26T20:45:04.751Z",
				"SignatureVersion": "1",
				"Signature":        "EXAMPLEpH+DcEwjAPg8O9mY8dReBSwksfg2S=",
				"SigningCertURL":   "https://sns.ap-northeast-1.amazonaws.com/invalid.pem",
				"UnsubscribeURL":   "https://localhost/unsubscribe",
			},
		}

		Convey("It should return a SNSError", func() {
			actual := validator.verifySignature()

			So(actual, ShouldNotBeNil)
			So(actual, ShouldHaveSameTypeAs, snserrors.New("Type", "Message"))

			Convey("Returned SNSError should be of type ErrInvalidCert", func() {
				So(actual.Type(), ShouldEqual, ErrInvalidCert)
			})
		})
	})

	Convey("Given SNSValidator of message with corrupted base64 signature", t, func() {
		gock.New("https://sns.ap-northeast-1.amazonaws.com").
			Get("cert.pem").
			Reply(200).
			BodyString(certData)

		validator := SNSValidator{
			MessageMap: map[string]string{
				"Type":             "Notification",
				"MessageId":        "165545c9-2a5c-472c-8df2-7ff2be2b3b1b",
				"Token":            "token-is-unnecessary-in-notification-message",
				"TopicArn":         "arn:aws:sns:us-west-2:123456789012:MyTopic",
				"Message":          "Test notification",
				"SubscribeURL":     "",
				"Subject":          "Test subject",
				"Timestamp":        "2012-04-26T20:45:04.751Z",
				"SignatureVersion": "1",
				"Signature":        "EXAMPLEpH+DcEwjAPg8O9mY8dReBSwksfg2S=",
				"SigningCertURL":   "https://sns.ap-northeast-1.amazonaws.com/cert.pem",
				"UnsubscribeURL":   "https://localhost/unsubscribe",
			},
		}

		Convey("It should return a SNSError", func() {
			actual := validator.verifySignature()

			So(actual, ShouldNotBeNil)
			So(actual, ShouldHaveSameTypeAs, snserrors.New("Type", "Message"))

			Convey("Returned SNSError should be of type ErrIncorrectSignature", func() {
				So(actual.Type(), ShouldEqual, ErrIncorrectSignature)
			})
			Convey("Returned SNSError should be about base64 decode error", func() {
				So(actual.Error(), ShouldEqual, "Could not base64 decode the signature")
			})
		})
	})

	Convey("Given SNSValidator of message with incorrect signature", t, func() {
		gock.New("https://sns.ap-northeast-1.amazonaws.com").
			Get("cert.pem").
			Reply(200).
			BodyString(certData)

		validator := SNSValidator{
			MessageMap: map[string]string{
				"Type":             "Notification",
				"MessageId":        "165545c9-2a5c-472c-8df2-7ff2be2b3b1b",
				"Token":            "token-is-unnecessary-in-notification-message",
				"TopicArn":         "arn:aws:sns:us-west-2:123456789012:MyTopic",
				"Message":          "Test notification",
				"SubscribeURL":     "",
				"Subject":          "Test subject",
				"Timestamp":        "2012-04-26T20:45:04.751Z",
				"SignatureVersion": "1",
				"Signature":        "aW52YWxpZC1zaWduYXR1cmU=",
				"SigningCertURL":   "https://sns.ap-northeast-1.amazonaws.com/cert.pem",
				"UnsubscribeURL":   "https://localhost/unsubscribe",
			},
		}

		Convey("It should return a SNSError", func() {
			actual := validator.verifySignature()

			So(actual, ShouldNotBeNil)
			So(actual, ShouldHaveSameTypeAs, snserrors.New("Type", "Message"))

			Convey("Returned SNSError should be of type ErrIncorrectSignature", func() {
				So(actual.Type(), ShouldEqual, ErrIncorrectSignature)
			})
		})
	})

	Convey("Given SNSValidator of message with valid signature", t, func() {
		gock.New("https://sns.ap-northeast-1.amazonaws.com").
			Get("cert.pem").
			Reply(200).
			BodyString(certData)

		validator := SNSValidator{
			MessageMap: map[string]string{
				"Type":             "Notification",
				"MessageId":        "165545c9-2a5c-472c-8df2-7ff2be2b3b1b",
				"Token":            "token-is-unnecessary-in-notification-message",
				"TopicArn":         "arn:aws:sns:us-west-2:123456789012:MyTopic",
				"Message":          "Test notification",
				"SubscribeURL":     "",
				"Subject":          "Test subject",
				"Timestamp":        "2012-04-26T20:45:04.751Z",
				"SignatureVersion": "1",
				"Signature":        "ol5x/KiU+7dWKRuyD6Y1EntwXo+orXlVgQbq4JDy5uh/+EBBz/mfWQ0X0LXyyxkXXCykDakEz1F0h9y9xV9UitLlYA/tEMzI7WU9ob9d9L8YTCZVaHZUtCu4S0p0eCFzT69q+ijPuH9N1znuZOzDogsJIf8E9/8owtRmi6M50Co=",
				"SigningCertURL":   "https://sns.ap-northeast-1.amazonaws.com/cert.pem",
				"UnsubscribeURL":   "https://localhost/unsubscribe",
			},
		}

		Convey("It should return nil", func() {
			actual := validator.verifySignature()

			So(actual, ShouldBeNil)
		})
	})

}
