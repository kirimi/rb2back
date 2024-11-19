package api

import (
	pb "github.com/kirimi/rb2backend/user_service/pkg/generated/user_service"
)

type userAPI struct {
	pb.UnimplementedUserServiceServer
}

func NewUserAPI() pb.UserServiceServer {
	return &userAPI{}
}
