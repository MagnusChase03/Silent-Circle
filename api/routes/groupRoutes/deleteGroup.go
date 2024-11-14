/* =========================================================================
*  File Name: routes/userRoutes/deleteUser.go
*  Description: Handler for deleting users.
*  Author: Matthew-Basinger
*  =======================================================================*/
package groupRoutes

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/MagnusChase03/CS4389-Project/controllers/groupControllers"
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
func DeleteGroupHandler(w http.ResponseWriter, r *http.Request) {
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

	groupStr := r.FormValue("group")
	if groupStr == "" {
		fmt.Printf("[ERROR] groupname empty.\n")
		utils.SendBadRequest(w)
		return
	}

	groupID, err := strconv.ParseInt(groupStr, 10, 64)
	if groupStr == "" {
		fmt.Printf("[ERROR] Invalid group.\n")
		utils.SendBadRequest(w)
		return
	}

	resp, err := groupControllers.DeleteGroupController(int(groupID), userID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] %v\n", err)
	}

	if err := utils.SendResponse(w, resp); err != nil {
		utils.SendInternalServerError(w, err)
	}
}
