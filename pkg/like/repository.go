package like

import (
	"github.com/trinhdaiphuc/social-network/internal/logger"
)

type LikeRepository interface {
}

type repository struct {
	Logger *logger.AppLog
}

func NewLikeRepository( log *logger.AppLog) LikeRepository {
	return &repository{
		Logger:     log,
	}
}