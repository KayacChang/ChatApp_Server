# Chat Service 

### Payload Structure

```json
{
    "type": "<TYPE>",
    "action": "<ACTION>",
    "status": "OK" | "ERROR",
    "data": "<DATA>"
}
```

## API

### User

#### Join

##### Request
```json
{
    "type": "USER",
    "action": "JOIN",
    "data": { 
        "username": "<USER_NAME>"
    }
}
```

##### Response
```json
{
    "type": "USER",
    "action": "JOIN",
    "status": "OK",
    "data": {
        "username": "<USER_NAME>",
        "message": "User Join Success"
    }
}
```

#### Leave

##### Request
```json
{
    "type": "USER",
    "action": "LEAVE",
}
```

##### Response
```json
{
    "type": "USER",
    "action": "LEAVE",
    "status": "OK",
    "data": {
        "username": "<USER_NAME>",
        "message":"User Leave Success"
    }
}
```

### Room 

#### Join

##### Request
```json
{
    "type": "ROOM",
    "action": "JOIN",
    "data": {
        "room_id": "<ROOM_ID>"
    }
}
```

##### Response
```json
{
    "type": "ROOM",
    "action": "JOIN",
    "status": "OK",
    "data": {
        "room_id": "<ROOM_ID>",
        "message": "Room Leave Success"
    }
}
```

#### Leave

##### Request
```json
{
    "type": "ROOM",
    "action": "LEAVE",
}
```

##### Response
```json
{
    "type": "ROOM",
    "action": "LEAVE",
    "status": "OK",
    "data": {
        "room_id": "<ROOM_ID>",
        "message":"Room Leave Success"
    }
}
```

### Message

#### Send

```json
{
    "type": "MSG",
    "action": "SEND",
    "data": {
        "from": "<USER_NAME>",
        "message": "<MSG>"
    }
}
```

#### Receive

```json
{
    "type": "MSG",
    "action": "RECEIVE",
    "status": "OK",
    "data": {
        "id": "<UUID>",
        "name": "<USER_NAME>",
        "message": "<MSG>",
        "time": "<TIME_STRING>"
    }
}
```
