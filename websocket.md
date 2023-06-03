# JSON structure for Websocket messages

## chat message
```JSON
{
    "type": "new_message",
    "data": {
        "recipient_id": 123, // 0 if group chat
        "grouop_id": 123, // 0 if private chat
        "content": "message content",
        "timestamp": "2006-01-02T15:04:05Z07:00"
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