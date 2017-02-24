---
title: Popcube iner API alpha

toc_footers:
- <a href='http://github.com/tripit/slate'>Documentation Powered by Slate</a>

search: true
---

# Popcube iner API alpha
> ### Consumes  
> `application/json`  

> ### Produces
> `application/json`  

**Schemes**: `http, https`

**Host**: `api.popcube`

**Base path**: `/alpha`


## 

```http
GET /alpha/ HTTP/1.1
```
	
```http
HTTP/1.1 200 OK
```
```http
HTTP/1.1 [default] 
```

Hello World


### Responses
Http code | Type | Description
--- | --- | ---
200 | no content | 
default | no content | AppError Type used to structure error reporting for popcube chat project.


## Get avatars

```http
GET /alpha/avatar HTTP/1.1
```
	
```http
HTTP/1.1 200 OK
```
```http
HTTP/1.1 503 Service Unavailable
```
```http
HTTP/1.1 [default] 
```

This will get all the avatars available in the organisation.


### Responses
Http code | Type | Description
--- | --- | ---
200 | no content | 
503 | no content | AppError Type used to structure error reporting for popcube chat project.
default | no content | AppError Type used to structure error reporting for popcube chat project.

## New avatar

```http
POST /alpha/avatar HTTP/1.1
```
	
```http
HTTP/1.1 200 OK
```
```http
HTTP/1.1 503 Service Unavailable
```
```http
HTTP/1.1 [default] 
```

This will create an avatar for organisation avatars library.


### Responses
Http code | Type | Description
--- | --- | ---
200 | no content | 
503 | no content | AppError Type used to structure error reporting for popcube chat project.
default | no content | AppError Type used to structure error reporting for popcube chat project.


## Get avatars

```http
GET /alpha/avatar/all HTTP/1.1
```
	
```http
HTTP/1.1 200 OK
```
```http
HTTP/1.1 503 Service Unavailable
```
```http
HTTP/1.1 [default] 
```

This will get all the avatars available in the organisation.


### Responses
Http code | Type | Description
--- | --- | ---
200 | no content | 
503 | no content | AppError Type used to structure error reporting for popcube chat project.
default | no content | AppError Type used to structure error reporting for popcube chat project.


## New avatar

```http
POST /alpha/avatar/new HTTP/1.1
```
	
```http
HTTP/1.1 200 OK
```
```http
HTTP/1.1 503 Service Unavailable
```
```http
HTTP/1.1 [default] 
```

This will create an avatar for organisation avatars library.


### Responses
Http code | Type | Description
--- | --- | ---
200 | no content | 
503 | no content | AppError Type used to structure error reporting for popcube chat project.
default | no content | AppError Type used to structure error reporting for popcube chat project.


## Should result in 500

```http
GET /alpha/panic HTTP/1.1
```
	
```http
HTTP/1.1 500 Internal Server Error
```
```http
HTTP/1.1 [default] 
```

Test panic cautching


### Responses
Http code | Type | Description
--- | --- | ---
500 | no content | 
default | no content | AppError Type used to structure error reporting for popcube chat project.


## Pong

```http
GET /alpha/ping HTTP/1.1
```
	
```http
HTTP/1.1 200 OK
```
```http
HTTP/1.1 [default] 
```

Test api ping


### Responses
Http code | Type | Description
--- | --- | ---
200 | no content | 
default | no content | AppError Type used to structure error reporting for popcube chat project.



# Models
## Avatar
```json
{
    "id": "integer",
    "link": "string",
    "name": "string"
}
```

Avatar object store default avatar images you can use within an association.
Required apply only for creation of the object. Id is required only on update.

	
### Fields
Name | Type | Format | Description
--- | --- | --- | ---
id | integer | uint64 | id of the avatar
link<span title="required" class="required">&nbsp;*&nbsp;</span> | string |  | Path into server
name<span title="required" class="required">&nbsp;*&nbsp;</span> | string |  | Avatar name

	
## Channel
```json
{
    "avatar": "string",
    "description": "string",
    "id": "integer",
    "last_update": "integer",
    "name": "string",
    "private": "boolean",
    "subject": "string",
    "type": "string",
    "web_id": "string"
}
```

