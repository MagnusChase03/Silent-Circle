/* =========================================================================
*  File Name: routes/groupRoutes/acceptGroup.go
*  Description: Handler for accepting a group invite request.
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
*  Handles the control flow for accepting an invite request.
*
*  Arguments:
*      - w (http.ResponseWriter): The object that is used to write a response.
*      - r (*http.Request): The request being made from the client.
*
*  Returns:
*      - N/A
 */
func AcceptGroupInviteHandler(w http.ResponseWriter, r *http.Request) {
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

	groupIDStr := r.FormValue("group")
	if groupIDStr == "" {
		fmt.Printf("[ERROR] Group is empty.\n")
		utils.SendBadRequest(w)
		return
	}

	groupID, err := strconv.ParseInt(groupIDStr, 10, 64)
	if groupIDStr == "" {
		fmt.Printf("[ERROR] Group is invalid.\n")
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

	resp, err := groupControllers.AcceptGroupInviteController(userID, int(groupID))
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] %v\n", err)
	}

	if err := utils.SendResponse(w, resp); err != nil {
		utils.SendInternalServerError(w, err)
	}
}
