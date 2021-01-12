package common

import "github.com/trinhdaiphuc/social-network/pkg/models"

var (
	AddChannelObserver = make(map[string]chan *models.Post)
)
