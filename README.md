# Blog Aggregator Project

## Commands

### Database Queries
- `sqlc generate` - Generate Go code from SQL queries

### Application Commands
- `go run . reset` - Reset the application
- `go run . register <username>` - Register a new user
- `go run . login <username>` - Login as an existing user
- `go run . users` - List all users with current user marked

## Development Notes
- Remember to run `sqlc generate` after adding or modifying SQL queries
- Experienced some sort of bug between sqlc and go which creates a memory overflow and panic. Used docker as a workaround. Use the following instead of sqlc generate:
            sudo docker run --rm -v $(pwd):/src -w /src kjconroy/sqlc generate
- The current user is stored in the `CurrentUserName` field of the `Config` struct