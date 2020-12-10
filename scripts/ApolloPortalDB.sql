/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

# Create Database
# ------------------------------------------------------------
CREATE DATABASE IF NOT EXISTS dida_apollo_plus_portal DEFAULT CHARACTER SET = utf8mb4;

Use dida_apollo_plus_portal;

# Dump of table app
# ------------------------------------------------------------

DROP TABLE IF EXISTS `Role`;

CREATE TABLE `Role` (
  `Id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `AppId` varchar(500) NOT NULL DEFAULT 'default' COMMENT 'AppID',
  `UserId` varchar(100) NOT NULL DEFAULT 'default' COMMENT 'UserId',
  `Namespace` varchar(100) NOT NULL DEFAULT 'application' COMMENT 'UserName',
  `Cluster`   varchar(64) NOT NULL DEFAULT 'default' COMMENT 'Cluster',
  `UserName` varchar(100) NOT NULL DEFAULT 'default' COMMENT 'UserName',
  `Level` int(10) NOT NULL DEFAULT 0 COMMENT '权限级别',
  `IsDeleted` tinyint(1) NOT NULL DEFAULT b'0' COMMENT '1: deleted, 0: normal',
  `DataChange_CreatedBy` varchar(32) NOT NULL DEFAULT 'default' COMMENT '创建人邮箱前缀',
  `DataChange_CreatedTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`Id`),
  KEY `AppId`(`AppId`),
  KEY `UserId` (`UserId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='权限表';





use dida_apollo_plus_portal;
show tables ;
desc Role;
select * from Role;
