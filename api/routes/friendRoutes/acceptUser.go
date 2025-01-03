/* =========================================================================
*  File Name: routes/friendRoutes/acceptUser.go
*  Description: Handler for accepting a friend request to a user.
*  Author: MagnusChase03
*  =======================================================================*/
package friendRoutes

import (
	"fmt"
	"net/http"
	"os"

	"github.com/MagnusChase03/CS4389-Project/controllers/friendControllers"
	"github.com/MagnusChase03/CS4389-Project/session"
	"github.com/MagnusChase03/CS4389-Project/utils"
)

/*
*  Handles the control flow for accepting a friend request.
*
*  Arguments:
*      - w (http.ResponseWriter): The object that is used to write a response.
*      - r (*http.Request): The request being made from the client.
*
*  Returns:
*      - N/A
 */
func AcceptFriendRequestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		utils.SendBadRequest(w)
		return
	}

	err := r.ParseForm()
	if err != nil {
		fmt.Printf("[ERROR] Failed to parse form.\n")
		utils.SendBadRequest(w)
		return
	}

	username := r.FormValue("username")
	if username == "" {
		fmt.Printf("[ERROR] username is empty.\n")
		utils.SendBadRequest(w)
		return
	}

	cookie, err := r.Cookie("authCookie")
	if err != nil {
		utils.SendUnauthorizedRequest(w)
		return
	}

	userID, _, err := session.ParseUserCookie(cookie.Value)
	if err != nil {
		utils.SendUnauthorizedRequest(w)
		return
	}

	resp, err := friendControllers.AcceptFriendRequestController(userID, username)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] %v\n", err)
	}

	if err := utils.SendResponse(w, resp); err != nil {
		utils.SendInternalServerError(w, err)
	}
}
