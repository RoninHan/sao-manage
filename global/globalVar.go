package global

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-redis/redis"
	"github.com/minio/minio-go"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"sao-manage/config"
)

var (Settings config.ServerConfig)
var (Lg *zap.Logger)
var (Trans ut.Translator)
var ( DB    *gorm.DB)
var ( Redis  *redis.Client)
var (MinioClient *minio.Client)