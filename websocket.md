# JSON structure for Websocket messages

## chat message

```JSON
{
    "type": "message",
    "data": {
        "id": 1, // message id
        "sender_id": 1,
        "sender_name" : "sdfs", // username (if exists) or first name last name
        "recipient_id": 123, // 0 if group chat
        "recipient_name": "somename", // username (if exists) or first name last name
        "group_id": 123, // 0 if private chat
        "group_name": "name", //empty if private chat
        "body": "message content",
        "timestamp": "2006-01-02T15:04:05Z07:00" //won't be sending from frontend, but still need to receive it
    }
}
```

## request chatlist

```JSON
{
    "type": "request_chatlist",
}
```

## chatlist

```JSON
{
    "type": "chatlist",
    "data": {
        "userid": 1, //own id
        "users" : [{"userid": 123, // 0 if group
        "group_id": 123, // 0 if user
        "name": "username", // username (if exists) or firstname lastname || group name if group
        "timestamp": "2006-01-02T15:04:05Z07:00", // date of last message in the chat if any, might use it to sort chats by last message
        "avatarImage": "link" // empty if no image or group
        }]
    }
}
```

## request message history

```JSON
{
    "type": "request_message_history",
    "data": {
        "id": 1, //0 if group chat
        "group_id": 1, // 0 if private chat
    }
}
```

## message history

```JSON
{
    "type": "message_history",
    "data": {
        "messages" : [{
            "id": 1, //message id
            "sender_id": 123, // 0 if group
            "sender_name": "username", // either a  username (if exists) or firstname and lastname
            "recipient_id": 1, // 0 if group
            "recipient_name": 1, // either a username (if   exists) or firstname and lastname && empty if     group
            "group_id": 123, // 0 if user
            "group_name": "name", //empty if user
            "body": "message",
            "timestamp": "2023-06-05 16:01:00.303095707 +03:00",
        }]
    }
}
```

## follow request

```JSON
{
    "type": "follow_request",
    "data": {
        "id": 1, //notification id
        "following_id": 123,
        "name": "something" // either a username (if exists) or firstname and lastname
    }
}
```

## group invite - someone inviting you

```JSON
{
    "type": "group_invite",
    "data": {
        "id": 1, //notification id
        "group_id": 123,
        "group_name": "somename",
        "sender_id": 123,
        "sender_name": "adsad", // either a username (if exists) or firstname and lastname
    }
}
```

## group join - someone wants to join a group you created

```JSON
{
    "type": "group_join",
    "data": {
        "id": 1, //notification id
        "group_id": 123,
        "group_name": "somename",
        "sender_id": 123,
        "sender_name": "adsad", // either a username (if exists) or firstname and lastname
    }
}
```

## event

```JSON
{
    "type": "event",
    "data": {
        "id": 1, //notification id
        "event_id": 123,
        "event_name": "somename",
        "event_datetime": "2023-06-05 16:01:00.303095707+03:00" //time of start
    }
}
```

## response

```JSON
{
    "type": "response",
    "data": {
        "id": 1, // notification id
        "reaction": true || false,
    }
}
```

## follow accept

```JSON
{
    "type": "follow_accept",
    "data": {
        "follower_id": 123,
    }
}
```

## follow reject

```JSON
{
    "type": "follow_reject",
    "data": {
        "follower_id": 123,
    }
}
```

## unfollow

```JSON
{
    "type": "unfollow",
    "data": {
        "following_id": 123,
    }
}
```

## group invite accept

```JSON
{
    "type": "group_accept",
    "data": {
        "group_id": 123,
    }
}
```

## group invite reject

```JSON
{
    "type": "group_reject",
    "data": {
        "group_id": 123,
    }
}
```
