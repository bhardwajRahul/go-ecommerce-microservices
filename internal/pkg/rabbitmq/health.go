package rabbitmq

import (
	"context"

	"github.com/mehdihadeli/go-food-delivery-microservices/internal/pkg/health/contracts"
	"github.com/mehdihadeli/go-food-delivery-microservices/internal/pkg/rabbitmq/types"

	"emperror.dev/errors"
)

type gormHealthChecker struct {
	connection types.IConnection
}

func NewRabbitMQHealthChecker(connection types.IConnection) contracts.Health {
	return &gormHealthChecker{connection}
}

func (g gormHealthChecker) CheckHealth(ctx context.Context) error {
	if g.connection.IsConnected() {
		return nil
	} else {
		return errors.New("rabbitmq is not available")
	}
}

func (g gormHealthChecker) GetHealthName() string {
	return "rabbitmq"
}
