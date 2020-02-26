package main

import (
	"errors"
	"time"

	"github.com/getsentry/sentry-go"
)

func main() {
	sentry.Init(sentry.ClientOptions{
		Dsn: "http://b626dc50edf44700a32a19a104e41e31@sentry.yuzhua-dev.com/2",
	})

	sentry.CaptureException(errors.New("write some error message for testing in go !"))
	// Since sentry emits events in the background we need to make sure
	// they are sent before we shut down
	sentry.Flush(time.Second * 5)
}
