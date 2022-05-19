package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/OmarAouini/employee-manager/constants"
	"github.com/OmarAouini/employee-manager/core"
	"github.com/OmarAouini/employee-manager/utils"
	"github.com/Shopify/sarama"
)

type KafkaMessage struct {
	Action string      `json:"action"`
	Issuer string      `json:"issuer"`
	Data   interface{} `json:"data"`
}

//PRODUCER
////////////////////////////////////////////////

func CreateTopic(topicName string) (bool, error) {
	log.Printf("INTO CREATE TOPIC WITH TOPIC NAME: %s", topicName)

	//creating kafka topic with cluster admin struct
	config := *sarama.NewConfig()
	admin, err := sarama.NewClusterAdmin(constants.BROKERS, &config)
	if err != nil {
		return false, fmt.Errorf("error during create kafka cluster admin for create topic, err: " + err.Error())
	}
	defer func() { _ = admin.Close() }()

	//topic configuration map
	cleanupPolicy := "delete"
	deleteRetention := "180000"
	retentionMs := "180000"
	segmentBytes := "10485760"
	segmentMs := "300000"
	fileDeleteDelayMs := "1"

	configMap := make(map[string]*string)
	configMap["cleanup.policy"] = &cleanupPolicy
	configMap["delete.retention.ms"] = &deleteRetention
	configMap["retention.ms"] = &retentionMs
	configMap["segment.bytes"] = &segmentBytes
	configMap["segment.ms"] = &segmentMs
	configMap["file.delete.delay.ms"] = &fileDeleteDelayMs

	topicInfo := sarama.TopicDetail{
		NumPartitions:     8,
		ReplicationFactor: 2,
		ConfigEntries:     configMap,
	}
	err = admin.CreateTopic(topicName, &topicInfo, false)
	if err != nil {
		log.Printf("error during create new topic on kafka cluster: %s", err.Error())
		return false, fmt.Errorf("error during create new topic on kafka cluster, err: " + err.Error())
	}

	log.Printf("topic %s created!", topicName)
	return true, nil
}

func newProducer() (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(constants.BROKERS, config)
	if err != nil {
		return nil, fmt.Errorf("error during create kafka producer: %s", err.Error())
	}

	return producer, nil
}

func prepareMessage(topic, message string) *sarama.ProducerMessage {
	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Partition: -1,
		Value:     sarama.StringEncoder(message),
	}
	return msg
}

func SendMessage(topic string, message KafkaMessage) error {
	log.Printf("sending message to kafka to topic %s with body: %s", topic, utils.PrettyPrint(message))
	kafkaSyncProd, err := newProducer()
	if err != nil {
		log.Printf("error during creating kafka producer, err: %s", err.Error())
		return fmt.Errorf("error during creating kafka producer, err: %s", err.Error())
	}

	messageJson, _ := json.Marshal(message)
	preparedMessage := prepareMessage(topic, string(messageJson))
	_, _, err = kafkaSyncProd.SendMessage(preparedMessage)
	if err != nil {
		log.Printf("error during sending message to kakfa on topic %s, err: %s", topic, err.Error())
		return fmt.Errorf("error during sending message to kakfa on topic %s, err: %s", topic, err.Error())
	}
	log.Println("message to kafka sent.")

	return nil
}

//CONSUMER
////////////////////////////////////////////////

type KafkaConsumer struct {
	MsgConsumerGroup MsgConsumerGroup
}

type MsgConsumerGroup struct {
	Core   core.Core
	Topics []string
}

func (MsgConsumerGroup) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (MsgConsumerGroup) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h MsgConsumerGroup) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		var parsedMsg KafkaMessage

		err := json.Unmarshal(msg.Value, &parsedMsg)
		if err != nil {
			log.Printf("error unmarshalling job queued, message: %s", msg.Value)
			//Mark message as consumed and continue
			sess.MarkMessage(msg, "")
			continue
		}

		//spawn goroutines for each operation
		switch msg.Topic {
		case "a":
			go func() {

			}()
		case "b":
			go func() {

			}()
		default:
			//Mark message as consumed and continue
			sess.MarkMessage(msg, "")
			continue
		}

		//Mark message as consumed
		sess.MarkMessage(msg, "")
	}

	return nil
}

//this need to be ALWAYS running in separate goroutine, defined in main function
//in order to keep the subscribe to the specified topics

func SubscribeTopics(MsgConsumerGroup MsgConsumerGroup) {
	log.Println("trying to subscribe to encoding queue")
	consumerConfig := sarama.NewConfig()
	consumerConfig.Version = sarama.V2_8_0_0 // version
	consumerConfig.Consumer.Return.Errors = false
	//consumerConfig.Consumer.Offsets.AutoCommit.Enable = true      //  Disable auto submit , Change to manual
	//consumerConfig.Consumer.Offsets.AutoCommit.Interval = time.Second * 1 //  test 3 Seconds automatically submit
	consumerConfig.Consumer.Offsets.Initial = sarama.OffsetNewest

	cGroup, err := sarama.NewConsumerGroup(constants.BROKERS, "employee-manager-api", consumerConfig)
	if err != nil {
		log.Println("Err, cannot create consumer group")
		panic(err)
	}

	log.Println("consumer group created.")

	for {
		err := cGroup.Consume(context.Background(), MsgConsumerGroup.Topics, MsgConsumerGroup)
		if err != nil {
			log.Println(err.Error())
		}
	}

	//defer cGroup.Close()
}
