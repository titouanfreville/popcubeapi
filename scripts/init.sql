-- MySQL Script generated by MySQL Workbench
-- sam. 11 mars 2017 16:00:09 CET
-- Model: New Model    Version: 1.0
-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='TRADITIONAL,ALLOW_INVALID_DATES';

-- -----------------------------------------------------
-- Schema popcube
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema popcube
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `popcube` DEFAULT CHARACTER SET utf8 ;
-- -----------------------------------------------------
-- Schema popcube_test
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema popcube_test
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `popcube_test` ;
-- -----------------------------------------------------
-- Schema popcube_dev
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema popcube_dev
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `popcube_dev` ;
USE `popcube` ;

-- -----------------------------------------------------
-- Table `popcube`.`organisations`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `popcube`.`organisations` (
  `idOrganisation` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `dockerStack` INT UNSIGNED NOT NULL,
  `organisationName` VARCHAR(45) NOT NULL,
  `public` TINYINT(1) NOT NULL DEFAULT 0,
  `description` VARCHAR(45) NULL,
  `avatar` VARCHAR(45) NULL,
  `domain` VARCHAR(45) NULL,
  PRIMARY KEY (`idOrganisation`))
ENGINE = InnoDB
COMMENT = 'Table to store Organisation related informations\n';

CREATE UNIQUE INDEX `idOrganisation_UNIQUE` ON `popcube`.`organisations` (`idOrganisation` ASC);

CREATE UNIQUE INDEX `dockerStack_UNIQUE` ON `popcube`.`organisations` (`dockerStack` ASC);

CREATE UNIQUE INDEX `name_UNIQUE` ON `popcube`.`organisations` (`organisationName` ASC);


-- -----------------------------------------------------
-- Table `popcube`.`roles`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `popcube`.`roles` (
  `idRole` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `roleName` VARCHAR(45) NOT NULL,
  `canUsePrivate` TINYINT(1) NOT NULL DEFAULT 1 COMMENT 'Can create private channels\n',
  `canModerate` TINYINT(1) NOT NULL DEFAULT 0 COMMENT 'Can moderate channel\nuser did not create\n',
  `canArchive` TINYINT(1) NOT NULL DEFAULT 0 COMMENT 'Can archive channel \nuser did not create\n',
  `canInvite` TINYINT(1) NOT NULL DEFAULT 0 COMMENT 'Can invite new member\nin organisation',
  `canManage` TINYINT(1) NOT NULL DEFAULT 0 COMMENT 'Can manage organisation\n(update information,\nadd bots, add plugins).\n',
  `canManageUser` TINYINT(1) NOT NULL DEFAULT 0 COMMENT 'Role name ~~',
  PRIMARY KEY (`idRole`))
ENGINE = InnoDB
COMMENT = 'Contain all roles within an organisation and their rights';

CREATE UNIQUE INDEX `idRoles_UNIQUE` ON `popcube`.`roles` (`idRole` ASC);

CREATE UNIQUE INDEX `roleName_UNIQUE` ON `popcube`.`roles` (`roleName` ASC);


-- -----------------------------------------------------
-- Table `popcube`.`channels`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `popcube`.`channels` (
  `idChannel` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `webId` VARCHAR(45) NOT NULL,
  `channelName` VARCHAR(45) NOT NULL,
  `type` VARCHAR(45) NOT NULL,
  `private` TINYINT(1) NOT NULL DEFAULT 0,
  `lastUpdate` BIGINT NOT NULL,
  `description` VARCHAR(45) NULL,
  `avatar` VARCHAR(45) NULL,
  `subject` VARCHAR(45) NULL,
  PRIMARY KEY (`idChannel`))
ENGINE = InnoDB
COMMENT = 'Channel Management';

CREATE UNIQUE INDEX `idChannel_UNIQUE` ON `popcube`.`channels` (`idChannel` ASC);

CREATE UNIQUE INDEX `channelName_UNIQUE` ON `popcube`.`channels` (`channelName` ASC);

CREATE UNIQUE INDEX `WebId_UNIQUE` ON `popcube`.`channels` (`webId` ASC);


