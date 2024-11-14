/* =========================================================================
*  File Name: controller/groupControllers/chat.go
*  Description: Controller for real time chat with websocket from redis pub/sub
*  Author: MagnusChase03
*  =======================================================================*/
package groupControllers

import (
	"fmt"

	"github.com/MagnusChase03/CS4389-Project/db"
	"github.com/MagnusChase03/CS4389-Project/models"
	"github.com/gorilla/websocket"
)

/*
*  Sends data from redis pub/sub to websocket about user group invite requests.
*
*  Arguments:
*      - client (*websocket.Conn): The userID of the sending user.
*      - userID (int): The userID of the sending user.
*
*  Returns:
*      - error: An error if any occurred.
*
 */
func ChatController(client *websocket.Conn, userID int, groupID int) error {
	db, err := db.GetRedisDB()
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to get redis db. %w", err)
	}

	user, err := models.GetUserByID(userID)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to get user. %w", err)
	}

	groups, err := models.GetGroups(userID)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to get user groups. %w", err)
	}

	found := false
	for _, g := range groups {
		if g.GroupID == groupID {
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("[ERROR] User is not apart of group. %w", err)
	}

	subscriber := db.Connection.Subscribe(db.Ctx, fmt.Sprintf("chat-%d", groupID))
	defer subscriber.Close()

	go func() {
		for {
			msg, err := subscriber.ReceiveMessage(db.Ctx)
			if err != nil {
				fmt.Printf("[ERROR] Failed to receive message. %v\n", err)
				return
			}

			var data struct {
				Message string
			}
			data.Message = msg.Payload

			err = client.WriteJSON(data)
			if err != nil {
				fmt.Printf("[ERROR] Failed to send message to client. %v\n", err)
				return
			}
		}
	}()

	for {
		var data struct {
			Message string
		}
		err = client.ReadJSON(&data)
		if err != nil {
			return fmt.Errorf("[ERROR] Failed to read message from client. %w", err)
		}

		err = db.Connection.Publish(
			db.Ctx,
			fmt.Sprintf("chat-%d", groupID),
			fmt.Sprintf("%s-%s", user.Username, data.Message),
		).Err()
		if err != nil {
			return fmt.Errorf("[ERROR] Failed to publish message from client. %w", err)
		}

		err = models.InsertMessage(userID, groupID, data.Message)
		if err != nil {
			return fmt.Errorf("[ERROR] Failed to publish message from client. %w", err)
		}
	}

	return nil
}
