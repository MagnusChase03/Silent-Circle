/* =========================================================================
*  File Name: models/group.go
*  Description: Handles database interactions related to groups.
*  Author: Zachary, Matthew-Basinger
*  =======================================================================*/
package models

import (
	"fmt"
	"time"

	"github.com/MagnusChase03/CS4389-Project/db"
)

type Group struct {
	GroupID   int
	CreatorID int
	GroupName string
}

type Message struct {
	UserID           int
	GroupID          int
	Timestamp        time.Time
	EncryptedMessage string
}

/*
*  Returns a group with given GroupID
*
*  Arguments
*      - id (int): The GroupID to find
*
*  Returns:
*      - Group: the group information
*      - error: An error if any occured
 */
func GetGroupByID(id int) (Group, error) {
	var group Group
	instance, err := db.GetMariaDB()
	if err != nil {
		return group, fmt.Errorf("[ERROR] Failed to get mariadb instance. %w", err)
	}

	query, err := instance.Connection.Prepare("SELECT * FROM Groups WHERE GroupID = ?")
	if err != nil {
		return group, fmt.Errorf("[ERROR] Failed to get parse SQL query. %w", err)
	}
	defer query.Close()

	err = query.QueryRow(id).Scan(
		&group.GroupID,
		&group.CreatorID,
		&group.GroupName,
	)
	if err != nil {
		return group, fmt.Errorf("[ERROR] Failed to find group with GroupID %d. %w", id, err)
	}

	return group, nil
}

/*
*  Create a group in the database with given attributes.
*
*  Arguments:
*      - groupName (string): The name of the group.
*      - creatorID (int): The ID of the creator.
*
*  Returns:
*      - error: An error if any occurred.
*
 */
func CreateGroup(creatorID int, groupName string) (int, error) {
	instance, err := db.GetMariaDB()
	if err != nil {
		return 0, fmt.Errorf("[ERROR] Failed to get mariadb instance. %w", err)
	}

	insertStatement, err := instance.Connection.Prepare(
		"INSERT INTO Groups(CreatorID, GroupName) VALUES (?, ?)",
	)
	if err != nil {
		return 0, fmt.Errorf("[ERROR] Failed to parse SQL query. %w", err)
	}
	defer insertStatement.Close()

	_, err = insertStatement.Exec(creatorID, groupName)
	if err != nil {
		return 0, fmt.Errorf("[ERROR] Failed to create group. %w", err)
	}

	groupID, err := GetGroupIDByGroupName(groupName)
	if err != nil {
		err = DeleteGroup(creatorID, groupID)
		return 0, fmt.Errorf("[ERROR] Failed to create group. %w", err)
	}

	err = AddCreatorUserGroup(creatorID, groupID)
	if err != nil {
		err = DeleteGroup(creatorID, groupID)
		return 0, fmt.Errorf("[ERROR] Failed to add creator to group. %w", err)
	}

	return groupID, nil

}

/*
*  Delete a group in the database with given ID.
*
*  Arguments:
*      - groupID (int): The groupID.
*      - userID (int): The userID.
*
*  Returns:
*      - error: An error if any occurred.
*
 */
func DeleteGroup(userID int, groupID int) error {
	instance, err := db.GetMariaDB()
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to get mariadb instance. %w", err)
	}

	deleteStatement, err := instance.Connection.Prepare(
		"DELETE FROM Groups WHERE GroupID = ? AND CreatorID = ?",
	)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to parse SQL query. %w", err)
	}
	defer deleteStatement.Close()

	_, err = deleteStatement.Exec(groupID, userID)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to delete group. %w", err)
	}

	return nil
}

/*
*  Determines the creatorID of a group in the database with given groupname.
*
*  Arguments:
*      - groupname (string): The groupname.
*
*  Returns:
*      - error: An error if any occurred.
*
 */
func GetCreatorIDByGroupName(groupname string) (int, error) {
	var group Group
	instance, err := db.GetMariaDB()
	if err != nil {
		return 0, fmt.Errorf("[ERROR] Failed to get mariadb instance. %w", err)
	}

	query, err := instance.Connection.Prepare("SELECT * FROM Groups WHERE GroupName = ?")
	if err != nil {
		return 0, fmt.Errorf("[ERROR] Failed to get parse SQL query. %w", err)
	}
	defer query.Close()

	err = query.QueryRow(groupname).Scan(
		&group.GroupID,
		&group.CreatorID,
		&group.GroupName,
	)
	if err != nil {
		return 0, fmt.Errorf("[ERROR] Failed to find group. %w", err)
	}

	return group.CreatorID, nil
}

/*
*  Determines the groupID of a group in the database with given groupname.
*
*  Arguments:
*      - groupname (string): The groupname.
*
*  Returns:
*      - error: An error if any occurred.
*
 */
