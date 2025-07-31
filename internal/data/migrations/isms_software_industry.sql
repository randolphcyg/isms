CREATE TABLE `isms_software_industry` (
  `software_id` INT UNSIGNED NOT NULL COMMENT '软件ID（关联isms_software.id）',
  `industry_id` SMALLINT UNSIGNED NOT NULL COMMENT '行业ID（关联isms_industry.id）',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '关联创建时间',
  PRIMARY KEY (`software_id`, `industry_id`) COMMENT '联合主键防重复',
  KEY `idx_industry` (`industry_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='软件与行业多对多关联表';