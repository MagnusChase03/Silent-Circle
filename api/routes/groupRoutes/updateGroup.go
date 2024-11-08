/* =========================================================================
*  File Name: routes/groupRoutes/updateGroup.go
*  Description: Handler for updating user information.
*  Author: Matthew-Basinger
*  =======================================================================*/
package groupRoutes

import (
    "os"
    "fmt"
    "net/http"
    "strconv"

    "github.com/MagnusChase03/CS4389-Project/utils"
    "github.com/MagnusChase03/CS4389-Project/session"
    "github.com/MagnusChase03/CS4389-Project/controllers/groupControllers"
)

/*
*  Handles the control flow for the update group route.
*
*  Arguments:
*      - w (http.ResponseWriter): The object that is used to write a response.
*      - r (*http.Request): The request being made from the client.
* 
*  Returns:
*      - N/A
*/
func UpdateGroupHandler(w http.ResponseWriter, r *http.Request) { 
    if r.Method != "POST" {
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

    err = r.ParseForm();
    if err != nil {
        fmt.Printf("[ERROR] Failed to parse form.\n");
        utils.SendBadRequest(w);
        return;
    }
	//
	groupname := r.FormValue("groupname");
    groupID,err := strconv.Atoi(r.FormValue("groupID"));
    if groupname == ""{
        fmt.Printf("[ERROR] groupname empty.\n");
        utils.SendBadRequest(w);
        return;
    }

    resp, err := groupControllers.UpdateGroupController(userID, groupname, groupID);
    if err != nil {
        fmt.Fprintf(os.Stderr, "[ERROR] %v\n", err);
    }

    if err := utils.SendResponse(w, resp); err != nil {
        utils.SendInternalServerError(w, err);
    }
}
