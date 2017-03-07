# PopCube API

*This repository contain the GO API for PopCube projet*

## Aims

Popcube api is a simple api to manage database communication for the chat project PopCube. It contains data models and methods to manage the database. The api in itself provide basics methods we are using to manage data.

## How to use

### By go get

To use this project, you can just make a `go get github.com/titouanfreville/popcubeapi`

### From source

From source, you just have to clone the project, then `run godep get` to install all the dependencies.

## Routes

<details>
<summary>`/`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/**
	- _GET_
		- [basicRoutes.func1](/api/api.go#L103)

</details>
<details>
<summary>`/avatar`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/avatar**
	- [github.com/goware/jwtauth.(*JwtAuth).Verifier-fm](/api/avatar_route.go#L23)
	- [Authenticator](https://github.com/goware/jwtauth/jwtauth.go#L196)
	- **/**
		- _GET_
			- [getAllAvatar](/api/avatar_route.go#L149)
		- _POST_
			- [newAvatar](/api/avatar_route.go#L186)

</details>
<details>
<summary>`/avatar/:avatarID/delete`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/avatar**
	- [github.com/goware/jwtauth.(*JwtAuth).Verifier-fm](/api/avatar_route.go#L23)
	- [Authenticator](https://github.com/goware/jwtauth/jwtauth.go#L196)
	- **/:avatarID**
		- [avatarContext](/api/avatar_route.go#L133)
		- **/delete**
			- _DELETE_
				- [deleteAvatar](/api/avatar_route.go#L239)

</details>
<details>
<summary>`/avatar/:avatarID/update`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/avatar**
	- [github.com/goware/jwtauth.(*JwtAuth).Verifier-fm](/api/avatar_route.go#L23)
	- [Authenticator](https://github.com/goware/jwtauth/jwtauth.go#L196)
	- **/:avatarID**
		- [avatarContext](/api/avatar_route.go#L133)
		- **/update**
			- _PUT_
				- [updateAvatar](/api/avatar_route.go#L212)

</details>
<details>
<summary>`/avatar/all`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/avatar**
	- [github.com/goware/jwtauth.(*JwtAuth).Verifier-fm](/api/avatar_route.go#L23)
	- [Authenticator](https://github.com/goware/jwtauth/jwtauth.go#L196)
	- **/all**
		- _GET_
			- [getAllAvatar](/api/avatar_route.go#L149)

</details>
<details>
<summary>`/avatar/link/:avatarLink`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/avatar**
	- [github.com/goware/jwtauth.(*JwtAuth).Verifier-fm](/api/avatar_route.go#L23)
	- [Authenticator](https://github.com/goware/jwtauth/jwtauth.go#L196)
	- **/link**
		- **/:avatarLink**
			- [avatarContext](/api/avatar_route.go#L133)
			- **/**
				- _GET_
					- [getAvatarFromLink](/api/avatar_route.go#L173)

</details>
<details>
<summary>`/avatar/name/:avatarName`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/avatar**
	- [github.com/goware/jwtauth.(*JwtAuth).Verifier-fm](/api/avatar_route.go#L23)
	- [Authenticator](https://github.com/goware/jwtauth/jwtauth.go#L196)
	- **/name**
		- **/:avatarName**
			- [avatarContext](/api/avatar_route.go#L133)
			- **/**
				- _GET_
					- [getAvatarFromName](/api/avatar_route.go#L160)

</details>
<details>
<summary>`/avatar/new`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/avatar**
	- [github.com/goware/jwtauth.(*JwtAuth).Verifier-fm](/api/avatar_route.go#L23)
	- [Authenticator](https://github.com/goware/jwtauth/jwtauth.go#L196)
	- **/new**
		- _POST_
			- [newAvatar](/api/avatar_route.go#L186)

</details>
<details>
<summary>`/channel`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/channel**
	- [github.com/goware/jwtauth.(*JwtAuth).Verifier-fm](/api/avatar_route.go#L23)
	- [Authenticator](https://github.com/goware/jwtauth/jwtauth.go#L196)
	- **/**
		- _GET_
			- [getAllChannel](/api/channel_route.go#L170)
		- _POST_
			- [newChannel](/api/channel_route.go#L224)

</details>
<details>
<summary>`/channel/:channelID/delete`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/channel**
	- [github.com/goware/jwtauth.(*JwtAuth).Verifier-fm](/api/avatar_route.go#L23)
	- [Authenticator](https://github.com/goware/jwtauth/jwtauth.go#L196)
	- **/:channelID**
		- [channelContext](/api/channel_route.go#L154)
		- **/delete**
			- _DELETE_
				- [deleteChannel](/api/channel_route.go#L277)

</details>
<details>
<summary>`/channel/:channelID/update`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/channel**
	- [github.com/goware/jwtauth.(*JwtAuth).Verifier-fm](/api/avatar_route.go#L23)
	- [Authenticator](https://github.com/goware/jwtauth/jwtauth.go#L196)
	- **/:channelID**
		- [channelContext](/api/channel_route.go#L154)
		- **/update**
			- _PUT_
				- [updateChannel](/api/channel_route.go#L250)

</details>
<details>
<summary>`/channel/all`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/channel**
	- [github.com/goware/jwtauth.(*JwtAuth).Verifier-fm](/api/avatar_route.go#L23)
	- [Authenticator](https://github.com/goware/jwtauth/jwtauth.go#L196)
	- **/all**
		- _GET_
			- [getAllChannel](/api/channel_route.go#L170)

</details>
<details>
<summary>`/channel/name/:channelName`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/channel**
	- [github.com/goware/jwtauth.(*JwtAuth).Verifier-fm](/api/avatar_route.go#L23)
	- [Authenticator](https://github.com/goware/jwtauth/jwtauth.go#L196)
	- **/name**
		- **/:channelName**
			- [channelContext](/api/channel_route.go#L154)
			- **/**
				- _GET_
					- [getChannelFromName](/api/channel_route.go#L206)

</details>
<details>
<summary>`/channel/new`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/channel**
	- [github.com/goware/jwtauth.(*JwtAuth).Verifier-fm](/api/avatar_route.go#L23)
	- [Authenticator](https://github.com/goware/jwtauth/jwtauth.go#L196)
	- **/new**
		- _POST_
			- [newChannel](/api/channel_route.go#L224)

</details>
<details>
<summary>`/channel/private`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/channel**
	- [github.com/goware/jwtauth.(*JwtAuth).Verifier-fm](/api/avatar_route.go#L23)
	- [Authenticator](https://github.com/goware/jwtauth/jwtauth.go#L196)
	- **/private**
		- _GET_
			- [getPrivateChannel](/api/channel_route.go#L194)

</details>
<details>
<summary>`/channel/public`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/channel**
	- [github.com/goware/jwtauth.(*JwtAuth).Verifier-fm](/api/avatar_route.go#L23)
	- [Authenticator](https://github.com/goware/jwtauth/jwtauth.go#L196)
	- **/public**
		- _GET_
			- [getPublicChannel](/api/channel_route.go#L182)

</details>
<details>
<summary>`/channel/type/:channelType`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/channel**
	- [github.com/goware/jwtauth.(*JwtAuth).Verifier-fm](/api/avatar_route.go#L23)
	- [Authenticator](https://github.com/goware/jwtauth/jwtauth.go#L196)
	- **/type**
		- **/:channelType**
			- [channelContext](/api/channel_route.go#L154)
			- **/**
				- _GET_
					- [getChannelFromType](/api/channel_route.go#L215)

</details>
<details>
<summary>`/emoji`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/emoji**
	- [github.com/goware/jwtauth.(*JwtAuth).Verifier-fm](/api/avatar_route.go#L23)
	- [Authenticator](https://github.com/goware/jwtauth/jwtauth.go#L196)
	- **/**
		- _GET_
			- [getAllEmoji](/api/emojis_route.go#L167)
		- _POST_
			- [newEmoji](/api/emojis_route.go#L206)

</details>
<details>
<summary>`/emoji/:emojiID/delete`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/emoji**
	- [github.com/goware/jwtauth.(*JwtAuth).Verifier-fm](/api/avatar_route.go#L23)
	- [Authenticator](https://github.com/goware/jwtauth/jwtauth.go#L196)
	- **/:emojiID**
		- [emojiContext](/api/emojis_route.go#L149)
		- **/delete**
			- _DELETE_
				- [deleteEmoji](/api/emojis_route.go#L259)

</details>
<details>
<summary>`/emoji/:emojiID/update`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/emoji**
	- [github.com/goware/jwtauth.(*JwtAuth).Verifier-fm](/api/avatar_route.go#L23)
	- [Authenticator](https://github.com/goware/jwtauth/jwtauth.go#L196)
	- **/:emojiID**
		- [emojiContext](/api/emojis_route.go#L149)
		- **/update**
			- _PUT_
				- [updateEmoji](/api/emojis_route.go#L232)

</details>
<details>
<summary>`/emoji/all`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/emoji**
	- [github.com/goware/jwtauth.(*JwtAuth).Verifier-fm](/api/avatar_route.go#L23)
	- [Authenticator](https://github.com/goware/jwtauth/jwtauth.go#L196)
	- **/all**
		- _GET_
			- [getAllEmoji](/api/emojis_route.go#L167)

</details>
<details>
<summary>`/emoji/link/:emojiLink`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/emoji**
	- [github.com/goware/jwtauth.(*JwtAuth).Verifier-fm](/api/avatar_route.go#L23)
	- [Authenticator](https://github.com/goware/jwtauth/jwtauth.go#L196)
	- **/link**
		- **/:emojiLink**
			- [emojiContext](/api/emojis_route.go#L149)
			- **/**
				- _GET_
					- [getEmojiFromLink](/api/emojis_route.go#L197)

</details>
<details>
<summary>`/emoji/name/:emojiName`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/emoji**
	- [github.com/goware/jwtauth.(*JwtAuth).Verifier-fm](/api/avatar_route.go#L23)
	- [Authenticator](https://github.com/goware/jwtauth/jwtauth.go#L196)
	- **/name**
		- **/:emojiName**
			- [emojiContext](/api/emojis_route.go#L149)
			- **/**
				- _GET_
					- [getEmojiFromName](/api/emojis_route.go#L179)

</details>
<details>
<summary>`/emoji/new`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/emoji**
	- [github.com/goware/jwtauth.(*JwtAuth).Verifier-fm](/api/avatar_route.go#L23)
	- [Authenticator](https://github.com/goware/jwtauth/jwtauth.go#L196)
	- **/new**
		- _POST_
			- [newEmoji](/api/emojis_route.go#L206)

</details>
<details>
<summary>`/emoji/shortcut/:emojiShortcut`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/emoji**
	- [github.com/goware/jwtauth.(*JwtAuth).Verifier-fm](/api/avatar_route.go#L23)
	- [Authenticator](https://github.com/goware/jwtauth/jwtauth.go#L196)
	- **/shortcut**
		- **/:emojiShortcut**
			- [emojiContext](/api/emojis_route.go#L149)
			- **/**
				- _GET_
					- [getEmojiFromShortcut](/api/emojis_route.go#L188)

</details>
<details>
<summary>`/folder`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/folder**
	- [github.com/goware/jwtauth.(*JwtAuth).Verifier-fm](/api/avatar_route.go#L23)
	- [Authenticator](https://github.com/goware/jwtauth/jwtauth.go#L196)
	- **/**
		- _GET_
			- [getAllFolder](/api/folder_route.go#L179)
		- _POST_
			- [newFolder](/api/folder_route.go#L240)

</details>
<details>
<summary>`/folder/:folderID/delete`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/folder**
	- [github.com/goware/jwtauth.(*JwtAuth).Verifier-fm](/api/avatar_route.go#L23)
	- [Authenticator](https://github.com/goware/jwtauth/jwtauth.go#L196)
	- **/:folderID**
		- [folderContext](/api/folder_route.go#L161)
		- **/delete**
			- _DELETE_
				- [deleteFolder](/api/folder_route.go#L293)

</details>
<details>
<summary>`/folder/:folderID/update`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/folder**
	- [github.com/goware/jwtauth.(*JwtAuth).Verifier-fm](/api/avatar_route.go#L23)
	- [Authenticator](https://github.com/goware/jwtauth/jwtauth.go#L196)
	- **/:folderID**
		- [folderContext](/api/folder_route.go#L161)
		- **/update**
			- _PUT_
				- [updateFolder](/api/folder_route.go#L266)

</details>
<details>
<summary>`/folder/all`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/folder**
	- [github.com/goware/jwtauth.(*JwtAuth).Verifier-fm](/api/avatar_route.go#L23)
	- [Authenticator](https://github.com/goware/jwtauth/jwtauth.go#L196)
	- **/all**
		- _GET_
			- [getAllFolder](/api/folder_route.go#L179)

</details>
<details>
<summary>`/folder/link/:folderLink`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/folder**
	- [github.com/goware/jwtauth.(*JwtAuth).Verifier-fm](/api/avatar_route.go#L23)
	- [Authenticator](https://github.com/goware/jwtauth/jwtauth.go#L196)
	- **/link**
		- **/:folderLink**
			- [folderContext](/api/folder_route.go#L161)
			- **/**
				- _GET_
					- [getFolderFromLink](/api/folder_route.go#L209)

</details>
<details>
<summary>`/folder/message`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/folder**
	- [github.com/goware/jwtauth.(*JwtAuth).Verifier-fm](/api/avatar_route.go#L23)
	- [Authenticator](https://github.com/goware/jwtauth/jwtauth.go#L196)
	- **/message**
		- _POST_
			- [getFolderFromMessage](/api/folder_route.go#L218)

</details>
<details>
<summary>`/folder/name/:folderName`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/folder**
	- [github.com/goware/jwtauth.(*JwtAuth).Verifier-fm](/api/avatar_route.go#L23)
	- [Authenticator](https://github.com/goware/jwtauth/jwtauth.go#L196)
	- **/name**
		- **/:folderName**
			- [folderContext](/api/folder_route.go#L161)
			- **/**
				- _GET_
					- [getFolderFromName](/api/folder_route.go#L191)

</details>
<details>
<summary>`/folder/new`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/folder**
	- [github.com/goware/jwtauth.(*JwtAuth).Verifier-fm](/api/avatar_route.go#L23)
	- [Authenticator](https://github.com/goware/jwtauth/jwtauth.go#L196)
	- **/new**
		- _POST_
			- [newFolder](/api/folder_route.go#L240)

</details>
<details>
<summary>`/folder/type/:folderType`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/folder**
	- [github.com/goware/jwtauth.(*JwtAuth).Verifier-fm](/api/avatar_route.go#L23)
	- [Authenticator](https://github.com/goware/jwtauth/jwtauth.go#L196)
	- **/type**
		- **/:folderType**
			- [folderContext](/api/folder_route.go#L161)
			- **/**
				- _GET_
					- [getFolderFromType](/api/folder_route.go#L200)

</details>
<details>
<summary>`/heartbeat`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/heartbeat**
	- _GET_
		- [basicRoutes.func3](/api/api.go#L118)

</details>
<details>
<summary>`/login`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/login**
	- _POST_
		- [loginMiddleware](/api/api.go#L146)

</details>
<details>
<summary>`/message`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/message**
	- [github.com/goware/jwtauth.(*JwtAuth).Verifier-fm](/api/avatar_route.go#L23)
	- [Authenticator](https://github.com/goware/jwtauth/jwtauth.go#L196)
	- **/**
		- _GET_
			- [getAllMessage](/api/message_route.go#L154)
		- _POST_
			- [newMessage](/api/message_route.go#L219)

</details>
<details>
<summary>`/message/:messageID/delete`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/message**
	- [github.com/goware/jwtauth.(*JwtAuth).Verifier-fm](/api/avatar_route.go#L23)
	- [Authenticator](https://github.com/goware/jwtauth/jwtauth.go#L196)
	- **/:messageID**
		- [messageContext](/api/message_route.go#L140)
		- **/delete**
			- _DELETE_
				- [deleteMessageFunction](/api/message_route.go#L272)

</details>
<details>
<summary>`/message/:messageID/update`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/message**
	- [github.com/goware/jwtauth.(*JwtAuth).Verifier-fm](/api/avatar_route.go#L23)
	- [Authenticator](https://github.com/goware/jwtauth/jwtauth.go#L196)
	- **/:messageID**
		- [messageContext](/api/message_route.go#L140)
		- **/update**
			- _PUT_
				- [updateMessage](/api/message_route.go#L245)

</details>
<details>
<summary>`/message/all`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/message**
	- [github.com/goware/jwtauth.(*JwtAuth).Verifier-fm](/api/avatar_route.go#L23)
	- [Authenticator](https://github.com/goware/jwtauth/jwtauth.go#L196)
	- **/all**
		- _GET_
			- [getAllMessage](/api/message_route.go#L154)

</details>
<details>
<summary>`/message/channel`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/message**
	- [github.com/goware/jwtauth.(*JwtAuth).Verifier-fm](/api/avatar_route.go#L23)
	- [Authenticator](https://github.com/goware/jwtauth/jwtauth.go#L196)
	- **/channel**
		- _POST_
			- [getMessageFromChannel](/api/message_route.go#L197)

</details>
<details>
<summary>`/message/creator`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/message**
	- [github.com/goware/jwtauth.(*JwtAuth).Verifier-fm](/api/avatar_route.go#L23)
	- [Authenticator](https://github.com/goware/jwtauth/jwtauth.go#L196)
	- **/creator**
		- _POST_
			- [getMessageFromUser](/api/message_route.go#L175)

</details>
<details>
<summary>`/message/date/:messageDate`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/message**
	- [github.com/goware/jwtauth.(*JwtAuth).Verifier-fm](/api/avatar_route.go#L23)
	- [Authenticator](https://github.com/goware/jwtauth/jwtauth.go#L196)
	- **/date**
		- **/:messageDate**
			- [messageContext](/api/message_route.go#L140)
			- **/**
				- _GET_
					- [getMessageFromDate](/api/message_route.go#L166)

</details>
<details>
<summary>`/message/new`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/message**
	- [github.com/goware/jwtauth.(*JwtAuth).Verifier-fm](/api/avatar_route.go#L23)
	- [Authenticator](https://github.com/goware/jwtauth/jwtauth.go#L196)
	- **/new**
		- _POST_
			- [newMessage](/api/message_route.go#L219)

</details>
<details>
<summary>`/organisation`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/organisation**
	- [github.com/goware/jwtauth.(*JwtAuth).Verifier-fm](/api/avatar_route.go#L23)
	- [Authenticator](https://github.com/goware/jwtauth/jwtauth.go#L196)
	- **/**
		- _GET_
			- [getAllOrganisation](/api/organisation_route.go#L98)
		- _POST_
			- [newOrganisation](/api/organisation_route.go#L110)

</details>
<details>
<summary>`/organisation/:organisationID/update`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/organisation**
	- [github.com/goware/jwtauth.(*JwtAuth).Verifier-fm](/api/avatar_route.go#L23)
	- [Authenticator](https://github.com/goware/jwtauth/jwtauth.go#L196)
	- **/:organisationID**
		- [organisationContext](/api/organisation_route.go#L86)
		- **/update**
			- _PUT_
				- [updateOrganisation](/api/organisation_route.go#L136)

</details>
<details>
<summary>`/organisation/all`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/organisation**
	- [github.com/goware/jwtauth.(*JwtAuth).Verifier-fm](/api/avatar_route.go#L23)
	- [Authenticator](https://github.com/goware/jwtauth/jwtauth.go#L196)
	- **/all**
		- _GET_
			- [getAllOrganisation](/api/organisation_route.go#L98)

</details>
<details>
<summary>`/organisation/new`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/organisation**
	- [github.com/goware/jwtauth.(*JwtAuth).Verifier-fm](/api/avatar_route.go#L23)
	- [Authenticator](https://github.com/goware/jwtauth/jwtauth.go#L196)
	- **/new**
		- _POST_
			- [newOrganisation](/api/organisation_route.go#L110)

</details>
<details>
<summary>`/panic`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/panic**
	- _GET_
		- [basicRoutes.func4](/api/api.go#L128)

</details>
<details>
<summary>`/parameter`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/parameter**
	- [github.com/goware/jwtauth.(*JwtAuth).Verifier-fm](/api/avatar_route.go#L23)
	- [Authenticator](https://github.com/goware/jwtauth/jwtauth.go#L196)
	- **/**
		- _POST_
			- [newParameter](/api/parameter_route.go#L111)
		- _GET_
			- [getAllParameter](/api/parameter_route.go#L99)

</details>
<details>
<summary>`/parameter/:parameterID/update`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/parameter**
	- [github.com/goware/jwtauth.(*JwtAuth).Verifier-fm](/api/avatar_route.go#L23)
	- [Authenticator](https://github.com/goware/jwtauth/jwtauth.go#L196)
	- **/:parameterID**
		- [parameterContext](/api/parameter_route.go#L87)
		- **/update**
			- _PUT_
				- [updateParameter](/api/parameter_route.go#L137)

</details>
<details>
<summary>`/parameter/all`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/parameter**
	- [github.com/goware/jwtauth.(*JwtAuth).Verifier-fm](/api/avatar_route.go#L23)
	- [Authenticator](https://github.com/goware/jwtauth/jwtauth.go#L196)
	- **/all**
		- _GET_
			- [getAllParameter](/api/parameter_route.go#L99)

</details>
<details>
<summary>`/parameter/new`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/parameter**
	- [github.com/goware/jwtauth.(*JwtAuth).Verifier-fm](/api/avatar_route.go#L23)
	- [Authenticator](https://github.com/goware/jwtauth/jwtauth.go#L196)
	- **/new**
		- _POST_
			- [newParameter](/api/parameter_route.go#L111)

</details>
<details>
<summary>`/ping`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/ping**
	- _GET_
		- [basicRoutes.func2](/api/api.go#L115)

</details>
<details>
<summary>`/role`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/role**
	- [github.com/goware/jwtauth.(*JwtAuth).Verifier-fm](/api/avatar_route.go#L23)
	- [Authenticator](https://github.com/goware/jwtauth/jwtauth.go#L196)
	- **/**
		- _POST_
			- [newRole](/api/role_route.go#L183)
		- _GET_
			- [getAllRole](/api/role_route.go#L140)

</details>
<details>
<summary>`/role/:roleID/delete`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/role**
	- [github.com/goware/jwtauth.(*JwtAuth).Verifier-fm](/api/avatar_route.go#L23)
	- [Authenticator](https://github.com/goware/jwtauth/jwtauth.go#L196)
	- **/:roleID**
		- [roleContext](/api/role_route.go#L126)
		- **/delete**
			- _DELETE_
				- [deleteRole](/api/role_route.go#L236)

</details>
<details>
<summary>`/role/:roleID/update`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/role**
	- [github.com/goware/jwtauth.(*JwtAuth).Verifier-fm](/api/avatar_route.go#L23)
	- [Authenticator](https://github.com/goware/jwtauth/jwtauth.go#L196)
	- **/:roleID**
		- [roleContext](/api/role_route.go#L126)
		- **/update**
			- _PUT_
				- [updateRole](/api/role_route.go#L209)

</details>
<details>
<summary>`/role/all`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/role**
	- [github.com/goware/jwtauth.(*JwtAuth).Verifier-fm](/api/avatar_route.go#L23)
	- [Authenticator](https://github.com/goware/jwtauth/jwtauth.go#L196)
	- **/all**
		- _GET_
			- [getAllRole](/api/role_route.go#L140)

</details>
<details>
<summary>`/role/name/:roleName`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/role**
	- [github.com/goware/jwtauth.(*JwtAuth).Verifier-fm](/api/avatar_route.go#L23)
	- [Authenticator](https://github.com/goware/jwtauth/jwtauth.go#L196)
	- **/name**
		- **/:roleName**
			- [roleContext](/api/role_route.go#L126)
			- **/**
				- _GET_
					- [getRoleFromName](/api/role_route.go#L152)

</details>
<details>
<summary>`/role/new`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/role**
	- [github.com/goware/jwtauth.(*JwtAuth).Verifier-fm](/api/avatar_route.go#L23)
	- [Authenticator](https://github.com/goware/jwtauth/jwtauth.go#L196)
	- **/new**
		- _POST_
			- [newRole](/api/role_route.go#L183)

</details>
<details>
<summary>`/role/rights`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/role**
	- [github.com/goware/jwtauth.(*JwtAuth).Verifier-fm](/api/avatar_route.go#L23)
	- [Authenticator](https://github.com/goware/jwtauth/jwtauth.go#L196)
	- **/rights**
		- _POST_
			- [getRoleFromRight](/api/role_route.go#L161)

</details>
<details>
<summary>`/user`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/user**
	- [github.com/goware/jwtauth.(*JwtAuth).Verifier-fm](/api/avatar_route.go#L23)
	- [Authenticator](https://github.com/goware/jwtauth/jwtauth.go#L196)
	- **/**
		- _GET_
			- [getAllUser](/api/user_route.go#L233)
		- _POST_
			- [newUser](/api/user_route.go#L333)

</details>
<details>
<summary>`/user/:userID/delete`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/user**
	- [github.com/goware/jwtauth.(*JwtAuth).Verifier-fm](/api/avatar_route.go#L23)
	- [Authenticator](https://github.com/goware/jwtauth/jwtauth.go#L196)
	- **/:userID**
		- [userContext](/api/user_route.go#L209)
		- **/delete**
			- _DELETE_
				- [deleteUser](/api/user_route.go#L386)

</details>
<details>
<summary>`/user/:userID/update`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/user**
	- [github.com/goware/jwtauth.(*JwtAuth).Verifier-fm](/api/avatar_route.go#L23)
	- [Authenticator](https://github.com/goware/jwtauth/jwtauth.go#L196)
	- **/:userID**
		- [userContext](/api/user_route.go#L209)
		- **/update**
			- _PUT_
				- [updateUser](/api/user_route.go#L359)

</details>
<details>
<summary>`/user/all`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/user**
	- [github.com/goware/jwtauth.(*JwtAuth).Verifier-fm](/api/avatar_route.go#L23)
	- [Authenticator](https://github.com/goware/jwtauth/jwtauth.go#L196)
	- **/all**
		- _GET_
			- [getAllUser](/api/user_route.go#L233)

</details>
<details>
<summary>`/user/date`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/user**
	- [github.com/goware/jwtauth.(*JwtAuth).Verifier-fm](/api/avatar_route.go#L23)
	- [Authenticator](https://github.com/goware/jwtauth/jwtauth.go#L196)
	- **/date**
		- _GET_
			- [getOrderedByDate](/api/user_route.go#L302)

</details>
<details>
<summary>`/user/deleted`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/user**
	- [github.com/goware/jwtauth.(*JwtAuth).Verifier-fm](/api/avatar_route.go#L23)
	- [Authenticator](https://github.com/goware/jwtauth/jwtauth.go#L196)
	- **/deleted**
		- _GET_
			- [getDeletedUser](/api/user_route.go#L245)

</details>
<details>
<summary>`/user/email/:userEmail`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/user**
	- [github.com/goware/jwtauth.(*JwtAuth).Verifier-fm](/api/avatar_route.go#L23)
	- [Authenticator](https://github.com/goware/jwtauth/jwtauth.go#L196)
	- **/email**
		- **/:userEmail**
			- [userContext](/api/user_route.go#L209)
			- **/**
				- _GET_
					- [getUserFromEmail](/api/user_route.go#L293)

</details>
<details>
<summary>`/user/firstname/:firstName`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/user**
	- [github.com/goware/jwtauth.(*JwtAuth).Verifier-fm](/api/avatar_route.go#L23)
	- [Authenticator](https://github.com/goware/jwtauth/jwtauth.go#L196)
	- **/firstname**
		- **/:firstName**
			- [userContext](/api/user_route.go#L209)
			- **/**
				- _GET_
					- [getUserFromFirstName](/api/user_route.go#L275)

</details>
<details>
<summary>`/user/lastname/:lastName`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/user**
	- [github.com/goware/jwtauth.(*JwtAuth).Verifier-fm](/api/avatar_route.go#L23)
	- [Authenticator](https://github.com/goware/jwtauth/jwtauth.go#L196)
	- **/lastname**
		- **/:lastName**
			- [userContext](/api/user_route.go#L209)
			- **/**
				- _GET_
					- [getUserFromLastName](/api/user_route.go#L284)

</details>
<details>
<summary>`/user/new`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/user**
	- [github.com/goware/jwtauth.(*JwtAuth).Verifier-fm](/api/avatar_route.go#L23)
	- [Authenticator](https://github.com/goware/jwtauth/jwtauth.go#L196)
	- **/new**
		- _POST_
			- [newUser](/api/user_route.go#L333)

</details>
<details>
<summary>`/user/nickname/:nickName`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/user**
	- [github.com/goware/jwtauth.(*JwtAuth).Verifier-fm](/api/avatar_route.go#L23)
	- [Authenticator](https://github.com/goware/jwtauth/jwtauth.go#L196)
	- **/nickname**
		- **/:nickName**
			- [userContext](/api/user_route.go#L209)
			- **/**
				- _GET_
					- [getUserFromNickName](/api/user_route.go#L266)

</details>
<details>
<summary>`/user/role`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/user**
	- [github.com/goware/jwtauth.(*JwtAuth).Verifier-fm](/api/avatar_route.go#L23)
	- [Authenticator](https://github.com/goware/jwtauth/jwtauth.go#L196)
	- **/role**
		- _POST_
			- [getUserFromRole](/api/user_route.go#L311)

</details>
<details>
<summary>`/user/username/:userName`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/user**
	- [github.com/goware/jwtauth.(*JwtAuth).Verifier-fm](/api/avatar_route.go#L23)
	- [Authenticator](https://github.com/goware/jwtauth/jwtauth.go#L196)
	- **/username**
		- **/:userName**
			- [userContext](/api/user_route.go#L209)
			- **/**
				- _GET_
					- [getUserFromName](/api/user_route.go#L257)

</details>

## Models

### Avatar

The avatar model is used to store Avatar locations and names.

#### Structure

| Row      | Type   | Parameters                 | Database Column | JSON Key  |
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

##### AvatarFromJSON

AvatarFromJSON is a function who will try to parse a JSON object as Go Avatar struct.

Usage: `AvatarFromJSON(strings.NewReader(json))`

##### AvatarListToJSON

AvatarListToJSON is a function to convert the Go Avatar struct slice in JSON object. 

Usage: `AvatarListToJSON(avatarList)`

##### AvatarListFromJSON

AvatarListFromJSON is a function who will try to parse a JSON object as Go Avatar struct slice.

Usage: `AvatarListFromJSON(strings.NewReader(json))`
