/* =========================================================================
*  File Name: controller/userController/deleteUser.go
*  Description: Controller for deleting a user.
*  Author: Matthew-Basinger
*  =======================================================================*/
package groupControllers

import (
	"fmt"

	"github.com/MagnusChase03/CS4389-Project/models"
	"github.com/MagnusChase03/CS4389-Project/utils"
)

/*
*  Attempts to create a new user with given attributes.
*
*  Arguments:
*      - groupname (string): The groupname.
*	   - userID (int): The ID of the user
*
*  Returns:
*      - utils.JSONResponse: The response to be made to the client.
*      - error: An error if any occurred.
*
 */
func DeleteGroupController(groupID int, userID int) (utils.JSONResponse, error) {
	err := models.DeleteGroup(userID, groupID)
	if err != nil {
		return utils.JSONResponse{
			StatusCode: 401,
			Data:       "Failed to delete group.",
		}, fmt.Errorf("[ERROR] Failed to delete group. %w", err)
	}

	return utils.JSONResponse{
		StatusCode: 200,
		Data:       "Ok",
	}, nil
}
