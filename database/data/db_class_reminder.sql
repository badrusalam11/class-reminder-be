-- --------------------------------------------------------
-- Host:                         127.0.0.1
-- Server version:               8.0.30 - MySQL Community Server - GPL
-- Server OS:                    Win64
-- HeidiSQL Version:             12.1.0.6537
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


-- Dumping database structure for db_class_reminder
CREATE DATABASE IF NOT EXISTS `db_class_reminder` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `db_class_reminder`;

-- Dumping structure for table db_class_reminder.tbl_event
CREATE TABLE IF NOT EXISTS `tbl_event` (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(50) DEFAULT NULL,
  `description` varchar(50) DEFAULT NULL,
  `schedule` time DEFAULT NULL,
  `job_every` varchar(100) DEFAULT NULL,
  `id_event_type` int DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=32 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table db_class_reminder.tbl_event: ~10 rows (approximately)
INSERT INTO `tbl_event` (`id`, `title`, `description`, `schedule`, `job_every`, `id_event_type`) VALUES
	(1, 'Algoritma Pemrograman', 'Algoritma Pemrograman', '17:53:06', '{"MON":true,"SUN":true,"TUE":true}', 1),
	(2, 'Seminar terbuka', 'Seminar terbuka', '17:53:06', '{"MON":true,"SUN":true,"TUE":true}', 2),
	(22, 'matematika teknik dasar', 'description', '10:19:00', '{"MON":true,"SUN":true,"TUE":true}', 1),
	(23, 'matematika teknik dasar', 'description', '10:21:01', '{"MON":true,"SUN":true,"TUE":true}', 1),
	(24, 'seminar khusus', 'description', '10:26:01', '{"MON":true,"SUN":true,"TUE":true}', 2),
	(27, 'seminar khusus', 'description', '10:26:01', '{"MON":true,"SUN":true,"TUE":true}', 2),
	(28, 'seminar khusus', 'description', '10:26:01', '{"MON":true,"SUN":true,"TUE":true}', 2),
	(29, 'seminar khusus', 'description', '11:26:01', '{"MON":true,"SUN":true,"TUE":true}', 2),
	(30, 'seminar khusus', 'description', '11:35:01', '{"MON":true,"SUN":true,"TUE":true}', 2),
	(31, 'seminar khusus', 'description', '09:39:01', '{"MON":true,"SUN":true,"TUE":true}', 2);

-- Dumping structure for table db_class_reminder.tbl_event_type
CREATE TABLE IF NOT EXISTS `tbl_event_type` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `is_specific_user` int DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table db_class_reminder.tbl_event_type: ~2 rows (approximately)
INSERT INTO `tbl_event_type` (`id`, `name`, `is_specific_user`) VALUES
	(1, 'course', 1),
	(2, 'seminar', 0);

-- Dumping structure for table db_class_reminder.tbl_user
CREATE TABLE IF NOT EXISTS `tbl_user` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `password` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `token_key` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `last_login` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `username` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table db_class_reminder.tbl_user: ~2 rows (approximately)
INSERT INTO `tbl_user` (`id`, `username`, `password`, `token_key`, `last_login`) VALUES
	(1, 'usertes', '$2a$12$FUiKbGDrpc.xs6lHM/LIEuYA3/p1hQTaUAuILfYoZ3cay7w6Ampie', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTgwMzYxOTQsInVzZXJuYW1lIjoidXNlcnRlcyJ9.Gd0ObuzIpk1h6WBCfWvJNiaJvo6CINa5Jawh_YJ9YK0', '2023-10-23 02:43:14'),
	(2, 'usertes2', '$2a$12$FUiKbGDrpc.xs6lHM/LIEuYA3/p1hQTaUAuILfYoZ3cay7w6Ampie', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTc5NzY4NzMsInVzZXJuYW1lIjoidXNlcnRlczIifQ.RI_QqkCTm_NlbS60DZjR40-jW_9FqJcC9xajWipFvMg', '2023-10-22 12:14:23');

-- Dumping structure for table db_class_reminder.tbl_user_event
CREATE TABLE IF NOT EXISTS `tbl_user_event` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(50) DEFAULT NULL,
  `id_event` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `username` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table db_class_reminder.tbl_user_event: ~1 rows (approximately)
INSERT INTO `tbl_user_event` (`id`, `username`, `id_event`) VALUES
	(1, 'usertes', 1);

-- Dumping structure for table db_class_reminder.tbl_user_notif
CREATE TABLE IF NOT EXISTS `tbl_user_notif` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(50) DEFAULT NULL,
  `notif_id` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `last_update` timestamp NULL DEFAULT NULL,
  `is_allowed` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `username` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table db_class_reminder.tbl_user_notif: ~2 rows (approximately)
INSERT INTO `tbl_user_notif` (`id`, `username`, `notif_id`, `last_update`, `is_allowed`) VALUES
	(1, 'usertes', 'eZV3faglj1TWBNGJ9ChLb7:APA91bE1IMlRUXhC5sR2e81n9brLW0cTmhYmb9WRVOiphxQFyQqga3L4B820WK5JqMdSjy5qyEjwkmp4aAFbmuS0MITGWNY3CVX3VG4YD25m5xhX-3T7s591_P4kHcIoHY2yWW2VeC4G', '2023-10-23 02:43:23', 0),
	(2, 'usertes2', 'eu-_w6Ipl4QDzkhFAj7XOc:APA91bGbndp-crXIRCidajiEY1zOwUpJSgNgx0eWgylcJiaGGxFnzhwuJW3RIIauAomYJYy7GweVRC_Idp467-ATP7mcXXY80mHUKjjunFfD5JZ5t5ZEUTdTJ38VBT59FncK1M-1H6lW', '2023-10-22 12:00:32', 0);

-- Dumping structure for table db_class_reminder.tbl_user_student
CREATE TABLE IF NOT EXISTS `tbl_user_student` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `nim` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `major` varchar(50) DEFAULT NULL,
  `total_sks` int DEFAULT NULL,
  `ipk` float DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `username` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table db_class_reminder.tbl_user_student: ~1 rows (approximately)
INSERT INTO `tbl_user_student` (`id`, `username`, `nim`, `name`, `major`, `total_sks`, `ipk`) VALUES
	(1, 'usertes', '3332170078', 'Saipul jamil', 'S1 - Sistem informasi', 138, 3.15);

/*!40103 SET TIME_ZONE=IFNULL(@OLD_TIME_ZONE, 'system') */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
