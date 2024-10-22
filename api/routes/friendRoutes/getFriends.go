/* =========================================================================
*  File Name: routes/friendRoutes/getFriends.go
*  Description: Handler for getting a users friends.
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
)

/*
*  Handles the control flow for getting a users friends.
*
*  Arguments:
*      - w (http.ResponseWriter): The object that is used to write a response.
*      - r (*http.Request): The request being made from the client.
* 
*  Returns:
*      - N/A
*/
func GetFriendHandler(w http.ResponseWriter, r *http.Request) { 
    if r.Method != "GET" {
        utils.SendBadRequest(w);
        return;
    }

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

    resp, err := friendControllers.GetFriendController(userID);
    if err != nil {
        fmt.Fprintf(os.Stderr, "[ERROR] %v\n", err);
    }

    if err := utils.SendResponse(w, resp); err != nil {
        utils.SendInternalServerError(w, err);
    }
}
