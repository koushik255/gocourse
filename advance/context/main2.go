package main

import (
	"context"
	"fmt"
	"net/http"
)

// key is a custom type to avoid key collisions in context
type key string

const userIDKey key = "userID"

// getUserInfo simulates retrieving user information from a database.
func getUserInfo(ctx context.Context) (string, error) {
	// Retrieve the user ID from the context.
	// This allows us to access request-specific data without using global variables.
	userID := ctx.Value(userIDKey).(string)

	// Simulate a database query to get user information.
	// In a real application, this would involve querying a database.
	return fmt.Sprintf("User Info for ID: %s", userID), nil
}

// processRequest handles incoming HTTP requests.
func processRequest(w http.ResponseWriter, r *http.Request) {
	// Create a context with a user ID.
	// This context is derived from the incoming request's context and includes a value.
	ctx := context.WithValue(r.Context(), userIDKey, "user123")

	// Call the function to get user information, passing the context.
	// This allows the getUserInfo function to access the user ID stored in the context.
	userInfo, err := getUserInfo(ctx)
	if err != nil {
		// If there is an error retrieving user info, respond with an error message.
		http.Error(w, "Error retrieving user info", http.StatusInternalServerError)
		return
	}

	// Respond to the client with the user information retrieved from the context.
	fmt.Fprintln(w, userInfo)
}

func main() {
	// Set up the HTTP server to handle requests.
	// The processRequest function will be called for each incoming request.
	http.HandleFunc("/", processRequest)
	fmt.Println("Server is running on :8080")
	// Start the server and listen for incoming requests.
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server error:", err)
	}
}