Channel is the place where user can speak.
Required apply only for creation of the object.

	
### Fields
Name | Type | Format | Description
--- | --- | --- | ---
avatar | string |  | Photo :O
description | string |  | Describe the channel$<br/>max lenght: 1024
id | integer | uint64 | id of the channel
last_update | integer | int64 | Last time channel information where updated
name<span title="required" class="required">&nbsp;*&nbsp;</span> | string |  | 
private | boolean |  | Channel is private ?
subject | string |  | What we are speaking about<br/>max lenght: 250
type | string |  | Set if channel is text, video, audio or direct
web_id | string |  | web id for the user used only for cache and cookie purpose

	
## Emoji
```json
{
    "id": "integer",
    "link": "string",
    "name": "string",
    "shortcut": "string"
}
```

Emoji object describe the emoji available in the organisation and theire shortcuts.
Required apply only for creation of the object.

	
### Fields
Name | Type | Format | Description
--- | --- | --- | ---
id | integer | uint64 | id of the emoji
link<span title="required" class="required">&nbsp;*&nbsp;</span> | string |  | path to emoji into server
name<span title="required" class="required">&nbsp;*&nbsp;</span> | string |  | 
shortcut<span title="required" class="required">&nbsp;*&nbsp;</span> | string |  | 

	
## Folder
```json
{
    "id": "integer",
    "id_message": "integer",
    "link": "string",
    "name": "string",
    "type": "string"
}
```

A folder is a file larger than a simple message.

	
### Fields
Name | Type | Format | Description
--- | --- | --- | ---
id | integer | uint64 | id of the folder
id_message<span title="required" class="required">&nbsp;*&nbsp;</span> | integer | uint64 | id of the message folder is in
link<span title="required" class="required">&nbsp;*&nbsp;</span> | string |  | path to the folder in the server
name<span title="required" class="required">&nbsp;*&nbsp;</span> | string |  | folder name
type<span title="required" class="required">&nbsp;*&nbsp;</span> | string |  | type if the folder (extension, snippet)

	
## Member
```json
{
    "id_channel": "integer",
    "id_role": "integer",
    "id_user": "integer"
}
```

Member is the link between an User and a Channel. It also state the role of the user
in the channel if it is channel specific.

	
### Fields
Name | Type | Format | Description
--- | --- | --- | ---
id_channel<span title="required" class="required">&nbsp;*&nbsp;</span> | integer | uint64 | 
id_role | integer | uint64 | 
id_user<span title="required" class="required">&nbsp;*&nbsp;</span> | integer | uint64 | 

	
## Message
```json
{
    "content": "string",
    "date": "integer",
    "id": "integer",
    "id_channel": "integer",
    "id_user": "integer"
}
```

Message informations and content

	
### Fields
Name | Type | Format | Description
--- | --- | --- | ---
content | string |  | Content of the message
date<span title="required" class="required">&nbsp;*&nbsp;</span> | integer | int64 | Date the message was sent at
id | integer | uint64 | id of the message
id_channel<span title="required" class="required">&nbsp;*&nbsp;</span> | integer | uint64 | Channel reference id
id_user<span title="required" class="required">&nbsp;*&nbsp;</span> | integer | uint64 | User reference id

	
## Organisation
```json
{
    "avatar": "string",
    "description": "string",
    "docker_stack": "integer",
    "domain": "string",
    "id": "integer",
    "name": "string"
}
```

Describe organisation you are in. It is an unique object in the database.

	
### Fields
Name | Type | Format | Description
--- | --- | --- | ---
avatar | string |  | 
description | string |  | 
docker_stack<span title="required" class="required">&nbsp;*&nbsp;</span> | integer | int64 | Stack into docker swarm
domain | string |  | Domain name of the organisation
id | integer | uint64 | id of the organisation
name<span title="required" class="required">&nbsp;*&nbsp;</span> | string |  | 

	
## Parameter
```json
{
    "id": "integer",
    "local": "string",
    "sleep_end": "integer",
    "sleep_start": "integer",
    "time_zone": "string"
}
```

