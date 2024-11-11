/* =========================================================================
*  File Name: controller/friendControllers/acceptUser.go
*  Description: Controller for accepting a friend request.
*  Author: MagnusChase03
*  =======================================================================*/
package friendControllers

import (
	"fmt"

	"github.com/MagnusChase03/CS4389-Project/models"
	"github.com/MagnusChase03/CS4389-Project/utils"
)

/*
*  Attempts to accept a friend request to a given user.
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
func AcceptFriendRequestController(userID int, username string) (utils.JSONResponse, error) {
	err := models.AcceptFriendRequestUser(userID, username)
	if err != nil {
		return utils.JSONResponse{
			StatusCode: 400,
			Data:       "Failed to accept friend request.",
		}, fmt.Errorf("[ERROR] Failed to accpet friend request. %w", err)
	}

	return utils.JSONResponse{
		StatusCode: 200,
		Data:       "Ok",
	}, nil
}
