package installevents

import (
	"github.com/RiotMingle/newrelic-client-go/internal/http"
	"github.com/RiotMingle/newrelic-client-gogo/pkg/config"
	"github.com/RiotMingle/newrelic-client-gogo/pkg/logging"
)

// Installevents is used to communicate with the Install Events Service.
type Installevents struct {
	client http.Client
	logger logging.Logger
}

// New returns a new client for sending Install Events.
func New(config config.Config) Installevents {
	return Installevents{
		client: http.NewClient(config),
		logger: config.GetLogger(),
	}
}
