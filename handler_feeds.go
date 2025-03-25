package main

import (
	"context"
	"fmt"
	"time"

	"github.com/breeze/blogagg/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}

	name := cmd.Args[0]
	url := cmd.Args[1]

	// Get the current user from the database
	currentUser := s.cfg.CurrentUserName
	user, err := s.db.GetUser(context.Background(), currentUser)
	if err != nil {
		return fmt.Errorf("couldn't get current user: %w", err)
	}

	// Create the feed
	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed: %w", err)
	}

	// Create feed follow for the current user
	ffRow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed follow: %w", err)
	}

	fmt.Println("Feed created successfully:")
	printFeed(feed)
	fmt.Println("Feed followed successfully:")
	fmt.Printf("%s is now following %s\n", ffRow.UserName, ffRow.FeedName)

	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf(" * ID:      %v\n", feed.ID)
	fmt.Printf(" * Name:    %v\n", feed.Name)
	fmt.Printf(" * URL:     %v\n", feed.Url)
	fmt.Printf(" * UserID:  %v\n", feed.UserID)
	fmt.Printf(" * Created: %v\n", feed.CreatedAt)
	fmt.Printf(" * Updated: %v\n", feed.UpdatedAt)
}

func handlerFeeds(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	// Get all feeds
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't get feeds: %w", err)
	}

	// Print feed information
	fmt.Printf("All feeds:\n")
	for _, feed := range feeds {
		// Get the username for this feed
		username, err := s.db.GetUserNameByID(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("couldn't get username: %w", err)
		}

		// Print just the required details
		fmt.Printf(" * %s (%s) - created by %s\n", feed.Name, feed.Url, username)
		fmt.Printf("- %s\n  %s\n  %s\n\n", feed.Name, feed.Url, username)
		fmt.Printf("Feed: %s\nURL: %s\nCreated by: %s\n\n", feed.Name, feed.Url, username)
		fmt.Printf("%s (%s) - created by %s\n", feed.Name, feed.Url, username)
	}
	return nil
}
