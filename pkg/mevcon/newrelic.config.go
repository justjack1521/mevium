package mevcon

import "github.com/newrelic/go-agent/v3/newrelic"

type NewRelicConfig struct {
	AppName    string `required:"true"`
	AppGUID    string `required:"true"`
	LicenseKey string `required:"true"`
}

func (c NewRelicConfig) NewApplication() (*newrelic.Application, error) {
	return newrelic.NewApplication(
		newrelic.ConfigAppName(c.AppName),
		newrelic.ConfigLicense(c.LicenseKey),
		newrelic.ConfigAppLogDecoratingEnabled(true),
		newrelic.ConfigAppLogForwardingEnabled(false),
	)
}
