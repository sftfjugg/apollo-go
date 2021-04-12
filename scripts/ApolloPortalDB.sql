/*!40101 SET @OLD_CHARACTER_SET_CLIENT = @@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS = @@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION = @@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS = @@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS = 0 */;
/*!40101 SET @OLD_SQL_MODE = @@SQL_MODE, SQL_MODE = 'NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES = @@SQL_NOTES, SQL_NOTES = 0 */;

# Create Database
# ------------------------------------------------------------
CREATE DATABASE IF NOT EXISTS dida_apollo_plus_portal DEFAULT CHARACTER SET = utf8mb4;

Use dida_apollo_plus_portal;

# Dump of table app
# ------------------------------------------------------------

DROP TABLE IF EXISTS `Role`;

CREATE TABLE `Role`
(
    `Id`                     int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
    `AppId`                  varchar(64)      NOT NULL DEFAULT 'default' COMMENT 'AppID',
    `UserId`                 varchar(100)     NOT NULL DEFAULT 'default' COMMENT 'UserId',
    `Namespace`              varchar(100)     NOT NULL DEFAULT 'application' COMMENT 'UserName',
    `Env`                    varchar(64)      NOT NULL DEFAULT 'TEST' COMMENT '环境',
    `Cluster`                varchar(64)      NOT NULL DEFAULT 'default' COMMENT 'Cluster',
    `UserName`               varchar(100)     NOT NULL DEFAULT 'default' COMMENT 'UserName',
    `Level`                  int(10)          NOT NULL DEFAULT 0 COMMENT '权限级别',
    `IsDeleted`              tinyint(1)       NOT NULL DEFAULT b'0' COMMENT '1: deleted, 0: normal',
    `DataChange_CreatedBy`   varchar(32)      NOT NULL DEFAULT 'default' COMMENT '创建人邮箱前缀',
    `DataChange_CreatedTime` timestamp        NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`Id`),
    KEY `AppId` (`AppId`),
    KEY `UserId` (`UserId`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='权限表';

DROP TABLE IF EXISTS `History`;

CREATE TABLE `History`
(
    `Id`                     int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
    `AppId`                  varchar(64)      NOT NULL DEFAULT 'default' COMMENT 'AppID',
    `UserId`                 varchar(100)     NOT NULL DEFAULT 'default' COMMENT 'UserId',
    `UserName`               varchar(100)     NOT NULL DEFAULT 'default' COMMENT 'UserName',
    `IsDeleted`              tinyint(1)       NOT NULL DEFAULT b'0' COMMENT '1: deleted, 0: normal',
    `DataChange_CreatedTime` timestamp        NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`Id`),
    KEY `AppId` (`AppId`),
    KEY `UserId` (`UserId`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='浏览记录';


DROP TABLE IF EXISTS `Dingding`;

CREATE TABLE `Dingding`
(
    `Id`                     int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
    `Name`                   varchar(64)      NOT NULL DEFAULT 'default' COMMENT 'Name',
    `AppId`                  varchar(64)      NOT NULL DEFAULT 'default' COMMENT 'AppID',
    `Env`                    varchar(64)      NOT NULL DEFAULT 'TEST' COMMENT 'Env',
    `DeptName`               varchar(64)      NOT NULL DEFAULT 'default' COMMENT 'DeptName',
    `Type`                   varchar(64)      NOT NULL DEFAULT 'default' COMMENT 'Type',
    `Token`                  varchar(100)     NOT NULL DEFAULT '' COMMENT 'Token',
    `Level`                  int(10)          NOT NULL DEFAULT '0' COMMENT 'Level',
    `IsDeleted`              tinyint(1)       NOT NULL DEFAULT b'0' COMMENT '1: deleted, 0: normal',
    `DataChange_CreatedBy`   varchar(32)      NOT NULL DEFAULT 'default' COMMENT '创建人邮箱前缀',
    `DataChange_CreatedTime` timestamp        NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`Id`),
    KEY `Level` (`Level`),
    Unique `Name` (`Name`),
    KEY `Type` (`Type`),
    KEY `AppId` (`AppId`),
    KEY `DeptName` (`DeptName`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='钉钉';



use dida_apollo_plus_portal;
select *
from Role
where AppId = 'root';
show tables;
desc Role;
select *
from Role;
delete
from Role
where AppId = 'root';



USE dida_apollo_plus_config;
alter table AppNamespace
    add DeptName  varchar(64) DEFAULT '' COMMENT '部门名字' after LaneName,
    add IsDisplay tinyint(1) NOT NULL DEFAULT b'1' COMMENT '0:Hide , 1: Dispaly' after IsDeleted

USE dida_apollo_plus_portal;
select *
from Role
where (AppId = 'aim-mapboundary' and UserId = 'wangkun' and IsDeleted = 0 and Cluster = 'default' and Env = 'ALIYUN')
   or (AppId = 'root' and IsDeleted = 0 and UserId = 'wangkun')

show databases;
use plat_metis;
Use dida_sentinel_role;
show tables;
select *
from Role;
select *
from History;
select *
from `Role`;

SELECT a.id + 1 AS START, MIN(b.id) - 1 AS END
FROM Role AS a,
     Role AS b
WHERE a.id < b.id
GROUP BY a.id
HAVING START < MIN(b.id)

