/* =========================================================================
*  File Name: controller/userController/createUser.go
*  Description: Controller for the login route.
*  Author: MagnusChase03
*  =======================================================================*/
package userControllers

import (
    "fmt"

    "github.com/MagnusChase03/CS4389-Project/models"
    "github.com/MagnusChase03/CS4389-Project/utils"
)

/*
*  Attempts to create a new user with given attributes.
*
*  Arguments:
*      - username (string): The username to login to.
*      - passwordHash (string): The password hash to use to login.
*      - publicKey (string): The public key of the user.
* 
*  Returns:
*      - utils.JSONResponse: The response to be made to the client.
*      - error: An error if any occurred.
*
*/
func CreateUserController(username string, passwordHash string, publicKey string) (utils.JSONResponse, error) { 
    err := models.CreateUser(username, passwordHash, publicKey);
    if err != nil {
        return utils.JSONResponse{
            StatusCode: 401,
            Data: "Failed to create user.",
        }, fmt.Errorf("[ERROR] Failed to create user. %w", err);
    }

    return utils.JSONResponse{
        StatusCode: 200,
        Data: "Ok",
    }, nil;
}
