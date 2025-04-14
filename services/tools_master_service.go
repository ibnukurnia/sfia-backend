package services

import (
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"sv-sfia/dto"
	"sv-sfia/dto/requests"
	"sv-sfia/dto/responses"
	"sv-sfia/models"

	"github.com/minio/minio-go/v7"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ToolsMasterService struct {
	db          *gorm.DB
	minioClient *minio.Client
}

func newToolsMasterService(db *gorm.DB, minioClient *minio.Client) *ToolsMasterService {
	return &ToolsMasterService{db: db, minioClient: minioClient}
}

func (service ToolsMasterService) GetToolList() ([]responses.ToolsResponse, *dto.ApiError) {
	tools := []models.Tool{}

	err := service.db.
		Order("created_at asc").
		Find(&tools).Error

	if err != nil {
		zap.L().Error("error query tool list", zap.Error(err))
		return []responses.ToolsResponse{}, dto.InternalError(err)
	}

	results := []responses.ToolsResponse{}

	for _, tool := range tools {
		urlPhoto := fmt.Sprintf("http://%s:%s/%s/%s/%s", os.Getenv("MINIO_IP"), os.Getenv("MINIO_PORT_API"), os.Getenv("MINIO_BUCKET_NAME"), os.Getenv("MINIO_TOOLS_PATH"), tool.Url)
		results = append(results, responses.ToolsResponse{
			Id:   tool.Base.Uuid.String(),
			Name: tool.Name,
			Url:  urlPhoto,
		})
	}

	return results, nil
}

func (service ToolsMasterService) AddTools(req requests.AddToolsRequest, fileHeader *multipart.FileHeader) *dto.ApiError {

	fileExt := filepath.Ext(fileHeader.Filename)

	if fileHeader.Size > 2*1024*1024 {
		zap.L().Error("file too large or too small")
		return dto.BadRequestError(errors.New("file too large or too small"))
	}
	if fileHeader.Header["Content-Type"][0] != "image/jpeg" {
		if fileExt != "jpg" && fileExt != "jpeg" && fileExt != "png" {
			zap.L().Error("file format is not allowed")
			return dto.BadRequestError(errors.New("file format is not allowed"))
		}
	}

	file, err := fileHeader.Open()

	if err != nil {
		zap.L().Error("failed open file")
		return dto.BadRequestError(errors.New("failed open file"))
	}
	defer file.Close()

	fileName := strings.ReplaceAll(req.Name, " ", "_") + fileExt
	objectName := os.Getenv("MINIO_TOOLS_PATH") + "/" + fileName
	contentType := fileHeader.Header.Get("Content-Type")
	if err := UploadFile(service.minioClient, os.Getenv("MINIO_BUCKET_NAME"), objectName, contentType, file, fileHeader.Size); err != nil {
		zap.L().Error("failed upload file")
		return dto.BadRequestError(errors.New("failed upload file"))
	}

	tool := models.Tool{
		Name: req.Name,
		Url:  fileName,
	}

	err = service.db.Create(&tool).Error
	if err != nil {
		zap.L().Error("error create tool", zap.Error(err))
		return dto.InternalError(err)
	}

	return nil
}

func (service ToolsMasterService) UpdateTools(req requests.UpdateToolsRequest, fileHeader *multipart.FileHeader) *dto.ApiError {
	tools := models.Tool{}

	err := service.db.
		Where("uuid = ?", req.ToolsId).
		First(&tools).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		zap.L().Warn("tools master not found", zap.String("uuid", req.ToolsId))
		return dto.NotFoundError(err)
	}

	if err != nil {
		zap.L().Error("error get tools", zap.Error(err))
		return dto.InternalError(err)
	}

	fileExt := filepath.Ext(fileHeader.Filename)

	if fileHeader.Size > 2*1024*1024 {
		zap.L().Error("file too large or too small")
		return dto.BadRequestError(errors.New("file too large or too small"))
	}
	if fileHeader.Header["Content-Type"][0] != "image/jpeg" {
		if fileExt != "jpg" && fileExt != "jpeg" && fileExt != "png" {
			zap.L().Error("file format is not allowed")
			return dto.BadRequestError(errors.New("file format is not allowed"))
		}
	}

	file, err := fileHeader.Open()

	if err != nil {
		zap.L().Error("failed open file")
		return dto.BadRequestError(errors.New("failed open file"))
	}

	defer file.Close()

	oldFile := os.Getenv("MINIO_TOOLS_PATH") + "/" + tools.Url
	RemoveFile(service.minioClient, os.Getenv("MINIO_BUCKET_NAME"), oldFile)

	fileName := strings.ReplaceAll(req.Name, " ", "_") + fileExt
	objectName := os.Getenv("MINIO_TOOLS_PATH") + "/" + fileName
	contentType := fileHeader.Header.Get("Content-Type")
	if err := UploadFile(service.minioClient, os.Getenv("MINIO_BUCKET_NAME"), objectName, contentType, file, fileHeader.Size); err != nil {
		zap.L().Error("failed upload file")
		return dto.BadRequestError(errors.New("failed upload file"))
	}

	tools.Name = req.Name
	tools.Url = fileName

	err = service.db.Save(&tools).Error
	if err != nil {
		zap.L().Error("error update tools", zap.Error(err))
		return dto.InternalError(err)
	}

	return nil
}

func (service ToolsMasterService) DeleteTools(toolsId string) *dto.ApiError {
	tools := models.Tool{}

	err := service.db.
		Where("uuid = ?", toolsId).
		First(&tools).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		zap.L().Warn("tools master not found", zap.String("uuid", toolsId))
		return dto.NotFoundError(err)
	}

	if err != nil {
		zap.L().Error("error get tools master", zap.Error(err))
		return dto.InternalError(err)
	}

	oldFile := os.Getenv("MINIO_TOOLS_PATH") + "/" + tools.Url
	RemoveFile(service.minioClient, os.Getenv("MINIO_BUCKET_NAME"), oldFile)

	err = service.db.Delete(&tools).Error
	if err != nil {
		zap.L().Error("error delete tools", zap.Error(err))
		return dto.InternalError(err)
	}

	return nil
}

func RemoveFile(minioClient *minio.Client, bucketName, objectName string) (err error) {
	option := minio.RemoveObjectOptions{}
	ctx := context.Background()

	if err = minioClient.RemoveObject(ctx, bucketName, objectName, option); err != nil {
		return err
	}

	return nil
}

func UploadFile(minioClient *minio.Client, bucketName, objectName, contentType string, file io.Reader, objectSize int64) (err error) {
	option := minio.PutObjectOptions{ContentType: contentType}
	ctx := context.Background()

	_, err = minioClient.PutObject(ctx, bucketName, objectName, file, objectSize, option)

	if err != nil {
		return err
	}

	return nil
}
