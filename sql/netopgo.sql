-- MySQL dump 10.13  Distrib 5.5.47, for debian-linux-gnu (x86_64)
--
-- Host: localhost    Database: netopgo
-- ------------------------------------------------------
-- Server version	5.5.47-0ubuntu0.14.04.1

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
-- Table structure for table `all_size`
--

DROP TABLE IF EXISTS `all_size`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `all_size` (
  `timestamp` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `size` decimal(18,4) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `all_size`
--

LOCK TABLES `all_size` WRITE;
/*!40000 ALTER TABLE `all_size` DISABLE KEYS */;
INSERT INTO `all_size` VALUES ('2016-04-30 07:00:00',33.1200),('2016-05-01 07:00:00',34.2300),('2016-05-02 07:00:00',34.7800),('2015-05-02 07:00:00',32.0000),('2016-05-03 07:00:00',35.0000),('2016-05-04 07:00:00',34.2800),('2016-05-06 07:00:00',35.2100);
/*!40000 ALTER TABLE `all_size` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `db`
--

DROP TABLE IF EXISTS `db`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `db` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '',
  `uuid` varchar(255) NOT NULL DEFAULT '',
  `comment` varchar(255) NOT NULL DEFAULT '',
  `created` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `db`
--

LOCK TABLES `db` WRITE;
/*!40000 ALTER TABLE `db` DISABLE KEYS */;
INSERT INTO `db` VALUES (1,'dbmaster_conf','udb-rwsdk','lens_conf','2016-05-02 10:36:18'),(2,'dbslave1_conf','udc-skjdf9','lens_conf','2016-05-02 10:36:37');
/*!40000 ALTER TABLE `db` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `group`
--

DROP TABLE IF EXISTS `group`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `group` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '',
  `conment` varchar(255) NOT NULL DEFAULT '',
  `created` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `group`
--

LOCK TABLES `group` WRITE;
/*!40000 ALTER TABLE `group` DISABLE KEYS */;
INSERT INTO `group` VALUES (1,'flume','flume','2016-05-02 08:55:17'),(3,'amoeba','amoeba','2016-05-02 10:34:12'),(4,'mycat','mycat','2016-05-02 10:34:21');
/*!40000 ALTER TABLE `group` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `host`
--

DROP TABLE IF EXISTS `host`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `host` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '',
  `ip` varchar(255) NOT NULL DEFAULT '',
  `cpu` varchar(255) NOT NULL DEFAULT '',
  `mem` varchar(255) NOT NULL DEFAULT '',
  `disk` varchar(255) NOT NULL DEFAULT '',
  `idc` varchar(255) NOT NULL DEFAULT '',
  `rootpwd` varchar(255) NOT NULL DEFAULT '',
  `readpwd` varchar(255) NOT NULL DEFAULT '',
  `group` varchar(255) NOT NULL DEFAULT '',
  `created` datetime NOT NULL,
  `root` varchar(255) NOT NULL DEFAULT '',
  `read` varchar(255) NOT NULL DEFAULT '',
  `comment` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `host`
--

LOCK TABLES `host` WRITE;
/*!40000 ALTER TABLE `host` DISABLE KEYS */;
INSERT INTO `host` VALUES (1,'localhost','127.0.0.1','4核','8GB','1TB','Ucloud','88aJdGLcDQ==','88aJdGLcDQ==','flume','2016-05-02 08:55:17','','',''),(3,'amoeba','192.168.2.17','1核','2GB','1TB','百度云','6NSYLCGVSQ==','6NSYLCGVSQ==','amoeba','2016-05-02 17:33:56','root','root','amoeba'),(4,'mycat','192.168.2.18','1核','2GB','50GB','阿里云','6NSYLCGVSQ==','6NSYLCGVSQ==','mycat','2016-05-02 10:35:23','root','root','mycat');
/*!40000 ALTER TABLE `host` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `partition`
--

DROP TABLE IF EXISTS `partition`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `partition` (
  `schema` char(20) NOT NULL,
  `instance` varchar(50) NOT NULL,
  `timestamp` datetime NOT NULL,
  `count` int(11) DEFAULT NULL,
  `type` char(10) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `partition`
--

LOCK TABLES `partition` WRITE;
/*!40000 ALTER TABLE `partition` DISABLE KEYS */;
INSERT INTO `partition` VALUES ('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min'),('lens_conf','db0master_conf','2016-04-30 21:56:28',84,'min'),('lens_conf','db1master_conf','2016-04-30 21:56:34',83,'min');
/*!40000 ALTER TABLE `partition` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `schema`
--

DROP TABLE IF EXISTS `schema`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `schema` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `addr` varchar(255) NOT NULL DEFAULT '',
  `port` varchar(255) NOT NULL DEFAULT '',
  `name` varchar(255) NOT NULL DEFAULT '',
  `comment` varchar(255) NOT NULL DEFAULT '',
  `user` varchar(255) NOT NULL DEFAULT '',
  `passwd` varchar(255) NOT NULL DEFAULT '',
  `d_b_name` varchar(255) NOT NULL DEFAULT '',
  `partition` bigint(20) NOT NULL DEFAULT '0',
  `status` bigint(20) NOT NULL DEFAULT '0',
  `created` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `schema`
--

LOCK TABLES `schema` WRITE;
/*!40000 ALTER TABLE `schema` DISABLE KEYS */;
INSERT INTO `schema` VALUES (1,'192.168.2.83','3306','lens_conf','newlens','root','6NSYLCGVSQ==','newlens',0,1,'2016-05-02 23:41:05'),(2,'127.0.0.1','3306','netopgo','netopgo app','root','6NSYLCGVSQ==','netopgo',0,1,'2016-05-02 10:37:27'),(3,'192.168.2.17','8066','lens_mobapp_data','moapp data','lens','6NSYLCGVSQ==','lens_mobapp_data',84,2,'2016-05-02 10:38:21');
/*!40000 ALTER TABLE `schema` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '',
  `passwd` varchar(255) NOT NULL DEFAULT '',
  `email` varchar(255) NOT NULL DEFAULT '',
  `dept` varchar(255) NOT NULL DEFAULT '',
  `created` datetime NOT NULL,
  `auth` bigint(20) NOT NULL DEFAULT '0',
  `tel` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (1,'admin','bmJzMjAxMA==','admin@tingyun.com','op','2016-05-02 08:55:17',1,'18202808939'),(2,'dba','bmJzMjAxMA==','dba@tingyun.com','op','2016-05-02 08:55:17',2,'18202808939'),(3,'guest','bmJzMjAxMA==','guest@tingyun.com','op','2016-05-02 08:55:17',3,'18202808939');
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2016-05-02  3:39:17
