/* =========================================================================
*  File Name: routes/groupRoutes/chat.go
*  Description: Handler for real time chatting.
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

	"github.com/gorilla/websocket"
)

/*
*  Handles the control flow for real time chat via redis pub/sub
*
*  Arguments:
*      - w (http.ResponseWriter): The object that is used to write a response.
*      - r (*http.Request): The request being made from the client.
*
*  Returns:
*      - N/A
 */
func ChatHandler(w http.ResponseWriter, r *http.Request) {
	groupIDStr := r.URL.Query().Get("group")
	if groupIDStr == "" {
		utils.SendBadRequest(w)
		return
	}

	groupID, err := strconv.ParseInt(groupIDStr, 10, 64)
	if err != nil {
		utils.SendBadRequest(w)
		return
	}

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

	err = groupControllers.ChatController(client, userID, int(groupID))
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] %v\n", err)
	}
}
