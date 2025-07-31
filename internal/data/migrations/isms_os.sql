CREATE TABLE `isms_os` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `name` varchar(200) NOT NULL COMMENT '系统名称（如：Microsoft Windows、Ubuntu）',
  `version` varchar(50) NOT NULL COMMENT '系统版本（如：7、10、20.04 LTS）',
  `architecture` ENUM('x86', 'x64', 'arm32', 'arm64', 'ppc64le', 's390x', 'riscv64') NOT NULL COMMENT '硬件架构',
  `manufacturer` varchar(200) DEFAULT NULL COMMENT '系统开发商（如：Microsoft、Canonical）',
  `release_year` SMALLINT UNSIGNED DEFAULT NULL COMMENT '发布年份',
  `description` text COMMENT '系统说明（如包含的细分版本）',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_name_version_arch` (`name`,`version`,`architecture`) COMMENT '同一系统+版本+架构唯一'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='操作系统表';