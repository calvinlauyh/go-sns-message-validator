# Go SNS Message Validator

**Go SMS Message Validator** is an unofficial Go library for validating AWS SNS message. It validates an incoming SNS Message authenticity and integrity by validating the message structure and verifying the message signature. It is standalone and does not require AWS SDK to work.

This project is still in its early development and a more stable release will be coming out in a few weeks.

## Features
- Follows AWS documentation on message validation
- Validate [message structure](http://docs.aws.amazon.com/sns/latest/dg/json-formats.html)
- [Verify message signature](http://docs.aws.amazon.com/sns/latest/dg/SendMessageToHttp.verify.signature.html)

## Installation
```sh
$ go get -u github.com/yuhlau/go-sns-message-validator
```

## Usage

### Verifying a SNS message
```go
import (
	"fmt"

	"github.com/yuhlau/go-sns-message-validator/snsmessage"
)

// Retrieve SNS message in JSON format from HTTP request and store into variable messageJSON
...

message, err := snsmessage.NewFromJSON([]byte(messageJSON))
if err != nil {
	fmt.Println(err)
}
if err := message.GetValidator().ValidateMessage(); err != nil {
	fmt.Println(err)
}
// The SNS message is validated now
```

## Test
Most of the code are covered by test. Test coverage is about 99.5% right now. The only remaining part is an I/O error handling which requires special data to cover it in the test.

More testing with the real AWS platform is in progress and the first stable `1.0.0` release will wait until all tests are finished.

#### The test cases are written using

#### Goconvey [[Github](https://github.com/smartystreets/goconvey)] as testing framework
```sh
$ go get github.com/smartystreets/goconvey
```
#### gock [[Github](https://github.com/h2non/gock)] for HTTP mocking

```sh
$ go get -u gopkg.in/h2non/gock.v1
```

## License
MIT

## Version History

18 Sep 2017, v0.0.1
- Initial release