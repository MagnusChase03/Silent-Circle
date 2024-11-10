/* =========================================================================
*  File Name: routes/groupRoutes/listen.go
*  Description: Handler for listening for group invite requests.
*  Author: MagnusChase03
*  =======================================================================*/
package groupRoutes

import (
	"fmt"
	"net/http"
	"os"

	"github.com/MagnusChase03/CS4389-Project/controllers/groupControllers"
	"github.com/MagnusChase03/CS4389-Project/session"
	"github.com/MagnusChase03/CS4389-Project/utils"

	"github.com/gorilla/websocket"
)

/*
*  Handles the control flow for listening for group invite requests via redis pub/sub
*
*  Arguments:
*      - w (http.ResponseWriter): The object that is used to write a response.
*      - r (*http.Request): The request being made from the client.
*
*  Returns:
*      - N/A
 */
func GroupInviteListenerHandler(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{CheckOrigin: utils.WebsocketOriginCheck}
	client, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		utils.SendBadRequest(w)
		return
	}
	defer client.Close()

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

	err = groupControllers.GroupInviteListenerController(client, userID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] %v\n", err)
	}
}
