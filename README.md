# Chat Service 

### Payload Structure

```json
{
    "type": "<TYPE>",
    "action": "<ACTION>",
    "from": "<FROM>",
    "message": "<MSG>"
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
    "from": "<User ID>"
}
```

##### Response
```json
{
    "type": "USER",
    "action": "JOIN",
    "from": "SERVER",
    "message": "User Join Success"
}
```

### Room 

#### Join

##### Request
```json
{
    "type": "ROOM",
    "action": "JOIN",
    "from": "<User ID>",
    "message": "<Room ID>"
}
```

##### Response
```json
{
    "type": "ROOM",
    "action": "JOIN",
    "from": "SERVER",
    "message": "Room Join Success"
}
```

#### Leave

##### Request
```json
{
    "type": "ROOM",
    "action": "LEAVE",
    "from": "<User ID>"
}
```

##### Response
```json
{
    "type": "ROOM",
    "action": "LEAVE",
    "from": "SERVER",
    "message": "Room Leave Success"
}
```

### Message

#### Send

```json
{
    "type": "MSG",
    "action": "SEND",
    "from": "<User ID>",
    "message": "<MSG>"
}
```

#### Receive

```json
{
    "type": "MSG",
    "action": "RECEIVE",
    "from": "<User ID>",
    "message": "<MSG>"
}
```
