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

type InvokeRequest struct {
	Data     map[string]json.RawMessage
	Metadata map[string]json.RawMessage
}

type InvokeResponse struct {
	Outputs     map[string]any
	Logs        []string
	ReturnValue any
}

func (s *Server) azureFunctionsHandler(ctx echo.Context) error {
	var body InvokeRequest
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
	case len(body.Data["mySbMsg"]) > 0:
		jsonBody, err := body.Data["mySbMsg"].MarshalJSON()
		if err != nil {
			s.logger.Error("failed to marshal mySbMsg", zap.Error(err))
			return ctx.String(http.StatusBadRequest, "failed to marshal mySbMsg")
		}
		err = json.Unmarshal(jsonBody, &bodyData)
		if err != nil {
			s.logger.Error("failed to unmarshal mySbMsg", zap.Error(err))
			return ctx.String(http.StatusBadRequest, "failed to unmarshal mySbMsg")
		}
		s.logger.Info("mySbMsg", zap.Any("bodyData", bodyData), zap.Any("jsonBody", jsonBody))
	default:
		for k, v := range body.Data {
			s.logger.Info("data", zap.String("key", k), zap.Any("value", v))
		}
		return ctx.String(http.StatusBadRequest, "no data found")
	}

	s.logger.Info("azureFunctionsHandler", zap.Any("bodyData", bodyData))

	err = describer.DescribeHandler(ctx.Request().Context(), s.logger, describer.TriggeredByAzureFunction, bodyData)
	if err != nil {
		s.logger.Error("failed to run describer", zap.Error(err), zap.Any("bodyData", bodyData))
		return ctx.String(http.StatusInternalServerError, "failed to run describer")
	}

	return ctx.JSON(http.StatusOK, InvokeResponse{})
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
