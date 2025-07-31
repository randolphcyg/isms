CREATE TABLE `isms_country` (
  `id` SMALLINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `name_zh` varchar(50) NOT NULL COMMENT '国家中文名称（如：中国）',
  `name_en` varchar(100) NOT NULL COMMENT '国家英文名称（如：China）',
  `iso_code` char(2) NOT NULL COMMENT 'ISO 3166-1 alpha-2两位编码（如：CN）',
  `continent` varchar(20) DEFAULT NULL COMMENT '所属大洲（如：亚洲、欧洲）',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_iso_code` (`iso_code`) COMMENT 'ISO编码唯一',
  UNIQUE KEY `uk_name_zh` (`name_zh`) COMMENT '中文名称唯一'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='国家/地区表';