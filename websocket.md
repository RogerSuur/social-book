# JSON structure for Websocket messages

## chat message

```JSON
{
    "type": "message",
    "data": {
        "recipient_id": 123, // 0 if group chat
        "group_id": 123, // 0 if private chat
        "body": "message content",
        "timestamp": "2006-01-02T15:04:05Z07:00" //won't be sending from frontend, but still need to receive it
    }
}
```

## chatlist

```JSON
{
    "type": "chatlist",
    "data": {
        "userid": 123, // 0 if group
        "group_id": 123, // 0 if user
        "username": "username", //group name if group
        "first_name": "first name", //omit the field if group
        "last_name": "last name", //omit the field if group
        "timestamp": "2006-01-02T15:04:05Z07:00", // date of last message in the chat if any, might use it to sort chats by last message
        "avatarImage": "link" // if will be used in chatlist
    }
}
```

## follow request

```JSON
{
    "type": "follow_request",
    "data": {
        "following_id": 123,
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

## send group invite

```JSON
{
    "type": "group_invite",
    "data": {
        "group_id": 123,
        "recipient_id": 123,
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
