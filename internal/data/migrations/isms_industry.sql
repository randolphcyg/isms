-- 工业行业分类表
CREATE TABLE `isms_industry` (
  id SMALLINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `category_code` char(1) NOT NULL COMMENT '大类代码(B/C/D/E等)',
  `category_name` varchar(50) NOT NULL COMMENT '大类名称(采矿业/制造业等)',
  `subcategory_code` char(2) NOT NULL COMMENT '小类代码(06/07/13等)',
  `subcategory_name` varchar(100) NOT NULL COMMENT '小类名称(煤炭开采和洗选业等)',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_category_subcategory` (`category_code`,`subcategory_code`),
  KEY `idx_category` (`category_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='第二产业行业分类表';
