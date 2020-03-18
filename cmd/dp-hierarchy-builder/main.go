package main

import (
	"context"
	"github.com/ONSdigital/go-ns/server"
	"github.com/gorilla/mux"
	"os"
	"os/signal"
	"syscall"

	"github.com/ONSdigital/dp-graph/graph"
	"github.com/ONSdigital/dp-healthcheck/healthcheck"
	"github.com/ONSdigital/dp-hierarchy-builder/config"
	"github.com/ONSdigital/dp-hierarchy-builder/event"
	"github.com/ONSdigital/dp-hierarchy-builder/hierarchy"
	kafka "github.com/ONSdigital/dp-kafka"
	"github.com/ONSdigital/dp-reporter-client/reporter"
	"github.com/ONSdigital/go-ns/log"
)

var (
	// BuildTime represents the time in which the service was built
	BuildTime string
	// GitCommit represents the commit (SHA-1) hash of the service that is running
	GitCommit string
	// Version represents the version of the service that is running
	Version string
)

func main() {
	log.Namespace = "dp-hierarchy-builder"
	ctx := context.Background()

	log.Debug("Starting hierarchy builder", nil)

	cfg, err := config.Get()
	if err != nil {
		log.Error(err, nil)
		os.Exit(1)
	}

	// sensitive fields are omitted from cfg.String().
	log.Debug("loaded config", log.Data{"cfg": cfg})

	// a channel used to signal a graceful exit is required.
	errorChannel := make(chan error)

	kafkaBrokers := cfg.KafkaAddr
	cgChannels := kafka.CreateConsumerGroupChannels(true)
	kafkaConsumer, err := kafka.NewConsumerGroup(
		ctx,
		kafkaBrokers,
		cfg.ConsumerTopic,
		cfg.ConsumerGroup,
		kafka.OffsetNewest,
		true,
		cgChannels,
	)
	exitIfError(err)

	pChannels := kafka.CreateProducerChannels()
	useDefaultMaxMessageSize := 0 // pass zero to use the default
	kafkaProducer, err := kafka.NewProducer(ctx, kafkaBrokers, cfg.ProducerTopic, useDefaultMaxMessageSize, pChannels)
	exitIfError(err)

	errorProducerChannels := kafka.CreateProducerChannels()
	kafkaErrorProducer, err := kafka.NewProducer(ctx, kafkaBrokers, cfg.ErrorProducerTopic, useDefaultMaxMessageSize, errorProducerChannels)
	exitIfError(err)

	avroProducer := event.NewAvroProducer(kafkaProducer)

	db, err := graph.NewHierarchyStore(ctx)
	exitIfError(err)

	// when errors occur - we send a message on an error topic.
	errorHandler, err := reporter.NewImportErrorReporter(kafkaErrorProducer, log.Namespace)
	exitIfError(err)

	eventHandler := event.NewDataImportCompleteHandler(&hierarchy.Store{db}, avroProducer)

	eventConsumer := event.NewConsumer()
	eventConsumer.Consume(kafkaConsumer, eventHandler, errorHandler)

	hc := startHealthChecks(ctx, cfg, kafkaConsumer, kafkaProducer, kafkaErrorProducer, db)

	apiErrors, httpServer := startApi(hc, cfg)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	// this will block (main) until a fatal error occurs
	select {
	case err := <-apiErrors:
		log.ErrorC("http server error", err, nil)
	case err := <-kafkaConsumer.Channels().Errors:
		log.ErrorC("kafka consumer", err, nil)
	case err := <-kafkaProducer.Channels().Errors:
		log.ErrorC("kafka result producer", err, nil)
	case err := <-kafkaErrorProducer.Channels().Errors:
		log.ErrorC("kafka error producer", err, nil)
	case err := <-errorChannel:
		log.ErrorC("error channel", err, nil)
	case <-signals:
		log.Debug("os signal received", nil)
	}

	ctx, cancel := context.WithTimeout(ctx, cfg.GracefulShutdownTimeout)

	// gracefully dispose resources
	err = httpServer.Close(ctx)
	logIfError(err)

	err = eventConsumer.Close(ctx)
	logIfError(err)

	err = kafkaConsumer.Close(ctx)
	logIfError(err)

	err = kafkaProducer.Close(ctx)
	logIfError(err)

	err = kafkaErrorProducer.Close(ctx)
	logIfError(err)

	err = db.Close(ctx)
	logIfError(err)

	hc.Stop()

	// cancel the timer in the shutdown context
	cancel()

	log.Debug("graceful shutdown was successful", nil)
	os.Exit(1)
}

func startApi(hc healthcheck.HealthCheck, cfg *config.Config) (chan error, *server.Server) {
	router := mux.NewRouter()
	router.HandleFunc("/health", hc.Handler)
	apiErrors := make(chan error, 1)

	httpServer := server.New(cfg.BindAddr, router)
	// Disable this here to allow main to manage graceful shutdown of the entire app.
	httpServer.HandleOSSignals = false

	go func() {
		log.Debug("starting api", nil)
		if err := httpServer.ListenAndServe(); err != nil {
			log.Error(err, log.Data{})
			hc.Stop()
			apiErrors <- err
		}
	}()
	return apiErrors, httpServer
}

func startHealthChecks(
	ctx context.Context,
	cfg *config.Config,
	kafkaConsumer *kafka.ConsumerGroup,
	kafkaProducer *kafka.Producer,
	kafkaErrorProducer *kafka.Producer,
	db *graph.DB) healthcheck.HealthCheck {

	versionInfo, err := healthcheck.NewVersionInfo(BuildTime, GitCommit, Version)
	exitIfError(err)
	hc := healthcheck.New(versionInfo, cfg.HealthCheckRecoveryInterval, cfg.HealthCheckInterval)

	err = hc.AddCheck("Kafka Consumer", kafkaConsumer.Checker)
	exitIfError(err)

	err = hc.AddCheck("Kafka Producer", kafkaProducer.Checker)
	exitIfError(err)

	err = hc.AddCheck("Kafka Error Producer", kafkaErrorProducer.Checker)
	exitIfError(err)

	err = hc.AddCheck("GraphDB", db.Checker)
	exitIfError(err)

	hc.Start(ctx)
	return hc
}

func exitIfError(err error) {
	if err != nil {
		log.Error(err, nil)
		os.Exit(1)
	}
}

func logIfError(err error) {
	if err != nil {
		log.Error(err, nil)
	}
}
