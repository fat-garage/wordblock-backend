package service

import (
	"context"
	"github.com/fat-garage/wordblock-backend/models"
)

type Block interface {
	GetDIDBlockList(did string) ([]models.Block, error)
	AddBlock(ctx context.Context, did, content string) (string, error)
}
