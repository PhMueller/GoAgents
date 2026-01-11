package services

import (
	"context"

	"examples.com/assistants/db/repository"
)

type AuthService struct {
	queries repository.Querier
}

func NewAuthService(queries repository.Querier) *AuthService {
	return &AuthService{
		queries: queries,
	}
}

func (a *AuthService) GenerateToken(ctx context.Context, clientID string, clientSecret string) (string, error) {
	return "", nil
}

func (a *AuthService) ValidateToken(ctx context.Context, token string) (bool, error) {
	/* Validate the structure of the token */
	return true, nil
}

func (a *AuthService) VerifyToken(ctx context.Context, token string) (bool, error) {
	/* Verify that the user providing the token has access */
	return true, nil
}

func (a *AuthService) CreateSession(ctx context.Context, token string) error {
	/* Create a session for the user associated with the token */
	return nil
}

func (a *AuthService) DeleteSession(ctx context.Context, token string) error {
	/* Delete the session associated with the token */
	return nil
}

func (a *AuthService) CreateClient(ctx context.Context, clientID string, clientSecret string) error {
	/* ADMIN: Create a new client */
	return nil
}

func (a *AuthService) UpdateClient(ctx context.Context, clientID string, clientSecret string) error {
	/* ADMIN: Update an existing client */
	return nil
}

func (a *AuthService) DeleteClient(ctx context.Context, clientID string) error {
	/* ADMIN: Delete an existing client */
	return nil
}
