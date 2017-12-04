# ************************************************************
# Sequel Pro SQL dump
# Version 4541
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 127.0.0.1 (MySQL 5.7.20)
# Database: ChannelX
# Generation Time: 2017-12-04 19:35:33 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table ALIAS
# ------------------------------------------------------------

DROP TABLE IF EXISTS `ALIAS`;

CREATE TABLE `ALIAS` (
  `alias_id` int(11) NOT NULL AUTO_INCREMENT,
  `val` varchar(50) NOT NULL DEFAULT '',
  `is_user_defined` tinyint(1) NOT NULL,
  PRIMARY KEY (`alias_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `ALIAS` WRITE;
/*!40000 ALTER TABLE `ALIAS` DISABLE KEYS */;

INSERT INTO `ALIAS` (`alias_id`, `val`, `is_user_defined`)
VALUES
    (1,'Black Monk',0),
    (2,'Boring Panda',0),
    (3,'Boring Panda',0),
    (4,'Boring Panda',0),
    (5,'Boring Panda',0);

/*!40000 ALTER TABLE `ALIAS` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table CHANNEL
# ------------------------------------------------------------

DROP TABLE IF EXISTS `CHANNEL`;

CREATE TABLE `CHANNEL` (
  `channel_id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL,
  PRIMARY KEY (`channel_id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `CHANNEL` WRITE;
/*!40000 ALTER TABLE `CHANNEL` DISABLE KEYS */;

INSERT INTO `CHANNEL` (`channel_id`, `name`)
VALUES
    (4,'DENEME'),
    (13,'IdsaNdSfooC32NM'),
    (11,'INdSfC32NM'),
    (12,'INdSfooC32NM'),
    (7,'INSC2NM'),
    (8,'INSC32NM'),
    (5,'INSCNM'),
    (9,'INSfC32NM');

/*!40000 ALTER TABLE `CHANNEL` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table CHANNEL_USER
# ------------------------------------------------------------

DROP TABLE IF EXISTS `CHANNEL_USER`;

CREATE TABLE `CHANNEL_USER` (
  `channel_user_id` int(11) NOT NULL AUTO_INCREMENT,
  `channel_id` int(11) NOT NULL,
  `comm_id` int(11) NOT NULL,
  `alias_id` int(11) NOT NULL,
  `is_owner` tinyint(1) NOT NULL,
  `user_id` int(11) NOT NULL,
  PRIMARY KEY (`channel_user_id`),
  KEY `CHANNEL_USER_fk0` (`channel_id`),
  KEY `CHANNEL_USER_fk1` (`comm_id`),
  KEY `CHANNEL_USER_fk2` (`alias_id`),
  KEY `user_id` (`is_owner`),
  KEY `user_id_2` (`user_id`),
  CONSTRAINT `CHANNEL_USER_fk0` FOREIGN KEY (`channel_id`) REFERENCES `CHANNEL` (`channel_id`),
  CONSTRAINT `CHANNEL_USER_fk1` FOREIGN KEY (`comm_id`) REFERENCES `COMM` (`comm_id`),
  CONSTRAINT `CHANNEL_USER_fk2` FOREIGN KEY (`alias_id`) REFERENCES `ALIAS` (`alias_id`),
  CONSTRAINT `channel_user_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `USERS` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `CHANNEL_USER` WRITE;
/*!40000 ALTER TABLE `CHANNEL_USER` DISABLE KEYS */;

INSERT INTO `CHANNEL_USER` (`channel_user_id`, `channel_id`, `comm_id`, `alias_id`, `is_owner`, `user_id`)
VALUES
    (13,4,19,1,1,31),
    (14,9,19,2,1,32),
    (15,11,19,3,1,32),
    (16,12,19,4,1,32),
    (17,13,19,5,1,32);

/*!40000 ALTER TABLE `CHANNEL_USER` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table COMM
# ------------------------------------------------------------

DROP TABLE IF EXISTS `COMM`;

CREATE TABLE `COMM` (
  `comm_id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `type_id` int(11) NOT NULL,
  `val` varchar(50) NOT NULL DEFAULT '',
  PRIMARY KEY (`comm_id`),
  KEY `COMM_fk0` (`user_id`),
  KEY `COMM_fk1` (`type_id`),
  CONSTRAINT `COMM_fk0` FOREIGN KEY (`user_id`) REFERENCES `USERS` (`user_id`),
  CONSTRAINT `COMM_fk1` FOREIGN KEY (`type_id`) REFERENCES `COMM_TYPE` (`type_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `COMM` WRITE;
/*!40000 ALTER TABLE `COMM` DISABLE KEYS */;

INSERT INTO `COMM` (`comm_id`, `user_id`, `type_id`, `val`)
VALUES
    (19,30,2,'cicekhu@gmail.com'),
    (20,30,1,'5389486646'),
    (21,31,2,'keskinsaf@gmail.com'),
    (22,32,2,'cicekhu@itu.edu.tr'),
    (23,33,2,'a@ta.kan');

/*!40000 ALTER TABLE `COMM` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table COMM_TYPE
# ------------------------------------------------------------

DROP TABLE IF EXISTS `COMM_TYPE`;

CREATE TABLE `COMM_TYPE` (
  `type_id` int(11) NOT NULL AUTO_INCREMENT,
  `val` varchar(256) NOT NULL DEFAULT '',
  PRIMARY KEY (`type_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `COMM_TYPE` WRITE;
/*!40000 ALTER TABLE `COMM_TYPE` DISABLE KEYS */;

INSERT INTO `COMM_TYPE` (`type_id`, `val`)
VALUES
    (1,'SMS'),
    (2,'EMAIL'),
    (3,'');

/*!40000 ALTER TABLE `COMM_TYPE` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table INTER
# ------------------------------------------------------------

DROP TABLE IF EXISTS `INTER`;

CREATE TABLE `INTER` (
  `interval_id` int(11) NOT NULL AUTO_INCREMENT,
  `preference_id` int(11) DEFAULT NULL,
  `start_time_in_minutes` int(11) NOT NULL,
  `duration` int(11) NOT NULL,
  PRIMARY KEY (`interval_id`),
  KEY `INTERVAL_fk0` (`preference_id`),
  CONSTRAINT `INTERVAL_fk0` FOREIGN KEY (`preference_id`) REFERENCES `PREFERENCE` (`preference_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `INTER` WRITE;
/*!40000 ALTER TABLE `INTER` DISABLE KEYS */;

INSERT INTO `INTER` (`interval_id`, `preference_id`, `start_time_in_minutes`, `duration`)
VALUES
    (1,16,1440,60),
    (2,16,2700,30),
    (3,21,21323,2132),
    (4,22,21323,2132),
    (5,23,21323,2132),
    (6,24,21323,2132),
    (7,25,21323,2132),
    (8,26,21323,2132);

/*!40000 ALTER TABLE `INTER` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table PREFERENCE
# ------------------------------------------------------------

DROP TABLE IF EXISTS `PREFERENCE`;

CREATE TABLE `PREFERENCE` (
  `preference_id` int(11) NOT NULL AUTO_INCREMENT,
  `duration_days` int(11) NOT NULL DEFAULT '7',
  `user_id` int(11) DEFAULT NULL,
  `channel_id` int(11) DEFAULT NULL,
  `start_date` bigint(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`preference_id`),
  KEY `PREFERENCE_fk1` (`channel_id`),
  KEY `PREFERENCE_fk0` (`user_id`),
  CONSTRAINT `PREFERENCE_fk0` FOREIGN KEY (`user_id`) REFERENCES `USERS` (`user_id`),
  CONSTRAINT `PREFERENCE_fk1` FOREIGN KEY (`channel_id`) REFERENCES `CHANNEL` (`channel_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `PREFERENCE` WRITE;
/*!40000 ALTER TABLE `PREFERENCE` DISABLE KEYS */;

INSERT INTO `PREFERENCE` (`preference_id`, `duration_days`, `user_id`, `channel_id`, `start_date`)
VALUES
    (16,7,30,NULL,0),
    (17,7,31,NULL,0),
    (18,7,32,NULL,0),
    (19,7,NULL,4,21323),
    (20,7,33,NULL,0),
    (21,323,NULL,7,2131231321),
    (22,323,NULL,8,2131231321),
    (23,323,NULL,9,2131231321),
    (24,323,NULL,11,2131231321),
    (25,323,NULL,12,2131231321),
    (26,323,NULL,13,2131231321);

/*!40000 ALTER TABLE `PREFERENCE` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table RESOURCE
# ------------------------------------------------------------

DROP TABLE IF EXISTS `RESOURCE`;

CREATE TABLE `RESOURCE` (
  `resource_id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL,
  `val` varchar(50) NOT NULL DEFAULT '',
  `language` varchar(50) NOT NULL,
  PRIMARY KEY (`resource_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table RESTRICTION
# ------------------------------------------------------------

DROP TABLE IF EXISTS `RESTRICTION`;

CREATE TABLE `RESTRICTION` (
  `restriction_id` int(11) NOT NULL AUTO_INCREMENT,
  `preference_id` int(11) NOT NULL,
  `type_id` int(11) NOT NULL,
  `val` varchar(256) NOT NULL DEFAULT '',
  `cont_op` varchar(256) NOT NULL DEFAULT '',
  `cont_type` varchar(256) NOT NULL DEFAULT '',
  PRIMARY KEY (`restriction_id`),
  KEY `RESTRICTION_fk0` (`preference_id`),
  KEY `RESTRICTION_fk1` (`type_id`),
  CONSTRAINT `RESTRICTION_fk0` FOREIGN KEY (`preference_id`) REFERENCES `PREFERENCE` (`preference_id`),
  CONSTRAINT `RESTRICTION_fk1` FOREIGN KEY (`type_id`) REFERENCES `COMM_TYPE` (`type_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `RESTRICTION` WRITE;
/*!40000 ALTER TABLE `RESTRICTION` DISABLE KEYS */;

INSERT INTO `RESTRICTION` (`restriction_id`, `preference_id`, `type_id`, `val`, `cont_op`, `cont_type`)
VALUES
    (2,19,1,'@gmail.com','=','END'),
    (3,23,1,'@itu.edu.tr','=','END'),
    (4,24,1,'@itu.edu.tr','=','END'),
    (5,25,1,'@itu.edu.tr','=','END'),
    (6,26,1,'@itu.edu.tr','=','END');

/*!40000 ALTER TABLE `RESTRICTION` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table TOKENS
# ------------------------------------------------------------

DROP TABLE IF EXISTS `TOKENS`;

CREATE TABLE `TOKENS` (
  `token_id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `token` varchar(256) NOT NULL DEFAULT '',
  PRIMARY KEY (`token_id`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `tokens_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `USERS` (`user_id`) ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table USERS
# ------------------------------------------------------------

DROP TABLE IF EXISTS `USERS`;

CREATE TABLE `USERS` (
  `user_id` int(11) NOT NULL AUTO_INCREMENT,
  `password` varchar(256) NOT NULL DEFAULT '',
  `username` varchar(256) DEFAULT '',
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `USERS` WRITE;
/*!40000 ALTER TABLE `USERS` DISABLE KEYS */;

INSERT INTO `USERS` (`user_id`, `password`, `username`)
VALUES
    (30,'$2a$10$imKx78kI49ltNe6EVsPa/uUXRcLrRQ6LzN5oPttpFMbCRAaN5D3uO','Atakan'),
    (31,'$2a$10$XP9014P4tx697ug78u/Y9.KpgYrTxSqaJzSdZXYWojTdCRtlEZJkq','Safa'),
    (32,'$2a$10$Lj0zeZR7wsfiz89.x8oO5uGLv0M4WO2MQNXRuBMZCq0eU7OMqHsqG','Low'),
    (33,'$2a$10$7eY0SKxH6jkLWXKZgDy1M.L8wNGE0foUCTae8IrQ5pOH5.5CGCfRW','');

/*!40000 ALTER TABLE `USERS` ENABLE KEYS */;
UNLOCK TABLES;



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
