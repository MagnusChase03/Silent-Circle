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
*  Attempts to retrieve all groups a user is part of
*
*  Arguments:
*      - creatorID (int): The ID of the creator of the group.
*
*  Returns:
*      - utils.JSONResponse: The response to be made to the client.
*      - error: An error if any occurred.
*
*/
func GetGroupController(creatorID int) (utils.JSONResponse, error) {
	groupNameList, groupIDList, err := models.GetGroups(creatorID)
	if err != nil {
		return utils.JSONResponse{
			StatusCode: 401,
			Data:       "Failed to retrieve groups.",
		}, fmt.Errorf("[ERROR] Failed to retrieve group. %w", err)
	}

	var responseStruct struct {
		GroupNames []string
		GroupIDs []int
	}
	responseStruct.GroupNames = groupNameList
	responseStruct.GroupIDs = groupIDList
	return utils.JSONResponse{
		StatusCode: 200,
		Data:       responseStruct,
	}, nil
}
