/* =========================================================================
*  File Name: controller/userController/group.go
*  Description: Controller for getting user groups.
*  Author: MagnusChase03
*  =======================================================================*/
package userControllers

import (
	"fmt"

	"github.com/MagnusChase03/CS4389-Project/models"
	"github.com/MagnusChase03/CS4389-Project/utils"
)

/*
*  Attempts to get the groups a user is in.
*
*  Arguments:
*      - N/A
*
*  Returns:
*      - utils.JSONResponse: The response to be made to the client.
*      - error: An error if any occurred.
*
 */
func GetUserGroupsController(userID int) (utils.JSONResponse, error) {
	groups, err := models.GetGroups(userID)
	if err != nil {
		return utils.JSONResponse{
			StatusCode: 400,
			Data:       "Failed to get user groups.",
		}, fmt.Errorf("[ERROR] Failed to get user groups. %w", err)
	}

	var responseStruct struct {
		Groups []models.Group
	}
	responseStruct.Groups = groups

	return utils.JSONResponse{
		StatusCode: 200,
		Data:       responseStruct,
	}, nil
}
