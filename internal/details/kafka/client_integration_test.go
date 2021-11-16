package kafka_test

// import (
// 	"fmt"
// 	"log"
// 	"testing"

// 	"github.com/Shopify/sarama"
// 	"github.com/stretchr/testify/assert"
// )

// func TestNewKafkaClient(t *testing.T) {
// 	if testing.Short() {
// 		t.Skip()
// 	}

// 	suite := newSuite(t)
// 	suite.setup()
// 	defer suite.teardown()

// 	tests := map[string]func(*testing.T){
// 		"should return error when kafka is down": func(t *testing.T) {
// 			// given
// 			// when
// 			// then
// 		},
// 	}

// 	for name, run := range tests {
// 		t.Run(name, func(t *testing.T) {
// 			suite.migrationsUp()
// 			run(t)
// 			suite.migrationsDown()
// 		})
// 	}
// }

// type suite struct {
// 	t        *testing.T
// 	producer sarama.SyncProducer
// 	// consumer sarama.Consumer
// }

// func newSuite(t *testing.T) *suite {
// 	return &suite{
// 		t: t,
// 	}
// }

// func (s *suite) setup() {
// 	testHosts := []string{"localhost:9092"}

// 	testCfg := sarama.NewConfig()
// 	testCfg.Producer.RequiredAcks = sarama.WaitForAll
// 	testCfg.Producer.Retry.Max = 5
// 	testCfg.Producer.Return.Successes = true

// 	testProducer, err := sarama.NewSyncProducer(testHosts, testCfg)
// 	assert.NoError(s.t, err, "failed to setup test client")

// 	s.producer = testProducer
// }

// func (s *suite) teardown() {
// 	if err := s.producer.Close(); err != nil {
// 		log.Fatalln(err)
// 	}
// }

// func (s *suite) migrationsUp() {
// 	fmt.Println("migrationsUp...")

// 	kafkaMessage := sarama.ProducerMessage{
// 		Topic:     "test",
// 		Key:       sarama.StringEncoder("somekey"),
// 		Value:     sarama.StringEncoder("somevalue"),
// 		Partition: 1,
// 	}

// 	partition, offset, err := s.producer.SendMessage(&kafkaMessage)
// 	fmt.Println("MESSAGE SENT >>>> ", partition, offset, err)
// }

// func (s *suite) migrationsDown() {
// 	fmt.Println("migrationsDown...")
// }