func GetGroupIDByGroupName(groupname string) (int, error) {
	var group Group
	instance, err := db.GetMariaDB()
	if err != nil {
		return 0, fmt.Errorf("[ERROR] Failed to get mariadb instance. %w", err)
	}

	query, err := instance.Connection.Prepare("SELECT * FROM Groups WHERE GroupName = ?")
	if err != nil {
		return 0, fmt.Errorf("[ERROR] Failed to get parse SQL query. %w", err)
	}
	defer query.Close()

	err = query.QueryRow(groupname).Scan(
		&group.GroupID,
		&group.CreatorID,
		&group.GroupName,
	)
	if err != nil {
		return 0, fmt.Errorf("[ERROR] Failed to find group with GroupID %d. %w", groupname, err)
	}

	return group.GroupID, nil
}

/*
*  Update a group in the database with given attributes.
*
*  Arguments:
*      - userID (int): The userID to update.
*      - groupname (string): The new groupname for the group.
*      - groupID (string): The groupID of the group.
*
*  Returns:
*      - error: An error if any occurred.
*
 */
func UpdateGroup(userID int, groupname string, groupID int) error {
	instance, err := db.GetMariaDB()
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to get mariadb instance. %w", err)
	}
	if groupname == "" {
		return fmt.Errorf("[ERROR] Failed to parse SQL query. %w", err)
	} else {
		updateStatement, err := instance.Connection.Prepare(
			"UPDATE Groups SET GroupName=?, WHERE GroupID=?",
		)
		if err != nil {
			return fmt.Errorf("[ERROR] Failed to parse SQL query. %w", err)
		}
		defer updateStatement.Close()

		_, err = updateStatement.Exec(groupname, groupID)
		if err != nil {
			return fmt.Errorf("[ERROR] Failed to update user. %w", err)
		}
	}

	return nil
}

/*
*  Adds the group a user creates to their UserGroup.
*
*  Arguments:
*      - creatorID (int): The ID of the creator.
*      - groupID (int): The ID of the group.
*
*  Returns:
*      - error: An error if any occurred.
*
 */
func AddCreatorUserGroup(creatorID int, groupID int) error {
	instance, err := db.GetMariaDB()
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to get mariadb instance. %w", err)
	}

	insertStatement, err := instance.Connection.Prepare(
		"INSERT INTO UserGroup(UserID, GroupID) VALUES (?, ?)",
	)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to parse SQL query. %w", err)
	}
	defer insertStatement.Close()

	_, err = insertStatement.Exec(creatorID, groupID)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to add creator to group. %w", err)
	}

	return nil
}

/*
*  Invites a given user to the given group.
*
*  Arguments:
*      - userID (int): The userID sending the invite.
*      - username (string): The new groupname for the group.
*      - encryptedKey (string): The encrypted group session key.
*      - groupID (int): The groupID of the group.
*
*  Returns:
*      - error: An error if any occurred.
*
 */
func SendGroupInvite(userID int, username string, encryptedKey string, groupID int) error {
	instance, err := db.GetMariaDB()
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to get mariadb instance. %w", err)
	}

	// Ensure user is owner of group
	query, err := instance.Connection.Prepare(
		"SELECT Users.UserID FROM Users JOIN Groups ON Users.UserID = Groups.CreatorID WHERE Groups.GroupID = ?",
	)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to get parse SQL query. %w", err)
	}
	defer query.Close()

	var creatorID int
	err = query.QueryRow(groupID).Scan(&creatorID)
	if err != nil {
		return fmt.Errorf(
			"[ERROR] Could not find group. %w",
			err,
		)
	}

	// Get invitee ID
	query, err = instance.Connection.Prepare(
		"SELECT UserID FROM Users WHERE Username = ?",
	)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to get parse SQL query. %w", err)
	}
	defer query.Close()

	var user2ID int
	err = query.QueryRow(username).Scan(&user2ID)
	if err != nil {
		return fmt.Errorf(
			"[ERROR] Could not find user. %w",
			err,
		)
	}

	// Inserting request
	insertStatement, err := instance.Connection.Prepare(
		"INSERT INTO GroupInvites(UserID, GroupID, EncryptedKey) VALUES (?, ?, ?)",
	)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to parse SQL query. %w", err)
	}
	defer insertStatement.Close()

	_, err = insertStatement.Exec(user2ID, groupID, encryptedKey)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to create group. %w", err)
	}

	return nil
}

/*
*  Accepts the given group invite.
*
*  Arguments:
*      - userID (int): The userID sending the invite.
*      - groupID (int): The groupID of the group.
*
*  Returns:
*      - string: The encrypted key from the invite.
*      - error: An error if any occurred.
*
 */
