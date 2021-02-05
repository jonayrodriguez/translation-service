package service

import (
	"context"
	"fmt"

	empty "github.com/golang/protobuf/ptypes/empty"
	pb "github.com/jonayrodriguez/translation-service/api/translation/v1"
	"github.com/jonayrodriguez/translation-service/internal/translation/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
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
	repository repository.TranslationRepository
}

// NewTranslationServiceService creates new translate service
func NewTranslateService(repo repository.TranslationRepository) (*TranslationServiceImp, error) {
	return &TranslationServiceImp{repository: repo}, nil
}

// GetTranslation to retrieve the translated message.
func (ts *TranslationServiceImp) GetTranslation(ctx context.Context, in *pb.GetTranslationRequest) (*pb.GetTranslationResponse, error) {
	t, err := ts.repository.GetTranslation(in.GetLanguage(), in.GetScope(), in.GetKeyPattern())
	if err != nil {
		return nil, err
	}
	if len(t) == 0 {
		return nil, status.Error(codes.NotFound,
			fmt.Sprintf("there is not any translated message for the [key] %s [language]: %s and [scope]: %s ", in.GetKeyPattern(), in.GetLanguage(), in.GetScope()))
	}
	messages := []*pb.KeyMessage{}
	for index := range t {
		messages = append(messages, &pb.KeyMessage{Key: t[index].Key, Message: t[index].Message})
	}
	return &pb.GetTranslationResponse{Language: in.GetLanguage(), Scope: in.GetScope(), Messages: messages}, nil
}

// AddTranslation to add a new translated message.
func (ts *TranslationServiceImp) AddTranslation(ctx context.Context, in *pb.AddTranslationRequest) (*pb.AddTranslationResponse, error) {
	t, err := ts.repository.AddTranslation(in.GetLanguage(), in.GetScope(), in.GetKey(), in.GetMessage())
	if err != nil {
		return nil, err
	}
	return &pb.AddTranslationResponse{Language: t.LanguageName, Scope: t.Scope, Key: t.Key, Message: t.Message}, nil
}

// UpdateTranslation to update an existing translated message.
func (ts *TranslationServiceImp) UpdateTranslation(ctx context.Context, in *pb.UpdateTranslationRequest) (*pb.UpdateTranslationResponse, error) {
	t, err := ts.repository.UpdateTranslation(in.GetLanguage(), in.GetScope(), in.GetKey(), in.GetMessage())
	if err != nil {
		return nil, err
	}
	return &pb.UpdateTranslationResponse{Language: t.LanguageName, Scope: t.Scope, Key: t.Key, Message: t.Message}, nil
}

// DeleteTranslation to delete an existing translated message.
func (ts *TranslationServiceImp) DeleteTranslation(ctx context.Context, in *pb.DeleteTranslationRequest) (*empty.Empty, error) {
	if err := ts.repository.DeleteTranslation(in.GetLanguage(), in.GetScope(), in.GetKey()); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
