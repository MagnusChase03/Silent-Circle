/* =========================================================================
*  File Name: controller/friendControllers/inviteUser.go
*  Description: Controller for sending a friend request.
*  Author: MagnusChase03
*  =======================================================================*/
package friendControllers

import (
    "fmt"

    "github.com/MagnusChase03/CS4389-Project/db"
    "github.com/MagnusChase03/CS4389-Project/models"
    "github.com/MagnusChase03/CS4389-Project/utils"
)

/*
*  Attempts to send a friend request to a given user.
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
func FriendRequestController(userID int, username string) (utils.JSONResponse, error) { 
    db, err := db.GetRedisDB();
    if err != nil {
        return utils.JSONResponse{
            StatusCode: 400,
            Data: "Failed to send friend request.",
        }, fmt.Errorf("[ERROR] Failed to get redis db. %w", err);
    }

    err = models.FriendRequestUser(userID, username);
    if err != nil {
        return utils.JSONResponse{
            StatusCode: 400,
            Data: "Failed to send friend request.",
        }, fmt.Errorf("[ERROR] Failed to send friend request. %w", err);
    }

    user, err := models.GetUserByUsername(username);
    if err != nil {
        return utils.JSONResponse{
            StatusCode: 400,
            Data: "Failed to send friend request.",
        }, fmt.Errorf("[ERROR] Failed to get user. %w", err);
    }

    sendingUser, err := models.GetUserByID(userID);
    if err != nil {
        return utils.JSONResponse{
            StatusCode: 400,
            Data: "Failed to send friend request.",
        }, fmt.Errorf("[ERROR] Failed to get user. %w", err);
    }

    err = db.Connection.Publish(db.Ctx, fmt.Sprintf("fr-%d", user.UserID), sendingUser.Username).Err();
    if err != nil {
        return utils.JSONResponse{
            StatusCode: 400,
            Data: "Failed to send friend request.",
        }, fmt.Errorf("[ERROR] Failed to publish friend request. %w", err);
    }

    return utils.JSONResponse{
        StatusCode: 200,
        Data: "Ok",
    }, nil;
}
