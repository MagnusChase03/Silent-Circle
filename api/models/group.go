/* =========================================================================
*  File Name: models/group.go
*  Description: Handles database interactions related to groups.
*  Author: Matthew-Basinger
*  =======================================================================*/
package models

import (
	"fmt"

	"github.com/MagnusChase03/CS4389-Project/db"
)

type Group struct {
	GroupID   int
	CreatorID int
	GroupName string
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
func CreateGroup(creatorID int, groupName string) error {
	instance, err := db.GetMariaDB()
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to get mariadb instance. %w", err)
	}

	insertStatement, err := instance.Connection.Prepare(
		"INSERT INTO Groups(CreatorID, GroupName) VALUES (?, ?)",
	)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to parse SQL query. %w", err)
	}
	defer insertStatement.Close()

	_, err = insertStatement.Exec(creatorID, groupName)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to create group. %w", err)
	}

	return nil
}

/*
*  Delete a group in the database with given ID.
*
*  Arguments:
*      - groupID (int): The groupID.
*
*  Returns:
*      - error: An error if any occurred.
*
 */
func DeleteGroup(groupname string) error {
	instance, err := db.GetMariaDB()
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to get mariadb instance. %w", err)
	}

	deleteStatement, err := instance.Connection.Prepare(
		"DELETE FROM Groups WHERE GroupName = ?",
	)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to parse SQL query. %w", err)
	}
	defer deleteStatement.Close()

	_, err = deleteStatement.Exec(groupname)
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
