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

INSERT INTO organisations (dockerStack, organisationName, description, avatar, domain)
VALUES (1, "%org_organisationName%", "%org_description%", "%org_avatar%", "%org_domain%");
