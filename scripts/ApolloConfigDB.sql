/*!40101 SET @OLD_CHARACTER_SET_CLIENT = @@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS = @@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION = @@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS = @@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS = 0 */;
/*!40101 SET @OLD_SQL_MODE = @@SQL_MODE, SQL_MODE = 'NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES = @@SQL_NOTES, SQL_NOTES = 0 */;

# Create Database
# ------------------------------------------------------------
CREATE DATABASE IF NOT EXISTS dida_apollo_plus_config DEFAULT CHARACTER SET = utf8mb4;


# Dump of table appnamespace
# ------------------------------------------------------------

USE dida_apollo_plus_config;

DROP TABLE IF EXISTS `AppNamespace`;

CREATE TABLE `AppNamespace`
(
    `Id`                        int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
    `Name`                      varchar(64)      NOT NULL DEFAULT '' COMMENT 'namespace名字',
    `AppName`                   varchar(64)      NOT NULL DEFAULT '' COMMENT '项目名字',
    `AppId`                     varchar(64)      NOT NULL DEFAULT '' COMMENT 'app id',
    `Format`                    varchar(32)      NOT NULL DEFAULT 'properties' COMMENT 'namespace的format类型',
    `IsPublic`                  tinyint(1)       NOT NULL DEFAULT b'0' COMMENT 'namespace是否为公共',
    `ClusterName`               varchar(64)      NOT NULL DEFAULT 'default' COMMENT 'Cluster Name',
    `LaneName`                  varchar(64)      NOT NULL DEFAULT 'default' COMMENT '泳道名字',
    `DeptName`                  varchar(64)      NOT NULL DEFAULT 'default' COMMENT '部门名字',
    `Comment`                   varchar(500)     NOT NULL DEFAULT '' COMMENT '注释',
    `IsDeleted`                 tinyint(1)       NOT NULL DEFAULT b'0' COMMENT '1: deleted, 0: normal',
    `IsOperate`                 tinyint(1)       NOT NULL DEFAULT b'0' COMMENT '是否是operate创建的服务',
    `IsDisplay`                 tinyint(1)       NOT NULL DEFAULT b'0' COMMENT '1: deleted, 0: normal',
    `IsRelease`                 tinyint(1)       NOT NULL DEFAULT b'0' COMMENT '0: 未发布, 1: 已发布',
    `DataChange_CreatedBy`      varchar(32)      NOT NULL DEFAULT '' COMMENT '创建人邮箱前缀',
    `DataChange_CreatedTime`    timestamp        NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `DataChange_LastModifiedBy` varchar(32)               DEFAULT '' COMMENT '最后修改人邮箱前缀',
    `DataChange_LastTime`       timestamp        NULL     DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
    PRIMARY KEY (`Id`),
    KEY `IX_AppId` (`AppId`),
    KEY `ClusterName` (`ClusterName`(64)),
    KEY `IsPublic` (`IsPublic`),
    KEY `Name_AppId` (`Name`, `AppId`),
    KEY `LaneName` (`ClusterName`(64)),
    KEY `AppId_ClusterName_Name` (`AppId`, `ClusterName`(64), `Name`),
    KEY `DataChange_LastTime` (`DataChange_LastTime`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='应用namespace定义';



# Dump of table item
# ------------------------------------------------------------

DROP TABLE IF EXISTS `Item`;

CREATE TABLE `Item`
(
    `Id`                        int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增Id',
    `NamespaceId`               int(10) unsigned NOT NULL DEFAULT '0' COMMENT '集群NamespaceId',
    `Key`                       varchar(128)     NOT NULL DEFAULT 'default' COMMENT '配置项Key',
    `Value`                     longtext         NOT NULL COMMENT '配置项值',
    `ReleaseValue`              longtext         NOT NULL COMMENT '发布的配置项值',
    `Status`                    tinyint(1)       NOT NULL DEFAULT b'0' COMMENT '当前状态，0：未发布(新增),1：已发布,2修改（未发布）3：删除',
    `Comment`                   varchar(1024)             DEFAULT '' COMMENT '标签',
    `Describe`                  varchar(1024)             DEFAULT '' COMMENT '详细描述',
    `IsDeleted`                 tinyint(1)       NOT NULL DEFAULT b'0' COMMENT '1: deleted, 0: normal',
    `DataChange_CreatedBy`      varchar(32)      NOT NULL DEFAULT 'default' COMMENT '创建人邮箱前缀',
    `DataChange_CreatedTime`    timestamp        NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `DataChange_LastModifiedBy` varchar(32)               DEFAULT '' COMMENT '最后修改人邮箱前缀',
    `DataChange_LastTime`       timestamp        NULL     DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
    PRIMARY KEY (`Id`),
    KEY `Key` (`Key`),
    KEY `IX_GroupId` (`NamespaceId`),
    KEY `DataChange_LastTime` (`DataChange_LastTime`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='配置项目';


# Dump of table release
# ------------------------------------------------------------

DROP TABLE IF EXISTS `Release`;

CREATE TABLE `Release`
(
    `Id`                        int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
    `ReleaseKey`                varchar(64)      NOT NULL DEFAULT '' COMMENT '发布的Key',
    `Name`                      varchar(64)      NOT NULL DEFAULT 'default' COMMENT '发布名字',
    `Comment`                   varchar(256)              DEFAULT NULL COMMENT '发布说明',
    `AppId`                     varchar(500)     NOT NULL DEFAULT 'default' COMMENT 'AppID',
    `ClusterName`               varchar(64)      NOT NULL DEFAULT 'default' COMMENT 'ClusterName',
    `NamespaceName`             varchar(64)      NOT NULL DEFAULT 'default' COMMENT 'namespaceName',
    `LaneName`                  varchar(64)      NOT NULL DEFAULT 'default' COMMENT '灰度名字',
    `Configurations`            longtext         NOT NULL COMMENT '发布配置',
    `IsAbandoned`               tinyint(1)       NOT NULL DEFAULT b'0' COMMENT '是否废弃',
    `IsDeleted`                 tinyint(1)       NOT NULL DEFAULT b'0' COMMENT '1: deleted, 0: normal',
    `DataChange_CreatedBy`      varchar(32)      NOT NULL DEFAULT 'default' COMMENT '创建人邮箱前缀',
    `DataChange_CreatedTime`    timestamp        NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `DataChange_LastModifiedBy` varchar(32)               DEFAULT '' COMMENT '最后修改人邮箱前缀',
    `DataChange_LastTime`       timestamp        NULL     DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
    PRIMARY KEY (`Id`),
    KEY `AppId_ClusterName_GroupName` (`AppId`(64), `ClusterName`(64), `NamespaceName`(64)),
    KEY `DataChange_LastTime` (`DataChange_LastTime`),
    KEY `IX_ReleaseKey` (`ReleaseKey`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='发布';


# Dump of table releasehistory
# ------------------------------------------------------------

DROP TABLE IF EXISTS `ReleaseHistory`;

CREATE TABLE `ReleaseHistory`
(
    `Id`                        int(11) unsigned    NOT NULL AUTO_INCREMENT COMMENT '自增Id',
    `AppId`                     varchar(64)         NOT NULL DEFAULT 'default' COMMENT 'AppID',
    `ClusterName`               varchar(64)         NOT NULL DEFAULT 'default' COMMENT 'ClusterName',
    `NamespaceName`             varchar(64)         NOT NULL DEFAULT 'default' COMMENT 'namespaceName',
    `BranchName`                varchar(32)         NOT NULL DEFAULT 'default' COMMENT '保留字段',
    `LaneName`                  varchar(64)         NOT NULL DEFAULT 'default' COMMENT '灰度名字',
    `ReleaseId`                 int(11) unsigned    NOT NULL DEFAULT '0' COMMENT '保留字段',
    `PreviousReleaseId`         int(11) unsigned    NOT NULL DEFAULT '0' COMMENT '保留字段',
    `Operation`                 tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '发布类型，0: 普通发布，1: 灰度发布，2: 灰度全量发布',
    `OperationContext`          longtext            NOT NULL COMMENT '发布上下文信息（只展示不同的）',
    `ReleaseContext`            longtext            NOT NULL COMMENT '发布上下文信息（展示全部发布内容）',
    `IsDeleted`                 tinyint(1)          NOT NULL DEFAULT b'0' COMMENT '1: deleted, 0: normal',
    `DataChange_CreatedBy`      varchar(32)         NOT NULL DEFAULT 'default' COMMENT '创建人邮箱前缀',
    `DataChange_CreatedTime`    timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `DataChange_LastModifiedBy` varchar(32)                  DEFAULT '' COMMENT '最后修改人邮箱前缀',
    `DataChange_LastTime`       timestamp           NULL     DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
    PRIMARY KEY (`Id`),
    KEY `IX_Namespace` (`AppId`, `ClusterName`, `NamespaceName`, `BranchName`),
    KEY `IX_ReleaseId` (`ReleaseId`),
    KEY `IX_DataChange_LastTime` (`DataChange_LastTime`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='发布历史';


# Dump of table releasemessage
# ------------------------------------------------------------

DROP TABLE IF EXISTS `ReleaseMessage`;

CREATE TABLE `ReleaseMessage`
(
    `Id`                  int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
    `Message`             varchar(1024)    NOT NULL DEFAULT '' COMMENT '发布的消息内容',
    `DataChange_LastTime` timestamp        NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
    PRIMARY KEY (`Id`),
    KEY `DataChange_LastTime` (`DataChange_LastTime`),
    KEY `IX_Message` (`Message`(191))
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='发布消息';



# Dump of table serverconfig
# ------------------------------------------------------------

DROP TABLE IF EXISTS `ServerConfig`;

CREATE TABLE `ServerConfig`
(
    `Id`                        int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增Id',
    `Key`                       varchar(64)      NOT NULL DEFAULT 'default' COMMENT '配置项Key',
    `Cluster`                   varchar(32)      NOT NULL DEFAULT 'default' COMMENT '配置对应的集群，default为不针对特定的集群',
    `Value`                     varchar(2048)    NOT NULL DEFAULT 'default' COMMENT '配置项值',
    `Comment`                   varchar(1024)             DEFAULT '' COMMENT '注释',
    `IsDeleted`                 tinyint(1)       NOT NULL DEFAULT b'0' COMMENT '1: deleted, 0: normal',
    `DataChange_CreatedBy`      varchar(32)      NOT NULL DEFAULT 'default' COMMENT '创建人邮箱前缀',
    `DataChange_CreatedTime`    timestamp        NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `DataChange_LastModifiedBy` varchar(32)               DEFAULT '' COMMENT '最后修改人邮箱前缀',
    `DataChange_LastTime`       timestamp        NULL     DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
    PRIMARY KEY (`Id`),
    KEY `IX_Key` (`Key`),
    KEY `DataChange_LastTime` (`DataChange_LastTime`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='配置服务自身配置';


replace into Item(`Id`, `NamespaceId`, `Key`, `Value`, `ReleaseValue`, `Status`, `Comment`, `Describe`,
                  `DataChange_CreatedBy`, `DataChange_LastModifiedBy`, `DataChange_CreatedTime`, `DataChange_LastTime`)
values ('15408', '3429', 'logger.apollo.enable', 'true', 'true', '1', '', '', 'liulonglong', 'liulonglong',
        '2021-01-15 19:06:23', '2021-07-21 17:36:07'),
       ('15409', '3429', 'logging.level.root', 'INFO', 'INFO', '1', '', '', 'liulonglong', 'liulonglong',
        '2021-01-15 19:06:23', '2021-07-21 17:36:07'),
       ('15410', '3429', 'logging.level.com.didapinche.im', 'DEBUG', 'DEBUG', '1', '', '', 'liulonglong', 'liulonglong',
        '2021-01-15 19:06:23', '2021-07-21 17:36:07'),
       ('15411', '3429', 'logging.level.com.didapinche.zeus', 'ERROR', 'ERROR', '1', '', '', 'liulonglong',
        'liulonglong', '2021-01-15 19:06:23', '2021-07-21 17:36:07'),
       ('15412', '3429', 'server.port', '0', '0', '1', '', '', 'liulonglong', 'liulonglong', '2021-01-15 19:06:23',
        '2021-07-21 17:36:07'),
       ('15413', '3429', 'endpoints.prometheus.id', 'metrics', 'metrics', '1', '', '', 'liulonglong', 'liulonglong',
        '2021-01-15 19:06:23', '2021-07-21 17:36:07'),
       ('15414', '3429', 'management.security.enabled', 'false', 'false', '1', '', '', 'liulonglong', 'liulonglong',
        '2021-01-15 19:06:23', '2021-07-21 17:36:07'),
       ('0', '3429', 'zeus.bossThread', '1', '', '0', '', '', 'lihang', 'lihang', '2021-07-21 17:36:07',
        '2021-07-21 17:36:07'),
       ('0', '3429', 'zeus.workThread', '20', '', '0', '', '', 'lihang', 'lihang', '2021-07-21 17:36:07',
        '2021-07-21 17:36:07'),
       ('0', '3429', 'zeus.timeout', '500', '', '0', '', '', 'lihang', 'lihang', '2021-07-21 17:36:07',
        '2021-07-21 17:36:07'),
       ('15418', '3429', 'rocket.enable', 'true', 'true', '1', '', '', 'liulonglong', 'liulonglong',
        '2021-01-15 19:06:23', '2021-07-21 17:36:07'),
       ('15419', '3429', 'rocket.nameServer', '192.168.1.201:9876;192.168.1.197:9876',
        '192.168.1.201:9876;192.168.1.197:9876', '1', '', '', 'liulonglong', 'liulonglong', '2021-01-15 19:06:23',
        '2021-07-21 17:36:07'),
       ('15420', '3429', 'rocket.producer.group', 'P_IM_LOGIC', 'P_IM_LOGIC', '1', '', '', 'liulonglong', 'liulonglong',
        '2021-01-15 19:06:23', '2021-07-21 17:36:07'),
       ('15421', '3429', 'rocket.producer.timeout', '3000', '3000', '1', '', '', 'liulonglong', 'liulonglong',
        '2021-01-15 19:06:23', '2021-07-21 17:36:07'),
       ('15422', '3429', 'rocket.conflict.topic', 't_im_conflict', 't_im_conflict', '1', '', '', 'liulonglong',
        'liulonglong', '2021-01-15 19:06:23', '2021-07-21 17:36:07'),
       ('15423', '3429', 'rocket.conflict.group', 'C_IM_CONFLICT', 'C_IM_CONFLICT', '1', '', '', 'liulonglong',
        'liulonglong', '2021-01-15 19:06:23', '2021-07-21 17:36:07'),
       ('15424', '3429', 'rocket.systemChat.topic', 't_im_system_chat', 't_im_system_chat', '1', '', '', 'liulonglong',
        'liulonglong', '2021-01-15 19:06:23', '2021-07-21 17:36:07'),
       ('15425', '3429', 'rocket.systemChat.group', 'C_IM_SYSTEM_CHAT', 'C_IM_SYSTEM_CHAT', '1', '', '', 'liulonglong',
        'liulonglong', '2021-01-15 19:06:23', '2021-07-21 17:36:07'),
       ('15426', '3429', 'rocket.userChat.topic', 't_im_chat', 't_im_chat', '1', '', '', 'liulonglong', 'liulonglong',
        '2021-01-15 19:06:23', '2021-07-21 17:36:07'),
       ('15427', '3429', 'rocket.userChat.group', 'C_IM_PUSH_CHAT_ASYNC', 'C_IM_PUSH_CHAT_ASYNC', '1', '', '',
        'liulonglong', 'liulonglong', '2021-01-15 19:06:23', '2021-07-21 17:36:07'),
       ('15428', '3429', 'rocket.userState.topic', 't_im_user_state', 't_im_user_state', '1', '', '', 'liulonglong',
        'liulonglong', '2021-01-15 19:06:23', '2021-07-21 17:36:07'),
       ('15429', '3429', 'rocket.userState.group', 'C_IM_USER_STATE', 'C_IM_USER_STATE', '1', '', '', 'liulonglong',
        'liulonglong', '2021-01-15 19:06:23', '2021-07-21 17:36:07'),
       ('15430', '3429', 'rocket.biddingTaxi.group', 'C_IM_TAXI_BIDDING', 'C_IM_TAXI_BIDDING', '1', '', '',
        'liulonglong', 'liulonglong', '2021-01-15 19:06:23', '2021-07-21 17:36:07'),
       ('15431', '3429', 'rocket.biddingTaxi.topic', 't_taxi_bidding_message', 't_taxi_bidding_message', '1', '', '',
        'liulonglong', 'liulonglong', '2021-01-15 19:06:23', '2021-07-21 17:36:07'),
       ('15432', '3429', 'rocket.biddingMeter.group', 'C_IM_METER_BIDDING', 'C_IM_METER_BIDDING', '1', '', '',
        'liulonglong', 'liulonglong', '2021-01-15 19:06:23', '2021-07-21 17:36:07'),
       ('15433', '3429', 'rocket.biddingMeter.topic', 't_taxi_meter_bidding_message', 't_taxi_meter_bidding_message',
        '1', '', '', 'liulonglong', 'liulonglong', '2021-01-15 19:06:23', '2021-07-21 17:36:07'),
       ('15434', '3429', 'rocket.bidding.group', 'C_IM_BIDDIING', 'C_IM_BIDDIING', '1', '', '', 'liulonglong',
        'liulonglong', '2021-01-15 19:06:23', '2021-07-21 17:36:07'),
       ('15435', '3429', 'rocket.bidding.topic', 't_bidding_message', 't_bidding_message', '1', '', '', 'liulonglong',
        'liulonglong', '2021-01-15 19:06:23', '2021-07-21 17:36:07'),
       ('15436', '3429', 'redis.chat.syncKey.appId', '10008', '10008', '1', '', '', 'liulonglong', 'liulonglong',
        '2021-01-15 19:06:23', '2021-07-21 17:36:07'),
       ('15437', '3429', 'redis.chat.syncKey.url',
        'http://192.168.1.201:8585/cache/client/redis/cluster/10008.jsonclientVersion=1.0-SNAPSHOT',
        'http://192.168.1.201:8585/cache/client/redis/cluster/10008.jsonclientVersion=1.0-SNAPSHOT', '1', '', '',
        'liulonglong', 'liulonglong', '2021-01-15 19:06:23', '2021-07-21 17:36:07'),
       ('15438', '3429', 'redis.chat.syncKey.connTimeout', '1000', '1000', '1', '', '', 'liulonglong', 'liulonglong',
        '2021-01-15 19:06:23', '2021-07-21 17:36:07'),
       ('15439', '3429', 'redis.chat.syncKey.soTimeout', '100', '100', '1', '', '', 'liulonglong', 'liulonglong',
        '2021-01-15 19:06:23', '2021-07-21 17:36:07'),
       ('15440', '3429', 'redis.chat.syncKey.maxRedirections', '6', '6', '1', '', '', 'liulonglong', 'liulonglong',
        '2021-01-15 19:06:23', '2021-07-21 17:36:07'),
       ('15441', '3429', 'redis.chat.syncKey.pool.maxTotal', '200', '200', '1', '', '', 'liulonglong', 'liulonglong',
        '2021-01-15 19:06:23', '2021-07-21 17:36:07'),
       ('15442', '3429', 'redis.chat.syncKey.pool.maxIdle', '20', '20', '1', '', '', 'liulonglong', 'liulonglong',
        '2021-01-15 19:06:23', '2021-07-21 17:36:07'),
       ('15443', '3429', 'redis.chat.message.appId', '10008', '10008', '1', '', '', 'liulonglong', 'liulonglong',
        '2021-01-15 19:06:23', '2021-07-21 17:36:07'),
       ('15444', '3429', 'redis.chat.message.url',
        'http://192.168.1.201:8585/cache/client/redis/cluster/10008.jsonclientVersion=1.0-SNAPSHOT',
        'http://192.168.1.201:8585/cache/client/redis/cluster/10008.jsonclientVersion=1.0-SNAPSHOT', '1', '', '',
        'liulonglong', 'liulonglong', '2021-01-15 19:06:23', '2021-07-21 17:36:07'),
       ('15445', '3429', 'redis.chat.message.connTimeout', '1000', '1000', '1', '', '', 'liulonglong', 'liulonglong',
        '2021-01-15 19:06:23', '2021-07-21 17:36:07'),
       ('15446', '3429', 'redis.chat.message.soTimeout', '100', '100', '1', '', '', 'liulonglong', 'liulonglong',
        '2021-01-15 19:06:23', '2021-07-21 17:36:07'),
       ('15447', '3429', 'redis.chat.message.maxRedirections', '6', '6', '1', '', '', 'liulonglong', 'liulonglong',
        '2021-01-15 19:06:23', '2021-07-21 17:36:07'),
       ('15448', '3429', 'redis.chat.message.pool.maxTotal', '200', '200', '1', '', '', 'liulonglong', 'liulonglong',
        '2021-01-15 19:06:23', '2021-07-21 17:36:07'),
       ('15449', '3429', 'redis.chat.message.pool.maxIdle', '20', '20', '1', '', '', 'liulonglong', 'liulonglong',
        '2021-01-15 19:06:23', '2021-07-21 17:36:07'),
       ('15450', '3429', 'redis.ctob.appId', '10008', '10008', '1', '', '', 'liulonglong', 'liulonglong',
        '2021-01-15 19:06:23', '2021-07-21 17:36:07'),
       ('15451', '3429', 'redis.ctob.url',
        'http://192.168.1.201:8585/cache/client/redis/cluster/10008.jsonclientVersion=1.0-SNAPSHOT',
        'http://192.168.1.201:8585/cache/client/redis/cluster/10008.jsonclientVersion=1.0-SNAPSHOT', '1', '', '',
        'liulonglong', 'liulonglong', '2021-01-15 19:06:23', '2021-07-21 17:36:07'),
       ('15452', '3429', 'redis.ctob.connTimeout', '3000', '3000', '1', '', '', 'liulonglong', 'liulonglong',
        '2021-01-15 19:06:23', '2021-07-21 17:36:07'),
       ('15453', '3429', 'redis.ctob.soTimeout', '100', '100', '1', '', '', 'liulonglong', 'liulonglong',
        '2021-01-15 19:06:23', '2021-07-21 17:36:07'),
       ('15454', '3429', 'redis.ctob.maxRedirections', '6', '6', '1', '', '', 'liulonglong', 'liulonglong',
        '2021-01-15 19:06:23', '2021-07-21 17:36:07'),
       ('15455', '3429', 'redis.ctob.pool.maxTotal', '100', '100', '1', '', '', 'liulonglong', 'liulonglong',
        '2021-01-15 19:06:23', '2021-07-21 17:36:07'),
       ('15456', '3429', 'redis.ctob.pool.maxIdle', '10', '10', '1', '', '', 'liulonglong', 'liulonglong',
        '2021-01-15 19:06:23', '2021-07-21 17:36:07'),
       ('15457', '3429', 'beforeRide.relationship.expSeconds', '300', '300', '1', '', '', 'liulonglong', 'liulonglong',
        '2021-01-15 19:06:23', '2021-07-21 17:36:07'),
       ('15458', '3429', 'beforeRide.chatSet.sizeLimit', '20', '20', '1', '', '', 'liulonglong', 'liulonglong',
        '2021-01-15 19:06:23', '2021-07-21 17:36:07'),
       ('15459', '3429', 'beforeRide.unAnswer.countLimit', '3', '3', '1', '', '', 'liulonglong', 'liulonglong',
        '2021-01-15 19:06:23', '2021-07-21 17:36:07'),
       ('15460', '3429', 'taxi.version', '2.3.0', '2.3.0', '1', '', '', 'liulonglong', 'liulonglong',
        '2021-01-15 19:06:23', '2021-07-21 17:36:07'),
       ('15461', '3429', 'voice.interval', '60', '60', '1', '', '', 'liulonglong', 'liulonglong', '2021-01-15 19:06:23',
        '2021-07-21 17:36:07'),
       ('18153', '3429', 'rocket.lastEta.group', 'C_IM_LAST_ETA_ASYNC', 'C_IM_LAST_ETA_ASYNC', '1', '', '',
        'liulonglong', 'liulonglong', '2021-03-08 17:38:32', '2021-07-21 17:36:07'),
       ('18154', '3429', 'rocket.lastEta.topic', 't_im_last_eta', 't_im_last_eta', '1', '', '', 'liulonglong',
        'liulonglong', '2021-03-08 17:38:32', '2021-07-21 17:36:07'),
       ('23706', '3429', 'beforeRide.textTips.fisrtSend', '对方回复前，你只能发送3条消息', '对方回复前，你只能发送3条消息', '1', '', '',
        'liulonglong', 'liulonglong',
        '2021-06-22 11:),('23763','3429','beforeRide.textTips.twoSide','{"content":"<p>沟通中请注意：为提供更愉悦更安全的出行体验，如遇对方言语恶劣、要求线下交易等违规情况时，请点击<a href{"content":"<p>沟通中请注意：为提供更愉悦更安全的出行体验，如遇对方言语恶劣、要求线下交易等违规情况时，请点击<a href='didapinche://IMComplaint'>立即举报</a></p>"}',':07:29','2021-07-21 17:36:07'),('27381','3429','juno.application-name','im-async','im-async','1','','','liulonglong','liulonglong','2021-07-21 13:18:50','2021-07-21 17:36:07'),('27382','3429','zeus.enable','true','true','1','','','liulonglong','liulonglong','2021-07-21 13:19:06','2021-07-21 17:36:07'),('27383','3429','zeus.application.name','im-async','im-async','1','','','liulonglong','liulonglong','2021-07-21 13:19:26','2021-07-21 17:36:07'),('27384','3429','zeus.client.timeout-ms','500','500','1','','','liulonglong','liulonglong','2021-07-21 13:19:57','2021-07-21 17:36:07'),('27385','3429','zeus.client.load-balance','cidLoadBalance','cidLoadBalance','1','','','liulonglong','liulonglong','2021-07-21 13:20:15','2021-07-21 17:36:07'),('27386','3429','zeus.clients.ZeusImCometService.disable-lane-router','true','true','1','','','liulonglong','liulonglong','2021-07-21 13:20:34','2021-07-21 17:36:07'),('27387','3429','zeus.clients.ZeusImCometService.disable-priority-data-center-router','true','true','1','','','liulonglong','liulonglong','2021-07-21 13:20:47','2021-07-21 17:36:07'),('27388','3429','zeus.clients.ZeusImCometService.load-balance','cometLoadBalance','cometLoadBalance','1','','','liulonglong','liulonglong','2021-07-21 13:20:59','2021-07-21 17:36:07');
