# PopCube API
*This repository contain the GO API for PopCube projet*

## Aims

Popcube api is a simple api to manage database communication for the chat project PopCube. It contains data models and methods to manage the database. The api in itself provide basics methods we are using to manage data.

## How to use

### By go get

To use this project, you can just make a `go get github.com/titouanfreville/popcubeapi`

### From source

From source, you just have to clone the project, then `run godep get` to install all the dependencies.

## Models

### Avatar

The avatar model is used to store Avatar locations and names.

#### Structure

| Row      | Type   | Parameters                 | Database Column | JSON Key  |
| -------- | ------ | -------------------------- | --------------- | --------- |
| IDAvatar | uint64 | primary key, autoincrement | idAvatar        | UNDEFINED |
| Name     | string | not null, unique           | name            | name      |
| Link     | string | not null, unique           | link            | link      |

#### Functions

##### IsValid

IsValid is a function to check the integrity of the provided avatar before sending it to the database. It will ensure that the link is not empty as well as the name.

Usage: `avatar.IsValid()`

##### ToJSON

ToJSON is a function to convert the Go Avatar struct in JSON object. 

Usage: `avatar.ToJSON()`

##### AvatarFromJSON

AvatarFromJSON is a function who will try to parse a JSON object as Go Avatar struct.

Usage: `AvatarFromJSON(strings.NewReader(json))`

##### AvatarListToJSON

AvatarListToJSON is a function to convert the Go Avatar struct slice in JSON object. 

Usage: `AvatarListToJSON(avatarList)`

##### AvatarListFromJSON

AvatarListFromJSON is a function who will try to parse a JSON object as Go Avatar struct slice.

Usage: `AvatarListFromJSON(strings.NewReader(json))`
