/* =========================================================================
*  File Name: controller/friendControllers/rejectUser.go
*  Description: Controller for rejecting a friend request.
*  Author: MagnusChase03
*  =======================================================================*/
package friendControllers

import (
    "fmt"

    "github.com/MagnusChase03/CS4389-Project/models"
    "github.com/MagnusChase03/CS4389-Project/utils"
)

/*
*  Attempts to reject a friend request to a given user.
*
*  Arguments:
*      - userID (int): The userID of the sending user.
*      - username (string): The username of the user receiving user.
* 
*  Returns:
*      - utils.JSONResponse: The response to be made to the client.
*      - error: An error if any occurred.
*
*/
func RejectFriendRequestController(userID int, username string) (utils.JSONResponse, error) { 
    err := models.RejectFriendRequestUser(userID, username);
    if err != nil {
        return utils.JSONResponse{
            StatusCode: 400,
            Data: "Failed to reject friend request.",
        }, fmt.Errorf("[ERROR] Failed to reject friend request. %w", err);
    }

    return utils.JSONResponse{
        StatusCode: 200,
        Data: "Ok",
    }, nil;
}
