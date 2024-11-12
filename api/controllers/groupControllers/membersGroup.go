/* =========================================================================
*  File Name: controller/groupController/membersGroup.go
*  Description: Controller for listing all members in a group.
*  Author: Matthew-Basinger
*  =======================================================================*/
package groupControllers

import (
	"fmt"

	"github.com/MagnusChase03/CS4389-Project/models"
	"github.com/MagnusChase03/CS4389-Project/utils"
)

/*
*  Attempts to retrieve all members in a group
*
*  Arguments:
*      - groupID (int): The ID of the group.
*
*  Returns:
*      - utils.JSONResponse: The response to be made to the client.
*      - error: An error if any occurred.
*
*/
func MembersGroupController(groupID int) (utils.JSONResponse, error) {
	membersNameList, err := models.GetMembers(groupID)
	if err != nil {
		return utils.JSONResponse{
			StatusCode: 400,
			Data:       "Failed to retrieve group members.",
		}, fmt.Errorf("[ERROR] Failed to retrieve group. %w", err)
	}

	var responseStruct struct {
		MemberNames []string
	}
	responseStruct.MemberNames = membersNameList
	return utils.JSONResponse{
		StatusCode: 200,
		Data:       responseStruct,
	}, nil
}
