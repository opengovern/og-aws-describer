package main

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/opengovern/og-aws-describer/describer"
	"github.com/opengovern/og-util/pkg/describe"
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

type InvokeRequest struct {
	Data     map[string]string
	Metadata map[string]any
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
		jsonString := body.Data["mySbMsg"]
		unescaped, err := strconv.Unquote(jsonString)
		if err != nil {
			s.logger.Error("failed to unquote mySbMsg", zap.Error(err))
			return ctx.String(http.StatusBadRequest, "failed to unquote mySbMsg")
		}
		err = json.Unmarshal([]byte(unescaped), &bodyData)
		if err != nil {
			s.logger.Error("failed to unmarshal eventHubMessages", zap.Error(err))
			return ctx.String(http.StatusBadRequest, "failed to unmarshal eventHubMessages")
		}
	case len(body.Data["mySbMsg"]) > 0:
		jsonString := body.Data["mySbMsg"]
		unescaped, err := strconv.Unquote(string(jsonString))
		if err != nil {
			s.logger.Error("failed to unquote mySbMsg", zap.Error(err))
			return ctx.String(http.StatusBadRequest, "failed to unquote mySbMsg")
		}
		err = json.Unmarshal([]byte(unescaped), &bodyData)
		if err != nil {
			s.logger.Error("failed to unmarshal mySbMsg", zap.Error(err))
			return ctx.String(http.StatusBadRequest, "failed to unmarshal mySbMsg")
		}
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
