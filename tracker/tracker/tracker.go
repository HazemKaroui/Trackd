package tracker

import (
	context "context"
	//"tracker/models"
	"tracker/usecase"
)

type Server struct {
	trackedUsecase usecase.TrackedUsecase
	UnimplementedTrackedServiceServer
}

func InitServer(trackedUsecase usecase.TrackedUsecase) Server {
	return Server{
		trackedUsecase,
		UnimplementedTrackedServiceServer{},
	}
}

func (s *Server) CreateTracked(ctx context.Context, request *CreateTrackedRequest) (*CreateTrackedResponse, error) {
	return &CreateTrackedResponse{Success: true, Message: "Success to create tracked"}, nil
}

func (s *Server) GetTrackeds(ctx context.Context, request *GetTrackedsRequest) (*GetTrackedsResponse, error) {
	return &GetTrackedsResponse{Success: true, Message: "Success to get trackeds", Data: ""}, nil
}

func (s *Server) UpdateTracked(ctx context.Context, request *UpdateTrackedRequest) (*UpdateTrackedResponse, error) {
	return &UpdateTrackedResponse{Success: true, Message: "Success to update tracked"}, nil
}

func (s *Server) DeleteTracked(ctx context.Context, request *DeleteTrackedRequest) (*DeleteTrackedResponse, error) {
	return &DeleteTrackedResponse{Success: true, Message: "Success to delete tracked"}, nil
}
