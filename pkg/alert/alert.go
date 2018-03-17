package alert

import (
	sentryhook "github.com/evalphobia/logrus_sentry"
	sentry "github.com/getsentry/raven-go"
	"github.com/sirupsen/logrus"
)

func init() {
	c, err := sentry.New("https://0ae4997acba549f18cf03b2ef7ce54d9@sentry.io/304990")
	if err != nil {
		panic(err)
	}
	hook, err := sentryhook.NewWithClientSentryHook(c, []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
	})
	if err != nil {
		panic(err)
	}
	logrus.AddHook(hook)
}
