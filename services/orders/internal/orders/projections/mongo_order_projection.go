package projections

import (
	"context"
	"fmt"

	"emperror.dev/errors"
	attribute2 "go.opentelemetry.io/otel/attribute"

	"github.com/mehdihadeli/store-golang-microservice-sample/pkg/es/contracts/projection"
	"github.com/mehdihadeli/store-golang-microservice-sample/pkg/es/models"
	"github.com/mehdihadeli/store-golang-microservice-sample/pkg/http/http_errors/custom_errors"
	"github.com/mehdihadeli/store-golang-microservice-sample/pkg/logger"
	"github.com/mehdihadeli/store-golang-microservice-sample/pkg/mapper"
	"github.com/mehdihadeli/store-golang-microservice-sample/pkg/messaging/producer"
	"github.com/mehdihadeli/store-golang-microservice-sample/pkg/otel/tracing"
	"github.com/mehdihadeli/store-golang-microservice-sample/pkg/otel/tracing/attribute"
	"github.com/mehdihadeli/store-golang-microservice-sample/services/orders/internal/orders/contracts/repositories"
	"github.com/mehdihadeli/store-golang-microservice-sample/services/orders/internal/orders/dtos/v1"
	createOrderDomainEventsV1 "github.com/mehdihadeli/store-golang-microservice-sample/services/orders/internal/orders/features/creating_order/v1/events/domain_events"
	createOrderIntegrationEventsV1 "github.com/mehdihadeli/store-golang-microservice-sample/services/orders/internal/orders/features/creating_order/v1/events/integration_events"
	"github.com/mehdihadeli/store-golang-microservice-sample/services/orders/internal/orders/models/orders/read_models"
)

type mongoOrderProjection struct {
	mongoOrderRepository repositories.OrderReadRepository
	rabbitmqProducer     producer.Producer
	logger               logger.Logger
}

func NewMongoOrderProjection(mongoOrderRepository repositories.OrderReadRepository, rabbitmqProducer producer.Producer, logger logger.Logger) projection.IProjection {
	return &mongoOrderProjection{mongoOrderRepository: mongoOrderRepository, rabbitmqProducer: rabbitmqProducer, logger: logger}
}

func (m mongoOrderProjection) ProcessEvent(ctx context.Context, streamEvent *models.StreamEvent) error {
	// Handling and projecting event to elastic read model
	switch evt := streamEvent.Event.(type) {
	case *createOrderDomainEventsV1.OrderCreatedV1:
		return m.onOrderCreated(ctx, evt)
	}

	return nil
}

func (m *mongoOrderProjection) onOrderCreated(ctx context.Context, evt *createOrderDomainEventsV1.OrderCreatedV1) error {
	ctx, span := tracing.Tracer.Start(ctx, "mongoOrderProjection.onOrderCreated")
	span.SetAttributes(attribute.Object("Event", evt))
	span.SetAttributes(attribute2.String("OrderId", evt.OrderId.String()))
	defer span.End()

	items, err := mapper.Map[[]*read_models.ShopItemReadModel](evt.ShopItems)
	if err != nil {
		return errors.WrapIf(err, "[mongoOrderProjection_onOrderCreated.Map] error in mapping shopItems")
	}

	orderRead := read_models.NewOrderReadModel(evt.OrderId, items, evt.AccountEmail, evt.DeliveryAddress, evt.DeliveredTime)
	_, err = m.mongoOrderRepository.CreateOrder(ctx, orderRead)
	if err != nil {
		return tracing.TraceErrFromSpan(span, errors.WrapIf(err, "[mongoOrderProjection_onOrderCreated.CreateOrder] error in creating order with mongoOrderRepository"))
	}

	orderReadDto, err := mapper.Map[*dtosV1.OrderReadDto](orderRead)
	if err != nil {
		return tracing.TraceErrFromSpan(span, customErrors.NewApplicationErrorWrap(err, "[mongoOrderProjection_onOrderCreated.Map] error in mapping OrderReadDto"))
	}

	orderCreatedEvent := createOrderIntegrationEventsV1.NewOrderCreatedV1(orderReadDto)

	err = m.rabbitmqProducer.PublishMessage(ctx, orderCreatedEvent, nil)
	if err != nil {
		return tracing.TraceErrFromSpan(span, customErrors.NewApplicationErrorWrap(err, "[mongoOrderProjection_onOrderCreated.PublishMessage] error in publishing OrderCreated integration_events event"))
	}

	m.logger.Infow(fmt.Sprintf("[mongoOrderProjection.onOrderCreated] OrderCreated message with messageId `%s` published to the rabbitmq broker", orderCreatedEvent.MessageId), logger.Fields{"MessageId": orderCreatedEvent.MessageId, "Id": orderCreatedEvent.OrderId})

	m.logger.Infow(fmt.Sprintf("[mongoOrderProjection.onOrderCreated] order with id '%s' created", orderCreatedEvent.Id), logger.Fields{"Id": orderRead.Id, "MessageId": orderCreatedEvent.MessageId})

	return nil
}
