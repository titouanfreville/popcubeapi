-- ROLE INITIALISATION ------------------------------------------------------------------------
INSERT INTO roles (roleName, canUsePrivate, canModerate, canArchive, canInvite, canManage, canManageUser) 
VALUES ("owner", true, true, true, true, true, true);

INSERT INTO roles (roleName, canUsePrivate, canModerate, canArchive, canInvite, canManage, canManageUser) 
VALUES ("admin", true, true, true, true, true, true);

INSERT INTO roles (roleName, canUsePrivate, canModerate, canArchive, canInvite, canManage, canManageUser) 
VALUES ("standard", true, true, true, false, false, false);

INSERT INTO roles (roleName, canUsePrivate, canModerate, canArchive, canInvite, canManage, canManageUser) 
VALUES ("guest", false, false, false, false, false, false);

-- CHANNEL INITIALISATION ------------------------------------------------------------------------
INSERT INTO channels (webId, channelName, type, lastUpdate, private, avatar, description, subject) 
VALUES ("generaltextchannel", "general", "text", 1, false, "defaultAvatar", "Speak on general subjects", "General");

INSERT INTO channels (webId, channelName, type, lastUpdate, private, avatar,description) 
VALUES ("randomtextchannel", "random", "text", 1, false, "defaultAvatar","Speak about any thing");

INSERT INTO channels (webId, channelName, type, lastUpdate, private, avatar) 
VALUES ("generalvocchannel", "general - voc", "audio", 1, false, "defaultAvatar");

INSERT INTO channels (webId, channelName, type, lastUpdate, private, avatar) 
VALUES ("randomvocchannel", "random - voc", "audio", 1, false, "defaultAvatar");

INSERT INTO channels (webId, channelName, type, lastUpdate, private, avatar) 
VALUES ("generalvidchannel", "general - vid", "video", 1, false, "defaultAvatar");

INSERT INTO channels (webId, channelName, type, lastUpdate, private, avatar) 
VALUES ("randomvidchannel", "random - vid", "video", 1, false, "defaultAvatar");

-- PARMETER INITIALISATION ------------------------------------------------------------------------
INSERT INTO parameters (local, timeZone) 
VALUES ("en_EN", "UTC-1");

-- UNCOMMENT THE FOLLOWINGS FOR LOCAL DEV TEST ---------------------------------------------------

-- ROLE INITIALISATION <><><><><><><><><><><><><><><><><><><><><><><><><><><><><><><><><><><><><><>
-- INSERT INTO `popcube_dev`.`roles` (roleName, canUsePrivate, canModerate, canArchive, canInvite, canManage, canManageUser) 
-- VALUES ("owner", true, true, true, true, true, true);

-- INSERT INTO `popcube_dev`.`roles` (roleName, canUsePrivate, canModerate, canArchive, canInvite, canManage, canManageUser) 
-- VALUES ("admin", true, true, true, true, true, true);

-- INSERT INTO `popcube_dev`.`roles` (roleName, canUsePrivate, canModerate, canArchive, canInvite, canManage, canManageUser) 
-- VALUES ("standard", true, true, true, false, false, false);

-- INSERT INTO `popcube_dev`.`roles` (roleName, canUsePrivate, canModerate, canArchive, canInvite, canManage, canManageUser) 
-- VALUES ("guest", false, false, false, false, false, false);

-- SELECT * FROM `popcube_dev`.`roles`;

-- CHANNEL INITIALISATION <><><><><><><><><><><><><><><><><><><><><><><><><><><><><><><><><><><><><><>
-- INSERT INTO `popcube_dev`.`channels` (webId, channelName, type, lastUpdate, private, avatar, description, subject) 
-- VALUES ("generaltextchannel", "general", "text", 1, false, "defaultAvatar", "Speak on general subjects", "General");

-- INSERT INTO `popcube_dev`.`channels` (webId, channelName, type, lastUpdate, private, avatar,description) 
-- VALUES ("randomtextchannel", "random", "text", 1, false, "defaultAvatar","Speak about any thing");

-- INSERT INTO `popcube_dev`.`channels` (webId, channelName, type, lastUpdate, private, avatar) 
-- VALUES ("generalvocchannel", "general - voc", "audio", 1, false, "defaultAvatar");

-- INSERT INTO `popcube_dev`.`channels` (webId, channelName, type, lastUpdate, private, avatar) 
-- VALUES ("randomvocchannel", "random - voc", "audio", 1, false, "defaultAvatar");

-- INSERT INTO `popcube_dev`.`channels` (webId, channelName, type, lastUpdate, private, avatar) 
-- VALUES ("generalvidchannel", "general - vid", "video", 1, false, "defaultAvatar");

-- INSERT INTO `popcube_dev`.`channels` (webId, channelName, type, lastUpdate, private, avatar) 
-- VALUES ("randomvidchannel", "random - vid", "video", 1, false, "defaultAvatar");
-- PARMETER INITIALISATION <><><><><><><><><><><><><><><><><><><><><><><><><><><><><><><><><><><><><><>
-- INSERT INTO `popcube_dev`.`parameters` (local, timeZone) 
-- VALUES ("en_EN", "UTC-1");

-- USER INITIALISATION <><><><><><><><><><><><><><><><><><><><><><><><><><><><><><><><><><><><><><>
INSERT INTO `popcube_dev`.`users` (webId, userName, email, lastUpdate, password, idRole, avatar, nickName, firstName, lastName)
VALUES ("TestUserOwner", "devowner", "devowner@popcube.xyz", 1105154015461, "$2a$10$IU8oU9dseYZytHcr54VXj.H9tX78hS2xUuPrzMeVN6rFG7k89i6EW", 1, "user/owned.svg", "owner", "owner", "dev");


