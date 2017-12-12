# ************************************************************
# Sequel Pro SQL dump
# Version 4541
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 127.0.0.1 (MySQL 5.7.20)
# Database: ChannelX
# Generation Time: 2017-12-10 12:57:38 +0000
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

CREATE TABLE `ALIAS` (
  `alias_id` int(11) NOT NULL AUTO_INCREMENT,
  `val` varchar(50) NOT NULL DEFAULT '',
  `is_user_defined` tinyint(1) NOT NULL,
  PRIMARY KEY (`alias_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table CHANNEL
# ------------------------------------------------------------

CREATE TABLE `CHANNEL` (
  `channel_id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL,
  PRIMARY KEY (`channel_id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table CHANNEL_USER
# ------------------------------------------------------------

CREATE TABLE `CHANNEL_USER` (
  `channel_user_id` int(11) NOT NULL AUTO_INCREMENT,
  `channel_id` int(11) NOT NULL,
  `comm_id` int(11) NOT NULL,
  `alias_id` int(11) NOT NULL,
  `is_owner` tinyint(1) NOT NULL,
  `user_id` int(11) NOT NULL,
  `token` varchar(256) DEFAULT NULL,
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



# Dump of table COMM
# ------------------------------------------------------------

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



# Dump of table COMM_TYPE
# ------------------------------------------------------------

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
	(2,'EMAIL');

/*!40000 ALTER TABLE `COMM_TYPE` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table INTER
# ------------------------------------------------------------

CREATE TABLE `INTER` (
  `interval_id` int(11) NOT NULL AUTO_INCREMENT,
  `preference_id` int(11) DEFAULT NULL,
  `start_time_in_minutes` int(11) NOT NULL,
  `duration` int(11) NOT NULL,
  PRIMARY KEY (`interval_id`),
  KEY `INTERVAL_fk0` (`preference_id`),
  CONSTRAINT `INTERVAL_fk0` FOREIGN KEY (`preference_id`) REFERENCES `PREFERENCE` (`preference_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table PREFERENCE
# ------------------------------------------------------------

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



# Dump of table RESOURCE
# ------------------------------------------------------------

CREATE TABLE `RESOURCE` (
  `resource_id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL,
  `val` varchar(50) NOT NULL DEFAULT '',
  `language` varchar(50) NOT NULL,
  PRIMARY KEY (`resource_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table RESTRICTION
# ------------------------------------------------------------

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



# Dump of table TOKENS
# ------------------------------------------------------------

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

CREATE TABLE `USERS` (
  `user_id` int(11) NOT NULL AUTO_INCREMENT,
  `password` varchar(256) NOT NULL DEFAULT '',
  `username` varchar(256) DEFAULT '',
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;




/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
