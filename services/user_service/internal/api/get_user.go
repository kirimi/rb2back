package api

import (
	"context"
	pb "github.com/kirimi/rb2backend/user_service/pkg/generated/user_service"
	"github.com/rs/zerolog/log"
)

func (o *userAPI) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserConfigDto, error) {
	log.Info().Msgf("Get user request received %s", req.String())
	return &pb.UserConfigDto{
		IsEnabled:   false,
		PremiumType: pb.PremiumTypeDto_FREE,
	}, nil
}
