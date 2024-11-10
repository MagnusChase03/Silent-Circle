/* =========================================================================
*  File Name: controller/groupControllers/rejectGroup.go
*  Description: Controller for rejecting a friend request.
*  Author: MagnusChase03
*  =======================================================================*/
package groupControllers

import (
	"fmt"

	"github.com/MagnusChase03/CS4389-Project/models"
	"github.com/MagnusChase03/CS4389-Project/utils"
)

/*
*  Attempts to reject an invite request.
*
*  Arguments:
*      - userID (int): The userID of the sending user.
*      - groupID (int): The groupID of the group.
*
*  Returns:
*      - utils.JSONResponse: The response to be made to the client.
*      - error: An error if any occurred.
*
 */
func RejectGroupInviteController(userID int, groupID int) (utils.JSONResponse, error) {
	err := models.RejectGroupInvite(userID, groupID)
	if err != nil {
		return utils.JSONResponse{
			StatusCode: 400,
			Data:       "Failed to reject invite request.",
		}, fmt.Errorf("[ERROR] Failed to reject invite request. %w", err)
	}

	return utils.JSONResponse{
		StatusCode: 200,
		Data:       "Ok",
	}, nil
}
