-- MySQL dump 10.13  Distrib 5.7.33, for Win64 (x86_64)
--
-- Host: localhost    Database: xyz_multifinance
-- ------------------------------------------------------
-- Server version	5.7.24

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `customer_installments`
--

DROP TABLE IF EXISTS `customer_installments`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `customer_installments` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `customer_transaction_id` int(11) DEFAULT NULL,
  `customer_limit_id` int(11) DEFAULT NULL,
  `tenor` int(11) DEFAULT NULL,
  `total_amounts` float DEFAULT NULL,
  `remaining_amounts` float DEFAULT NULL,
  `created_at` timestamp NOT NULL,
  `updated_at` timestamp NOT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `customer_transaction_id` (`customer_transaction_id`),
  KEY `customer_limit_id` (`customer_limit_id`),
  CONSTRAINT `customer_installments_ibfk_1` FOREIGN KEY (`customer_transaction_id`) REFERENCES `customer_transactions` (`id`),
  CONSTRAINT `customer_installments_ibfk_2` FOREIGN KEY (`customer_limit_id`) REFERENCES `customer_limits` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `customer_installments`
--

LOCK TABLES `customer_installments` WRITE;
/*!40000 ALTER TABLE `customer_installments` DISABLE KEYS */;
/*!40000 ALTER TABLE `customer_installments` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `customer_limits`
--

DROP TABLE IF EXISTS `customer_limits`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `customer_limits` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `customer_id` int(11) DEFAULT NULL,
  `limit_1` float DEFAULT NULL,
  `limit_2` float DEFAULT NULL,
  `limit_3` float DEFAULT NULL,
  `limit_4` float DEFAULT NULL,
  `remaining_limit` float DEFAULT NULL,
  `created_at` timestamp NOT NULL,
  `updated_at` timestamp NOT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `customer_id` (`customer_id`),
  CONSTRAINT `customer_limits_ibfk_1` FOREIGN KEY (`customer_id`) REFERENCES `customers` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `customer_limits`
--

LOCK TABLES `customer_limits` WRITE;
/*!40000 ALTER TABLE `customer_limits` DISABLE KEYS */;
INSERT INTO `customer_limits` VALUES (1,1,100000,200000,500000,700000,NULL,'2023-07-26 16:37:47','2023-07-26 16:37:47',NULL),(2,2,1000000,1200000,1500000,2000000,NULL,'2023-07-26 16:37:47','2023-07-26 16:37:47',NULL);
/*!40000 ALTER TABLE `customer_limits` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `customer_transactions`
--

DROP TABLE IF EXISTS `customer_transactions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `customer_transactions` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `customer_id` int(11) DEFAULT NULL,
  `contract_number` varchar(255) DEFAULT NULL,
  `otr_price` float DEFAULT NULL,
  `admin_fee` float DEFAULT NULL,
  `installment_amount` float DEFAULT NULL,
  `interest_amount` float DEFAULT NULL,
  `asset_name` varchar(255) DEFAULT NULL,
  `status` varchar(50) DEFAULT NULL,
  `created_at` timestamp NOT NULL,
  `updated_at` timestamp NOT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `customer_id` (`customer_id`),
  CONSTRAINT `customer_transactions_ibfk_1` FOREIGN KEY (`customer_id`) REFERENCES `customers` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `customer_transactions`
--

LOCK TABLES `customer_transactions` WRITE;
/*!40000 ALTER TABLE `customer_transactions` DISABLE KEYS */;
INSERT INTO `customer_transactions` VALUES (1,1,'trx-1',10000,5000,105000,105000,'Kipas angin','PENDING','2023-07-26 23:05:32','2023-07-26 23:05:32',NULL),(2,1,'trx-1',1000,5000,105000,105000,'Kipas angin','PENDING','2023-07-26 23:10:09','2023-07-26 23:10:09',NULL),(3,1,'trx-1',1000,5000,105000,105000,'Kipas angin','PENDING','2023-07-26 23:12:30','2023-07-26 23:12:30',NULL);
/*!40000 ALTER TABLE `customer_transactions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `customers`
--

DROP TABLE IF EXISTS `customers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `customers` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `nik` varchar(25) NOT NULL,
  `full_name` varchar(200) NOT NULL,
  `legal_name` varchar(200) NOT NULL,
  `birth_place` varchar(200) NOT NULL,
  `birth_date` date NOT NULL,
  `salary` int(11) NOT NULL,
  `ktp_photo` varchar(255) NOT NULL,
  `selfie_photo` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL,
  `updated_at` timestamp NOT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `customers`
--

LOCK TABLES `customers` WRITE;
/*!40000 ALTER TABLE `customers` DISABLE KEYS */;
INSERT INTO `customers` VALUES (1,'123456789123','Budi Budiman','Budi Budiman','Jakarta','1990-10-10',15000000,'url_ktp_photo','url_selfie_photo','2023-07-26 16:37:46','2023-07-26 16:37:46',NULL),(2,'789789123','Annisa Nissa','Annisa Nissa','Surabata','1997-07-01',10000000,'url_ktp_photo','url_selfie_photo','2023-07-26 16:37:46','2023-07-26 16:37:46',NULL);
/*!40000 ALTER TABLE `customers` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping routines for database 'xyz_multifinance'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-07-27 14:10:31
