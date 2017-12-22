package servers

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/lvzhihao/silk/models"
	pbs "github.com/lvzhihao/silk/pbs"

	"go.uber.org/zap"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type Server struct {
	DB     *gorm.DB
	Logger *zap.Logger
}

func (c *Server) gdefer() {
	if r := recover(); r != nil {
		c.Logger.Error("Recovered Error", zap.Any("r", r))
		grpc.Errorf(codes.Internal, fmt.Errorf("%+v", r).Error())
	}
}

func (c *Server) CreateAccount(ctx context.Context, in *pbs.AccountRequest) (*pbs.AccountResponse, error) {
	defer c.gdefer()
	account := &models.Account{
		Platform:  in.GetPlatform(),
		AccountId: in.GetAccountId(),
		SerialNo:  in.GetSerialNo(),
		NickName:  in.GetNickName(),
		HeadImage: in.GetHeadImage(),
		QrCode:    in.GetQrCode(),
		Metadata:  in.GetMetadata(),
	}
	c.Logger.Debug("Request CreateAccount", zap.Any("ctx", ctx), zap.Any("request", in), zap.Any("account", account))
	err := account.Create(c.DB)
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, err.Error())
	} else {
		return &pbs.AccountResponse{
			Id: uint64(account.Model.ID),
		}, nil //grpc.Errorf(codes.AlreadyExists, "this account exists")
	}
}
