package service

import (
	"context"

	"github.com/Michael19s/go_first_api/config"
	"github.com/Michael19s/go_first_api/ent"
)

// Structure for handling user-related data manipulation
type userService struct {
	ctx    context.Context
	client *ent.Client
}

// Create a new service
func NewUserService(ctx context.Context) *userService {
	return &userService{
		ctx:    ctx,
		client: config.GetClient(),
	}
}

// Get the information of all users
func (receiverService *userService) UserGetAll() ([]*ent.User, error) {
	users, err := receiverService.client.User.Query().All(receiverService.ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

// Get user information
func (receiverService *userService) UserGetById(id int) (*ent.User, error) {
	user, err := receiverService.client.User.Get(receiverService.ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Create a new user
func (receiverService *userService) UserCreate(newUser ent.User) (*ent.User, error) {
	user, err := receiverService.client.User.Create().
		SetName(newUser.Name).
		SetAge(newUser.Age).
		Save(receiverService.ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Update an existing user
func (receiverService *userService) UserUpdate(updatedUser ent.User) (*ent.User, error) {
	user, err := receiverService.client.User.UpdateOneID(updatedUser.ID).
		SetName(updatedUser.Name).
		SetAge(updatedUser.Age).
		Save(receiverService.ctx)

	if err != nil {
		return nil, err
	}
	return user, nil
}

// Delete a user
func (receiverService *userService) UserDelete(id int) (int, error) {
	err := receiverService.client.User.DeleteOneID(id).Exec(receiverService.ctx)
	if err != nil {
		return 0, err
	}
	return id, nil
}
