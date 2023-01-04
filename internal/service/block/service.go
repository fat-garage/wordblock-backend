package block

import (
	"context"
	"github.com/fat-garage/wordblock-backend/internal/dao"
	"github.com/fat-garage/wordblock-backend/models"
	"github.com/fat-garage/wordblock-backend/pkg/gipfs"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"time"
)

type srv struct {
	dao *dao.BlockDao
}

func newSrv(db *gorm.DB) *srv {
	return &srv{
		dao: dao.NewUserDao(db),
	}
}

// GetDIDBlockList .
func (s *srv) GetDIDBlockList(did string) ([]models.Block, error) {
	return s.dao.GetDIDBlockList(dao.Query{
		DID: did,
	})
}

// AddBlock .
func (s *srv) AddBlock(ctx context.Context, did, content string) (string, error) {
	fileName := did + "-" + cast.ToString(time.Now().Unix())
	put, err := gipfs.W3sClient.Put(ctx, gipfs.NewMyFile(fileName, []byte(content)))
	if err != nil {
		return "", errors.New("upload to ipfs failed")
	}
	if err := s.dao.AddBlock(did, fileName, put.String()); err != nil {
		return "", err
	}
	return put.String(), nil
}
