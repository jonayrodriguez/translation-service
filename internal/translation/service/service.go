package service

import (
	"context"

	empty "github.com/golang/protobuf/ptypes/empty"
	pb "github.com/jonayrodriguez/translation-service/api/translation/v1"
)

// TranslationService is the interface for the Email Adaptor
type TranslationService interface {
	// GetTranslation to retrieve the translated message.
	GetTranslation(ctx context.Context, in *pb.GetTranslationRequest) (*pb.GetTranslationResponse, error)
	// AddTranslation to add a new translated message.
	AddTranslation(ctx context.Context, in *pb.AddTranslationRequest) (*pb.AddTranslationResponse, error)
	// UpdateTranslation to update an existing translated message.
	UpdateTranslation(ctx context.Context, in *pb.UpdateTranslationRequest) (*pb.UpdateTranslationResponse, error)
	// DeleteTranslation to delete an existing translated message.
	DeleteTranslation(ctx context.Context, in *pb.DeleteTranslationRequest) (*empty.Empty, error)
}

// TranslationServiceImp struct used for the Translate Service implementation
type TranslationServiceImp struct {
	// TODO - This will be replaced once we know what client is used to connect to the email server.
	client interface{}
}

// NewTranslationServiceService creates new translate service
func NewTranslateService(client interface{}) (*TranslationServiceImp, error) {
	return &TranslationServiceImp{client: client}, nil
}

// GetTranslation to retrieve the translated message.
func (ts *TranslationServiceImp) GetTranslation(ctx context.Context, in *pb.GetTranslationRequest) (*pb.GetTranslationResponse, error) {
	return nil, nil
}

// AddTranslation to add a new translated message.
func (ts *TranslationServiceImp) AddTranslation(ctx context.Context, in *pb.AddTranslationRequest) (*pb.AddTranslationResponse, error) {
	return nil, nil
}

// UpdateTranslation to update an existing translated message.
func (ts *TranslationServiceImp) UpdateTranslation(ctx context.Context, in *pb.UpdateTranslationRequest) (*pb.UpdateTranslationResponse, error) {
	return nil, nil
}

// DeleteTranslation to delete an existing translated message.
func (ts *TranslationServiceImp) DeleteTranslation(ctx context.Context, in *pb.DeleteTranslationRequest) (*empty.Empty, error) {
	return nil, nil
}
