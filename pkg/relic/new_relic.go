package relic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/justjack1521/mevium/pkg/mevcon"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/sirupsen/logrus"
	"net/http"
)

type NewRelic struct {
	LicenseKey  string
	EntityGUID  string
	EntityName  string
	Application *newrelic.Application
}

func NewRelicApplication(config mevcon.NewRelicConfig) *NewRelic {
	rel, err := config.NewApplication()
	if err != nil {
		panic(err)
	}
	return &NewRelic{
		LicenseKey:  config.LicenseKey,
		EntityGUID:  config.AppGUID,
		EntityName:  config.AppName,
		Application: rel,
	}
}

func (a *NewRelic) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (a *NewRelic) Fire(entry *logrus.Entry) error {
	evt := map[string]interface{}{
		"timestamp":   entry.Time.Unix(),
		"message":     entry.Message,
		"logtype":     entry.Level.String(),
		"entity.name": a.EntityName,
		"entity.guid": a.EntityGUID,
	}
	for k, v := range entry.Data {
		evt[k] = v
	}
	body, err := json.Marshal(evt)
	if err != nil {
		return err
	}
	req, err := http.Post(
		fmt.Sprintf("https://log-api.newrelic.com/log/v1?Api-Key=%s", a.LicenseKey),
		"application/json",
		bytes.NewBuffer(body),
	)
	defer req.Body.Close()
	if err != nil {
		return err
	}
	return nil
}
