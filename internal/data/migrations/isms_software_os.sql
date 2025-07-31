CREATE TABLE `isms_software_os` (
  `software_id` INT UNSIGNED NOT NULL COMMENT '软件ID（关联isms_software.id）',
  `os_id` INT UNSIGNED NOT NULL COMMENT '操作系统ID（关联isms_operating_system.id）',
  `note` varchar(500) DEFAULT NULL COMMENT '备注（如仅支持专业版）',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '关联创建时间',
  PRIMARY KEY (`software_id`, `os_id`) COMMENT '联合主键防重复',
  KEY `idx_os` (`os_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='软件与操作系统多对多关联表';