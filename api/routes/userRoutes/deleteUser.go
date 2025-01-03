/* =========================================================================
*  File Name: routes/userRoutes/deleteUser.go
*  Description: Handler for deleting users.
*  Author: MagnusChase03
*  =======================================================================*/
package userRoutes

import (
	"fmt"
	"net/http"
	"os"

	"github.com/MagnusChase03/CS4389-Project/controllers/userControllers"
	"github.com/MagnusChase03/CS4389-Project/session"
	"github.com/MagnusChase03/CS4389-Project/utils"
)

/*
*  Handles the control flow for the create user route.
*
*  Arguments:
*      - w (http.ResponseWriter): The object that is used to write a response.
*      - r (*http.Request): The request being made from the client.
*
*  Returns:
*      - N/A
 */
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
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

	resp, err := userControllers.DeleteUserController(userID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] %v\n", err)
	}
	session.DeleteUserCookie(w)

	if err := utils.SendResponse(w, resp); err != nil {
		utils.SendInternalServerError(w, err)
	}
}
