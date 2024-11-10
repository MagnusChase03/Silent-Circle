/* =========================================================================
*  File Name: controller/groupControllers/inviteUser.go
*  Description: Controller for sending a group invite
*  Author: MagnusChase03
*  =======================================================================*/
package groupControllers

import (
	"fmt"

	"github.com/MagnusChase03/CS4389-Project/db"
	"github.com/MagnusChase03/CS4389-Project/models"
	"github.com/MagnusChase03/CS4389-Project/utils"
)

/*
*  Attempts to invite a user to a given group.
*
*  Arguments:
*      - userID (int): The userID of the sending user.
*      - username (string): The username of the user to be invited.
*      - encryptedKey (string): The encrypted session key for the user.
*      - groupID (int): The group the user is getting invited to.
*
*  Returns:
*      - utils.JSONResponse: The response to be made to the client.
*      - error: An error if any occurred.
*
 */
func GroupInviteController(userID int, username string, encryptedKey string, groupID int) (utils.JSONResponse, error) {
	db, err := db.GetRedisDB()
	if err != nil {
		return utils.JSONResponse{
			StatusCode: 400,
			Data:       "Failed to send invite request.",
		}, fmt.Errorf("[ERROR] Failed to get redis db. %w", err)
	}

	err = models.SendGroupInvite(userID, username, encryptedKey, groupID)
	if err != nil {
		return utils.JSONResponse{
			StatusCode: 400,
			Data:       "Failed to send invite request.",
		}, fmt.Errorf("[ERROR] Failed to send invite request. %w", err)
	}

	user, err := models.GetUserByUsername(username)
	if err != nil {
		return utils.JSONResponse{
			StatusCode: 400,
			Data:       "Failed to send invite request.",
		}, fmt.Errorf("[ERROR] Failed to get user. %w", err)
	}

	group, err := models.GetGroupByID(groupID)
	if err != nil {
		return utils.JSONResponse{
			StatusCode: 400,
			Data:       "Failed to send invite request.",
		}, fmt.Errorf("[ERROR] Failed to get user. %w", err)
	}

	err = db.Connection.Publish(db.Ctx, fmt.Sprintf("gr-%d", user.UserID), fmt.Sprintf("%d-%s", groupID, group.GroupName)).Err()
	if err != nil {
		return utils.JSONResponse{
			StatusCode: 400,
			Data:       "Failed to send invite request.",
		}, fmt.Errorf("[ERROR] Failed to publish invite request. %w", err)
	}

	return utils.JSONResponse{
		StatusCode: 200,
		Data:       "Ok",
	}, nil
}
