package main

import (
	"encoding/json"
	"github.com/kaytu-io/kaytu-aws-describer/describer"
	"github.com/kaytu-io/kaytu-util/pkg/describe"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Server struct {
	echoServer *echo.Echo
	logger     *zap.Logger
}

type EventHubTriggerBody struct {
	Data struct {
		EventHubMessages string `json:"eventHubMessages"`
	}
	Metadata any
}

func (s *Server) azureFunctionsHandler(ctx echo.Context) error {
	var body EventHubTriggerBody
	err := ctx.Bind(&body)
	if err != nil {
		s.logger.Error("failed to bind request body", zap.Error(err))
		return ctx.String(http.StatusBadRequest, "failed to bind request body")
	}

	unescaped, err := strconv.Unquote(body.Data.EventHubMessages)
	if err != nil {
		s.logger.Error("failed to unquote eventHubMessages", zap.Error(err))
		return ctx.String(http.StatusBadRequest, "failed to unquote eventHubMessages")
	}

	body.Data.EventHubMessages = unescaped

	var bodyData []describe.DescribeWorkerInput
	err = json.Unmarshal([]byte(body.Data.EventHubMessages), &bodyData)
	if err != nil {
		s.logger.Error("failed to unmarshal eventHubMessages", zap.Error(err))
		return ctx.String(http.StatusBadRequest, "failed to unmarshal eventHubMessages")

	}
	s.logger.Info("azureFunctionsHandler", zap.Any("bodyData", bodyData))

	if len(bodyData) == 0 {
		return ctx.NoContent(http.StatusOK)
	}

	if len(bodyData) > 1 {
		s.logger.Error("multiple messages received", zap.Any("bodyData", bodyData))
		return ctx.String(http.StatusBadRequest, "multiple messages received")
	}

	err = describer.DescribeHandler(ctx.Request().Context(), s.logger, describer.TriggeredByAzureFunction, bodyData[0])
	if err != nil {
		s.logger.Error("failed to run describer", zap.Error(err))
		return ctx.String(http.StatusInternalServerError, "failed to run describer")
	}

	return ctx.NoContent(http.StatusOK)
}

func main() {
	listenAddr := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddr = ":" + val
	}
	logger := zap.NewNop()
	if val, ok := os.LookupEnv("DEBUG"); ok && strings.ToLower(val) == "true" {
		logger, _ = zap.NewProduction()
	}
	echoServer := echo.New()
	server := &Server{
		echoServer: echoServer,
		logger:     logger,
	}
	// the path is the trigger name e.g. POST /EventHubTrigger1
	echoServer.POST("/*", server.azureFunctionsHandler)
	logger.Info("Starting server", zap.String("addr", listenAddr))
	if err := echoServer.Start(listenAddr); err != nil {
		logger.Fatal("Failed to start server", zap.Error(err))
	}
}
