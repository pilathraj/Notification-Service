# API Contract

This is an API contract for the Library Management system.

## 1. Notification Service

### 1. User Notifications

- Endpoint: GET /api/notifications/user/:userId
- Description: Retrieves all notifications for a specific user.
- Path Parameter:
- userId (string): The ID of the user.
- Response:
```json
[
  {
    "id": "integer",
    "userId": "string",
    "message": "string",
    "read": "boolean",
    "createdAt": "string",
    "updatedAt": "string"
  }
]
```
- example
```json 
[
    {
        "id": 1,
        "userId": "1",
        "subject": "Test",
        "content": "Test Message",
        "type": "notification",
        "channel": "Email",
        "email": "",
        "mobileno": "",
        "isRead": true,
        "createdAt": "2025-05-04T15:46:42.24175Z",
        "readAt": "2025-05-05T07:36:10.232371Z"
    }
]
```
- 404 Not Found
```json
{
  "error": "User not found"
}
```
- 500 Internal Server Error:
```json
{
  "error": "Failed to fetch notifications"
}
```


### 2. Send Notification

- Endpoint: POST /api/notifications
- Description: Creates and sends a notification.
- Request Payload:
```json
{
  "userId": "string",
  "message": "string",
  "channel": "string" // e.g., "email", "sms", "push"
}
```
- Sample Payload:
```json
    {
        "userId": "3",
        "subject": "Test123",
        "content": "Test Message 3",
        "type": "notification",
        "channel": "Email"
    }
```
- Response:
- 201 Created:
```json
{
  "message": "Notification created successfully"
}
```
- 400 Bad Request:
```json
{
  "error": "Invalid input"
}
```
- 500 Internal Server Error:
 ```json
{
  "error": "Failed to create notification"
}
```


### 3. Mark Notification as Read

- Endpoint: PUT /api/notifications/:notificationId/read
- Description: Marks a specific notification as read.
- Path Parameter:
- notificationId (integer): The ID of the notification.
- Response:
- 200 OK
- Body:
```json
{
  "message": "Notification marked as read"
}
```
- 404 Not Found
```json
{
  "error": "Notification not found"
}
```
- 500 Internal Server Error
```json
{
  "error": "Failed to mark notification as read"
}
```



### 4. Delete Notification
- Endpoint: DELETE /api/notifications/:notificationId
- Description: Deletes a specific notification.
- Path Parameter:
- notificationId (integer): The ID of the notification.
- Response:
- 200 OK
```json
{
  "message": "Notification deleted successfully"
}
```
- 404 Not Found
```json
{
  "error": "Notification not found"
}
```
- 500 Internal Server Error
{
  "error": "Failed to delete notification"
}

### 5. User Preferences Update

- Endpoint: GET /api/notifications/preferences/:userId
- Description: Retrieves the notification preferences for a specific user.
- Path Parameter:
- userId (string): The ID of the user.
- Response:
- 200 OK:
```json
{
  "userId": "string",
  "dueReminders": "boolean",
  "overdueNotices": "boolean",
  "reservationNotices": "boolean",
  "fineNotices": "boolean",
  "preferredChannels": ["string"]
}
```
- Sample
```json
{
    "id": 2,
    "userId": "1",
    "dueReminders": true,
    "overdueNotices": false,
    "reservationNotices": false,
    "fineNotices": false,
    "preferredChannels": [
        "email",
        "Sms"
    ],
    "updatedAt": "2025-05-05T08:22:04.214938Z"
}
```
- 404 Not Found
```json
{
  "error": "Preferences not found"
}
```
- 500 Internal Server Error
```json
{
  "error": "Failed to fetch preferences"
}
```


### 6. Update User Preferences

- Endpoint: PUT /api/notifications/preferences/:userId
- Description: Updates the notification preferences for a specific user.
- Path Parameter:
- userId (string): The ID of the user.
- Request Payload
```json
{
  "dueReminders": "boolean",
  "overdueNotices": "boolean",
  "reservationNotices": "boolean",
  "fineNotices": "boolean",
  "preferredChannels": ["string"] // e.g., ["email", "sms"]
}
```
- Response
- 200 OK:
```json
{
  "message": "Preferences updated successfully"
}
```
- 400 Bad Request
```json
{
  "error": "Invalid input"
}
```
- 500 Internal Server Error
```json
{
  "error": "Failed to update preferences"
}
```


### Notify [#TODO]

- Method: GET
- URL: `http://localhost:8081/api/notifications/notify`
