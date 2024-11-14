/* =========================================================================
*  File Name: controller/groupControllers/removeUser.go
*  Description: Controller for removing a user from a group.
*  Author: MagnusChase03
*  =======================================================================*/
package groupControllers

import (
	"fmt"

	"github.com/MagnusChase03/CS4389-Project/models"
	"github.com/MagnusChase03/CS4389-Project/utils"
)

/*
*  Attempts to remove a user from a group.
*
*  Arguments:
*      - userID (int): The userID of the sending user.
*      - groupID (int): The groupID of the group.
*	   - username (string): The user to remove.
*
*  Returns:
*      - utils.JSONResponse: The response to be made to the client.
*      - error: An error if any occurred.
*
 */
func RemoveGroupUserController(userID int, groupID int, username string) (utils.JSONResponse, error) {
	err := models.RemoveGroupUser(userID, groupID, username)
	if err != nil {
		return utils.JSONResponse{
			StatusCode: 400,
			Data:       "Failed to remove user.",
		}, fmt.Errorf("[ERROR] Failed to remove user. %w", err)
	}

	return utils.JSONResponse{
		StatusCode: 200,
		Data:       "Ok",
	}, nil
}
