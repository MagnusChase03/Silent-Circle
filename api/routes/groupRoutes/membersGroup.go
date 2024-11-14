/* =========================================================================
*  File Name: routes/groupRoutes/membersGroup.go
*  Description: Handler for listing all members in a group.
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
*  Handles the control flow for the retrieval of group members routes.
*
*  Arguments:
*      - w (http.ResponseWriter): The object that is used to write a response.
*      - r (*http.Request): The request being made from the client.
*
*  Returns:
*      - N/A
 */
func GroupMembersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		utils.SendBadRequest(w)
		return
	}

	cookie, err := r.Cookie("authCookie")
	if err != nil {
		utils.SendUnauthorizedRequest(w)
		return
	}

	_, _, err = session.ParseUserCookie(cookie.Value)
	if err != nil {
		utils.SendUnauthorizedRequest(w)
		return
	}

	err = r.ParseForm()
	if err != nil {
		fmt.Printf("[ERROR] Failed to parse form.\n")
		utils.SendBadRequest(w)
		return
	}

	groupID, err := strconv.Atoi(r.FormValue("groupID"))
	if groupID < 1 {
		fmt.Printf("[ERROR] groupID invalid.\n")
		utils.SendBadRequest(w)
		return
	}


	resp, err := groupControllers.MembersGroupController(groupID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] %v\n", err)
	}

	if err := utils.SendResponse(w, resp); err != nil {
		utils.SendInternalServerError(w, err)
	}
}
