/* =========================================================================
*  File Name: controller/friendControllers/getInvite.go
*  Description: Controller for getting a users friend invites.
*  Author: MagnusChase03
*  =======================================================================*/
package friendControllers

import (
	"fmt"

	"github.com/MagnusChase03/CS4389-Project/models"
	"github.com/MagnusChase03/CS4389-Project/utils"
)

/*
*  Attempts to get friend invites from given user.
*
*  Arguments:
*      - userID (int): The userID of the sending user.
*
*  Returns:
*      - utils.JSONResponse: The response to be made to the client.
*      - error: An error if any occurred.
*
 */
func GetFriendInvitesController(userID int) (utils.JSONResponse, error) {
	users, err := models.GetFriendInvites(userID)
	if err != nil {
		return utils.JSONResponse{
			StatusCode: 400,
			Data:       "Failed to find friend invites.",
		}, fmt.Errorf("[ERROR] Failed to find friend invites. %w", err)
	}

	var responseStruct struct {
		Users []string
	}
	responseStruct.Users = users

	return utils.JSONResponse{
		StatusCode: 200,
		Data:       responseStruct,
	}, nil
}
