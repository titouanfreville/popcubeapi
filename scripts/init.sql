-- MySQL Script generated by MySQL Workbench
-- dim. 12 févr. 2017 19:57:47 CET
-- Model: New Model    Version: 1.0
-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='TRADITIONAL,ALLOW_INVALID_DATES';

-- -----------------------------------------------------
-- Schema popcube_test
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema popcube_test
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `popcube_test` DEFAULT CHARACTER SET utf8 ;
USE `popcube_test` ;

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
  `lastUpdate` BIGINT(20) NOT NULL COMMENT 'Last known update (in MS)\n',
  `deleted` TINYINT(1) NOT NULL DEFAULT 0 COMMENT 'User deleted ? ',
  `password` VARCHAR(200) NOT NULL COMMENT 'User password. Stored encrypted\n',
  `lastPasswordUpdate` BIGINT(20) NOT NULL COMMENT 'Last time password was updated (in MS)\n',
  `failedAttemprs` INT NOT NULL DEFAULT 0 COMMENT 'Number of failed attempts for user login.\n',
  `locale` VARCHAR(5) NOT NULL DEFAULT 'fr_FR' COMMENT 'Langage code (fr_FR, en_US, en_EN .... )\n',
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

CREATE UNIQUE INDEX `nickName_UNIQUE` ON `popcube_test`.`users` (`nickName` ASC);


-- -----------------------------------------------------
-- Table `popcube_test`.`organisations`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `popcube_test`.`organisations` (
  `idOrganisation` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `dockerStack` INT UNSIGNED NOT NULL,
  `organisationName` VARCHAR(45) NOT NULL,
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
-- Table `popcube_test`.`channels`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `popcube_test`.`channels` (
  `idChannel` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `webId` VARCHAR(45) NOT NULL,
  `channelName` VARCHAR(45) NOT NULL,
  `type` VARCHAR(45) NOT NULL,
  `private` TINYINT(1) NOT NULL DEFAULT 0,
  `lastUpdate` BIGINT(20) NOT NULL,
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
-- Table `popcube_test`.`members`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `popcube_test`.`members` (
  `user` INT UNSIGNED NOT NULL,
  `channel` INT UNSIGNED NOT NULL,
  `role` INT UNSIGNED NULL,
  PRIMARY KEY (`user`, `channel`),
  CONSTRAINT `fk_Member_User`
    FOREIGN KEY (`user`)
    REFERENCES `popcube_test`.`users` (`idUser`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_Member_Role`
    FOREIGN KEY (`role`)
    REFERENCES `popcube_test`.`roles` (`idRole`)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT `fk_Member_Channel`
    FOREIGN KEY (`channel`)
    REFERENCES `popcube_test`.`channels` (`idChannel`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB
COMMENT = 'Store information about member of organisation.';

CREATE UNIQUE INDEX `User_UNIQUE` ON `popcube_test`.`members` (`user` ASC);

CREATE UNIQUE INDEX `Role_UNIQUE` ON `popcube_test`.`members` (`role` ASC);

CREATE UNIQUE INDEX `channel_UNIQUE` ON `popcube_test`.`members` (`channel` ASC);


-- -----------------------------------------------------
-- Table `popcube_test`.`messages`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `popcube_test`.`messages` (
  `idMessage` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `idUser` INT UNSIGNED NOT NULL,
  `idChannel` INT UNSIGNED NOT NULL COMMENT 'Channel FK',
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

CREATE INDEX `fk_Message_User_idx` ON `popcube_test`.`messages` (`idUser` ASC);

CREATE INDEX `fk_Message_Channel_idx` ON `popcube_test`.`messages` (`idChannel` ASC);


-- -----------------------------------------------------
-- Table `popcube_test`.`folders`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `popcube_test`.`folders` (
  `idFolder` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `message` INT UNSIGNED NOT NULL,
  `type` VARCHAR(3) NOT NULL DEFAULT 'svg' COMMENT 'File extension\n',
  `lien` VARCHAR(45) NOT NULL DEFAULT '/downloads/',
  `name` VARCHAR(45) NOT NULL DEFAULT 'file',
  PRIMARY KEY (`idFolder`, `message`),
  CONSTRAINT `fk_Fichier_Message`
    FOREIGN KEY (`message`)
    REFERENCES `popcube_test`.`messages` (`idMessage`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB;

CREATE UNIQUE INDEX `idFichier_UNIQUE` ON `popcube_test`.`folders` (`idFolder` ASC);

CREATE INDEX `fk_Fichier_Message_idx` ON `popcube_test`.`folders` (`message` ASC);


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


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
