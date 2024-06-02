package broker

import (
	"strings"
	"time"

	"auto_reconcile_service_v2/configs"

	"github.com/IBM/sarama"
	"github.com/sirupsen/logrus"
)

type ConsumerHandler func(message *sarama.ConsumerMessage)

func NewProducerSarama(conf *configs.Config, l *logrus.Logger) sarama.SyncProducer {
	broker := strings.Split(conf.Kafka.Server, ",")
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true
	config.Net.SASL.Enable = true
	config.Net.SASL.SCRAMClientGeneratorFunc = func() sarama.SCRAMClient { return &XDGSCRAMClient{HashGeneratorFcn: SHA512} }
	config.Net.SASL.User = conf.Kafka.Username
	config.Net.SASL.Password = conf.Kafka.Password
	config.Net.SASL.Mechanism = sarama.SASLMechanism(conf.Kafka.Mechanism)

	producer, err := sarama.NewSyncProducer(broker, config)
	if err != nil {
		l.WithError(err).Error("failed to start Samara producer")
		panic(err)
	}

	return producer
}

func NewConsumerSarama(conf *configs.Config, l *logrus.Logger) sarama.Consumer {
	broker := strings.Split(conf.Kafka.Server, ",")
	config := sarama.NewConfig()
	config.Net.SASL.Enable = true
	config.Consumer.Offsets.AutoCommit.Enable = true
	config.Consumer.Offsets.AutoCommit.Interval = 1 * time.Second
	config.Consumer.Offsets.Initial = sarama.OffsetNewest
	config.Consumer.Offsets.Retry.Max = 3
	config.Net.SASL.SCRAMClientGeneratorFunc = func() sarama.SCRAMClient { return &XDGSCRAMClient{HashGeneratorFcn: SHA512} }
	config.Net.SASL.User = conf.Kafka.Username
	config.Net.SASL.Password = conf.Kafka.Password
	config.Net.SASL.Mechanism = sarama.SASLMechanism(conf.Kafka.Mechanism)
	config.Consumer.IsolationLevel = sarama.ReadCommitted
	consumer, err := sarama.NewConsumer(broker, config)
	if err != nil {
		l.WithError(err).Error("failed to start samara consumer")
		panic(err)
	}
	return consumer
}
