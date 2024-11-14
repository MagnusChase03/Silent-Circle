/* =========================================================================
*  File Name: controller/groupControllers/getMessage.go
*  Description: Controller for getting old chat messages.
*  Author: MagnusChase03
*  =======================================================================*/
package groupControllers

import (
	"fmt"

	"github.com/MagnusChase03/CS4389-Project/models"
	"github.com/MagnusChase03/CS4389-Project/utils"
)

/*
*  Attempts to get old chat messages.
*
*  Arguments:
*      - userID (int): The userID of the sending user.
*      - groupID (int): The groupID of the group.
*	   - start (string): The start date.
*	   - end (string): The end date.
*
*  Returns:
*      - utils.JSONResponse: The response to be made to the client.
*      - error: An error if any occurred.
*
 */
func GetMessageController(userID int, groupID int, start string, end string) (utils.JSONResponse, error) {
	messages, err := models.GetMessages(userID, groupID, start, end)
	if err != nil {
		return utils.JSONResponse{
			StatusCode: 400,
			Data:       "Failed to get messages.",
		}, fmt.Errorf("[ERROR] Failed to get messages. %w", err)
	}

	var responseStruct struct {
		Messages []models.Message
	}
	responseStruct.Messages = messages

	return utils.JSONResponse{
		StatusCode: 200,
		Data:       responseStruct,
	}, nil
}
