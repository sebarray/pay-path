package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"paypath/config"
	"paypath/internal/helpers"
	"paypath/internal/user/model"

	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

type User interface {
	GetUser(ctx context.Context, userID string) (*model.User, error)
	CreateUser(ctx context.Context, user *model.User) (string, error)
	ConfirmUser(ctx context.Context, user *model.User) error
	Login(ctx context.Context, user *model.User) (string, error)
}

type repository struct {
	redisClient *redis.Client
	db          *sql.DB
}

func NewRepository() (User, error) {
	rdb, err := config.ConnectRedis()
	if err != nil {
		log.Println("error connecting to redis: ", err.Error())
		return nil, err
	}

	db, err := config.ConnectDB()
	if err != nil {
		log.Println("error connecting to redis: ", err.Error())
		return nil, err
	}
	return &repository{redisClient: rdb, db: db}, nil
}
func (r *repository) GetUser(ctx context.Context, userID string) (*model.User, error) {
	userData, err := r.redisClient.Get(ctx, userID).Result()
	if err == redis.Nil {
		return nil, fmt.Errorf("user not found")
	} else if err != nil {
		return nil, err
	}

	var user model.User
	err = json.Unmarshal([]byte(userData), &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *repository) CreateUser(ctx context.Context, user *model.User) (string, error) {
	userData, err := json.Marshal(user)
	if err != nil {
		return "", err
	}
	code, err := helpers.RandomAlphanumeric(6)
	if err != nil {
		return "", err
	}
	err = r.redisClient.Set(ctx, code, userData, 0).Err()
	if err != nil {
		return "", err
	}

	return code, nil
}
func (r *repository) ConfirmUser(ctx context.Context, user *model.User) error {
	query := "INSERT INTO users (id, password, email) VALUES (?, ?, ?)"
	_, err := r.db.ExecContext(ctx, query, user.ID, user.Password, user.Email)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) Login(ctx context.Context, user *model.User) (string, error) {
	query := "SELECT  password FROM users WHERE email=?"
	row := r.db.QueryRowContext(ctx, query, user.Email)
	var hashedPassword string

	err := row.Scan(&hashedPassword)
	if err == sql.ErrNoRows {
		return "", fmt.Errorf("user not found")
	} else if err != nil {
		return "", fmt.Errorf("error querying user: %v", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(user.Password))
	if err != nil {
		return "", fmt.Errorf("invalid password or email")
	}
	user.Password = ""

	return user.ID, nil
}
