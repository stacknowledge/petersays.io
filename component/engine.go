package component

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/stacknowledge/petersays.io/component/prophet"
	"github.com/stacknowledge/petersays.io/configs"
)

type (
	Application interface {
		Bootstrap()
		Run()
	}

	ApplicationWheel interface {
		RegisterHandlers(*mux.Router)
	}

	EngineError struct {
		Error string `json:"error"`
	}

	Engine struct {
		router *mux.Router
		config *configs.Config
		logger *logrus.Logger
	}
)

func (engine *Engine) Boot() {
	engine.router = mux.NewRouter()
	engine.config = configs.NewConfigurations()
	engine.logger = logrus.New()

	engine.registerWheel(prophet.NewWheel(engine.logger))
	engine.registerSpareWheels()

	engine.start()
}

func (engine *Engine) start() {
	engine.logger.WithFields(logrus.Fields{
		"Address": engine.config.Application.Address,
		"Port":    engine.config.Application.Port,
		"Status":  "Listening",
	}).Info("A " + engine.config.Application.Name + " node is now born and healthy.")

	if err := http.ListenAndServe(
		engine.config.Application.Address+":"+engine.config.Application.Port,
		engine.router,
	); err != nil {
		engine.logger.WithError(err).
			Fatal("Wall broken! An engine error was raised!")
	}
}

func (engine *Engine) registerWheel(wheel ApplicationWheel) {
	wheel.RegisterHandlers(engine.router)

	engine.logger.WithFields(logrus.Fields{
		"Wheel": strings.Split(fmt.Sprintf("%T", wheel), ".")[1],
	}).Info("Wheel Health Check | OK")
}

func (engine *Engine) registerSpareWheels() {
	engine.router.NotFoundHandler = http.HandlerFunc(handleNotFound)
	engine.router.MethodNotAllowedHandler = http.HandlerFunc(handleMethodNotAllowed)
}

func handleNotFound(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	writer.WriteHeader(http.StatusNotFound)

	json.NewEncoder(writer).Encode(&EngineError{"You are in the forest kid!"})
}

func handleMethodNotAllowed(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	writer.WriteHeader(http.StatusMethodNotAllowed)

	json.NewEncoder(writer).Encode(&EngineError{"Peters has no opinion about that request method."})
}
