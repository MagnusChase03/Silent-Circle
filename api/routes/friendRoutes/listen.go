/* =========================================================================
*  File Name: routes/friendRoutes/listen.go
*  Description: Handler for listening for friend requests.
*  Author: MagnusChase03
*  =======================================================================*/
package friendRoutes

import (
    "os"
    "fmt"
    "net/http"

    "github.com/MagnusChase03/CS4389-Project/session"
    "github.com/MagnusChase03/CS4389-Project/utils"
    "github.com/MagnusChase03/CS4389-Project/controllers/friendControllers"

    "github.com/gorilla/websocket"
)

/*
*  Handles the control flow for listening for friend requests via redis pub/sub
*
*  Arguments:
*      - w (http.ResponseWriter): The object that is used to write a response.
*      - r (*http.Request): The request being made from the client.
* 
*  Returns:
*      - N/A
*/
func FriendRequestListenerHandler(w http.ResponseWriter, r *http.Request) { 
    upgrader := websocket.Upgrader{CheckOrigin: utils.WebsocketOriginCheck}
    client, err := upgrader.Upgrade(w, r, nil);
    if err != nil {
        utils.SendBadRequest(w);
        return;
    }
    defer client.Close();

    cookie, err := r.Cookie("authCookie");
    if err != nil {
        utils.SendUnauthorizedRequest(w);
        return;
    }

    userID, _, err := session.ParseUserCookie(cookie.Value);
    if err != nil {
        utils.SendUnauthorizedRequest(w);
        return;
    }

    err = friendControllers.FriendRequestListenerController(client, userID);
    if err != nil {
        fmt.Fprintf(os.Stderr, "[ERROR] %v\n", err);
    }
}
