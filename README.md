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
		- [basicRoutes.func1](/api/api.go#L44)

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
	- **/:avatarID**
		- [avatarContext](/api/avatar_route.go#L42)
		- **/delete**
			- _DELETE_
				- [deleteAvatar](/api/avatar_route.go#L141)

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
	- **/:avatarID**
		- [avatarContext](/api/avatar_route.go#L42)
		- **/update**
			- _PUT_
				- [updateAvatar](/api/avatar_route.go#L114)

</details>
<details>
<summary>`/avatar/get`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/avatar**
	- **/get**
		- **/**
			- _GET_
				- [getAllAvatar](/api/avatar_route.go#L58)

</details>
<details>
<summary>`/avatar/get/all`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/avatar**
	- **/get**
		- **/all**
			- _GET_
				- [getAllAvatar](/api/avatar_route.go#L58)

</details>
<details>
<summary>`/avatar/get/fromlink/:avatarLink`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/avatar**
	- **/get**
		- **/fromlink**
			- **/:avatarLink**
				- [avatarContext](/api/avatar_route.go#L42)
				- **/**
					- _GET_
						- [getAvatarFromLink](/api/avatar_route.go#L79)

</details>
<details>
<summary>`/avatar/get/fromname/:avatarName`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/avatar**
	- **/get**
		- **/fromname**
			- **/:avatarName**
				- [avatarContext](/api/avatar_route.go#L42)
				- **/**
					- _GET_
						- [getAvatarFromName](/api/avatar_route.go#L70)

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
	- **/new**
		- _POST_
			- [newAvatar](/api/avatar_route.go#L88)

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
	- **/:channelID**
		- [channelContext](/api/channel_route.go#L44)
		- **/delete**
			- _DELETE_
				- [deleteChannel](/api/channel_route.go#L169)

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
	- **/:channelID**
		- [channelContext](/api/channel_route.go#L44)
		- **/update**
			- _PUT_
				- [updateChannel](/api/channel_route.go#L142)

</details>
<details>
<summary>`/channel/get`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/channel**
	- **/get**
		- **/**
			- _GET_
				- [getAllChannel](/api/channel_route.go#L62)

</details>
<details>
<summary>`/channel/get/all`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/channel**
	- **/get**
		- **/all**
			- _GET_
				- [getAllChannel](/api/channel_route.go#L62)

</details>
<details>
<summary>`/channel/get/fromname/:channelName`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/channel**
	- **/get**
		- **/fromname**
			- **/:channelName**
				- [channelContext](/api/channel_route.go#L44)
				- **/**
					- _GET_
						- [getChannelFromName](/api/channel_route.go#L98)

</details>
<details>
<summary>`/channel/get/fromtype/:channelType`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/channel**
	- **/get**
		- **/fromtype**
			- **/:channelType**
				- [channelContext](/api/channel_route.go#L44)
				- **/**
					- _GET_
						- [getChannelFromType](/api/channel_route.go#L107)

</details>
<details>
<summary>`/channel/get/private`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/channel**
	- **/get**
		- **/private**
			- _GET_
				- [getPrivateChannel](/api/channel_route.go#L86)

</details>
<details>
<summary>`/channel/get/public`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/channel**
	- **/get**
		- **/public**
			- _GET_
				- [getPublicChannel](/api/channel_route.go#L74)

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
	- **/new**
		- _POST_
			- [newChannel](/api/channel_route.go#L116)

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
	- **/:emojiID**
		- [emojiContext](/api/emojis_route.go#L48)
		- **/delete**
			- _DELETE_
				- [deleteEmoji](/api/emojis_route.go#L158)

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
	- **/:emojiID**
		- [emojiContext](/api/emojis_route.go#L48)
		- **/update**
			- _PUT_
				- [updateEmoji](/api/emojis_route.go#L131)

</details>
<details>
<summary>`/emoji/get`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/emoji**
	- **/get**
		- **/**
			- _GET_
				- [getAllEmoji](/api/emojis_route.go#L66)

</details>
<details>
<summary>`/emoji/get/all`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/emoji**
	- **/get**
		- **/all**
			- _GET_
				- [getAllEmoji](/api/emojis_route.go#L66)

</details>
<details>
<summary>`/emoji/get/fromlink/:emojiLink`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/emoji**
	- **/get**
		- **/fromlink**
			- **/:emojiLink**
				- [emojiContext](/api/emojis_route.go#L48)
				- **/**
					- _GET_
						- [getEmojiFromLink](/api/emojis_route.go#L96)

</details>
<details>
<summary>`/emoji/get/fromname/:emojiName`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/emoji**
	- **/get**
		- **/fromname**
			- **/:emojiName**
				- [emojiContext](/api/emojis_route.go#L48)
				- **/**
					- _GET_
						- [getEmojiFromName](/api/emojis_route.go#L78)

</details>
<details>
<summary>`/emoji/get/fromshortcut/:emojiShortcut`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/emoji**
	- **/get**
		- **/fromshortcut**
			- **/:emojiShortcut**
				- [emojiContext](/api/emojis_route.go#L48)
				- **/**
					- _GET_
						- [getEmojiFromShortcut](/api/emojis_route.go#L87)

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
	- **/new**
		- _POST_
			- [newEmoji](/api/emojis_route.go#L105)

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
	- **/:folderID**
		- [folderContext](/api/folder_route.go#L49)
		- **/delete**
			- _DELETE_
				- [deleteFolder](/api/folder_route.go#L181)

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
	- **/:folderID**
		- [folderContext](/api/folder_route.go#L49)
		- **/update**
			- _PUT_
				- [updateFolder](/api/folder_route.go#L154)

</details>
<details>
<summary>`/folder/get`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/folder**
	- **/get**
		- **/**
			- _GET_
				- [getAllFolder](/api/folder_route.go#L67)

</details>
<details>
<summary>`/folder/get/all`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/folder**
	- **/get**
		- **/all**
			- _GET_
				- [getAllFolder](/api/folder_route.go#L67)

</details>
<details>
<summary>`/folder/get/foldername/:folderName`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/folder**
	- **/get**
		- **/foldername**
			- **/:folderName**
				- [folderContext](/api/folder_route.go#L49)
				- **/**
					- _GET_
						- [getFolderFromName](/api/folder_route.go#L79)

</details>
<details>
<summary>`/folder/get/link/:folderLink`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/folder**
	- **/get**
		- **/link**
			- **/:folderLink**
				- [folderContext](/api/folder_route.go#L49)
				- **/**
					- _GET_
						- [getFolderFromLink](/api/folder_route.go#L97)

</details>
<details>
<summary>`/folder/get/message`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/folder**
	- **/get**
		- **/message**
			- _POST_
				- [getFolderFromMessage](/api/folder_route.go#L106)

</details>
<details>
<summary>`/folder/get/type/:folderType`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/folder**
	- **/get**
		- **/type**
			- **/:folderType**
				- [folderContext](/api/folder_route.go#L49)
				- **/**
					- _GET_
						- [getFolderFromType](/api/folder_route.go#L88)

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
	- **/new**
		- _POST_
			- [newFolder](/api/folder_route.go#L128)

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
	- **/:messageID**
		- [messageContext](/api/messge_route.go#L38)
		- **/delete**
			- _DELETE_
				- [deleteMessage](/api/messge_route.go#L170)

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
	- **/:messageID**
		- [messageContext](/api/messge_route.go#L38)
		- **/update**
			- _PUT_
				- [updateMessage](/api/messge_route.go#L143)

</details>
<details>
<summary>`/message/get`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/message**
	- **/get**
		- **/**
			- _GET_
				- [getAllMessage](/api/messge_route.go#L52)

</details>
<details>
<summary>`/message/get/all`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/message**
	- **/get**
		- **/all**
			- _GET_
				- [getAllMessage](/api/messge_route.go#L52)

</details>
<details>
<summary>`/message/get/channel`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/message**
	- **/get**
		- **/channel**
			- _POST_
				- [getMessageFromChannel](/api/messge_route.go#L95)

</details>
<details>
<summary>`/message/get/creator`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/message**
	- **/get**
		- **/creator**
			- _POST_
				- [getMessageFromUser](/api/messge_route.go#L73)

</details>
<details>
<summary>`/message/get/date/:date`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/message**
	- **/get**
		- **/date**
			- **/:date**
				- [messageContext](/api/messge_route.go#L38)
				- **/**
					- _GET_
						- [getMessageFromDate](/api/messge_route.go#L64)

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
	- **/new**
		- _POST_
			- [newMessage](/api/messge_route.go#L117)

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
	- **/:organisationID**
		- [organisationContext](/api/organisation_route.go#L29)
		- **/update**
			- _PUT_
				- [updateOrganisation](/api/organisation_route.go#L79)

</details>
<details>
<summary>`/organisation/get`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/organisation**
	- **/get**
		- **/**
			- _GET_
				- [getAllOrganisation](/api/organisation_route.go#L41)

</details>
<details>
<summary>`/organisation/get/all`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/organisation**
	- **/get**
		- **/all**
			- _GET_
				- [getAllOrganisation](/api/organisation_route.go#L41)

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
	- **/new**
		- _POST_
			- [newOrganisation](/api/organisation_route.go#L53)

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
		- [basicRoutes.func3](/api/api.go#L52)

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
	- **/:parameterID**
		- [parameterContext](/api/parameter_route.go#L29)
		- **/update**
			- _PUT_
				- [updateParameter](/api/parameter_route.go#L79)

</details>
<details>
<summary>`/parameter/get`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/parameter**
	- **/get**
		- **/**
			- _GET_
				- [getAllParameter](/api/parameter_route.go#L41)

</details>
<details>
<summary>`/parameter/get/all`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/parameter**
	- **/get**
		- **/all**
			- _GET_
				- [getAllParameter](/api/parameter_route.go#L41)

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
	- **/new**
		- _POST_
			- [newParameter](/api/parameter_route.go#L53)

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
		- [basicRoutes.func2](/api/api.go#L48)

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
	- **/:roleID**
		- [roleContext](/api/role_route.go#L37)
		- **/delete**
			- _DELETE_
				- [deleteRole](/api/role_route.go#L147)

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
	- **/:roleID**
		- [roleContext](/api/role_route.go#L37)
		- **/update**
			- _PUT_
				- [updateRole](/api/role_route.go#L120)

</details>
<details>
<summary>`/role/get`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/role**
	- **/get**
		- **/**
			- _GET_
				- [getAllRole](/api/role_route.go#L51)

</details>
<details>
<summary>`/role/get/all`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/role**
	- **/get**
		- **/all**
			- _GET_
				- [getAllRole](/api/role_route.go#L51)

</details>
<details>
<summary>`/role/get/fromname/:roleName`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/role**
	- **/get**
		- **/fromname**
			- **/:roleName**
				- [roleContext](/api/role_route.go#L37)
				- **/**
					- _GET_
						- [getRoleFromName](/api/role_route.go#L63)

</details>
<details>
<summary>`/role/get/fromrights`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/role**
	- **/get**
		- **/fromrights**
			- _POST_
				- [getRoleFromRight](/api/role_route.go#L72)

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
	- **/new**
		- _POST_
			- [newRole](/api/role_route.go#L94)

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
	- **/:userID**
		- [userContext](/api/user_route.go#L68)
		- **/delete**
			- _DELETE_
				- [deleteUser](/api/user_route.go#L245)

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
	- **/:userID**
		- [userContext](/api/user_route.go#L68)
		- **/update**
			- _PUT_
				- [updateUser](/api/user_route.go#L218)

</details>
<details>
<summary>`/user/get`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/user**
	- **/get**
		- **/**
			- _GET_
				- [getAllUser](/api/user_route.go#L92)

</details>
<details>
<summary>`/user/get/all`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/user**
	- **/get**
		- **/all**
			- _GET_
				- [getAllUser](/api/user_route.go#L92)

</details>
<details>
<summary>`/user/get/date/:date`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/user**
	- **/get**
		- **/date**
			- **/:date**
				- [userContext](/api/user_route.go#L68)
				- **/**
					- _GET_
						- [getUserFromDate](/api/user_route.go#L161)

</details>
<details>
<summary>`/user/get/deleted`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/user**
	- **/get**
		- **/deleted**
			- _GET_
				- [getDeletedUser](/api/user_route.go#L104)

</details>
<details>
<summary>`/user/get/email/:userEmail`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/user**
	- **/get**
		- **/email**
			- **/:userEmail**
				- [userContext](/api/user_route.go#L68)
				- **/**
					- _GET_
						- [getUserFromEmail](/api/user_route.go#L152)

</details>
<details>
<summary>`/user/get/firstname/:firstName`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/user**
	- **/get**
		- **/firstname**
			- **/:firstName**
				- [userContext](/api/user_route.go#L68)
				- **/**
					- _GET_
						- [getUserFromFirstName](/api/user_route.go#L134)

</details>
<details>
<summary>`/user/get/lastname/:lastName`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/user**
	- **/get**
		- **/lastname**
			- **/:lastName**
				- [userContext](/api/user_route.go#L68)
				- **/**
					- _GET_
						- [getUserFromLastName](/api/user_route.go#L143)

</details>
<details>
<summary>`/user/get/nickname/:nickName`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/user**
	- **/get**
		- **/nickname**
			- **/:nickName**
				- [userContext](/api/user_route.go#L68)
				- **/**
					- _GET_
						- [getUserFromNickName](/api/user_route.go#L125)

</details>
<details>
<summary>`/user/get/role`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/user**
	- **/get**
		- **/role**
			- _POST_
				- [getUserFromRole](/api/user_route.go#L170)

</details>
<details>
<summary>`/user/get/username/:userName`</summary>

- [RequestID](https://github.com/pressly/chi/middleware/request_id.go#L63)
- [RealIP](https://github.com/pressly/chi/middleware/realip.go#L29)
- [Logger](https://github.com/pressly/chi/middleware/logger.go#L26)
- [Recoverer](https://github.com/pressly/chi/middleware/recoverer.go#L16)
- [StripSlashes](https://github.com/pressly/chi/middleware/strip.go#L12)
- [Timeout.func1](https://github.com/pressly/chi/middleware/timeout.go#L33)
- [Heartbeat.func1](https://github.com/pressly/chi/middleware/heartbeat.go#L13)
- [CloseNotify](https://github.com/pressly/chi/middleware/closenotify17.go#L16)
- **/user**
	- **/get**
		- **/username**
			- **/:userName**
				- [userContext](/api/user_route.go#L68)
				- **/**
					- _GET_
						- [getUserFromName](/api/user_route.go#L116)

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
	- **/new**
		- _POST_
			- [newUser](/api/user_route.go#L192)

</details>

Total # of routes: 72

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