-- -----------------------------------------------------
-- Table `popcube`.`users`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `popcube`.`users` (
  `idUser` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `webId` VARCHAR(26) NOT NULL COMMENT 'Used to generate web storage keys\n',
  `userName` VARCHAR(64) NOT NULL COMMENT 'User Name. Can be used\nto login instead of mail adress\n',
  `email` VARCHAR(128) NOT NULL COMMENT 'User email. Can be used\nto login instead of userName\n',
  `emailVerified` TINYINT(1) NOT NULL DEFAULT 0 COMMENT 'Store email verification state\n(default false)\n',
  `lastUpdate` BIGINT UNSIGNED NOT NULL COMMENT 'Last known update (in MS)\n',
  `deleted` TINYINT(1) NOT NULL DEFAULT 0 COMMENT 'User deleted ? ',
  `password` VARCHAR(200) NOT NULL COMMENT 'User password. Stored encrypted\n',
  `lastPasswordUpdate` BIGINT NOT NULL COMMENT 'Last time password was updated (in MS)\n',
  `failedAttempts` INT NOT NULL DEFAULT 0 COMMENT 'Number of failed attempts for user login.\n',
  `idRole` INT UNSIGNED NOT NULL,
  `avatar` VARCHAR(45) NULL,
  `nickName` VARCHAR(45) NULL,
  `firstName` VARCHAR(45) NULL,
  `lastName` VARCHAR(45) NULL,
  PRIMARY KEY (`idUser`),
  CONSTRAINT `fk_User_Roles`
    FOREIGN KEY (`idRole`)
    REFERENCES `popcube`.`roles` (`idRole`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB
COMMENT = 'Table to store all user informations\n';

CREATE UNIQUE INDEX `webId_UNIQUE` ON `popcube`.`users` (`webId` ASC);

CREATE UNIQUE INDEX `idUser_UNIQUE` ON `popcube`.`users` (`idUser` ASC);

CREATE UNIQUE INDEX `userName_UNIQUE` ON `popcube`.`users` (`userName` ASC);

CREATE UNIQUE INDEX `email_UNIQUE` ON `popcube`.`users` (`email` ASC);

CREATE INDEX `fk_User_Roles_idx` ON `popcube`.`users` (`idRole` ASC);


-- -----------------------------------------------------
-- Table `popcube`.`messages`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `popcube`.`messages` (
  `idMessage` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `idUser` INT UNSIGNED NOT NULL,
  `idChannel` INT UNSIGNED NOT NULL,
  `date` INT NOT NULL,
  `content` LONGTEXT NULL DEFAULT NULL,
  PRIMARY KEY (`idMessage`),
  CONSTRAINT `fk_Message_Channel`
    FOREIGN KEY (`idChannel`)
    REFERENCES `popcube`.`channels` (`idChannel`)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT `fk_messages_User`
    FOREIGN KEY (`idUser`)
    REFERENCES `popcube`.`users` (`idUser`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB
COMMENT = 'Store message ';

CREATE INDEX `fk_Message_Channel_idx` ON `popcube`.`messages` (`idChannel` ASC);

CREATE INDEX `fk_messages_User_idx` ON `popcube`.`messages` (`idUser` ASC);


-- -----------------------------------------------------
-- Table `popcube`.`folders`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `popcube`.`folders` (
  `idFolder` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `idMessage` INT UNSIGNED NOT NULL,
  `type` VARCHAR(3) NOT NULL DEFAULT 'svg' COMMENT 'File extension\n',
  `link` VARCHAR(45) NOT NULL DEFAULT '/downloads/',
  `name` VARCHAR(45) NOT NULL DEFAULT 'file',
  PRIMARY KEY (`idFolder`, `idMessage`),
  CONSTRAINT `fk_Fichier_Message`
    FOREIGN KEY (`idMessage`)
    REFERENCES `popcube`.`messages` (`idMessage`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB;

CREATE UNIQUE INDEX `idFichier_UNIQUE` ON `popcube`.`folders` (`idFolder` ASC);

CREATE INDEX `fk_Fichier_Message_idx` ON `popcube`.`folders` (`idMessage` ASC);


-- -----------------------------------------------------
-- Table `popcube`.`parameters`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `popcube`.`parameters` (
  `idParameter` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `local` CHAR(5) NOT NULL DEFAULT 'fr_FR',
  `timeZone` CHAR(6) NOT NULL DEFAULT 'UTC-0',
  `sleepStart` INT NOT NULL DEFAULT '1200' COMMENT 'time in minute 24h format\n\n',
  `sleepEnd` INT NOT NULL DEFAULT '240' COMMENT 'time in minute 24h fi\n',
  PRIMARY KEY (`idParameter`))
ENGINE = InnoDB;

CREATE UNIQUE INDEX `idParameter_UNIQUE` ON `popcube`.`parameters` (`idParameter` ASC);

CREATE UNIQUE INDEX `local_UNIQUE` ON `popcube`.`parameters` (`local` ASC);

CREATE UNIQUE INDEX `timeZone_UNIQUE` ON `popcube`.`parameters` (`timeZone` ASC);

CREATE UNIQUE INDEX `sleepStart_UNIQUE` ON `popcube`.`parameters` (`sleepStart` ASC);

CREATE UNIQUE INDEX `sleepEnd_UNIQUE` ON `popcube`.`parameters` (`sleepEnd` ASC);


-- -----------------------------------------------------
-- Table `popcube`.`emojis`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `popcube`.`emojis` (
  `idEmoji` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NOT NULL,
  `shortcut` VARCHAR(45) NOT NULL,
  `link` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`idEmoji`))
ENGINE = InnoDB
COMMENT = 'What emoji can you use ;)';

CREATE UNIQUE INDEX `idEmojis_UNIQUE` ON `popcube`.`emojis` (`idEmoji` ASC);

CREATE UNIQUE INDEX `name_UNIQUE` ON `popcube`.`emojis` (`name` ASC);

CREATE UNIQUE INDEX `raccourcie_UNIQUE` ON `popcube`.`emojis` (`shortcut` ASC);

CREATE UNIQUE INDEX `lien_UNIQUE` ON `popcube`.`emojis` (`link` ASC);


-- -----------------------------------------------------
-- Table `popcube`.`avatars`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `popcube`.`avatars` (
  `idAvatar` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NOT NULL,
  `link` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`idAvatar`))
ENGINE = InnoDB;

CREATE UNIQUE INDEX `idAvatar_UNIQUE` ON `popcube`.`avatars` (`idAvatar` ASC);

CREATE UNIQUE INDEX `name_UNIQUE` ON `popcube`.`avatars` (`name` ASC);

CREATE UNIQUE INDEX `lien_UNIQUE` ON `popcube`.`avatars` (`link` ASC);


-- -----------------------------------------------------
-- Table `popcube`.`read`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `popcube`.`read` (
  `idUser` INT UNSIGNED NOT NULL,
  `idChannel` INT UNSIGNED NOT NULL,
  `idMessage` INT UNSIGNED NOT NULL,
  PRIMARY KEY (`idUser`, `idChannel`, `idMessage`),
  CONSTRAINT `fk_Read_Channel`
    FOREIGN KEY (`idChannel`)
    REFERENCES `popcube`.`channels` (`idChannel`)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT `fk_Read_Message`
    FOREIGN KEY (`idMessage`)
    REFERENCES `popcube`.`messages` (`idMessage`)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT `fk_read_User`
    FOREIGN KEY (`idUser`)
    REFERENCES `popcube`.`users` (`idUser`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB;

CREATE INDEX `fk_Read_Channel_idx` ON `popcube`.`read` (`idChannel` ASC);

CREATE INDEX `fk_Read_Message_idx` ON `popcube`.`read` (`idMessage` ASC);


-- -----------------------------------------------------
-- Table `popcube`.`allowed_web_mails`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `popcube`.`allowed_web_mails` (
  `idAllowedWebMails` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `domain` VARCHAR(45) NOT NULL,
  `provider` VARCHAR(45) NULL,
  `defaultRights` VARCHAR(45) NULL DEFAULT 'standard',
  PRIMARY KEY (`idAllowedWebMails`))
ENGINE = InnoDB
COMMENT = 'Table to manage webmail domain that can create an account on organisation without being invitated. ';

CREATE UNIQUE INDEX `idAllowedWebMails_UNIQUE` ON `popcube`.`allowed_web_mails` (`idAllowedWebMails` ASC);

CREATE UNIQUE INDEX `domain_UNIQUE` ON `popcube`.`allowed_web_mails` (`domain` ASC);


-- -----------------------------------------------------
-- Table `popcube`.`user_parameter`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `popcube`.`user_parameter` (
  `idUser` INT UNSIGNED NOT NULL,
  `parameterName` VARCHAR(45) NOT NULL,
  `local` CHAR(5) NULL,
  `timeZone` CHAR(4) NULL,
  `sleepStart` INT NULL,
  `sleepEnd` INT NULL,
  PRIMARY KEY (`parameterName`, `idUser`),
  CONSTRAINT `fk_userParameter_User`
    FOREIGN KEY (`idUser`)
    REFERENCES `popcube`.`users` (`idUser`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB;

CREATE INDEX `fk_userParameter_User_idx` ON `popcube`.`user_parameter` (`idUser` ASC);


-- -----------------------------------------------------
-- Table `popcube`.`members`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `popcube`.`members` (
  `idUser` INT UNSIGNED NOT NULL,
  `idChannel` INT UNSIGNED NOT NULL,
  `idRole` INT UNSIGNED NULL,
  PRIMARY KEY (`idUser`, `idChannel`),
  CONSTRAINT `fk_Member_User`
    FOREIGN KEY (`idUser`)
    REFERENCES `popcube`.`users` (`idUser`)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT `fk_Member_Role`
    FOREIGN KEY (`idRole`)
    REFERENCES `popcube`.`roles` (`idRole`)
    ON DELETE SET NULL
    ON UPDATE CASCADE,
  CONSTRAINT `fk_Member_Channel`
    FOREIGN KEY (`idChannel`)
    REFERENCES `popcube`.`channels` (`idChannel`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB
COMMENT = 'Store information about member of organisation.';

CREATE INDEX `fk_Member_Role_idx` ON `popcube`.`members` (`idRole` ASC);

CREATE INDEX `fk_Member_Channel_idx` ON `popcube`.`members` (`idChannel` ASC);

USE `popcube_test` ;

-- -----------------------------------------------------
-- Table `popcube_test`.`allowed_web_mails`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `popcube_test`.`allowed_web_mails` (
  `idAllowedWebMails` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `domain` VARCHAR(45) NOT NULL,
  `provider` VARCHAR(45) NULL,
  `defaultRights` VARCHAR(45) NULL DEFAULT 'standard',
  PRIMARY KEY (`idAllowedWebMails`))
ENGINE = InnoDB
COMMENT = 'Table to manage webmail domain that can create an account on organisation without being invitated. ';

CREATE UNIQUE INDEX `idAllowedWebMails_UNIQUE` ON `popcube_test`.`allowed_web_mails` (`idAllowedWebMails` ASC);

CREATE UNIQUE INDEX `domain_UNIQUE` ON `popcube_test`.`allowed_web_mails` (`domain` ASC);


-- -----------------------------------------------------
-- Table `popcube_test`.`avatars`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `popcube_test`.`avatars` (
  `idAvatar` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NOT NULL,
  `link` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`idAvatar`))
ENGINE = InnoDB;

CREATE UNIQUE INDEX `idAvatar_UNIQUE` ON `popcube_test`.`avatars` (`idAvatar` ASC);

CREATE UNIQUE INDEX `name_UNIQUE` ON `popcube_test`.`avatars` (`name` ASC);

CREATE UNIQUE INDEX `lien_UNIQUE` ON `popcube_test`.`avatars` (`link` ASC);


-- -----------------------------------------------------
-- Table `popcube_test`.`channels`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `popcube_test`.`channels` (
  `idChannel` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `webId` VARCHAR(45) NOT NULL,
  `channelName` VARCHAR(45) NOT NULL,
  `type` VARCHAR(45) NOT NULL,
  `private` TINYINT(1) NOT NULL DEFAULT 0,
  `lastUpdate` BIGINT NOT NULL,
  `description` VARCHAR(45) NULL,
  `avatar` VARCHAR(45) NULL,
  `subject` VARCHAR(45) NULL,
  PRIMARY KEY (`idChannel`))
ENGINE = InnoDB
COMMENT = 'Channel Management';

CREATE UNIQUE INDEX `idChannel_UNIQUE` ON `popcube_test`.`channels` (`idChannel` ASC);

CREATE UNIQUE INDEX `channelName_UNIQUE` ON `popcube_test`.`channels` (`channelName` ASC);

CREATE UNIQUE INDEX `WebId_UNIQUE` ON `popcube_test`.`channels` (`webId` ASC);


-- -----------------------------------------------------
-- Table `popcube_test`.`emojis`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `popcube_test`.`emojis` (
  `idEmoji` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NOT NULL,
  `shortcut` VARCHAR(45) NOT NULL,
  `link` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`idEmoji`))
ENGINE = InnoDB
COMMENT = 'What emoji can you use ;)';

CREATE UNIQUE INDEX `idEmojis_UNIQUE` ON `popcube_test`.`emojis` (`idEmoji` ASC);

CREATE UNIQUE INDEX `name_UNIQUE` ON `popcube_test`.`emojis` (`name` ASC);

CREATE UNIQUE INDEX `raccourcie_UNIQUE` ON `popcube_test`.`emojis` (`shortcut` ASC);

CREATE UNIQUE INDEX `lien_UNIQUE` ON `popcube_test`.`emojis` (`link` ASC);


-- -----------------------------------------------------
-- Table `popcube_test`.`roles`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `popcube_test`.`roles` (
  `idRole` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `roleName` VARCHAR(45) NOT NULL,
  `canUsePrivate` TINYINT(1) NOT NULL DEFAULT 1 COMMENT 'Can create private channels\n',
  `canModerate` TINYINT(1) NOT NULL DEFAULT 0 COMMENT 'Can moderate channel\nuser did not create\n',
  `canArchive` TINYINT(1) NOT NULL DEFAULT 0 COMMENT 'Can archive channel \nuser did not create\n',
  `canInvite` TINYINT(1) NOT NULL DEFAULT 0 COMMENT 'Can invite new member\nin organisation',
  `canManage` TINYINT(1) NOT NULL DEFAULT 0 COMMENT 'Can manage organisation\n(update information,\nadd bots, add plugins).\n',
  `canManageUser` TINYINT(1) NOT NULL DEFAULT 0 COMMENT 'Role name ~~',
  PRIMARY KEY (`idRole`))
ENGINE = InnoDB
COMMENT = 'Contain all roles within an organisation and their rights';

CREATE UNIQUE INDEX `idRoles_UNIQUE` ON `popcube_test`.`roles` (`idRole` ASC);

CREATE UNIQUE INDEX `roleName_UNIQUE` ON `popcube_test`.`roles` (`roleName` ASC);


-- -----------------------------------------------------
-- Table `popcube_test`.`users`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `popcube_test`.`users` (
  `idUser` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `webId` VARCHAR(26) NOT NULL COMMENT 'Used to generate web storage keys\n',
  `userName` VARCHAR(64) NOT NULL COMMENT 'User Name. Can be used\nto login instead of mail adress\n',
  `email` VARCHAR(128) NOT NULL COMMENT 'User email. Can be used\nto login instead of userName\n',
  `emailVerified` TINYINT(1) NOT NULL DEFAULT 0 COMMENT 'Store email verification state\n(default false)\n',
  `lastUpdate` BIGINT UNSIGNED NOT NULL COMMENT 'Last known update (in MS)\n',
  `deleted` TINYINT(1) NOT NULL DEFAULT 0 COMMENT 'User deleted ? ',
  `password` VARCHAR(200) NOT NULL COMMENT 'User password. Stored encrypted\n',
  `lastPasswordUpdate` BIGINT NOT NULL COMMENT 'Last time password was updated (in MS)\n',
  `failedAttempts` INT NOT NULL DEFAULT 0 COMMENT 'Number of failed attempts for user login.\n',
  `idRole` INT UNSIGNED NOT NULL,
  `avatar` VARCHAR(45) NULL,
  `nickName` VARCHAR(45) NULL,
  `firstName` VARCHAR(45) NULL,
  `lastName` VARCHAR(45) NULL,
  PRIMARY KEY (`idUser`),
  CONSTRAINT `fk_User_Roles`
    FOREIGN KEY (`idRole`)
    REFERENCES `popcube_test`.`roles` (`idRole`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB
COMMENT = 'Table to store all user informations\n';

CREATE UNIQUE INDEX `webId_UNIQUE` ON `popcube_test`.`users` (`webId` ASC);

CREATE UNIQUE INDEX `idUser_UNIQUE` ON `popcube_test`.`users` (`idUser` ASC);

CREATE UNIQUE INDEX `userName_UNIQUE` ON `popcube_test`.`users` (`userName` ASC);

CREATE UNIQUE INDEX `email_UNIQUE` ON `popcube_test`.`users` (`email` ASC);

CREATE INDEX `fk_User_Roles_idx` ON `popcube_test`.`users` (`idRole` ASC);


-- -----------------------------------------------------
-- Table `popcube_test`.`messages`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `popcube_test`.`messages` (
  `idMessage` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `idUser` INT UNSIGNED NOT NULL,
  `idChannel` INT UNSIGNED NOT NULL,
  `date` INT NOT NULL,
  `content` LONGTEXT NULL DEFAULT NULL,
  PRIMARY KEY (`idMessage`),
  CONSTRAINT `fk_Message_User`
    FOREIGN KEY (`idUser`)
    REFERENCES `popcube_test`.`users` (`idUser`)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT `fk_Message_Channel`
    FOREIGN KEY (`idChannel`)
    REFERENCES `popcube_test`.`channels` (`idChannel`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB
COMMENT = 'Store message ';

CREATE INDEX `fk_Message_Channel_idx` ON `popcube_test`.`messages` (`idChannel` ASC);

CREATE INDEX `fk_Message_User_idx` ON `popcube_test`.`messages` (`idUser` ASC);


-- -----------------------------------------------------
-- Table `popcube_test`.`folders`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `popcube_test`.`folders` (
  `idFolder` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `idMessage` INT UNSIGNED NOT NULL,
  `type` VARCHAR(3) NOT NULL DEFAULT 'svg' COMMENT 'File extension\n',
  `link` VARCHAR(45) NOT NULL DEFAULT '/downloads/',
  `name` VARCHAR(45) NOT NULL DEFAULT 'file',
  PRIMARY KEY (`idFolder`, `idMessage`),
  CONSTRAINT `fk_Fichier_Message`
    FOREIGN KEY (`idMessage`)
    REFERENCES `popcube_test`.`messages` (`idMessage`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB;

CREATE UNIQUE INDEX `idFichier_UNIQUE` ON `popcube_test`.`folders` (`idFolder` ASC);

CREATE INDEX `fk_Fichier_Message_idx` ON `popcube_test`.`folders` (`idMessage` ASC);


-- -----------------------------------------------------
-- Table `popcube_test`.`members`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `popcube_test`.`members` (
  `idUser` INT UNSIGNED NOT NULL,
  `idChannel` INT UNSIGNED NOT NULL,
  `idRole` INT UNSIGNED NULL,
  PRIMARY KEY (`idUser`, `idChannel`),
  CONSTRAINT `fk_Member_User`
    FOREIGN KEY (`idUser`)
    REFERENCES `popcube_test`.`users` (`idUser`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_Member_Role`
    FOREIGN KEY (`idRole`)
    REFERENCES `popcube_test`.`roles` (`idRole`)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT `fk_Member_Channel`
    FOREIGN KEY (`idChannel`)
    REFERENCES `popcube_test`.`channels` (`idChannel`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB
COMMENT = 'Store information about member of organisation.';

CREATE INDEX `fk_Member_Role_idx` ON `popcube_test`.`members` (`idRole` ASC);

CREATE INDEX `fk_Member_Channel_idx` ON `popcube_test`.`members` (`idChannel` ASC);


-- -----------------------------------------------------
-- Table `popcube_test`.`organisations`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `popcube_test`.`organisations` (
  `idOrganisation` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `dockerStack` INT UNSIGNED NOT NULL,
  `organisationName` VARCHAR(45) NOT NULL,
  `public` TINYINT(1) NOT NULL DEFAULT 0,
  `description` VARCHAR(45) NULL,
  `avatar` VARCHAR(45) NULL,
  `domain` VARCHAR(45) NULL,
  PRIMARY KEY (`idOrganisation`))
ENGINE = InnoDB
COMMENT = 'Table to store Organisation related informations\n';

CREATE UNIQUE INDEX `idOrganisation_UNIQUE` ON `popcube_test`.`organisations` (`idOrganisation` ASC);

CREATE UNIQUE INDEX `dockerStack_UNIQUE` ON `popcube_test`.`organisations` (`dockerStack` ASC);

CREATE UNIQUE INDEX `name_UNIQUE` ON `popcube_test`.`organisations` (`organisationName` ASC);


-- -----------------------------------------------------
-- Table `popcube_test`.`parameters`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `popcube_test`.`parameters` (
  `idParameter` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `local` CHAR(5) NOT NULL DEFAULT 'fr_FR',
  `timeZone` CHAR(6) NOT NULL DEFAULT 'UTC-0',
  `sleepStart` INT NOT NULL DEFAULT '1200' COMMENT 'time in minute 24h format\n\n',
  `sleepEnd` INT NOT NULL DEFAULT '240' COMMENT 'time in minute 24h fi\n',
  PRIMARY KEY (`idParameter`))
ENGINE = InnoDB;

CREATE UNIQUE INDEX `idParameter_UNIQUE` ON `popcube_test`.`parameters` (`idParameter` ASC);

CREATE UNIQUE INDEX `local_UNIQUE` ON `popcube_test`.`parameters` (`local` ASC);

CREATE UNIQUE INDEX `timeZone_UNIQUE` ON `popcube_test`.`parameters` (`timeZone` ASC);

CREATE UNIQUE INDEX `sleepStart_UNIQUE` ON `popcube_test`.`parameters` (`sleepStart` ASC);

CREATE UNIQUE INDEX `sleepEnd_UNIQUE` ON `popcube_test`.`parameters` (`sleepEnd` ASC);


-- -----------------------------------------------------
-- Table `popcube_test`.`read`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `popcube_test`.`read` (
  `idUser` INT UNSIGNED NOT NULL,
  `idChannel` INT UNSIGNED NOT NULL,
  `idMessage` INT UNSIGNED NOT NULL,
  PRIMARY KEY (`idUser`, `idChannel`, `idMessage`),
  CONSTRAINT `fk_Read_User`
    FOREIGN KEY (`idUser`)
    REFERENCES `popcube_test`.`users` (`idUser`)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT `fk_Read_Channel`
    FOREIGN KEY (`idChannel`)
    REFERENCES `popcube_test`.`channels` (`idChannel`)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT `fk_Read_Message`
    FOREIGN KEY (`idMessage`)
    REFERENCES `popcube_test`.`messages` (`idMessage`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB;

CREATE INDEX `fk_Read_Channel_idx` ON `popcube_test`.`read` (`idChannel` ASC);

CREATE INDEX `fk_Read_Message_idx` ON `popcube_test`.`read` (`idMessage` ASC);


-- -----------------------------------------------------
-- Table `popcube_test`.`user_parameter`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `popcube_test`.`user_parameter` (
  `idUser` INT UNSIGNED NOT NULL,
  `parameterName` VARCHAR(45) NOT NULL,
  `local` CHAR(5) NULL,
  `timeZone` CHAR(4) NULL,
  `sleepStart` INT NULL,
  `sleepEnd` INT NULL,
  PRIMARY KEY (`parameterName`, `idUser`),
  CONSTRAINT `fk_userParameter_User`
    FOREIGN KEY (`idUser`)
    REFERENCES `popcube_test`.`users` (`idUser`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB;

CREATE INDEX `fk_userParameter_User_idx` ON `popcube_test`.`user_parameter` (`idUser` ASC);

USE `popcube_dev` ;

-- -----------------------------------------------------
-- Table `popcube_dev`.`allowed_web_mails`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `popcube_dev`.`allowed_web_mails` (
  `idAllowedWebMails` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `domain` VARCHAR(45) NOT NULL,
  `provider` VARCHAR(45) NULL,
  `defaultRights` VARCHAR(45) NULL DEFAULT 'standard',
  PRIMARY KEY (`idAllowedWebMails`))
ENGINE = InnoDB
COMMENT = 'Table to manage webmail domain that can create an account on organisation without being invitated. ';

CREATE UNIQUE INDEX `idAllowedWebMails_UNIQUE` ON `popcube_dev`.`allowed_web_mails` (`idAllowedWebMails` ASC);

CREATE UNIQUE INDEX `domain_UNIQUE` ON `popcube_dev`.`allowed_web_mails` (`domain` ASC);


-- -----------------------------------------------------
-- Table `popcube_dev`.`avatars`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `popcube_dev`.`avatars` (
  `idAvatar` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NOT NULL,
  `link` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`idAvatar`))
ENGINE = InnoDB;

CREATE UNIQUE INDEX `idAvatar_UNIQUE` ON `popcube_dev`.`avatars` (`idAvatar` ASC);

CREATE UNIQUE INDEX `name_UNIQUE` ON `popcube_dev`.`avatars` (`name` ASC);

CREATE UNIQUE INDEX `lien_UNIQUE` ON `popcube_dev`.`avatars` (`link` ASC);


-- -----------------------------------------------------
-- Table `popcube_dev`.`channels`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `popcube_dev`.`channels` (
  `idChannel` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `webId` VARCHAR(45) NOT NULL,
  `channelName` VARCHAR(45) NOT NULL,
  `type` VARCHAR(45) NOT NULL,
  `private` TINYINT(1) NOT NULL DEFAULT 0,
  `lastUpdate` BIGINT NOT NULL,
  `description` VARCHAR(45) NULL,
  `avatar` VARCHAR(45) NULL,
  `subject` VARCHAR(45) NULL,
  PRIMARY KEY (`idChannel`))
ENGINE = InnoDB
COMMENT = 'Channel Management';

CREATE UNIQUE INDEX `idChannel_UNIQUE` ON `popcube_dev`.`channels` (`idChannel` ASC);

CREATE UNIQUE INDEX `channelName_UNIQUE` ON `popcube_dev`.`channels` (`channelName` ASC);

CREATE UNIQUE INDEX `WebId_UNIQUE` ON `popcube_dev`.`channels` (`webId` ASC);


-- -----------------------------------------------------
-- Table `popcube_dev`.`emojis`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `popcube_dev`.`emojis` (
  `idEmoji` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NOT NULL,
  `shortcut` VARCHAR(45) NOT NULL,
  `link` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`idEmoji`))
ENGINE = InnoDB
COMMENT = 'What emoji can you use ;)';

CREATE UNIQUE INDEX `idEmojis_UNIQUE` ON `popcube_dev`.`emojis` (`idEmoji` ASC);

CREATE UNIQUE INDEX `name_UNIQUE` ON `popcube_dev`.`emojis` (`name` ASC);

CREATE UNIQUE INDEX `raccourcie_UNIQUE` ON `popcube_dev`.`emojis` (`shortcut` ASC);

CREATE UNIQUE INDEX `lien_UNIQUE` ON `popcube_dev`.`emojis` (`link` ASC);


-- -----------------------------------------------------
-- Table `popcube_dev`.`roles`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `popcube_dev`.`roles` (
  `idRole` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `roleName` VARCHAR(45) NOT NULL,
  `canUsePrivate` TINYINT(1) NOT NULL DEFAULT 1 COMMENT 'Can create private channels\n',
  `canModerate` TINYINT(1) NOT NULL DEFAULT 0 COMMENT 'Can moderate channel\nuser did not create\n',
  `canArchive` TINYINT(1) NOT NULL DEFAULT 0 COMMENT 'Can archive channel \nuser did not create\n',
  `canInvite` TINYINT(1) NOT NULL DEFAULT 0 COMMENT 'Can invite new member\nin organisation',
  `canManage` TINYINT(1) NOT NULL DEFAULT 0 COMMENT 'Can manage organisation\n(update information,\nadd bots, add plugins).\n',
  `canManageUser` TINYINT(1) NOT NULL DEFAULT 0 COMMENT 'Role name ~~',
  PRIMARY KEY (`idRole`))
ENGINE = InnoDB
COMMENT = 'Contain all roles within an organisation and their rights';

CREATE UNIQUE INDEX `idRoles_UNIQUE` ON `popcube_dev`.`roles` (`idRole` ASC);

CREATE UNIQUE INDEX `roleName_UNIQUE` ON `popcube_dev`.`roles` (`roleName` ASC);


-- -----------------------------------------------------
-- Table `popcube_dev`.`users`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `popcube_dev`.`users` (
  `idUser` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `webId` VARCHAR(26) NOT NULL COMMENT 'Used to generate web storage keys\n',
  `userName` VARCHAR(64) NOT NULL COMMENT 'User Name. Can be used\nto login instead of mail adress\n',
  `email` VARCHAR(128) NOT NULL COMMENT 'User email. Can be used\nto login instead of userName\n',
  `emailVerified` TINYINT(1) NOT NULL DEFAULT 0 COMMENT 'Store email verification state\n(default false)\n',
  `lastUpdate` BIGINT UNSIGNED NOT NULL COMMENT 'Last known update (in MS)\n',
  `deleted` TINYINT(1) NOT NULL DEFAULT 0 COMMENT 'User deleted ? ',
  `password` VARCHAR(200) NOT NULL COMMENT 'User password. Stored encrypted\n',
  `lastPasswordUpdate` BIGINT NOT NULL COMMENT 'Last time password was updated (in MS)\n',
  `failedAttempts` INT NOT NULL DEFAULT 0 COMMENT 'Number of failed attempts for user login.\n',
  `idRole` INT UNSIGNED NOT NULL,
  `avatar` VARCHAR(45) NULL,
  `nickName` VARCHAR(45) NULL,
  `firstName` VARCHAR(45) NULL,
  `lastName` VARCHAR(45) NULL,
  PRIMARY KEY (`idUser`),
  CONSTRAINT `fk_User_Roles`
    FOREIGN KEY (`idRole`)
    REFERENCES `popcube_dev`.`roles` (`idRole`)
    ON DELETE RESTRICT
    ON UPDATE CASCADE)
ENGINE = InnoDB
COMMENT = 'Table to store all user informations\n';

CREATE UNIQUE INDEX `webId_UNIQUE` ON `popcube_dev`.`users` (`webId` ASC);

CREATE UNIQUE INDEX `idUser_UNIQUE` ON `popcube_dev`.`users` (`idUser` ASC);

CREATE UNIQUE INDEX `userName_UNIQUE` ON `popcube_dev`.`users` (`userName` ASC);

CREATE UNIQUE INDEX `email_UNIQUE` ON `popcube_dev`.`users` (`email` ASC);

CREATE INDEX `fk_User_Roles_idx` ON `popcube_dev`.`users` (`idRole` ASC);


-- -----------------------------------------------------
-- Table `popcube_dev`.`messages`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `popcube_dev`.`messages` (
  `idMessage` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `idUser` INT UNSIGNED NOT NULL,
  `idChannel` INT UNSIGNED NOT NULL,
  `date` INT NOT NULL,
  `content` LONGTEXT NULL DEFAULT NULL,
  PRIMARY KEY (`idMessage`),
  CONSTRAINT `fk_Message_User`
    FOREIGN KEY (`idUser`)
    REFERENCES `popcube_dev`.`users` (`idUser`)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT `fk_Message_Channel`
    FOREIGN KEY (`idChannel`)
    REFERENCES `popcube_dev`.`channels` (`idChannel`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB
COMMENT = 'Store message ';

CREATE INDEX `fk_Message_User_idx` ON `popcube_dev`.`messages` (`idUser` ASC);

CREATE INDEX `fk_Message_Channel_idx` ON `popcube_dev`.`messages` (`idChannel` ASC);


-- -----------------------------------------------------
-- Table `popcube_dev`.`folders`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `popcube_dev`.`folders` (
  `idFolder` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `idMessage` INT UNSIGNED NOT NULL,
  `type` VARCHAR(3) NOT NULL DEFAULT 'svg' COMMENT 'File extension\n',
  `link` VARCHAR(45) NOT NULL DEFAULT '/downloads/',
  `name` VARCHAR(45) NOT NULL DEFAULT 'file',
  PRIMARY KEY (`idFolder`, `idMessage`),
  CONSTRAINT `fk_Fichier_Message`
    FOREIGN KEY (`idMessage`)
    REFERENCES `popcube_dev`.`messages` (`idMessage`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB;

CREATE UNIQUE INDEX `idFichier_UNIQUE` ON `popcube_dev`.`folders` (`idFolder` ASC);

CREATE INDEX `fk_Fichier_Message_idx` ON `popcube_dev`.`folders` (`idMessage` ASC);


-- -----------------------------------------------------
-- Table `popcube_dev`.`members`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `popcube_dev`.`members` (
  `idUser` INT UNSIGNED NOT NULL,
  `idChannel` INT UNSIGNED NOT NULL,
  `idRole` INT UNSIGNED NULL,
  PRIMARY KEY (`idUser`, `idChannel`),
  CONSTRAINT `fk_Member_User`
    FOREIGN KEY (`idUser`)
    REFERENCES `popcube_dev`.`users` (`idUser`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_Member_Role`
    FOREIGN KEY (`idRole`)
    REFERENCES `popcube_dev`.`roles` (`idRole`)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT `fk_Member_Channel`
    FOREIGN KEY (`idChannel`)
    REFERENCES `popcube_dev`.`channels` (`idChannel`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB
COMMENT = 'Store information about member of organisation.';

CREATE INDEX `fk_Member_Role_idx` ON `popcube_dev`.`members` (`idRole` ASC);

CREATE INDEX `fk_Member_Channel_idx` ON `popcube_dev`.`members` (`idChannel` ASC);


-- -----------------------------------------------------
-- Table `popcube_dev`.`organisations`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `popcube_dev`.`organisations` (
  `idOrganisation` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `dockerStack` INT UNSIGNED NOT NULL,
  `organisationName` VARCHAR(45) NOT NULL,
  `public` TINYINT(1) NOT NULL DEFAULT 0,
  `description` VARCHAR(45) NULL,
  `avatar` VARCHAR(45) NULL,
  `domain` VARCHAR(45) NULL,
  PRIMARY KEY (`idOrganisation`))
ENGINE = InnoDB
COMMENT = 'Table to store Organisation related informations\n';

CREATE UNIQUE INDEX `idOrganisation_UNIQUE` ON `popcube_dev`.`organisations` (`idOrganisation` ASC);

CREATE UNIQUE INDEX `dockerStack_UNIQUE` ON `popcube_dev`.`organisations` (`dockerStack` ASC);

CREATE UNIQUE INDEX `name_UNIQUE` ON `popcube_dev`.`organisations` (`organisationName` ASC);


-- -----------------------------------------------------
-- Table `popcube_dev`.`parameters`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `popcube_dev`.`parameters` (
  `idParameter` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `local` CHAR(5) NOT NULL DEFAULT 'fr_FR',
  `timeZone` CHAR(6) NOT NULL DEFAULT 'UTC-0',
  `sleepStart` INT NOT NULL DEFAULT '1200' COMMENT 'time in minute 24h format\n\n',
  `sleepEnd` INT NOT NULL DEFAULT '240' COMMENT 'time in minute 24h fi\n',
  PRIMARY KEY (`idParameter`))
ENGINE = InnoDB;

CREATE UNIQUE INDEX `idParameter_UNIQUE` ON `popcube_dev`.`parameters` (`idParameter` ASC);

CREATE UNIQUE INDEX `local_UNIQUE` ON `popcube_dev`.`parameters` (`local` ASC);

CREATE UNIQUE INDEX `timeZone_UNIQUE` ON `popcube_dev`.`parameters` (`timeZone` ASC);

CREATE UNIQUE INDEX `sleepStart_UNIQUE` ON `popcube_dev`.`parameters` (`sleepStart` ASC);

CREATE UNIQUE INDEX `sleepEnd_UNIQUE` ON `popcube_dev`.`parameters` (`sleepEnd` ASC);


-- -----------------------------------------------------
-- Table `popcube_dev`.`read`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `popcube_dev`.`read` (
  `idUser` INT UNSIGNED NOT NULL,
  `idChannel` INT UNSIGNED NOT NULL,
  `idMessage` INT UNSIGNED NOT NULL,
  PRIMARY KEY (`idUser`, `idChannel`, `idMessage`),
  CONSTRAINT `fk_Read_User`
    FOREIGN KEY (`idUser`)
    REFERENCES `popcube_dev`.`users` (`idUser`)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT `fk_Read_Channel`
    FOREIGN KEY (`idChannel`)
    REFERENCES `popcube_dev`.`channels` (`idChannel`)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT `fk_Read_Message`
    FOREIGN KEY (`idMessage`)
    REFERENCES `popcube_dev`.`messages` (`idMessage`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB;

CREATE INDEX `fk_Read_Channel_idx` ON `popcube_dev`.`read` (`idChannel` ASC);

CREATE INDEX `fk_Read_Message_idx` ON `popcube_dev`.`read` (`idMessage` ASC);


-- -----------------------------------------------------
-- Table `popcube_dev`.`user_parameter`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `popcube_dev`.`user_parameter` (
  `idUser` INT UNSIGNED NOT NULL,
  `parameterName` VARCHAR(45) NOT NULL,
  `local` CHAR(5) NULL,
  `timeZone` CHAR(4) NULL,
  `sleepStart` INT NULL,
  `sleepEnd` INT NULL,
  PRIMARY KEY (`parameterName`, `idUser`),
  CONSTRAINT `fk_userParameter_User`
    FOREIGN KEY (`idUser`)
    REFERENCES `popcube_dev`.`users` (`idUser`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB;

CREATE INDEX `fk_userParameter_User_idx` ON `popcube_dev`.`user_parameter` (`idUser` ASC);


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
