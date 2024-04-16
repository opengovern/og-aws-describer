package main

import (
	"encoding/json"
	"github.com/kaytu-io/kaytu-aws-describer/describer"
	"github.com/kaytu-io/kaytu-util/pkg/describe"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
	"os"
	"strings"
)

type Server struct {
	echoServer *echo.Echo
	logger     *zap.Logger
}

type TriggerBody struct {
	Data     map[string]json.RawMessage
	Metadata map[string]json.RawMessage
}

func (s *Server) azureFunctionsHandler(ctx echo.Context) error {
	var body TriggerBody
	err := ctx.Bind(&body)
	if err != nil {
		s.logger.Error("failed to bind request body", zap.Error(err))
		return ctx.String(http.StatusBadRequest, "failed to bind request body")
	}
	var bodyData describe.DescribeWorkerInput
	switch {
	case len(body.Data["eventHubMessages"]) > 0:
		jsonBody, err := body.Data["eventHubMessages"].MarshalJSON()
		if err != nil {
			s.logger.Error("failed to marshal eventHubMessages", zap.Error(err))
			return ctx.String(http.StatusBadRequest, "failed to marshal eventHubMessages")
		}
		err = json.Unmarshal(jsonBody, &bodyData)
		if err != nil {
			s.logger.Error("failed to unmarshal eventHubMessages", zap.Error(err))
			return ctx.String(http.StatusBadRequest, "failed to unmarshal eventHubMessages")
		}
	//case body.Data.QueueItem != "":
	//	fmt.Println(zap.Any("QueueItem", body.Data.QueueItem).String)
	//	unescaped, err := strconv.Unquote(body.Data.EventHubMessages)
	//	if err != nil {
	//		s.logger.Error("failed to unquote eventHubMessages", zap.Error(err))
	//		return ctx.String(http.StatusBadRequest, "failed to unquote eventHubMessages")
	//	}
	//	body.Data.QueueItem = unescaped
	//	s.logger.Info(zap.Any("QueueItemUnescaped", body.Data.QueueItem).String)
	//	return nil
	default:
		s.logger.Info(zap.Any("body", body).String)
		return nil
	}

	s.logger.Info("azureFunctionsHandler", zap.Any("bodyData", bodyData))

	err = describer.DescribeHandler(ctx.Request().Context(), s.logger, describer.TriggeredByAzureFunction, bodyData)
	if err != nil {
		s.logger.Error("failed to run describer", zap.Error(err), zap.Any("bodyData", bodyData))
		return ctx.String(http.StatusInternalServerError, "failed to run describer")
	}

	return ctx.String(http.StatusOK, "OK")
}

func main() {
	listenAddr := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddr = ":" + val
	}
	logger, _ := zap.NewProduction(zap.IncreaseLevel(zap.WarnLevel))
	if val, ok := os.LookupEnv("DEBUG"); ok && strings.ToLower(val) == "true" {
		logger, _ = zap.NewProduction(zap.IncreaseLevel(zap.DebugLevel))
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
