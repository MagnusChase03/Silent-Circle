/* =========================================================================
*  File Name: main.go
*  Description: Starts the HTTP API and hands off requests.
*  Author: MagnusChase03, Matthew-Basinger
*  =======================================================================*/
package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/MagnusChase03/CS4389-Project/db"
	"github.com/MagnusChase03/CS4389-Project/middleware"
	"github.com/MagnusChase03/CS4389-Project/routes"
	"github.com/MagnusChase03/CS4389-Project/routes/authRoutes"
	"github.com/MagnusChase03/CS4389-Project/routes/friendRoutes"
	"github.com/MagnusChase03/CS4389-Project/routes/groupRoutes"
	"github.com/MagnusChase03/CS4389-Project/routes/userRoutes"
)

func main() {
	// Connect to databases
	_, err := db.GetRedisDB()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] Failed to connect to Redis DB. %v\n", err)
		return
	}

	_, err = db.GetMariaDB()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] Failed to connect to MariaDB. %v\n", err)
		return
	}

	mux := http.NewServeMux()

	// Assign routes
	mux.Handle("/healthcheck", middleware.HandleWithMiddleware(
		http.HandlerFunc(routes.HealthcheckHandler),
		middleware.CorsMiddleware,
		middleware.LogMiddleware,
	))

	mux.Handle("/login", middleware.HandleWithMiddleware(
		http.HandlerFunc(authRoutes.LoginHandler),
		middleware.CorsMiddleware,
		middleware.LogMiddleware,
	))

	mux.Handle("/logout", middleware.HandleWithMiddleware(
		http.HandlerFunc(authRoutes.LogoutHandler),
		middleware.AuthMiddleware,
		middleware.CorsMiddleware,
		middleware.LogMiddleware,
	))

	mux.Handle("/user/get", middleware.HandleWithMiddleware(
		http.HandlerFunc(userRoutes.GetUserHandler),
		middleware.CorsMiddleware,
		middleware.LogMiddleware,
	))

	mux.Handle("/user/update", middleware.HandleWithMiddleware(
		http.HandlerFunc(userRoutes.UpdateUserHandler),
		middleware.AuthMiddleware,
		middleware.CorsMiddleware,
		middleware.LogMiddleware,
	))

	mux.Handle("/user/create", middleware.HandleWithMiddleware(
		http.HandlerFunc(userRoutes.CreateUserHandler),
		middleware.CorsMiddleware,
		middleware.LogMiddleware,
	))

	mux.Handle("/user/delete", middleware.HandleWithMiddleware(
		http.HandlerFunc(userRoutes.DeleteUserHandler),
		middleware.AuthMiddleware,
		middleware.CorsMiddleware,
		middleware.LogMiddleware,
	))

	mux.Handle("/user/friend/invite", middleware.HandleWithMiddleware(
		http.HandlerFunc(friendRoutes.FriendRequestHandler),
		middleware.AuthMiddleware,
		middleware.CorsMiddleware,
		middleware.LogMiddleware,
	))

	mux.Handle("/user/friend/accept", middleware.HandleWithMiddleware(
		http.HandlerFunc(friendRoutes.AcceptFriendRequestHandler),
		middleware.AuthMiddleware,
		middleware.CorsMiddleware,
		middleware.LogMiddleware,
	))

	mux.Handle("/user/friend/reject", middleware.HandleWithMiddleware(
		http.HandlerFunc(friendRoutes.RejectFriendRequestHandler),
		middleware.AuthMiddleware,
		middleware.CorsMiddleware,
		middleware.LogMiddleware,
	))

	mux.Handle("/user/friend/remove", middleware.HandleWithMiddleware(
		http.HandlerFunc(friendRoutes.RemoveFriendHandler),
		middleware.AuthMiddleware,
		middleware.CorsMiddleware,
		middleware.LogMiddleware,
	))

	mux.Handle("/user/friend/get", middleware.HandleWithMiddleware(
		http.HandlerFunc(friendRoutes.GetFriendHandler),
		middleware.AuthMiddleware,
		middleware.CorsMiddleware,
		middleware.LogMiddleware,
	))

	mux.Handle("/user/friend/listen", middleware.HandleWithMiddleware(
		http.HandlerFunc(friendRoutes.FriendRequestListenerHandler),
		middleware.AuthMiddleware,
		middleware.CorsMiddleware,
		middleware.LogMiddleware,
	))

	mux.Handle("/user/groups", middleware.HandleWithMiddleware(
		http.HandlerFunc(userRoutes.GetUserGroupsHandler),
		middleware.AuthMiddleware,
		middleware.CorsMiddleware,
		middleware.LogMiddleware,
	))

	mux.Handle("/user/friend/invite/get", middleware.HandleWithMiddleware(
		http.HandlerFunc(friendRoutes.GetFriendInvitesHandler),
		middleware.AuthMiddleware,
		middleware.CorsMiddleware,
		middleware.LogMiddleware,
	))

	mux.Handle("/group/users", middleware.HandleWithMiddleware(
		http.HandlerFunc(groupRoutes.GetGroupUsersHandler),
		middleware.AuthMiddleware,
		middleware.CorsMiddleware,
		middleware.LogMiddleware,
	))

	mux.Handle("/group/create", middleware.HandleWithMiddleware(
		http.HandlerFunc(groupRoutes.CreateGroupHandler),
		middleware.AuthMiddleware,
		middleware.CorsMiddleware,
		middleware.LogMiddleware,
	))

	mux.Handle("/group/delete", middleware.HandleWithMiddleware(
		http.HandlerFunc(groupRoutes.DeleteGroupHandler),
		middleware.AuthMiddleware,
		middleware.CorsMiddleware,
		middleware.LogMiddleware,
	))

	mux.Handle("/group/update", middleware.HandleWithMiddleware(
		http.HandlerFunc(groupRoutes.UpdateGroupHandler),
		middleware.AuthMiddleware,
		middleware.CorsMiddleware,
		middleware.LogMiddleware,
	))

	mux.Handle("/group/invite", middleware.HandleWithMiddleware(
		http.HandlerFunc(groupRoutes.GroupInviteHandler),
		middleware.AuthMiddleware,
		middleware.CorsMiddleware,
		middleware.LogMiddleware,
	))

	mux.Handle("/group/invite/accept", middleware.HandleWithMiddleware(
		http.HandlerFunc(groupRoutes.AcceptGroupInviteHandler),
		middleware.AuthMiddleware,
		middleware.CorsMiddleware,
		middleware.LogMiddleware,
	))

	mux.Handle("/group/invite/reject", middleware.HandleWithMiddleware(
		http.HandlerFunc(groupRoutes.RejectGroupInviteHandler),
		middleware.AuthMiddleware,
		middleware.CorsMiddleware,
		middleware.LogMiddleware,
	))

	mux.Handle("/group/invite/listen", middleware.HandleWithMiddleware(
		http.HandlerFunc(groupRoutes.GroupInviteListenerHandler),
		middleware.AuthMiddleware,
		middleware.CorsMiddleware,
		middleware.LogMiddleware,
	))

	mux.Handle("/group/chat", middleware.HandleWithMiddleware(
		http.HandlerFunc(groupRoutes.ChatHandler),
		middleware.AuthMiddleware,
		middleware.CorsMiddleware,
		middleware.LogMiddleware,
	))

	mux.Handle("/group/messages", middleware.HandleWithMiddleware(
		http.HandlerFunc(groupRoutes.GetMessageHandler),
		middleware.AuthMiddleware,
		middleware.CorsMiddleware,
		middleware.LogMiddleware,
	))

	mux.Handle("/group/ban", middleware.HandleWithMiddleware(
		http.HandlerFunc(groupRoutes.RemoveGroupUserHandler),
		middleware.AuthMiddleware,
		middleware.CorsMiddleware,
		middleware.LogMiddleware,
	))

	mux.Handle("/group/invite/get", middleware.HandleWithMiddleware(
		http.HandlerFunc(groupRoutes.GetGroupInvitesHandler),
		middleware.AuthMiddleware,
		middleware.CorsMiddleware,
		middleware.LogMiddleware,
	))

	// Start HTTPS server on port 8080
	fmt.Printf("[LOG] Starting API server on 0.0.0.0:8080.\n")
	if err := http.ListenAndServeTLS("0.0.0.0:8080", "/certs/server.crt", "/certs/server.key", mux); err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] Failed to start API server. %v\n", err)
	}
}