func AcceptGroupInvite(userID int, groupID int) (string, error) {
	instance, err := db.GetMariaDB()
	if err != nil {
		return "", fmt.Errorf("[ERROR] Failed to get mariadb instance. %w", err)
	}

	// Ensure invite exist
	query, err := instance.Connection.Prepare(
		"SELECT EncryptedKey FROM GroupInvites WHERE UserID = ? AND GroupID = ?",
	)
	if err != nil {
		return "", fmt.Errorf("[ERROR] Failed to get parse SQL query. %w", err)
	}
	defer query.Close()

	var encryptedKey string
	err = query.QueryRow(userID, groupID).Scan(&encryptedKey)
	if err != nil {
		return "", fmt.Errorf(
			"[ERROR] Could not find invite. %w",
			err,
		)
	}

	// Delete invite
	deleteStatement, err := instance.Connection.Prepare(
		"DELETE FROM GroupInvites WHERE UserID = ? AND GroupID = ?",
	)
	if err != nil {
		return "", fmt.Errorf("[ERROR] Failed to parse SQL query. %w", err)
	}
	defer deleteStatement.Close()

	_, err = deleteStatement.Exec(userID, groupID)
	if err != nil {
		return "", fmt.Errorf("[ERROR] Failed to delete invite. %w", err)
	}

	// Add to group
	insertStatement, err := instance.Connection.Prepare(
		"INSERT INTO UserGroup(UserID, GroupID) VALUES (?, ?)",
	)
	if err != nil {
		return "", fmt.Errorf("[ERROR] Failed to parse SQL query. %w", err)
	}
	defer insertStatement.Close()

	_, err = insertStatement.Exec(userID, groupID)
	if err != nil {
		return "", fmt.Errorf("[ERROR] Failed to insert user in group. %w", err)
	}

	return encryptedKey, nil
}

/*
*  Rejects the given group invite.
*
*  Arguments:
*      - userID (int): The userID sending the invite.
*      - groupID (int): The groupID of the group.
*
*  Returns:
*      - error: An error if any occurred.
*
 */
func RejectGroupInvite(userID int, groupID int) error {
	instance, err := db.GetMariaDB()
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to get mariadb instance. %w", err)
	}

	// Delete invite
	deleteStatement, err := instance.Connection.Prepare(
		"DELETE FROM GroupInvites WHERE UserID = ? AND GroupID = ?",
	)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to parse SQL query. %w", err)
	}
	defer deleteStatement.Close()

	_, err = deleteStatement.Exec(userID, groupID)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to delete invite. %w", err)
	}

	return nil
}

/*
*  Inserts a new message into the database.
*
*  Arguments:
*      - userID (int): The userID sending the invite.
*      - groupID (int): The groupID of the group.
*	   - message (string): The message to be added.
*
*  Returns:
*      - error: An error if any occurred.
*
 */
func InsertMessage(userID int, groupID int, message string) error {
	instance, err := db.GetMariaDB()
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to get mariadb instance. %w", err)
	}

	insertStatement, err := instance.Connection.Prepare(
		"INSERT INTO Messages(UserID, GroupID, Timestamp, EncryptedMessage) VALUES (?, ?, ?, ?)",
	)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to parse SQL query. %w", err)
	}
	defer insertStatement.Close()

	_, err = insertStatement.Exec(userID, groupID, time.Now(), message)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to insert message. %w", err)
	}

	return nil
}

/*
*  Returns the list of message within the group.
*
*  Arguments:
*      - userID (int): The userID sending the invite.
*      - groupID (int): The groupID of the group.
*	   - start (string): The start date.
*	   - end (string): The end date.
*
*  Returns:
*      - error: An error if any occurred.
*
 */
func GetMessages(userID int, groupID int, start string, end string) ([]Message, error) {
	instance, err := db.GetMariaDB()
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Failed to get mariadb instance. %w", err)
	}

	query, err := instance.Connection.Prepare("SELECT UserID FROM UserGroup WHERE UserID = ? AND GroupID = ?")
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Failed to get parse SQL query. %w", err)
	}
	defer query.Close()

	var matchUserID int
	err = query.QueryRow(userID, groupID).Scan(&matchUserID)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Failed to find group assossiated with user. %w", err)
	}

	startTime, err := time.Parse(time.DateOnly, start)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Invalid start time. %w", err)
	}

	endTime, err := time.Parse(time.DateOnly, end)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Invalid end time. %w", err)
	}

	query, err = instance.Connection.Prepare(`
		SELECT UserID, GroupID, Timestamp, EncryptedMessage
		FROM Messages
		WHERE GroupID = ? AND
		Timestamp >= ? AND
		Timestamp <= ?
		ORDER BY Timestamp ASC
	`)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Failed to get parse SQL query. %w", err)
	}
	defer query.Close()

	rows, err := query.Query(groupID, startTime, endTime)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Failed to find messages. %w", err)
	}

	messages := make([]Message, 0)
	for rows.Next() {
		var m Message
		if err := rows.Scan(&m.UserID, &m.GroupID, &m.Timestamp, &m.EncryptedMessage); err != nil {
			return nil, fmt.Errorf("[ERROR] Failed to find message. %w", err)
		}
		messages = append(messages, m)
	}

	return messages, nil
}
