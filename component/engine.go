package component

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/stacknowledge/petersays.io/component/example"
	"github.com/stacknowledge/petersays.io/configs"
)

type (
	Application interface {
		Bootstrap()
		Run()
	}

	ApplicationWheel interface {
		RegisterHandlers(*http.ServeMux)
	}

	Engine struct {
		server *http.ServeMux
		config *configs.Config
		logger *logrus.Logger
	}
)

func (engine *Engine) Boot() {
	engine.server = http.NewServeMux()
	engine.config = configs.NewConfigurations()
	engine.logger = logrus.New()

	engine.registerWheel(example.NewWheel(engine.logger))

	engine.start()
}

func (engine *Engine) registerWheel(wheel ApplicationWheel) {
	wheel.RegisterHandlers(engine.server)

	engine.logger.WithFields(logrus.Fields{
		"Wheel": strings.Split(fmt.Sprintf("%T", wheel), ".")[1],
	}).Info("Wheel Health Check | OK")
}

func (engine *Engine) start() {
	engine.logger.WithFields(logrus.Fields{
		"Address": engine.config.Application.Address,
		"Port":    engine.config.Application.Port,
		"Status":  "Listening",
	}).Info("A " + engine.config.Application.Name + " node is now born and healthy.")
	if err := http.ListenAndServe(
		engine.config.Application.Address+":"+engine.config.Application.Port,
		engine.server,
	); err != nil {
		engine.logger.WithError(err).
			Fatal("Wall broken! An engine error was raised!")
	}

}
