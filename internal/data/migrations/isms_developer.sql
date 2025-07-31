CREATE TABLE `isms_developer` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `name_zh` varchar(200) NOT NULL COMMENT '开发商中文名称（如：施耐德电气）',
  `name_en` varchar(200) NOT NULL COMMENT '开发商英文名称（如：Schneider Electric）',
  `country_id` SMALLINT UNSIGNED NOT NULL COMMENT '所属国家ID（关联isms_country.id，通过代码逻辑维护关联）', -- 注释说明关联方式
  `website` varchar(500) DEFAULT NULL COMMENT '官方网站URL',
  `description` text COMMENT '开发商简介',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_name_zh` (`name_zh`) COMMENT '中文名称唯一',
  KEY `idx_country` (`country_id`) -- 保留索引，优化按国家查询的性能
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='软件开发商表（国家关联通过代码逻辑维护，非数据库外键）';