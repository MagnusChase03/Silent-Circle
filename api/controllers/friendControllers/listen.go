/* =========================================================================
*  File Name: controller/friendControllers/listen.go
*  Description: Controller for sending friend requests to websocket from redis pub/sub
*  Author: MagnusChase03
*  =======================================================================*/
package friendControllers

import (
	"fmt"

	"github.com/MagnusChase03/CS4389-Project/db"
	"github.com/gorilla/websocket"
)

/*
*  Sends data from redis pub/sub to websocket about user friend requests.
*
*  Arguments:
*      - client (*websocket.Conn): The userID of the sending user.
*      - userID (int): The userID of the sending user.
*
*  Returns:
*      - error: An error if any occurred.
*
 */
func FriendRequestListenerController(client *websocket.Conn, userID int) error {
	db, err := db.GetRedisDB()
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to get redis db. %w", err)
	}

	subscriber := db.Connection.Subscribe(db.Ctx, fmt.Sprintf("fr-%d", userID))
	defer subscriber.Close()

	for {
		msg, err := subscriber.ReceiveMessage(db.Ctx)
		if err != nil {
			return fmt.Errorf("[ERROR] Failed to receive message.")
		}

		var data struct {
			Message string
		}
		data.Message = msg.Payload

		err = client.WriteJSON(data)
		if err != nil {
			return fmt.Errorf("[ERROR] Failed to send message to client. %w", err)
		}
	}

	return nil
}