INSERT INTO `popcube_dev`.`users` (webId, userName, email, lastUpdate, password, idRole, avatar, nickName, firstName, lastName)
VALUES ("TestUserAdmin", "devadmin", "devadmin@popcube.xyz", 1105154015461, "$2a$10$IU8oU9dseYZytHcr54VXj.H9tX78hS2xUuPrzMeVN6rFG7k89i6EW", 2, "user/avatar.svg", "admin", "admin", "dev");

INSERT INTO `popcube_dev`.`users` (webId, userName, email, lastUpdate, password, idRole, avatar, nickName, firstName, lastName)
VALUES ("TestUserStandard", "devstandard", "devstandard@popcube.xyz", 1105154015461, "$2a$10$IU8oU9dseYZytHcr54VXj.H9tX78hS2xUuPrzMeVN6rFG7k89i6EW", 3, "user/avatar.svg", "standard", "standard", "dev");

INSERT INTO `popcube_dev`.`users` (webId, userName, email, lastUpdate, password, idRole, avatar, nickName, firstName, lastName)
VALUES ("TestUserGuest", "devguest", "devguest@popcube.xyz", 1105154015461, "$2a$10$IU8oU9dseYZytHcr54VXj.H9tX78hS2xUuPrzMeVN6rFG7k89i6EW", 4, "user/avatar.svg", "guest", "guest", "dev");
-- ORGANISATION INITIALISATION <><><><><><><><><><><><><><><><><><><><><><><><><><><><><><><><><>
INSERT INTO `popcube_dev`.`organisations` (dockerStack, organisationName, description, avatar, domain) 
VALUES (1, "Popcube Dev", "Test for popcube", "popcube.svg", "popcubedev.popbcube.xyz");

-- MEMBERS INITITIALISATION  <><><><><><><><><><><><><><><><><><><><><><><><><><><><><><><><><>
INSERT INTO `popcube_dev`.`members` (idUser, idChannel)
VALUES (1, 1);
INSERT INTO `popcube_dev`.`members` (idUser, idChannel)
VALUES (2, 1);
INSERT INTO `popcube_dev`.`members` (idUser, idChannel)
VALUES (3, 1);
INSERT INTO `popcube_dev`.`members` (idUser, idChannel)
VALUES (4, 1);

INSERT INTO `popcube_dev`.`members` (idUser, idChannel)
VALUES (1, 2);
INSERT INTO `popcube_dev`.`members` (idUser, idChannel)
VALUES (2, 2);
INSERT INTO `popcube_dev`.`members` (idUser, idChannel)
VALUES (3, 2);
INSERT INTO `popcube_dev`.`members` (idUser, idChannel)
VALUES (4, 2);

INSERT INTO `popcube_dev`.`members` (idUser, idChannel)
VALUES (1, 3);
INSERT INTO `popcube_dev`.`members` (idUser, idChannel, idRole)
VALUES (2, 3,3);
INSERT INTO `popcube_dev`.`members` (idUser, idChannel)
VALUES (3, 3);
INSERT INTO `popcube_dev`.`members` (idUser, idChannel)
VALUES (4, 3);

INSERT INTO `popcube_dev`.`members` (idUser, idChannel)
VALUES (1, 4);
INSERT INTO `popcube_dev`.`members` (idUser, idChannel)
VALUES (2, 4);
INSERT INTO `popcube_dev`.`members` (idUser, idChannel, idRole)
VALUES (3, 4, 2);
INSERT INTO `popcube_dev`.`members` (idUser, idChannel)
VALUES (4, 4);

-- MESSAGE INITIALISATION <><><><><><><><><><><><><><><><><><><><><><><><><><><><><><><><><>
INSERT INTO `popcube_dev`.`messages` (idUser, idChannel, date, content) 
VALUES (1, 1, 10210541, "Test message.");
INSERT INTO `popcube_dev`.`messages` (idUser, idChannel, date, content) 
VALUES (1, 1, 10210542, "");
INSERT INTO `popcube_dev`.`messages` (idUser, idChannel, date, content) 
VALUES (1, 1, 10210543, "Test message with folder.");

-- FOLDER INITIALISATION <><><><><><><><><><><><><><><><><><><><><><><><><><><><><><><><><>
INSERT INTO `popcube_dev`.`folders` (idMessage, type, link, name) 
VALUES (2, "txt", "folders/text.txt", "text");

INSERT INTO `popcube_dev`.`folders` (idMessage) 
VALUES (3);

-- EMOJIS INITIALISATION <><><><><><><><><><><><><><><><><><><><><><><><><><><><><><><><><>
INSERT INTO `popcube_dev`.`emojis` (name, shortcut, link)
VALUE ("troll face", ":troll:", "emojis/troll.svg");

INSERT INTO `popcube_dev`.`emojis` (name, shortcut, link)
VALUE ("love", "<3", "emojis/love.svg");

-- AVATARS INITIALISATION <><><><><><><><><><><><><><><><><><><><><><><><><><><><><><><><><>
INSERT INTO `popcube_dev`.`avatars` (name, link)
VALUE ("troll face", "emojis/troll.svg");

INSERT INTO `popcube_dev`.`avatars` (name, link)
VALUE ("Strawberrie", "emojis/straw.svg");