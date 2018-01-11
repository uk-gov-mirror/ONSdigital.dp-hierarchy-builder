package event_test

import (
	"context"
	"errors"
	"github.com/ONSdigital/dp-hierarchy-builder/event"
	"github.com/ONSdigital/dp-hierarchy-builder/event/eventtest"
	"github.com/ONSdigital/dp-reporter-client/reporter/reportertest"
	"github.com/ONSdigital/go-ns/kafka"
	"github.com/ONSdigital/go-ns/kafka/kafkatest"
	"github.com/ONSdigital/go-ns/log"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
	"github.com/ONSdigital/dp-import/events"
)

func TestConsume_UnmarshallError(t *testing.T) {
	Convey("Given an event consumer with an invalid schema and a valid schema", t, func() {

		messages := make(chan kafka.Message, 2)
		mockConsumer := kafkatest.NewMessageConsumer(messages)

		mockEventHandler := &eventtest.HandlerMock{
			HandleFunc: func(dataImportComplete *events.DataImportComplete) error {
				return nil
			},
		}

		expectedEvent := getExampleEvent()

		messages <- kafkatest.NewMessage([]byte("invalid schema"))
		messages <- kafkatest.NewMessage(marshal(*expectedEvent))

		consumer := event.NewConsumer()

		Convey("When consume messages is called", func() {

			consumer.Consume(mockConsumer, mockEventHandler, nil)
			waitForEventsToBeSentToHandler(mockEventHandler)

			Convey("Only the valid event is sent to the mockEventHandler ", func() {
				So(len(mockEventHandler.HandleCalls()), ShouldEqual, 1)

				event := mockEventHandler.HandleCalls()[0].DataImportComplete
				So(event.InstanceID, ShouldEqual, expectedEvent.InstanceID)
			})
		})
	})
}

func TestConsume(t *testing.T) {

	Convey("Given an event consumer with a valid schema", t, func() {

		messages := make(chan kafka.Message, 1)
		mockConsumer := kafkatest.NewMessageConsumer(messages)
		mockEventHandler := &eventtest.HandlerMock{
			HandleFunc: func(event *events.DataImportComplete) error {
				return nil
			},
		}

		expectedEvent := getExampleEvent()
		message := kafkatest.NewMessage(marshal(*expectedEvent))

		messages <- message

		consumer := event.NewConsumer()

		Convey("When consume is called", func() {

			consumer.Consume(mockConsumer, mockEventHandler, nil)
			waitForEventsToBeSentToHandler(mockEventHandler)

			Convey("A event is sent to the mockEventHandler ", func() {
				So(len(mockEventHandler.HandleCalls()), ShouldEqual, 1)

				event := mockEventHandler.HandleCalls()[0].DataImportComplete
				So(event.InstanceID, ShouldEqual, expectedEvent.InstanceID)
				So(event.CodeListID, ShouldEqual, expectedEvent.CodeListID)
				So(event.DimensionName, ShouldEqual, expectedEvent.DimensionName)
			})

			Convey("The message is committed", func() {
				So(message.Committed(), ShouldEqual, true)
			})
		})
	})
}

func TestConsume_HandlerError(t *testing.T) {

	Convey("Given an event consumer with a mock event handler that returns an error", t, func() {

		expectedError := errors.New("something bad happened in the event handler")

		messages := make(chan kafka.Message, 1)
		mockConsumer := kafkatest.NewMessageConsumer(messages)
		mockEventHandler := &eventtest.HandlerMock{
			HandleFunc: func(event *events.DataImportComplete) error {
				return expectedError
			},
		}

		mockErrorHandler := reportertest.NewImportErrorReporterMock(expectedError)

		expectedEvent := getExampleEvent()

		message := kafkatest.NewMessage(marshal(*expectedEvent))
		messages <- message

		consumer := event.NewConsumer()

		Convey("When consume is called", func() {

			consumer.Consume(mockConsumer, mockEventHandler, mockErrorHandler)
			waitForEventsToBeSentToHandler(mockEventHandler)

			Convey("The error handler is given the error returned from the event handler", func() {
				So(len(mockErrorHandler.NotifyCalls()), ShouldEqual, 1)
				So(mockErrorHandler.NotifyCalls()[0].Err, ShouldEqual, expectedError)
				So(mockErrorHandler.NotifyCalls()[0].ID, ShouldEqual, expectedEvent.InstanceID)
			})

			Convey("The message is committed", func() {
				So(message.Committed(), ShouldEqual, true)
			})
		})
	})
}

func TestClose(t *testing.T) {

	Convey("Given a consumer", t, func() {

		messages := make(chan kafka.Message, 1)
		mockConsumer := kafkatest.NewMessageConsumer(messages)
		mockEventHandler := &eventtest.HandlerMock{
			HandleFunc: func(event *events.DataImportComplete) error {
				return nil
			},
		}

		expectedEvent := getExampleEvent()
		message := kafkatest.NewMessage(marshal(*expectedEvent))

		messages <- message

		consumer := event.NewConsumer()

		consumer.Consume(mockConsumer, mockEventHandler, nil)

		Convey("When close is called", func() {

			err := consumer.Close(context.Background())

			Convey("Then no errors are returned", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

// marshal helper method to marshal a event into a []byte
func marshal(event events.DataImportComplete) []byte {
	bytes, err := events.DataImportCompleteSchema.Marshal(event)
	So(err, ShouldBeNil)
	return bytes
}

func getExampleEvent() *events.DataImportComplete {
	return &events.DataImportComplete{
		InstanceID: "7653",
		DimensionName: "geography",
		CodeListID: "543",
	}
}

func waitForEventsToBeSentToHandler(eventHandler *eventtest.HandlerMock) {

	start := time.Now()
	timeout := start.Add(time.Millisecond * 500)
	for {
		if len(eventHandler.HandleCalls()) > 0 {
			log.Debug("events have been sent to the handler", nil)
			break
		}

		if time.Now().After(timeout) {
			log.Debug("timeout hit", nil)
			break
		}

		time.Sleep(time.Millisecond * 10)
	}
}
