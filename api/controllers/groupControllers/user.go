/* =========================================================================
*  File Name: controller/groupController/user.go
*  Description: Controller for getting users within a group.
*  Author: MagnusChase03
*  =======================================================================*/
package groupControllers

import (
	"fmt"

	"github.com/MagnusChase03/CS4389-Project/models"
	"github.com/MagnusChase03/CS4389-Project/utils"
)

/*
*  Attempts to get the users that are within a group.
*
*  Arguments:
*      - userID (int): The ID of the user.
*	   - groupID (int): The ID of the group.
*
*  Returns:
*      - utils.JSONResponse: The response to be made to the client.
*      - error: An error if any occurred.
*
 */
func GetGroupUsersController(userID int, groupID int) (utils.JSONResponse, error) {
	users, err := models.GetGroupUsers(userID, groupID)
	if err != nil {
		return utils.JSONResponse{
			StatusCode: 400,
			Data:       "Failed to get group users.",
		}, fmt.Errorf("[ERROR] Failed to get group users. %w", err)
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
