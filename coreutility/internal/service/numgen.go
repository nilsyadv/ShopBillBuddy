package service

import (
	"github.com/go-redis/redis"
	wraperror "github.com/nilsyadv/ShopBillBuddy/common/pkg/error"
	"github.com/nilsyadv/ShopBillBuddy/common/pkg/logger"
	"github.com/nilsyadv/ShopBillBuddy/coreutility/internal/model"
)

type NumGenService struct {
	rds *redis.Client
	log logger.InterfaceLogger
}

func NewNumGenService(rds *redis.Client, log logger.InterfaceLogger) *NumGenService {
	return &NumGenService{
		rds: rds,
		log: log,
	}
}

func (numgenser NumGenService) SetKey(req model.NumberGeneratorRequest) *wraperror.WrappedError {
	return nil
}

func (numgenser NumGenService) NextKey(req model.NumberGeneratorRequest) (string, *wraperror.WrappedError) {
	return "", nil
}
