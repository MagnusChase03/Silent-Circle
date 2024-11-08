/* =========================================================================
*  File Name: controller/groupController/updateGroup.go
*  Description: Controller for updating a user.
*  Author: Matthew-Basinger
*  =======================================================================*/
package groupControllers

import (
    "fmt"

    "github.com/MagnusChase03/CS4389-Project/models"
    "github.com/MagnusChase03/CS4389-Project/utils"
)

/*
*  Attempts to update a group with given attributes.
*
*  Arguments:
*	   - userID (int): The unique ID of the user
*      - groupname (string): The name of the group.
*      - groupID (int): The unique ID of the group.
* 
*  Returns:
*      - utils.JSONResponse: The response to be made to the client.
*      - error: An error if any occurred.
*
*/
func UpdateGroupController( userID int, groupname string, groupID int) (utils.JSONResponse, error) { 
    creatorID, err := models.GetCreatorIDByGroupName(groupname);
    if err != nil {
		return utils.JSONResponse{
            StatusCode: 401,
            Data: "Failed to update group.",
        }, fmt.Errorf("[ERROR] Failed to update group. %w", err);
	} else if creatorID != userID{
		return utils.JSONResponse{
            StatusCode: 401,
            Data: "Failed to update group.",
        }, fmt.Errorf("[ERROR] Not group creator. %w", err);
	}
    err = models.UpdateGroup(userID, groupname, groupID);
    if err != nil {
        return utils.JSONResponse{
            StatusCode: 400,
            Data: "Failed to update group.",
        }, fmt.Errorf("[ERROR] Failed to update group. %w", err);
    }

    return utils.JSONResponse{
        StatusCode: 200,
        Data: "Ok",
    }, nil;
}
