package service

import (
	"github.com/google/wire"

	v1 "kratos-realworld/api/realworld/v1"
	"kratos-realworld/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewRealWorldService)

// GreeterService is a greeter service.
type RealWorldService struct {
	v1.UnimplementedRealWorldServer

	uu *biz.UserUsecase
}

// NewGreeterService new a greeter service.
func NewRealWorldService(uu *biz.UserUsecase) *RealWorldService {
	return &RealWorldService{uu: uu}
}
