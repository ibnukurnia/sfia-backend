package main

import (
	"fmt"
	"os"
	"strconv"
	"sv-sfia/db"
	"sv-sfia/middleware"
	"sv-sfia/routes"
	"sv-sfia/services"

	"github.com/gin-gonic/gin"
	"github.com/gookit/goutil/dump"
	"github.com/joho/godotenv"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func initZapLogger() *zap.Logger {
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	logger, _ := config.Build()
	return logger
}

func main() {
	zapLogger := initZapLogger()
	zap.ReplaceGlobals(zapLogger)

	defer zapLogger.Sync()

	logger := zapLogger.Sugar()

	err := godotenv.Load()
	if err != nil {
		logger.Fatal("Failed to load .env", zap.Error(err))
	}

	db, err := db.ConnectDb()

	if err != nil {
		logger.Fatal("Failed to connect database: ", zap.Error(err))
	}

	host := fmt.Sprintf("%s:%s", os.Getenv("MINIO_IP"), os.Getenv("MINIO_PORT_API"))
	dump.P(host)

	minioClient, err := minio.New(host, &minio.Options{
		Creds:  credentials.NewStaticV4(os.Getenv("MINIO_ACCESS_KEY"), os.Getenv("MINIO_SECRET_KEY"), ""),
		Secure: false,
	})

	if err != nil {
		logger.Fatalf("Failed to initialize MinIO client: %v", err)
	}

	serviceProvider := services.NewServiceProvider(db, minioClient)

	srv := gin.Default()

	srv.Use(middleware.Cors(), middleware.LogMiddleware(zapLogger))

	routes.InitApiRouter(srv, serviceProvider)

	appPort := 5000
	port := os.Getenv("APP_PORT")
	if port != "" {
		n, err := strconv.Atoi(port)
		if err == nil {
			appPort = n
		}
	}

	go func() {
		zap.L().Info("Starting HTTP server", zap.Int("port", appPort))
	}()

	if err := srv.Run(fmt.Sprintf(":%d", appPort)); err != nil {
		zap.L().Error("Failed to start server: ", zap.Error(err))
	}
}

// goose -dir migrations clickhouse "http://localhost:9005?database=sfia&username=default&password=" up
