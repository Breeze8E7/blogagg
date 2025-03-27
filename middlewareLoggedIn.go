package main

import (
	"context"
	"fmt"

	"github.com/breeze/blogagg/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		currentUser := s.cfg.CurrentUserName
		if currentUser == "" {
			return fmt.Errorf("no user currently logged in")
		}
		user, err := s.db.GetUser(context.Background(), currentUser)
		if err != nil {
			return fmt.Errorf("couldn't get current user: %w", err)
		}
		return handler(s, cmd, user)
	}
}
