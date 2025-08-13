INSERT INTO `isms_country` (
  `id`, `name_zh`, `name_en`, `iso_code`, `continent`, `created_at`, `updated_at`
) VALUES
  (1, '未知国家', 'Unknown Country', 'XX', '未知', NOW(), NOW()),

  -- 亚洲
  (2, '中国', 'China', 'CN', '亚洲', NOW(), NOW()),
  (3, '日本', 'Japan', 'JP', '亚洲', NOW(), NOW()),
  (4, '韩国', 'South Korea', 'KR', '亚洲', NOW(), NOW()),
  (5, '以色列', 'Israel', 'IL', '亚洲', NOW(), NOW()),
  (6, '印度', 'India', 'IN', '亚洲', NOW(), NOW()),
  (7, '新加坡', 'Singapore', 'SG', '亚洲', NOW(), NOW()),

  -- 欧洲
  (8, '德国', 'Germany', 'DE', '欧洲', NOW(), NOW()),
  (9, '法国', 'France', 'FR', '欧洲', NOW(), NOW()),
  (10, '英国', 'United Kingdom', 'GB', '欧洲', NOW(), NOW()),
  (11, '瑞士', 'Switzerland', 'CH', '欧洲', NOW(), NOW()),
  (12, '瑞典', 'Sweden', 'SE', '欧洲', NOW(), NOW()),
  (13, '荷兰', 'Netherlands', 'NL', '欧洲', NOW(), NOW()),
  (14, '意大利', 'Italy', 'IT', '欧洲', NOW(), NOW()),
  (15, '芬兰', 'Finland', 'FI', '欧洲', NOW(), NOW()),
  (16, '丹麦', 'Denmark', 'DK', '欧洲', NOW(), NOW()),
  (17, '比利时', 'Belgium', 'BE', '欧洲', NOW(), NOW()),
  (18, '奥地利', 'Austria', 'AT', '欧洲', NOW(), NOW()),
  (19, '挪威', 'Norway', 'NO', '欧洲', NOW(), NOW()),
  (20, '西班牙', 'Spain', 'ES', '欧洲', NOW(), NOW()),
  (21, '俄罗斯', 'Russia', 'RU', '欧洲', NOW(), NOW()),
  (22, '爱尔兰', 'Ireland', 'IE', '欧洲', NOW(), NOW()),
  (23, '捷克', 'Czech Republic', 'CZ', '欧洲', NOW(), NOW()),
  (24, '波兰', 'Poland', 'PL', '欧洲', NOW(), NOW()),
  (25, '乌克兰', 'Ukraine', 'UA', '欧洲', NOW(), NOW()),

  -- 美洲
  (26, '美国', 'United States', 'US', '美洲', NOW(), NOW()),
  (27, '加拿大', 'Canada', 'CA', '美洲', NOW(), NOW()),
  (28, '巴西', 'Brazil', 'BR', '美洲', NOW(), NOW()),
  (29, '墨西哥', 'Mexico', 'MX', '美洲', NOW(), NOW()),

  -- 大洋洲
  (30, '澳大利亚', 'Australia', 'AU', '大洋洲', NOW(), NOW()),
  (31, '新西兰', 'New Zealand', 'NZ', '大洋洲', NOW(), NOW()),

  -- 其他地区
  (32, '土耳其', 'Turkey', 'TR', '欧亚', NOW(), NOW()),
  (33, '南非', 'South Africa', 'ZA', '非洲', NOW(), NOW());