package delivery

import (
	"microservice/config"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

type (
	EngineInterface interface {
		Bootstrap()
	}

	Engine struct {
		settings config.ManagerInterface
	}
)

func (engine *Engine) Bootstrap() {
	log := logrus.New()
	log.Out = os.Stdout

	mux := http.NewServeMux()

	log.WithFields(logrus.Fields{
		"Address": "0.0.0.0:3000",
		"Status":  "Listening",
	}).Info("A petersays instance is now born and healthy.")

	if err := http.ListenAndServe(":3000", mux); err != nil {
		log.WithError(err).Fatal("Wall broken! An engine error was raised!")
	}
}
