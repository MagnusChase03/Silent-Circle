/* =========================================================================
*  File Name: controller/groupControllers/getInvites.go
*  Description: Controller for getting group invites.
*  Author: MagnusChase03
*  =======================================================================*/
package groupControllers

import (
	"fmt"

	"github.com/MagnusChase03/CS4389-Project/models"
	"github.com/MagnusChase03/CS4389-Project/utils"
)

/*
*  Attempts to get group invites.
*
*  Arguments:
*	   - userID (int): The ID of the requesting user.
*
*  Returns:
*      - utils.JSONResponse: The response to be made to the client.
*      - error: An error if any occurred.
*
 */
func GetGroupInvitesController(userID int) (utils.JSONResponse, error) {
	groups, err := models.GetGroupInvites(userID)
	if err != nil {
		return utils.JSONResponse{
			StatusCode: 400,
			Data:       "Failed to get group invites.",
		}, fmt.Errorf("[ERROR] Failed to get group invites. %w", err)
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
