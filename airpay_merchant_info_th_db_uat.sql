# ************************************************************
# Sequel Pro SQL dump
# Version 4541
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: db-master-airpay-uatsg-sg1-uat.airpaymobile.com (MySQL 5.7.21-20-log)
# Database: airpay_merchant_info_th_db
# Generation Time: 2020-11-10 10:51:17 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table address_info_tab
# ------------------------------------------------------------

DROP TABLE IF EXISTS `address_info_tab`;

CREATE TABLE `address_info_tab` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `relation_code` bigint(20) unsigned NOT NULL,
  `address_level` tinyint(4) unsigned NOT NULL,
  `name_en` varchar(256) NOT NULL DEFAULT '',
  `name_local` varchar(256) NOT NULL DEFAULT '',
  `postal_code` varchar(32) NOT NULL DEFAULT '',
  `create_time` int(10) unsigned NOT NULL DEFAULT '0',
  `update_time` int(10) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_relation_unique` (`relation_code`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table app_function_tab
# ------------------------------------------------------------

DROP TABLE IF EXISTS `app_function_tab`;

CREATE TABLE `app_function_tab` (
  `app_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `merchant_id` bigint(20) unsigned NOT NULL,
  `reference_type` tinyint(4) unsigned NOT NULL COMMENT 'reference type eg:top-merchant/merchant/outlet',
  `m_name` varchar(256) NOT NULL DEFAULT '',
  `m_description` varchar(1024) NOT NULL DEFAULT '',
  `logo` varchar(1024) NOT NULL DEFAULT '',
  `app_key` varchar(256) NOT NULL,
  `function_set` bigint(20) unsigned NOT NULL COMMENT 'function set with bit flag',
  `m_status` tinyint(4) unsigned NOT NULL,
  `get_detail_url` varchar(1024) NOT NULL DEFAULT '',
  `validate_url` varchar(1024) NOT NULL DEFAULT '',
  `notify_url` varchar(1024) NOT NULL DEFAULT '',
  `oauth_redirect_url` varchar(1024) NOT NULL DEFAULT '',
  `ip_white_list` varchar(1024) NOT NULL DEFAULT '',
  `mapping_mode` tinyint(4) unsigned NOT NULL DEFAULT '0',
  `create_time` int(10) unsigned NOT NULL,
  `update_time` int(10) unsigned NOT NULL,
  PRIMARY KEY (`app_id`),
  KEY `idx_merchant_app` (`app_id`,`merchant_id`,`reference_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table category_info_tab
# ------------------------------------------------------------

DROP TABLE IF EXISTS `category_info_tab`;

CREATE TABLE `category_info_tab` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `category` int(10) unsigned NOT NULL COMMENT 'category is unique in each region and source',
  `merchant_source` int(10) unsigned NOT NULL COMMENT 'source that merchant create from',
  `m_description` varchar(1024) NOT NULL DEFAULT '',
  `create_time` int(10) unsigned NOT NULL,
  `update_time` int(10) unsigned NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_category_source` (`category`,`merchant_source`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table collection_tab
# ------------------------------------------------------------

DROP TABLE IF EXISTS `collection_tab`;

CREATE TABLE `collection_tab` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `m_name` varchar(128) NOT NULL,
  `m_description` varchar(1024) NOT NULL DEFAULT '',
  `merchant_id_list` text,
  `create_time` int(10) unsigned NOT NULL,
  `update_time` int(10) unsigned NOT NULL,
  `m_status` tinyint(4) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table config_info_tab
# ------------------------------------------------------------

DROP TABLE IF EXISTS `config_info_tab`;

CREATE TABLE `config_info_tab` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `reference_id` bigint(20) unsigned NOT NULL COMMENT 'reference id this record belongs to',
  `reference_type` tinyint(4) unsigned NOT NULL COMMENT 'reference type eg:top-merchant/merchant/outlet',
  `config_type` tinyint(4) unsigned NOT NULL COMMENT 'config type',
  `config_value` text,
  `create_time` int(10) unsigned NOT NULL,
  `update_time` int(10) unsigned NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_reference_config` (`reference_id`,`reference_type`,`config_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table external_mapping_tab
# ------------------------------------------------------------

DROP TABLE IF EXISTS `external_mapping_tab`;

CREATE TABLE `external_mapping_tab` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `top_merchant_id` bigint(20) unsigned NOT NULL COMMENT 'top level merchant id, eg.ISV',
  `ex_merchant_id` varchar(128) NOT NULL DEFAULT '' COMMENT 'external merchant id',
  `ex_outlet_id` varchar(128) NOT NULL DEFAULT '' COMMENT 'external outlet id',
  `merchant_id` bigint(20) unsigned NOT NULL DEFAULT '0',
  `outlet_id` bigint(20) unsigned NOT NULL DEFAULT '0',
  `deleted_at` int(10) unsigned NOT NULL DEFAULT '0' COMMENT 'record detete time, 0 if valid',
  `create_time` int(10) unsigned NOT NULL DEFAULT '0',
  `update_time` int(10) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_source_unique` (`top_merchant_id`,`ex_merchant_id`,`ex_outlet_id`,`merchant_id`,`outlet_id`,`deleted_at`) USING BTREE,
  KEY `idx_merchant` (`merchant_id`),
  KEY `idx_outlet` (`outlet_id`),
  KEY `idx_ex_merchant` (`top_merchant_id`,`ex_merchant_id`),
  KEY `idx_ex_outlet` (`top_merchant_id`,`ex_outlet_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table finance_info_tab
# ------------------------------------------------------------

DROP TABLE IF EXISTS `finance_info_tab`;

CREATE TABLE `finance_info_tab` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `reference_id` bigint(20) unsigned NOT NULL COMMENT 'reference id this record belongs to',
  `reference_type` tinyint(4) unsigned NOT NULL COMMENT 'reference type eg:top-merchant/merchant/outlet',
  `settlement_method` tinyint(4) unsigned NOT NULL COMMENT 'instant settlement/T+1 settlement/manual...',
  `payment_method` tinyint(4) unsigned NOT NULL COMMENT 'bank account/wallet...',
  `aggregate_settlement` tinyint(4) unsigned NOT NULL COMMENT 'settle for merchant/outlet/top-merchant',
  `merchant_wallet` bigint(20) unsigned NOT NULL,
  `bank_channel_id` bigint(20) unsigned NOT NULL DEFAULT '0',
  `bank_name` varchar(64) NOT NULL DEFAULT '',
  `bank_branch` varchar(64) NOT NULL DEFAULT '',
  `bank_account` varchar(1024) NOT NULL DEFAULT '',
  `bank_account_name` varchar(1024) NOT NULL DEFAULT '',
  `vat` tinyint(4) unsigned NOT NULL COMMENT 'Value Added Tax',
  `vat_id` varchar(64) NOT NULL DEFAULT '' COMMENT 'Value Added Tax Id',
  `wht` tinyint(4) unsigned NOT NULL COMMENT 'With Holding Tax',
  `tax_id` varchar(64) NOT NULL DEFAULT '' COMMENT 'Tax Id',
  `create_time` int(10) unsigned NOT NULL,
  `update_time` int(10) unsigned NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_reference` (`reference_id`,`reference_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table legal_info_tab
# ------------------------------------------------------------

DROP TABLE IF EXISTS `legal_info_tab`;

CREATE TABLE `legal_info_tab` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `reference_id` bigint(20) unsigned NOT NULL COMMENT 'reference id this record belongs to',
  `reference_type` tinyint(4) unsigned NOT NULL COMMENT 'reference type eg:top-merchant/merchant/outlet',
  `certificate_name` varchar(128) NOT NULL,
  `certificate_no` varchar(64) NOT NULL COMMENT 'business certificate no',
  `certificate_photo` varchar(256) NOT NULL DEFAULT '',
  `signed_contract` varchar(128) NOT NULL,
  `id_no` varchar(1024) DEFAULT NULL,
  `bank_book` varchar(128) DEFAULT NULL,
  `pp20` varchar(128) DEFAULT NULL,
  `incorporation_date` varchar(10) DEFAULT NULL COMMENT 'company registe date',
  `shareholders` varchar(2048) DEFAULT NULL COMMENT 'company shareholder list',
  `create_time` int(10) unsigned NOT NULL,
  `update_time` int(10) unsigned NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_reference` (`reference_id`,`reference_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table main_person_tab
# ------------------------------------------------------------

DROP TABLE IF EXISTS `main_person_tab`;

CREATE TABLE `main_person_tab` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `reference_id` bigint(20) unsigned NOT NULL COMMENT 'reference id this record belongs to',
  `reference_type` tinyint(4) unsigned NOT NULL COMMENT 'reference type eg:top-merchant/merchant/outlet',
  `title` tinyint(4) unsigned NOT NULL COMMENT 'person title, eg:MR/MRS/MS',
  `first_name` varchar(128) NOT NULL DEFAULT '',
  `last_name` varchar(128) NOT NULL DEFAULT '',
  `first_name_local` varchar(256) NOT NULL DEFAULT '',
  `last_name_local` varchar(256) NOT NULL DEFAULT '',
  `birthday` varchar(10) NOT NULL DEFAULT '',
  `laser_field` varchar(12) NOT NULL DEFAULT '',
  `id_card` varchar(1024) NOT NULL DEFAULT '',
  `registered_address` varchar(2048) NOT NULL,
  `residential_address` varchar(2048) NOT NULL,
  `contact` varchar(16) NOT NULL,
  `occupation` int(10) NOT NULL,
  `company` varchar(128) NOT NULL,
  `create_time` int(10) unsigned NOT NULL,
  `update_time` int(10) unsigned NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_reference` (`reference_id`,`reference_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table merchant_info_tab
# ------------------------------------------------------------

DROP TABLE IF EXISTS `merchant_info_tab`;

CREATE TABLE `merchant_info_tab` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `m_name` varchar(256) NOT NULL,
  `m_status` tinyint(4) unsigned NOT NULL,
  `logo` varchar(1024) NOT NULL,
  `picture_list` varchar(1024) NOT NULL,
  `m_type` int(8) unsigned NOT NULL,
  `airpay_uid` bigint(20) unsigned NOT NULL DEFAULT '0',
  `contact` varchar(16) NOT NULL,
  `email_list` varchar(1024) NOT NULL DEFAULT '',
  `state` varchar(128) NOT NULL DEFAULT '',
  `city` varchar(128) NOT NULL DEFAULT '',
  `district` varchar(128) NOT NULL DEFAULT '',
  `sub_district` varchar(128) NOT NULL DEFAULT '',
  `address` varchar(1024) NOT NULL DEFAULT '',
  `postal_code` varchar(32) NOT NULL DEFAULT '',
  `province_id` bigint(20) unsigned NOT NULL DEFAULT '0',
  `district_id` bigint(20) unsigned NOT NULL DEFAULT '0',
  `subdistrict_id` bigint(20) unsigned NOT NULL DEFAULT '0',
  `m_description` varchar(1024) NOT NULL DEFAULT '',
  `mcc` int(8) unsigned NOT NULL COMMENT 'MCC',
  `business_category` int(10) unsigned NOT NULL DEFAULT '0' COMMENT 'a business category lower than MCC',
  `tc` varchar(1024) NOT NULL DEFAULT '' COMMENT 'T&C,TermsAndConditions',
  `group_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT 'fill in top-merchant id, if top-merchant is a group',
  `brand` varchar(32) NOT NULL DEFAULT '',
  `merchant_source` int(10) unsigned NOT NULL COMMENT 'source that merchant create from',
  `main_person_info_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT 'responsible person info id',
  `legal_info_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT 'merchant legal info id',
  `finance_info_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT 'merchant finance info id',
  `category` int(10) unsigned NOT NULL DEFAULT '0' COMMENT 'category',
  `create_time` int(10) unsigned NOT NULL,
  `update_time` int(10) unsigned NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_name` (`m_name`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table merchant_mapping_tab
# ------------------------------------------------------------

DROP TABLE IF EXISTS `merchant_mapping_tab`;

CREATE TABLE `merchant_mapping_tab` (
  `source_id` bigint(20) unsigned NOT NULL,
  `source_merchant_id` varchar(128) NOT NULL COMMENT 'external merchant id',
  `merchant_id` bigint(20) unsigned NOT NULL,
  `create_time` int(10) unsigned NOT NULL,
  `update_time` int(10) unsigned NOT NULL,
  UNIQUE KEY `idx_source` (`source_id`,`source_merchant_id`),
  KEY `idx_merchant` (`merchant_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table merchant_universal_config_tab
# ------------------------------------------------------------

DROP TABLE IF EXISTS `merchant_universal_config_tab`;

CREATE TABLE `merchant_universal_config_tab` (
  `merchant_id` bigint(20) unsigned NOT NULL,
  `key` varchar(64) NOT NULL,
  `value` text NOT NULL,
  `remark` varchar(64) NOT NULL DEFAULT '',
  `create_time` int(10) unsigned NOT NULL,
  `update_time` int(10) unsigned NOT NULL,
  PRIMARY KEY (`merchant_id`,`key`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table outlet_info_tab
# ------------------------------------------------------------

DROP TABLE IF EXISTS `outlet_info_tab`;

CREATE TABLE `outlet_info_tab` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `m_name` varchar(256) NOT NULL,
  `m_status` tinyint(4) unsigned NOT NULL,
  `merchant_id` bigint(20) unsigned NOT NULL,
  `logo` varchar(1024) NOT NULL,
  `picture_list` varchar(1024) NOT NULL,
  `airpay_uid` bigint(20) unsigned NOT NULL DEFAULT '0',
  `contact` varchar(16) NOT NULL,
  `state` varchar(128) NOT NULL DEFAULT '',
  `sub_district` varchar(128) NOT NULL DEFAULT '',
  `city` varchar(128) NOT NULL DEFAULT '',
  `district` varchar(128) NOT NULL DEFAULT '',
  `address` varchar(1024) NOT NULL DEFAULT '',
  `province_id` bigint(20) unsigned NOT NULL DEFAULT '0',
  `district_id` bigint(20) unsigned NOT NULL DEFAULT '0',
  `subdistrict_id` bigint(20) unsigned NOT NULL DEFAULT '0',
  `longitude` bigint(20) NOT NULL,
  `latitude` bigint(20) NOT NULL,
  `postal_code` varchar(32) NOT NULL DEFAULT '',
  `outlet_flags_set` bigint(20) NOT NULL DEFAULT '0',
  `remark` varchar(1024) NOT NULL DEFAULT '',
  `group_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT 'fill in top-merchant id, if top-merchant is a group',
  `main_person_info_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT 'responsible person info id',
  `legal_info_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT 'merchant legal info id',
  `finance_info_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT 'merchant finance info id',
  `config_info_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT 'merchant config info id',
  `qr_code_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT 'QR code id',
  `category` int(10) unsigned NOT NULL DEFAULT '0' COMMENT 'category',
  `create_time` int(10) unsigned NOT NULL,
  `update_time` int(10) unsigned NOT NULL,
  `block_function_set` bigint(20) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `idx_name` (`m_name`) USING BTREE,
  KEY `idx_merchant_id` (`merchant_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table outlet_mapping_tab
# ------------------------------------------------------------

DROP TABLE IF EXISTS `outlet_mapping_tab`;

CREATE TABLE `outlet_mapping_tab` (
  `source_id` bigint(20) unsigned NOT NULL,
  `source_outlet_id` varchar(128) NOT NULL COMMENT 'external outlet id',
  `merchant_id` bigint(20) unsigned NOT NULL,
  `outlet_id` bigint(20) unsigned NOT NULL,
  `create_time` int(10) unsigned NOT NULL,
  `update_time` int(10) unsigned NOT NULL,
  UNIQUE KEY `idx_source` (`source_id`,`source_outlet_id`),
  KEY `idx_merchant` (`merchant_id`,`outlet_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table portal_account_tab
# ------------------------------------------------------------

DROP TABLE IF EXISTS `portal_account_tab`;

CREATE TABLE `portal_account_tab` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `source_type` int(11) unsigned NOT NULL DEFAULT '0',
  `source_uid` bigint(20) unsigned DEFAULT NULL,
  `merchant_id` bigint(20) unsigned NOT NULL DEFAULT '0',
  `staff_id` bigint(20) unsigned NOT NULL DEFAULT '0',
  `login_mobile_no` varchar(16) DEFAULT '',
  `otp_mobile_no` varchar(16) DEFAULT NULL,
  `account_name` varchar(128) DEFAULT NULL,
  `secret_code` varchar(128) DEFAULT NULL,
  `last_login_time` int(10) NOT NULL,
  `state` tinyint(4) unsigned NOT NULL DEFAULT '0' COMMENT '1:active, 2:inactive, 3:invalid',
  `language` varchar(4) NOT NULL DEFAULT '',
  `deleted_at` int(10) unsigned NOT NULL DEFAULT '0',
  `create_time` int(10) unsigned NOT NULL DEFAULT '0',
  `update_time` int(10) unsigned NOT NULL DEFAULT '0',
  `push_notify` tinyint(4) unsigned NOT NULL DEFAULT '1',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_name` (`account_name`,`deleted_at`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table product_contract_tab
# ------------------------------------------------------------

DROP TABLE IF EXISTS `product_contract_tab`;

CREATE TABLE `product_contract_tab` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `reference_id` bigint(20) unsigned NOT NULL COMMENT 'reference id this record belongs to',
  `reference_type` tinyint(4) unsigned NOT NULL COMMENT 'reference type eg:top-merchant/merchant',
  `product_id` bigint(20) unsigned NOT NULL COMMENT 'payment product',
  `m_status` tinyint(4) unsigned NOT NULL,
  `create_time` int(10) unsigned NOT NULL,
  `update_time` int(10) unsigned NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_reference_product` (`reference_id`,`reference_type`,`product_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table request_info_tab
# ------------------------------------------------------------

DROP TABLE IF EXISTS `request_info_tab`;

CREATE TABLE `request_info_tab` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `request_id` varchar(256) NOT NULL COMMENT 'request ID is unique in merchant_source system',
  `merchant_source` int(10) unsigned NOT NULL COMMENT 'source that merchant create from',
  `reference_type` tinyint(4) unsigned NOT NULL COMMENT 'reference type eg:top-merchant/merchant/outlet',
  `assigned_value` text COMMENT 'value save',
  `create_time` int(10) unsigned NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_request_source` (`request_id`,`merchant_source`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table settlement_config_tab
# ------------------------------------------------------------

DROP TABLE IF EXISTS `settlement_config_tab`;

CREATE TABLE `settlement_config_tab` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `settle_entity_type` tinyint(4) unsigned NOT NULL COMMENT 'settle_entity_type, 0:none 1:top merchant 2:merchant',
  `settle_entity_id` bigint(20) unsigned NOT NULL,
  `product_type` tinyint(4) unsigned NOT NULL COMMENT 'product type, 0:unknown 1:deal 2:csb 3:bsc',
  `settlement_mode` tinyint(4) unsigned NOT NULL COMMENT 'settlement_mode, 0:none 1:top merchant 2:merchant',
  `settlement_to` tinyint(4) unsigned NOT NULL COMMENT 'settlement_to, 0:none 1:top merchant 2:merchant',
  `mdr_config` text NOT NULL COMMENT 'json string, mdr config',
  `settlement_method` tinyint(4) unsigned NOT NULL COMMENT 'settlement_method, 0:none 1:t0 2:t1 3:weekly 4:monthly',
  `payment_method` tinyint(4) unsigned NOT NULL COMMENT 'payment_method, 0:none 1:top merchant 2:merchant',
  `payment_way` tinyint(4) unsigned NOT NULL COMMENT 'payment_way, 0:none 1:top merchant 2:merchant',
  `merchant_account` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT 'finance info id|fk|airpay_merchant_info_th_db.finance_info_tab.id|1:1',
  `create_time` int(10) unsigned NOT NULL,
  `update_time` int(10) unsigned NOT NULL,
  `extra_data` text NOT NULL COMMENT 'json string',
  `config_flag` tinyint(4) NOT NULL DEFAULT '1',
  `email_choice` tinyint(4) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `idx_type_id_product` (`settle_entity_type`,`settle_entity_id`,`product_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='settlement config table';



# Dump of table staff_apa_config_tab
# ------------------------------------------------------------

DROP TABLE IF EXISTS `staff_apa_config_tab`;

CREATE TABLE `staff_apa_config_tab` (
  `staff_id` bigint(20) NOT NULL,
  `outlet_id` int(10) unsigned NOT NULL DEFAULT '0',
  `recv_voice_notify` tinyint(4) unsigned NOT NULL DEFAULT '0' COMMENT '0:false, 1:true',
  `recv_notify` tinyint(4) unsigned NOT NULL DEFAULT '0' COMMENT '0:false, 1:true',
  `view_txns` tinyint(4) unsigned NOT NULL DEFAULT '1' COMMENT '0:false, 1:true',
  `support_sms` tinyint(4) unsigned NOT NULL DEFAULT '0' COMMENT '0:false, 1:true',
  UNIQUE KEY `idx_staff_outlet` (`staff_id`,`outlet_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table staff_apc_config_tab
# ------------------------------------------------------------

DROP TABLE IF EXISTS `staff_apc_config_tab`;

CREATE TABLE `staff_apc_config_tab` (
  `staff_id` bigint(20) unsigned NOT NULL,
  `outlet_id` bigint(20) unsigned NOT NULL DEFAULT '0',
  `work_shift_start` varchar(16) NOT NULL DEFAULT '',
  `work_shift_end` varchar(16) NOT NULL DEFAULT '',
  UNIQUE KEY `idx_outlet_staff` (`staff_id`,`outlet_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table staff_info_tab
# ------------------------------------------------------------

DROP TABLE IF EXISTS `staff_info_tab`;

CREATE TABLE `staff_info_tab` (
  `staff_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `top_merchant_id` bigint(20) unsigned NOT NULL DEFAULT '0',
  `merchant_id` bigint(20) unsigned NOT NULL DEFAULT '0',
  `source_type` int(11) unsigned DEFAULT '0' COMMENT '9:APA_ADMIN, 70000:APA, 60000:APC_ADMIN, 60010:APC_APP',
  `source_uid` bigint(20) DEFAULT NULL,
  `source_mobile_no` varchar(32) DEFAULT '',
  `account_id` bigint(20) unsigned DEFAULT '0',
  `default_outlet_id` bigint(20) unsigned DEFAULT NULL,
  `state` tinyint(4) unsigned NOT NULL DEFAULT '0' COMMENT '1:active 2:inactive',
  `real_name` varchar(256) DEFAULT NULL,
  `nick_name` varchar(64) DEFAULT NULL,
  `mobile_no` varchar(16) DEFAULT NULL,
  `title` varchar(32) DEFAULT NULL,
  `gender` tinyint(4) unsigned DEFAULT '0',
  `email` varchar(128) DEFAULT NULL,
  `profile_picture` varchar(255) DEFAULT NULL,
  `work_remark` varchar(512) DEFAULT NULL,
  `deleted_at` int(10) unsigned NOT NULL DEFAULT '0',
  `create_time` int(10) unsigned NOT NULL DEFAULT '0',
  `update_time` int(10) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`staff_id`),
  KEY `idx_staff` (`staff_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table staff_login_device_tab
# ------------------------------------------------------------

DROP TABLE IF EXISTS `staff_login_device_tab`;

CREATE TABLE `staff_login_device_tab` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `staff_id` bigint(20) unsigned NOT NULL,
  `app_version_name` varchar(64) NOT NULL DEFAULT '',
  `app_version` int(10) unsigned NOT NULL DEFAULT '0',
  `device_id` varchar(128) NOT NULL DEFAULT '',
  `device_type` tinyint(3) NOT NULL DEFAULT '0',
  `device_os_version` varchar(32) NOT NULL DEFAULT '',
  `offline_notify_token` varbinary(512) NOT NULL DEFAULT '',
  `offline_notify_type` tinyint(3) unsigned NOT NULL DEFAULT '0',
  `login_state` tinyint(4) unsigned NOT NULL DEFAULT '0',
  `last_login_time` int(10) unsigned NOT NULL DEFAULT '0',
  `create_time` int(10) unsigned NOT NULL DEFAULT '0',
  `update_time` int(10) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_staff_device` (`staff_id`,`device_id`,`device_type`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table staff_relation_tab
# ------------------------------------------------------------

DROP TABLE IF EXISTS `staff_relation_tab`;

CREATE TABLE `staff_relation_tab` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `staff_id` bigint(20) unsigned NOT NULL DEFAULT '0',
  `top_merchant_id` bigint(20) unsigned NOT NULL DEFAULT '0',
  `merchant_id` bigint(20) unsigned NOT NULL DEFAULT '0',
  `outlet_id` bigint(20) unsigned NOT NULL DEFAULT '0',
  `role` tinyint(4) unsigned NOT NULL DEFAULT '0' COMMENT '1:shop_owner, 2:shop_staff, 3:merchant_owner, 4:merchant_staff',
  `state` tinyint(4) unsigned NOT NULL DEFAULT '0' COMMENT '1:active, 2:inactive, 3:invalid',
  `create_time` int(10) unsigned NOT NULL DEFAULT '0',
  `update_time` int(10) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table sync_info_tab
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sync_info_tab`;

CREATE TABLE `sync_info_tab` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `trace_id` varchar(255) DEFAULT NULL,
  `reference_id` bigint(20) unsigned NOT NULL,
  `reference_type` tinyint(4) unsigned NOT NULL,
  `sync_type` tinyint(4) unsigned NOT NULL,
  `sync_status` tinyint(4) unsigned NOT NULL,
  `result_code` int(10) unsigned NOT NULL,
  `result_msg` varchar(255) NOT NULL DEFAULT '',
  `create_time` int(10) unsigned NOT NULL DEFAULT '0',
  `update_time` int(10) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `idx_reference` (`reference_id`,`reference_type`,`trace_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table tag_info_tab
# ------------------------------------------------------------

DROP TABLE IF EXISTS `tag_info_tab`;

CREATE TABLE `tag_info_tab` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `type_id` bigint(20) unsigned NOT NULL DEFAULT '0',
  `tag_name` varchar(256) NOT NULL,
  `m_status` tinyint(4) unsigned NOT NULL DEFAULT '0',
  `m_description` varchar(1024) NOT NULL DEFAULT '',
  `create_way` tinyint(4) unsigned NOT NULL DEFAULT '0',
  `delete_time` int(10) unsigned NOT NULL,
  `create_time` int(10) unsigned NOT NULL,
  `update_time` int(10) unsigned NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_type_tag` (`type_id`,`tag_name`,`delete_time`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table tag_platform_tab
# ------------------------------------------------------------

DROP TABLE IF EXISTS `tag_platform_tab`;

CREATE TABLE `tag_platform_tab` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `platform` int(10) unsigned NOT NULL,
  `tag_id` bigint(20) unsigned NOT NULL,
  `logo` varchar(1024) NOT NULL DEFAULT '',
  `picture_list` text,
  `create_time` int(10) unsigned NOT NULL,
  `update_time` int(10) unsigned NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_platform_tag` (`platform`,`tag_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table tag_relation_tab
# ------------------------------------------------------------

DROP TABLE IF EXISTS `tag_relation_tab`;

CREATE TABLE `tag_relation_tab` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `reference_id` bigint(20) unsigned NOT NULL COMMENT 'reference id this record belongs to',
  `reference_type` tinyint(4) unsigned NOT NULL COMMENT 'reference type eg:top-merchant/merchant/outlet',
  `tag_id` bigint(20) unsigned NOT NULL,
  `type_id` bigint(20) unsigned NOT NULL,
  `delete_time` int(10) unsigned NOT NULL,
  `create_time` int(10) unsigned NOT NULL,
  `update_time` int(10) unsigned NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_reference_tag` (`reference_id`,`reference_type`,`tag_id`,`delete_time`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table tag_type_tab
# ------------------------------------------------------------

DROP TABLE IF EXISTS `tag_type_tab`;

CREATE TABLE `tag_type_tab` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `type_name` varchar(256) NOT NULL,
  `preset` tinyint(1) unsigned NOT NULL DEFAULT '0',
  `create_time` int(10) unsigned NOT NULL,
  `update_time` int(10) unsigned NOT NULL,
  `m_status` tinyint(4) unsigned NOT NULL DEFAULT '0',
  `color` varchar(64) NOT NULL DEFAULT '',
  `delete_time` int(10) unsigned NOT NULL DEFAULT '0',
  `create_way` tinyint(4) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_type_name` (`type_name`,`delete_time`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table top_merchant_info_tab
# ------------------------------------------------------------

DROP TABLE IF EXISTS `top_merchant_info_tab`;

CREATE TABLE `top_merchant_info_tab` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `m_name` varchar(256) NOT NULL,
  `m_status` tinyint(4) unsigned NOT NULL,
  `logo` varchar(1024) NOT NULL,
  `picture_list` text,
  `m_type` int(8) unsigned NOT NULL,
  `airpay_uid` bigint(20) unsigned NOT NULL DEFAULT '0',
  `contact` varchar(16) NOT NULL COMMENT 'top-merchant contact number',
  `email` varchar(512) NOT NULL DEFAULT '',
  `state` varchar(128) NOT NULL DEFAULT '',
  `city` varchar(128) NOT NULL DEFAULT '',
  `district` varchar(128) NOT NULL DEFAULT '',
  `sub_district` varchar(128) NOT NULL DEFAULT '',
  `address` varchar(1024) NOT NULL DEFAULT '',
  `province_id` bigint(20) unsigned NOT NULL DEFAULT '0',
  `district_id` bigint(20) unsigned NOT NULL DEFAULT '0',
  `subdistrict_id` bigint(20) unsigned NOT NULL DEFAULT '0',
  `m_description` varchar(1024) NOT NULL DEFAULT '',
  `mcc` int(8) unsigned NOT NULL COMMENT 'MCC',
  `tc` varchar(1024) NOT NULL DEFAULT '' COMMENT 'T&C,TermsAndConditions',
  `main_person_info_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT 'responsible person info id',
  `legal_info_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT 'merchant legal info id',
  `finance_info_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT 'merchant finance info id',
  `create_time` int(10) unsigned NOT NULL,
  `update_time` int(10) unsigned NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_name` (`m_name`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;




/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
