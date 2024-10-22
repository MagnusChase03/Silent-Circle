/* =========================================================================
*  File Name: controller/friendControllers/removeFriend.go
*  Description: Controller for removing a friend.
*  Author: MagnusChase03
*  =======================================================================*/
package friendControllers

import (
    "fmt"

    "github.com/MagnusChase03/CS4389-Project/models"
    "github.com/MagnusChase03/CS4389-Project/utils"
)

/*
*  Attempts to remove a friend from given user.
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
func RemoveFriendController(userID int, username string) (utils.JSONResponse, error) { 
    err := models.RemoveFriendUser(userID, username);
    if err != nil {
        return utils.JSONResponse{
            StatusCode: 400,
            Data: "Failed to remove friend.",
        }, fmt.Errorf("[ERROR] Failed to remove friend. %w", err);
    }

    return utils.JSONResponse{
        StatusCode: 200,
        Data: "Ok",
    }, nil;
}
