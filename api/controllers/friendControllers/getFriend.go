/* =========================================================================
*  File Name: controller/friendControllers/getFriend.go
*  Description: Controller for getting a users friends.
*  Author: MagnusChase03
*  =======================================================================*/
package friendControllers

import (
	"fmt"

	"github.com/MagnusChase03/CS4389-Project/models"
	"github.com/MagnusChase03/CS4389-Project/utils"
)

/*
*  Attempts to get friends from given user.
*
*  Arguments:
*      - userID (int): The userID of the sending user.
*
*  Returns:
*      - utils.JSONResponse: The response to be made to the client.
*      - error: An error if any occurred.
*
 */
func GetFriendController(userID int) (utils.JSONResponse, error) {
	friends, err := models.GetFriendUser(userID)
	if err != nil {
		return utils.JSONResponse{
			StatusCode: 400,
			Data:       "Failed to find friends.",
		}, fmt.Errorf("[ERROR] Failed to find friends. %w", err)
	}

	var responseStruct struct {
		Friends []string
	}
	responseStruct.Friends = friends

	return utils.JSONResponse{
		StatusCode: 200,
		Data:       responseStruct,
	}, nil
}
