package source_product

import (
	"context"
	"github.com/georgysavva/scany/v2/pgxscan"
	"go.uber.org/zap"
	"si-test/pkg/common"
)

type SourceProductRepository struct {
	Super  common.BaseCapsule
	Logger *zap.Logger
}

type ISourceProductRepository interface {
	GetAllData() ([]SourceProductDTO, error)
}

func (s SourceProductRepository) GetAllData() ([]SourceProductDTO, error) {
	var out []SourceProductDTO
	sql := `select * from source_product order by id`
	q, err := s.Super.Database.Query(context.Background(), sql)
	if err != nil {
		s.Logger.Error("error query get monthly enterprise", zap.Error(err))
		return nil, err
	}
	defer q.Close()
	err = pgxscan.ScanAll(&out, q)
	if err != nil {
		s.Logger.Error("error scan get monthly settlement", zap.Error(err))
		return nil, err
	}

	return out, nil
}

func NewSourceProductRepository(super common.BaseCapsule, logger *zap.Logger) ISourceProductRepository {
	return &SourceProductRepository{
		Super:  super,
		Logger: logger,
	}
}
