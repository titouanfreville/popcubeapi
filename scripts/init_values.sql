-- ROLE INITIALISATION ------------------------------------------------------------------------
INSERT INTO roles (roleName, canUsePrivate, canModerate, canArchive, canInvite, canManage, canManageUser) 
VALUES ("owner", true, true, true, true, true, true);

INSERT INTO roles (roleName, canUsePrivate, canModerate, canArchive, canInvite, canManage, canManageUser) 
VALUES ("admin", true, true, true, true, true, true);

INSERT INTO roles (roleName, canUsePrivate, canModerate, canArchive, canInvite, canManage, canManageUser) 
VALUES ("standart", true, true, true, false, false, false);

INSERT INTO roles (roleName, canUsePrivate, canModerate, canArchive, canInvite, canManage, canManageUser) 
VALUES ("guest", false, false, false, false, false, false);