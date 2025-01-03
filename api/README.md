# API

*This is the HTTP API that the frontend client interacts with to create the
functionallity of an end-to-end encrypted group messenger.*

## Table of Contents

**Auth**

- [Login](#login)
- [Logout](#logout)

**User**
- [Get User](#userget)
- [Update User](#userupdate)
- [Create User](#usercreate)
- [Delete User](#userdelete)
- [Send friend request](#userfriendinvite)
- [Accept friend request](#userfriendaccept)
- [Reject friend request](#userfriendreject)
- [Remove friend](#userfriendremove)
- [Get friend](#userfriendget)
- [Get groups](#usergroups)
- [Listen for friend request](#userfriendlisten)
- [Get Invites](#userfriendinviteget)

**Group**
- [Update Group](#groupupdate)
- [Create Group](#groupcreate)
- [Delete Group](#groupdelete)
- [Group Invite](#groupinvite)
- [Group Accept](#groupinviteaccept)
- [Group Reject](#groupinvitereject)
- [Listen for group invites](#groupinvitelisten)
- [Listen for group chat](#groupchat)
- [Get Messages](#groupmessages)
- [Get Users](#groupusers)
- [Remove User](#groupban)
- [Chat](#groupchat)
- [Get Invites](#groupinviteget)

**Misc.**

- [Healthcheck](#healthcheck)

## Build

To build the application:

```
$ sudo podman build -t localhost/cs4389-api -f ./Dockerfile ../
```

## Usage

To run the API:

```
$ sudo podman run --env-file ../.env --name cs4389-api -p 8081:8080 \
  -d localhost/cs4389-api
```

To stop the API:

```
$ sudo podman stop cs4389-api
```

## Documentation

### /healthcheck

*Route to test the current status of the API.*

**Method**: `GET`

**Parameters**: `N/A`

**Example**: `https://api.application.com/healthcheck`

**Returns**: `200`

```JSON
{
    "StatusCode": 200,
    "Data": "Ok"
}
```

### /login

*Route to login a user.*

**Method**: `POST`

**Body**: `username`, `password`

**Example**: `https://api.application.com/login`

**Returns**: `200`, `401`, `400`

```JSON
{
    "StatusCode": 200,
    "Data": "Ok"
}
```

```JSON
{
    "StatusCode": 400,
    "Data": "Bad Request"
}
```

```JSON
{
    "StatusCode": 401,
    "Data": "Invalid username or password."
}
```

### /logout

*Route to logout a user.*

**Method**: `POST`

**Body**: N/A

**Example**: `https://api.application.com/logout`

**Returns**: `200`, `401`, `400`

```JSON
{
    "StatusCode": 200,
    "Data": "Ok"
}
```

```JSON
{
    "StatusCode": 400,
    "Data": "Bad Request"
}
```

```JSON
{
    "StatusCode": 401,
    "Data": "Unauthorized"
}
```

### /user/get

*Route to get a users public key.*

**Method**: `POST`

**Body**: `username`

**Example**: `https://api.application.com/user/get`

**Returns**: `200`, `400`

```JSON
{
    "StatusCode": 200,
    "Data": {
        "PublicKey": "supersecretpublickey"
    }
}
```

```JSON
{
    "StatusCode": 400,
    "Data": "Bad Request"
}
```

### /user/update

*Route to update user information.*

**Method**: `POST`

**Body**: `password` (optional), `publicKey` (optional)

**Example**: `https://api.application.com/user/update`

**Returns**: `200`, `400`, `401`

```JSON
{
    "StatusCode": 200,
    "Data": "Ok"
}
```

```JSON
{
    "StatusCode": 400,
    "Data": "Bad Request"
}
```

```JSON
{
    "StatusCode": 401,
    "Data": "Unauthorized"
}
```

### /user/create

*Route to create a new user.*

**Method**: `POST`

**Body**: `username`, `password`, `publicKey`

**Example**: `https://api.application.com/user/create`

**Returns**: `200`, `400`

```JSON
{
    "StatusCode": 200,
    "Data": "Ok"
}
```

```JSON
{
    "StatusCode": 400,
    "Data": "Bad Request"
}
```

### /user/delete

*Route to create a delete a user.*

**Method**: `POST`

**Body**: N/A

**Example**: `https://api.application.com/user/delete`

**Returns**: `200`, `401`, `400`

```JSON
{
    "StatusCode": 200,
    "Data": "Ok"
}
```

```JSON
{
    "StatusCode": 400,
    "Data": "Bad Request"
}
```

```JSON
{
    "StatusCode": 401,
    "Data": "Unauthorized"
}
```

### /user/friend/invite

*Route to send a friend request to a user.*

**Method**: `POST`

**Body**: `username`

**Example**: `https://api.application.com/user/friend/invite`

**Returns**: `200`, `400`, `401`

```JSON
{
    "StatusCode": 200,
    "Data": "Ok"
}
```

```JSON
{
    "StatusCode": 400,
    "Data": "Bad Request"
}
```

```JSON
{
    "StatusCode": 401,
    "Data": "Unauthorized"
}
```

### /user/friend/accept

*Route to accept a friend request to a user.*

**Method**: `POST`

**Body**: `username`

**Example**: `https://api.application.com/user/friend/accept`

**Returns**: `200`, `400`, `401`

```JSON
{
    "StatusCode": 200,
    "Data": "Ok"
}
```

```JSON
{
    "StatusCode": 400,
    "Data": "Bad Request"
}
```

```JSON
{
    "StatusCode": 401,
    "Data": "Unauthorized"
}
```

### /user/friend/reject

*Route to reject a friend request to a user.*

**Method**: `POST`

**Body**: `username`

**Example**: `https://api.application.com/user/friend/reject`

**Returns**: `200`, `400`, `401`

```JSON
{
    "StatusCode": 200,
    "Data": "Ok"
}
```

```JSON
{
    "StatusCode": 400,
    "Data": "Bad Request"
}
```

```JSON
{
    "StatusCode": 401,
    "Data": "Unauthorized"
}
```

### /user/friend/remove

*Route to remove a friend from a user.*

**Method**: `POST`

**Body**: `username`

**Example**: `https://api.application.com/user/friend/remove`

**Returns**: `200`, `400`, `401`

```JSON
{
    "StatusCode": 200,
    "Data": "Ok"
}
```

```JSON
{
    "StatusCode": 400,
    "Data": "Bad Request"
}
```

```JSON
{
    "StatusCode": 401,
    "Data": "Unauthorized"
}
```

### /user/friend/get

*Route to get a list friend from a user.*

**Method**: `GET`

**Body**: `N/A`

**Example**: `https://api.application.com/user/friend/get`

**Returns**: `200`, `400`, `401`

```JSON
{
    "StatusCode":200,
    "Data": {
        "Friends": ["foo"]
    }
}
```

```JSON
{
    "StatusCode": 400,
    "Data": "Bad Request"
}
```

```JSON
{
    "StatusCode": 401,
    "Data": "Unauthorized"
}
```

### /user/friend/invite/get

*Route to get a users friend invites.*

**Method**: `POST`

**Body**: N/A

**Example**: `https://api.application.com/user/friend/invite/get`

**Returns**: `200`, `400`

```JSON
{
    "StatusCode": 200,
    "Data": {
        "Users": [
          "Foo"
        ]
    }
}
```

```JSON
{
    "StatusCode": 400,
    "Data": "Bad Request"
}
```

### /user/groups

*Route to get a list groups a user is in.*

**Method**: `POST`

**Body**: `N/A`

**Example**: `https://api.application.com/user/groups`

**Returns**: `200`, `400`, `401`

```JSON
{
    "StatusCode":200,
    "Data": {
        "Groups": [
          {
            "GroupID": 1,
            "CreatorID": 2,
            "GroupName": "Foo"
          }
        ]
    }
}
```

```JSON
{
    "StatusCode": 400,
    "Data": "Bad Request"
}
```

```JSON
{
    "StatusCode": 401,
    "Data": "Unauthorized"
}
```

### /user/friend/listen

*Websocket to listen for when friend requests occur for the user.*

**Method**: `N/A`

**Body**: `N/A`

**Example**: `wss://api.application.com/user/friend/listen`

**Returns**: `N/A`

```JSON
{
    "Message": "root"
}
```

### /group/create

*Route to create a new group.*

**Method**: `POST`

**Body**: `groupname`

**Example**: `https://api.application.com/group/create`

**Returns**: `200`, `400`

```JSON
{
    "StatusCode": 200,
    "Data": {
      "GroupID": 1
    }
}
```

```JSON
{
    "StatusCode": 400,
    "Data": "Bad Request"
}
```

### /group/delete

*Route to create a delete a group.*

**Method**: `POST`

**Body**: `group` (groupID)

**Example**: `https://api.application.com/group/delete`

**Returns**: `200`, `401`, `400`
```JSON
{
    "StatusCode": 200,
    "Data": "Ok"
}
```

```JSON
{
    "StatusCode": 400,
    "Data": "Bad Request"
}
```
```JSON
{
    "StatusCode": 401,
    "Data": "Unauthorized"
}
```

### /group/update

*Route to updates a group.*

**Method**: `POST`

**Body**: `groupname`, `groupID`

**Example**: `https://api.application.com/group/update`

**Returns**: `200`, `401`, `400`
```JSON
{
    "StatusCode": 200,
    "Data": "Ok"
}
```

```JSON
{
    "StatusCode": 400,
    "Data": "Bad Request"
}
```
```JSON
{
    "StatusCode": 401,
    "Data": "Unauthorized"
}
```

### /group/invite

*Route to send a group invite to a user.*

**Method**: `POST`

**Body**: `username`, `key`, `group` (groupID)

**Example**: `https://api.application.com/group/invite`

**Returns**: `200`, `400`, `401`

```JSON
{
    "StatusCode": 200,
    "Data": "Ok"
}
```

```JSON
{
    "StatusCode": 400,
    "Data": "Bad Request"
}
```

```JSON
{
    "StatusCode": 401,
    "Data": "Unauthorized"
}
```

### /group/invite/accept

*Route to accept a given group invite.*

**Method**: `POST`

**Body**: `group` (groupID)

**Example**: `https://api.application.com/group/invite/accept`

**Returns**: `200`, `400`, `401`

```JSON
{
    "StatusCode": 200,
    "Data": {
        "EncryptedKey": "supersecretkey"
    }
}
```

```JSON
{
    "StatusCode": 400,
    "Data": "Bad Request"
}
```

```JSON
{
    "StatusCode": 401,
    "Data": "Unauthorized"
}
```

### /group/invite/reject

*Route to reject a given group invite.*

**Method**: `POST`

**Body**: `group` (groupID)

**Example**: `https://api.application.com/group/invite/accept`

**Returns**: `200`, `400`, `401`

```JSON
{
    "StatusCode": 200,
    "Data": "Ok"
}
```

```JSON
{
    "StatusCode": 400,
    "Data": "Bad Request"
}
```

```JSON
{
    "StatusCode": 401,
    "Data": "Unauthorized"
}
```

### /group/invite/listen

*Websocket to listen for when group invite requests occur for the user.*

**Method**: `N/A`

**Body**: `N/A`

**Example**: `wss://api.application.com/group/invite/listen`

**Returns**: `N/A`

```JSON
{
    "Message": "1-FunnyGroup"
}
```

### /group/chat

*Websocket to handle real time chat.*

**Method**: `N/A`

**Body**: `N/A`

**Example**: `wss://api.application.com/group/chat?group=1`

**Returns**: `N/A`

`Read`
```JSON
{
    "Message": "root-<Encrypted Message>"
}
```

`Write`
```JSON
{
    "Message": "<Encrypted Message>"
}
```

### /group/messages

*Route to return messages from a group.*

**Method**: `POST`

**Body**: `group` (groupID), `start`, `end` (2024-01-01, inclusive range)

**Example**: `https://api.application.com/group/messages`

**Returns**: `200`, `400`, `401`

```JSON
{
    "StatusCode": 200,
    "Data": {
      "Messages": [
        {
          "UserID": 1,
          "GroupID": 1,
          "Timestamp": "<datetime>"
          "EncryptedMessage": "asdfasdfasdf"
        }
      ]
    }
}
```

```JSON
{
    "StatusCode": 400,
    "Data": "Bad Request"
}
```

```JSON
{
    "StatusCode": 401,
    "Data": "Unauthorized"
}
```

### /group/ban

*Remoes a user from the group*

**Method**: `POST`

**Body**: `group` (groupID), `username`

**Example**: `https://api.application.com/group/ban`

**Returns**: `200`, `400`, `401`

```JSON
{
    "StatusCode": 200,
    "Data": "Ok"
}
```

```JSON
{
    "StatusCode": 400,
    "Data": "Bad Request"
}
```

```JSON
{
    "StatusCode": 401,
    "Data": "Unauthorized"
}
```

### /group/users

*Route to get a list users in a group.*

**Method**: `POST`

**Body**: `group` (groupID)

**Example**: `https://api.application.com/group/users`

**Returns**: `200`, `400`, `401`

```JSON
{
    "StatusCode":200,
    "Data": {
        "Users": [
          "root"
        ]
    }
}
```

```JSON
{
    "StatusCode": 400,
    "Data": "Bad Request"
}
```

```JSON
{
    "StatusCode": 401,
    "Data": "Unauthorized"
}
```

### /group/invite/get

*Route to get a users group invites.*

**Method**: `POST`

**Body**: N/A

**Example**: `https://api.application.com/group/invite/get`

**Returns**: `200`, `400`

```JSON
{
    "StatusCode": 200,
    "Data": {
        "Groups": [
          "GroupID": 6,
          "CreatorID": 2,
          "GroupName": "FooBar"
        ]
    }
}
```

```JSON
{
    "StatusCode": 400,
    "Data": "Bad Request"
}
```
