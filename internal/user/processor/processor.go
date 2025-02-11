package processor

import (
	"context"
	"paypath/internal/helpers"
	"paypath/internal/user/model"
	"paypath/internal/user/repository"

	"github.com/google/uuid"
)

type User interface {
	CreateUser(user *model.User) error
	ConfirmUser(userID string) (string, error)
	Login(ctx context.Context, user *model.User) (string, error)
}
type processor struct {
	repo repository.User
}

func NewProcessor(repo repository.User) User {
	return &processor{repo: repo}
}

func (p *processor) CreateUser(user *model.User) error {

	ctx := context.Background()
	defer ctx.Done()
	user.Password, _ = helpers.HashPassword(user.Password)
	user.ID = uuid.New().String()
	code, err := p.repo.CreateUser(ctx, user)
	if err != nil {
		return err
	}
	err = helpers.ConfirmUserEmail(user.Email, code)
	if err != nil {
		return err
	}

	return nil
}
func (p *processor) ConfirmUser(userID string) (string, error) {

	ctx := context.Background()
	defer ctx.Done()

	user, err := p.repo.GetUser(ctx, userID)
	if err != nil {
		return "", err
	}
	err = p.repo.ConfirmUser(ctx, user)
	if err != nil {
		return "", err
	}

	tokenString, err := helpers.GenerateJWT(user)

	return tokenString, nil

}

func (p *processor) Login(ctx context.Context, user *model.User) (string, error) {

	ctx = context.Background()
	defer ctx.Done()

	userID, err := p.repo.Login(ctx, user)
	if err != nil {
		return "", err
	}
	user.ID = userID

	tokenString, err := helpers.GenerateJWT(user)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
