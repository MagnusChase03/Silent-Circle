/* =========================================================================
*  File Name: controller/groupController/createGroup.go
*  Description: Controller for creating a group.
*  Author: Matthew-Basinger
*  =======================================================================*/
package groupControllers

import (
	"fmt"

	"github.com/MagnusChase03/CS4389-Project/models"
	"github.com/MagnusChase03/CS4389-Project/utils"
)

/*
*  Attempts to create a new group with given attributes.
*
*  Arguments:
*      - creatorID (int): The ID of the creator of the group.
*      - groupName (string): The name of the group
*
*  Returns:
*      - utils.JSONResponse: The response to be made to the client.
*      - error: An error if any occurred.
*
 */
func CreateGroupController(creatorID int, groupname string) (utils.JSONResponse, error) {
	groupID, err := models.CreateGroup(creatorID, groupname)
	if err != nil {
		return utils.JSONResponse{
			StatusCode: 401,
			Data:       "Failed to create group.",
		}, fmt.Errorf("[ERROR] Failed to create group. %w", err)
	}

	var responseStruct struct {
		GroupID int
	}
	responseStruct.GroupID = groupID

	return utils.JSONResponse{
		StatusCode: 200,
		Data:       responseStruct,
	}, nil
}
