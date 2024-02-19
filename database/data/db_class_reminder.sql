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

-- Dumping structure for table db_class_reminder.tbl_blast_history
CREATE TABLE IF NOT EXISTS `tbl_blast_history` (
  `id` int NOT NULL AUTO_INCREMENT,
  `message` text NOT NULL,
  `user_success` int NOT NULL DEFAULT '0',
  `created_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=29 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table db_class_reminder.tbl_blast_history: ~8 rows (approximately)
INSERT INTO `tbl_blast_history` (`id`, `message`, `user_success`, `created_at`) VALUES
	(18, 'new data', 2, '2024-02-09 05:25:16'),
	(19, 'new blast', 1, '2024-02-09 05:28:02'),
	(20, 'new', 1, '2024-02-09 12:29:45'),
	(21, 'message new', 1, '2024-02-09 12:30:55'),
	(22, 'Selamat pagi,\n\nMau tahu banyak tentang dunia wirausaha? Ikuti seminar kewirausahaan GRATIS dengan tema “…” pada:\n\nTanggal: 31 Agustus 2023\n\nTempat: Ruang Auditorium Kampus B\n\nPembicara: Bunga Mawar (Pengusaha Muda omset 1 Milyar per bulan) dan Mama Eleanor (IRT sukses penghasilan 100 juta per bulan)\n\nDapatkan fasilitas penunjang seperti goodie bag langsung di lokasi! Salam.\n\nSilakan tambahkan emoticon yang perlu dan pantas saja ke dalam pesan tersebut.', 1, '2024-02-10 03:43:20'),
	(23, 'testing', 1, '2024-02-10 03:56:28'),
	(24, 'test', 1, '2024-02-10 04:02:18'),
	(25, 'hai ijab', 2, '2024-02-13 02:18:12'),
	(26, 'testing', 2, '2024-02-18 12:58:57'),
	(27, 'Selamat pagi,\n\nMau tahu banyak tentang dunia wirausaha? Ikuti seminar kewirausahaan GRATIS dengan tema “…” pada:\n\nTanggal: 31 Agustus 2023\n\nTempat: Ruang Auditorium Kampus B\n\nPembicara: Bunga Mawar (Pengusaha Muda omset 1 Milyar per bulan) dan Mama Eleanor (IRT sukses penghasilan 100 juta per bulan)\n\nDapatkan fasilitas penunjang seperti goodie bag langsung di lokasi! Salam.\n\nSilakan tambahkan emoticon yang perlu dan pantas saja ke dalam pesan tersebut.', 2, '2024-02-18 13:42:20'),
	(28, 'message new', 1, '2024-02-18 13:43:04');

-- Dumping structure for table db_class_reminder.tbl_content_notif
CREATE TABLE IF NOT EXISTS `tbl_content_notif` (
  `trx_type` varchar(50) NOT NULL DEFAULT '""',
  `title` varchar(50) DEFAULT NULL,
  `content` text,
  `additional_data` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`trx_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table db_class_reminder.tbl_content_notif: ~3 rows (approximately)
INSERT INTO `tbl_content_notif` (`trx_type`, `title`, `content`, `additional_data`) VALUES
	('Course', 'Notifikasi kelas', 'Hi $name ($nim), kelas $class mu akan dimulai besok pukul $time, yuk persiapkan kebutuhan dan bahan kuliahmu sekarang!', '{"name":"required","nim":"required","class":"required", "time":"required"}'),
	('Seminar', 'Notifikasi kelas', 'Hi, segera hadiri seminar $event besok dengan tema lingkungan.', '{"event":"required"}'),
	('TuitionFee', 'Notifikasi Pengingat UKT', 'Hai $name ($nim)! Yuk bayar tagihan kuliah mu:\r\nVA: $va_account\r\nJumlah tagihan: $bill\r\nJatuh tempo: $due_date\r\nHiraukan notifikasi ini apabila kamu sudah membayarnya.', '{"name":"required","nim":"required","va_account":"required","bill":"required","due_date":"required"}');

-- Dumping structure for table db_class_reminder.tbl_event
CREATE TABLE IF NOT EXISTS `tbl_event` (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(50) DEFAULT NULL,
  `description` varchar(50) DEFAULT NULL,
  `schedule` time DEFAULT NULL,
  `job_every` varchar(100) DEFAULT NULL,
  `event_day` varchar(100) DEFAULT NULL,
  `id_event_type` int DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=74 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table db_class_reminder.tbl_event: ~11 rows (approximately)
INSERT INTO `tbl_event` (`id`, `title`, `description`, `schedule`, `job_every`, `event_day`, `id_event_type`) VALUES
	(1, 'Algoritma Pemrograman', 'Algoritma Pemrograman', '17:53:06', 'TUE', 'WED', 1),
	(2, 'Seminar terbuka', 'Seminar terbuka', '17:53:06', 'TUE', 'WED', 2),
	(22, 'matematika teknik dasar kelas A', 'description', '10:19:00', 'TUE', 'WED', 1),
	(23, 'matematika teknik dasar kelas B', 'description', '10:21:01', 'TUE', 'WED', 1),
	(24, 'seminar khusus', 'description', '10:26:01', 'TUE', 'WED', 2),
	(27, 'seminar khusus', 'description', '10:26:01', 'TUE', 'WED', 2),
	(41, 'Bahasa Indonesia kelas A', 'Bahasa Indonesia kelas A', '23:17:00', 'TUE', 'WED', 1),
	(42, 'Bahasa Indonesia kelas B', 'Bahasa Indonesia kelas B', '23:20:00', 'TUE', 'WED', 1),
	(54, 'Pemrograman dasar kelas A', 'Pemrograman dasar kelas A', '08:35:00', 'THU', 'FRI', 1),
	(60, 'Tuition Fee Reminder', 'Tuition Fee Reminder', '09:39:01', '1', NULL, 2),
	(73, 'Payment reminder', 'Payment reminder', '12:12:00', '2', NULL, 3);

-- Dumping structure for table db_class_reminder.tbl_event_type
CREATE TABLE IF NOT EXISTS `tbl_event_type` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `is_specific_user` int DEFAULT '0',
  `trx_type` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `trx_type` (`trx_type`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table db_class_reminder.tbl_event_type: ~2 rows (approximately)
INSERT INTO `tbl_event_type` (`id`, `name`, `is_specific_user`, `trx_type`) VALUES
	(1, 'Course', 1, 'Course'),
	(2, 'Seminar', 0, 'Seminar'),
	(3, 'Tuition Fee reminder', 0, 'TuitionFee');

-- Dumping structure for table db_class_reminder.tbl_job
CREATE TABLE IF NOT EXISTS `tbl_job` (
  `id` int NOT NULL AUTO_INCREMENT,
  `job_name` varchar(300) DEFAULT NULL,
  `job_id` varchar(100) DEFAULT NULL,
  `id_event` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `id_event` (`id_event`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table db_class_reminder.tbl_job: ~13 rows (approximately)
INSERT INTO `tbl_job` (`id`, `job_name`, `job_id`, `id_event`) VALUES
	(9, '54:Pemrograman_dasar_kelas_A', '9a83cad2-3a57-4cb0-8b94-c035e34dec8a', 54),
	(10, '60:Tuition_Fee_Reminder', 'cc5492eb-6b2d-414b-ab8f-7c14395712fd', 60),
	(11, '63:job_scheduler', '0ce7ce54-3b6e-402a-bf76-c4c4fa945c67', 63),
	(12, '64:scheduler_new', '0db6213d-8bd6-42a9-b772-f528135a5399', 64),
	(13, '65:title_job', '4871c300-e6c0-44e8-98a1-308dbb644409', 65),
	(14, '66:Tuition_Fee_Reminder', '211ae823-fdd8-47ce-80a0-14144702a67e', 66),
	(15, '67:eifi', '66cf06d1-0f82-457d-a127-3be63e95b59b', 67),
	(16, '68:title', '78c6ba9e-ab46-4690-a994-52303132a884', 68),
	(17, '69:new', 'c3cb0214-4410-433c-b820-e5c65944e9ca', 69),
	(18, '70:tes', 'bbca3a26-7cbf-4ff3-bd1d-481f831f4eca', 70),
	(19, '71:title', '7a865915-5c05-449e-a07a-f1f3a1b730e0', 71),
	(20, '72:title', '469a70f4-a3ad-4166-aa95-4902ad39f434', 72),
	(21, '73:Payment_reminder', 'f03e6405-0985-4446-ba1d-e20be07e9338', 73);

-- Dumping structure for table db_class_reminder.tbl_trx_log
CREATE TABLE IF NOT EXISTS `tbl_trx_log` (
  `id` int NOT NULL AUTO_INCREMENT,
  `id_event` int DEFAULT NULL,
  `user_success` int DEFAULT NULL,
  `trx_type` varchar(50) DEFAULT NULL,
  `trx_date` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=26 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table db_class_reminder.tbl_trx_log: ~17 rows (approximately)
INSERT INTO `tbl_trx_log` (`id`, `id_event`, `user_success`, `trx_type`, `trx_date`) VALUES
	(8, 54, 2, 'Course', '2024-02-17 11:16:34'),
	(9, 54, 2, 'Course', '2024-02-17 12:13:01'),
	(10, 1, 5, 'Course', '2024-02-17 15:19:11'),
	(11, 1, 0, 'Course', '2024-02-17 17:00:15'),
	(12, 55, 1, 'TuitionFee', '2024-02-17 17:44:36'),
	(13, 55, 0, 'TuitionFee', '2024-02-17 17:46:04'),
	(14, 55, 1, 'TuitionFee', '2024-02-17 17:46:43'),
	(15, 55, 1, 'TuitionFee', '2024-02-17 17:49:55'),
	(16, 55, 1, 'TuitionFee', '2024-02-17 17:52:18'),
	(17, 55, 1, 'TuitionFee', '2024-02-17 19:15:12'),
	(18, 60, 1, 'TuitionFee', '2024-02-17 19:36:54'),
	(19, 64, 1, 'TuitionFee', '2024-02-18 05:20:05'),
	(20, 63, 2, 'Seminar', '2024-02-18 05:20:16'),
	(21, 73, 1, 'TuitionFee', '2024-02-18 07:26:01'),
	(22, 73, 1, 'TuitionFee', '2024-02-18 07:26:25'),
	(23, 73, 1, 'TuitionFee', '2024-02-18 07:32:22'),
	(24, 73, 1, 'TuitionFee', '2024-02-18 07:39:29'),
	(25, 1, 0, 'Course', '2024-02-18 15:19:03');

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
	(1, 'usertes', '$2a$12$FUiKbGDrpc.xs6lHM/LIEuYA3/p1hQTaUAuILfYoZ3cay7w6Ampie', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDgyNzc2NjYsInVzZXJuYW1lIjoiVXNlcnRlcyJ9.HpSOvmfKE61zoKD1dTb-2nU3XNjBLz7f8zG2dlFw1lE', '2024-02-18 15:34:26'),
	(2, 'usertes2', '$2a$12$FUiKbGDrpc.xs6lHM/LIEuYA3/p1hQTaUAuILfYoZ3cay7w6Ampie', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTc5NzY4NzMsInVzZXJuYW1lIjoidXNlcnRlczIifQ.RI_QqkCTm_NlbS60DZjR40-jW_9FqJcC9xajWipFvMg', '2023-10-22 12:14:23');

-- Dumping structure for table db_class_reminder.tbl_user_event
CREATE TABLE IF NOT EXISTS `tbl_user_event` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(50) DEFAULT NULL,
  `nim` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `id_event` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `username` (`username`),
  KEY `nim` (`nim`)
) ENGINE=InnoDB AUTO_INCREMENT=71 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table db_class_reminder.tbl_user_event: ~7 rows (approximately)
INSERT INTO `tbl_user_event` (`id`, `username`, `nim`, `id_event`) VALUES
	(1, 'usertes', NULL, 1),
	(3, 'usertes', NULL, 24),
	(4, 'usertes2', NULL, 1),
	(66, '3332170020', '3332170020', 41),
	(67, '3332170020', '3332170020', 54),
	(68, '3332170055', '3332170055', 1),
	(69, '3332170055', '3332170055', 54);

-- Dumping structure for table db_class_reminder.tbl_user_notif
CREATE TABLE IF NOT EXISTS `tbl_user_notif` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(50) DEFAULT NULL,
  `nim` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `notif_id` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `no_hp` varchar(20) DEFAULT NULL,
  `last_update` timestamp NULL DEFAULT NULL,
  `is_allowed` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `username` (`username`),
  KEY `nim` (`nim`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table db_class_reminder.tbl_user_notif: ~4 rows (approximately)
INSERT INTO `tbl_user_notif` (`id`, `username`, `nim`, `notif_id`, `no_hp`, `last_update`, `is_allowed`) VALUES
	(16, '3332170020', '3332170020', NULL, '087871855339', '2024-02-17 15:48:20', 1),
	(20, '3332170055', '3332170055', NULL, '0802391293312', '2024-02-17 15:48:53', 1);

-- Dumping structure for table db_class_reminder.tbl_user_payment
CREATE TABLE IF NOT EXISTS `tbl_user_payment` (
  `id` int NOT NULL AUTO_INCREMENT,
  `nim` varchar(50) DEFAULT NULL,
  `bill` int DEFAULT NULL,
  `va_account` varchar(50) DEFAULT NULL,
  `last_payment_date` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table db_class_reminder.tbl_user_payment: ~3 rows (approximately)
INSERT INTO `tbl_user_payment` (`id`, `nim`, `bill`, `va_account`, `last_payment_date`) VALUES
	(1, '3332170020', 4999988, '123123212', '2024-01-17 14:14:46'),
	(5, '3332170055', 4500000, '1231093019', '2024-02-17 15:48:53');

-- Dumping structure for table db_class_reminder.tbl_user_student
CREATE TABLE IF NOT EXISTS `tbl_user_student` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `nim` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `major` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `nim` (`nim`),
  KEY `username` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=39 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table db_class_reminder.tbl_user_student: ~3 rows (approximately)
INSERT INTO `tbl_user_student` (`id`, `username`, `nim`, `name`, `major`) VALUES
	(31, '3332170020', '3332170020', 'badru', 'elektro'),
	(37, '3332170055', '3332170055', 'rizal', 'IT');

/*!40103 SET TIME_ZONE=IFNULL(@OLD_TIME_ZONE, 'system') */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
