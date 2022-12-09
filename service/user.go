package service

import (
	"context"

	"github.com/TemurMannonov/medium_notification_service/config"
	pb "github.com/TemurMannonov/medium_notification_service/genproto/user_service"
)

type NotificationService struct {
	// pb.UnimplementedNotificationServiceServer
	cfg *config.Config
}

func NewNotificationService(cfg *config.Config) *NotificationService {
	return &NotificationService{
		cfg: cfg,
	}
}

func (s *NotificationService) Create(ctx context.Context, req *pb.User) (*pb.User, error) {

	// if err != nil {
	// 	return nil, status.Errorf(codes.Internal, "Internal server error: %v", err)
	// }
	return nil, nil

}