Global parameters to apply within organisation. unique object in database

	
### Fields
Name | Type | Format | Description
--- | --- | --- | ---
id | integer | uint64 | id of the parameter
local<span title="required" class="required">&nbsp;*&nbsp;</span> | string |  | Default langage
sleep_end<span title="required" class="required">&nbsp;*&nbsp;</span> | integer | int64 | Default end of non notification period
sleep_start<span title="required" class="required">&nbsp;*&nbsp;</span> | integer | int64 | Default start of non notification period
time_zone<span title="required" class="required">&nbsp;*&nbsp;</span> | string |  | Default time zone

	
## Role
```json
{
    "can_archive": "boolean",
    "can_invite": "boolean",
    "can_manage": "boolean",
    "can_manage_user": "boolean",
    "can_moderate": "boolean",
    "can_use_private": "boolean",
    "id": "integer",
    "name": "string"
}
```

Decribe rights linked to role

	
### Fields
Name | Type | Format | Description
--- | --- | --- | ---
can_archive<span title="required" class="required">&nbsp;*&nbsp;</span> | boolean |  | User can archive channels
can_invite<span title="required" class="required">&nbsp;*&nbsp;</span> | boolean |  | User can invite others to private channel or organisation
can_manage<span title="required" class="required">&nbsp;*&nbsp;</span> | boolean |  | User can manage organisation/channel parameters and data
can_manage_user<span title="required" class="required">&nbsp;*&nbsp;</span> | boolean |  | User can manage other organisation/channel user
can_moderate<span title="required" class="required">&nbsp;*&nbsp;</span> | boolean |  | User can moderate channels
can_use_private<span title="required" class="required">&nbsp;*&nbsp;</span> | boolean |  | User can use private channel
id | integer | uint64 | id of the role
name<span title="required" class="required">&nbsp;*&nbsp;</span> | string |  | 

	
## User
```json
{
    "avatar": "string",
    "deleted": "boolean",
    "email": "string",
    "email_verified": "boolean",
    "failed_attempts": "integer",
    "first_name": "string",
    "id": "integer",
    "id_role": "integer",
    "last_activity_at": "integer",
    "last_name": "string",
    "last_password_update": "integer",
    "last_update": "integer",
    "locale": "string",
    "nickname": "string",
    "password": "string",
    "username": "string",
    "web_id": "string"
}
```

An user is an account who have an access to a specific organisation. Each user is unique inside a given organisation, but users are not shared between
organisations. Required apply only for creation of the object.

	
### Fields
Name | Type | Format | Description
--- | --- | --- | ---
avatar | string |  | Avatar?used by user
deleted<span title="required" class="required">&nbsp;*&nbsp;</span> | boolean |  | User is deleted from organisation but still in database
email<span title="required" class="required">&nbsp;*&nbsp;</span> | string |  | User email
email_verified<span title="required" class="required">&nbsp;*&nbsp;</span> | boolean |  | State if email was verified
failed_attempts<span title="required" class="required">&nbsp;*&nbsp;</span> | integer | int64 | Number of attemps failed while loging in
first_name | string |  | First name
id | integer | uint64 | id of the user
id_role<span title="required" class="required">&nbsp;*&nbsp;</span> | integer | uint64 | Role key of user in the organisation
last_activity_at | integer | int64 | 
last_name | string |  | User Lastname
last_password_update<span title="required" class="required">&nbsp;*&nbsp;</span> | integer | int64 | Date of the last update of password from user
last_update | integer | int64 | Date of the last update from user
locale<span title="required" class="required">&nbsp;*&nbsp;</span> | string |  | User langage
nickname | string |  | User nickname
password<span title="required" class="required">&nbsp;*&nbsp;</span> | string |  | Encrypted user password
username<span title="required" class="required">&nbsp;*&nbsp;</span> | string |  | User name
web_id | string |  | web id for the user used only for cache and cookie purpose

	


<style>
    ul.enum {
        margin: 0;
        padding: 0 0 0 2px;
        list-style-position: inside;
    }
    .required {
        color: red;
        font-weight: bold;
    }
    a.pseudo {
        border-bottom:1px dashed; 
        text-decoration: none;
    }
</style>
