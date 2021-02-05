package controller

import (
	"context"

	"github.com/jonayrodriguez/translation-service/internal/log"
	"github.com/jonayrodriguez/translation-service/internal/translation/service"
	"go.uber.org/zap"

	empty "github.com/golang/protobuf/ptypes/empty"
	pb "github.com/jonayrodriguez/translation-service/api/translation/v1"
	"google.golang.org/grpc"
)

// TranslationController struct of the gRPC translation controller.
type TranslationController struct {
	pb.UnimplementedTranslationServiceServer
	service service.TranslationService
}

// Register the TranslationController server into the gRPC server instance
func (c *TranslationController) Register(server grpc.ServiceRegistrar) {
	log.Logger.Debug("Registering the TranslationController")
	pb.RegisterTranslationServiceServer(server, c)
}

// TranslationController returns a new instance of the controller
func NewTranslationService(svc service.TranslationService) *TranslationController {
	log.Logger.Debug("Adding the TranslationController to the controller")
	return &TranslationController{service: svc}
}

// GetTranslation to retrieve the translated message.
func (c *TranslationController) GetTranslation(ctx context.Context, in *pb.GetTranslationRequest) (*pb.GetTranslationResponse, error) {
	log.Logger.Debug("Retrieving the translation", zap.Any("KeyPattern", &in.KeyPattern), zap.Stringp("Language", &in.Language), zap.Stringp("Scope", &in.Scope))
	return c.service.GetTranslation(ctx, in)
}

// AddTranslation to add a new translated message.
func (c *TranslationController) AddTranslation(ctx context.Context, in *pb.AddTranslationRequest) (*pb.AddTranslationResponse, error) {
	log.Logger.Debug("Adding a new translation", zap.Any("Key", &in.Key), zap.Stringp("Language", &in.Language), zap.Stringp("Scope", &in.Scope), zap.Stringp("Message", &in.Message))
	return c.service.AddTranslation(ctx, in)
}

// UpdateTranslation to update an existing translated message.
func (c *TranslationController) UpdateTranslation(ctx context.Context, in *pb.UpdateTranslationRequest) (*pb.UpdateTranslationResponse, error) {
	log.Logger.Debug("Updating an existing translation", zap.Any("Key", &in.Key), zap.Stringp("Language", &in.Language),
		zap.Stringp("Scope", &in.Scope), zap.Stringp("Message", &in.Message))
	return c.service.UpdateTranslation(ctx, in)
}

// DeleteTranslation to delete an existing translated message.
func (c *TranslationController) DeleteTranslation(ctx context.Context, in *pb.DeleteTranslationRequest) (*empty.Empty, error) {
	log.Logger.Debug("Deleting an existing translation", zap.Any("Key", &in.Key), zap.Stringp("Language", &in.Language), zap.Stringp("Scope", &in.Scope))
	return c.service.DeleteTranslation(ctx, in)
}
