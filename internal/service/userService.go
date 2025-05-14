package service

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/VolkHackVH/todo-list/internal/db"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

type UserService struct {
	Db *db.Queries
}

type UserPublic struct {
	ID       int32  `json:"id"`
	Username string `json:"username"`
}

func NewUserService(db *db.Queries) *UserService {
	return &UserService{
		Db: db,
	}
}

func (s *UserService) CreateUser(ctx context.Context, username string, password string) (db.User, error) {
	if len(username) <= 3 {
		return db.User{}, errors.New("username cannot be less than 4 characters long")
	}

	if len(password) <= 8 {
		return db.User{}, errors.New("password cannot be less than 9 characters long")
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return db.User{}, errors.New("failed to hash password")
	}

	userRow, err := s.Db.CreateUser(ctx, db.CreateUserParams{
		Username: username,
		Password: string(hashPassword),
	})
	if err != nil {
		return db.User{}, err
	}

	user := db.User{
		ID:       userRow.ID,
		Username: userRow.Username,
	}

	return user, nil
}

func (s *UserService) GetUserByID(ctx context.Context, id int) (UserPublic, error) {
	user, err := s.Db.GetUserByID(ctx, int32(id))
	if err != nil {
		return UserPublic{}, err
	}

	publicUser := UserPublic{
		ID:       user.ID,
		Username: user.Username,
	}

	return publicUser, nil
}

func (s *UserService) ListUsers(ctx context.Context) ([]UserPublic, error) {
	users, err := s.Db.ListUsers(ctx)
	if err != nil {
		return nil, err
	}

	publicUsers := make([]UserPublic, 0, len(users))
	for _, u := range users {
		publicUsers = append(publicUsers, UserPublic{
			ID:       u.ID,
			Username: u.Username,
		})
	}

	return publicUsers, nil
}

func (s *UserService) DeleteUser(ctx context.Context, id int) error {
	if err := s.Db.DeleteUser(ctx, int32(id)); err != nil {
		return err
	}

	return nil
}

func (s *UserService) LoginUser(ctx context.Context, username string, password string) (string, error) {
	user, err := s.Db.GetUserByUsername(ctx, username)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid password")
	}

	claims := jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", errors.New("failed to token signed")
	}

	return signedToken, nil
}
