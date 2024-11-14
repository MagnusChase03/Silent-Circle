/* =========================================================================
*  File Name: routes/groupRoutes/removeUser.go
*  Description: Handler for removeing a user from a group.
*  Author: MagnusChase03
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
*  Handles the control flow for removing a user from a group.
*
*  Arguments:
*      - w (http.ResponseWriter): The object that is used to write a response.
*      - r (*http.Request): The request being made from the client.
*
*  Returns:
*      - N/A
 */
func RemoveGroupUserHandler(w http.ResponseWriter, r *http.Request) {
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

	groupStr := r.FormValue("group")
	username := r.FormValue("username")
	if username == "" || groupStr == "" {
		fmt.Printf("[ERROR] Username, or group is empty.\n")
		utils.SendBadRequest(w)
		return
	}

	groupID, err := strconv.ParseInt(groupStr, 10, 64)
	if err != nil {
		fmt.Printf("[ERROR] GroupID is invalid. %v\n", err)
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

	resp, err := groupControllers.RemoveGroupUserController(userID, int(groupID), username)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] %v\n", err)
	}

	if err := utils.SendResponse(w, resp); err != nil {
		utils.SendInternalServerError(w, err)
	}
}
