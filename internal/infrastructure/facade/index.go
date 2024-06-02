package facade

import (
	"echo-model/config"
	loggers "echo-model/pkg/logger"
	"github.com/sirupsen/logrus"
)

type EchoModelFacade struct {
	Logger *logrus.Logger
	Config *configs.Config
	//Redis                    repository.IRedis
	//Broker                   repository.IBroker
	//ActivityLogs             repository.IActivityLogs
	//WorkflowLogs             repository.IWorkflowLogs
	//KafkaConsumer            sarama.Consumer
	//Temporal                 temporal_client.Client
}

func NewEchoModelFacade(config *configs.Config) *EchoModelFacade {
	log := loggers.NewLogger()
	//var err error
	//temporal_log := loggers.NewLoggerTemporal(log)
	//client, err := redis.NewRedis(config, log)
	//if err != nil {
	//	log.WithError(err).Error("[facade] cannot connect redis")
	//	panic(err)
	//}
	//consumer := broker.NewConsumerSarama(config, log)
	//if err != nil {
	//	log.WithError(err).Error("[facade] cannot connect kafka consumer")
	//	panic(err)
	//}
	//producer := broker.NewProducerSarama(config, log)
	//if err != nil {
	//	log.WithError(err).Error("[facade] cannot connect kafka producer")
	//	panic(err)
	//}
	//pg := postgres.NewPostgres(config.Postgres)

	//var pgStag *postgres.Postgres

	//temporal, err := temporal_client.Dial(temporal_client.Options{
	//	HostPort:           config.TemporalConfig.Uri,
	//	ContextPropagators: []wf_temporal.ContextPropagator{ctxpropagation.NewContextPropagator()},
	//	Namespace:          config.TemporalConfig.Namespace,
	//	Logger:             temporal_log,
	//})
	//if err != nil {
	//	log.WithError(err).Error("[facade] unable to create temporal client")
	//	panic(err)
	//}

	f := &EchoModelFacade{
		Logger: log,
		Config: config,
	}

	return f
}
