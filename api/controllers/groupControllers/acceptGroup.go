/* =========================================================================
*  File Name: controller/groupControllers/acceptGroup.go
*  Description: Controller for accepting a friend request.
*  Author: MagnusChase03
*  =======================================================================*/
package groupControllers

import (
	"fmt"

	"github.com/MagnusChase03/CS4389-Project/models"
	"github.com/MagnusChase03/CS4389-Project/utils"
)

/*
*  Attempts to accept an invite request.
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
func AcceptGroupInviteController(userID int, groupID int) (utils.JSONResponse, error) {
	encryptedKey, err := models.AcceptGroupInvite(userID, groupID)
	if err != nil {
		return utils.JSONResponse{
			StatusCode: 400,
			Data:       "Failed to accept invite request.",
		}, fmt.Errorf("[ERROR] Failed to accpet invite request. %w", err)
	}

	var responseStruct struct {
		EncryptedKey string
	}
	responseStruct.EncryptedKey = encryptedKey

	return utils.JSONResponse{
		StatusCode: 200,
		Data:       responseStruct,
	}, nil
}
