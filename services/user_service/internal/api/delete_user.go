package api

import (
	"context"
	pb "github.com/kirimi/rb2backend/user_service/pkg/generated/user_service"
	"github.com/rs/zerolog/log"
)

func (o *userAPI) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	log.Info().Msgf("Delete user request received %s", req.String())
	return &pb.DeleteUserResponse{}, nil
}
