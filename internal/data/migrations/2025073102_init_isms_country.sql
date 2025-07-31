-- 国家表初始化（筛选有工业软件研发能力的国家，固定ID便于关联）
INSERT INTO `isms_country` (
  `id`, `name_zh`, `name_en`, `iso_code`, `continent`, `created_at`, `updated_at`
) VALUES
  -- 亚洲
  (1, '中国', 'China', 'CN', '亚洲', NOW(), NOW()),
  (2, '日本', 'Japan', 'JP', '亚洲', NOW(), NOW()),
  (3, '韩国', 'South Korea', 'KR', '亚洲', NOW(), NOW()),
  (4, '以色列', 'Israel', 'IL', '亚洲', NOW(), NOW()),
  (5, '新加坡', 'Singapore', 'SG', '亚洲', NOW(), NOW()),
  (6, '印度', 'India', 'IN', '亚洲', NOW(), NOW()),

  -- 欧洲
  (7, '德国', 'Germany', 'DE', '欧洲', NOW(), NOW()),
  (8, '法国', 'France', 'FR', '欧洲', NOW(), NOW()),
  (9, '英国', 'United Kingdom', 'GB', '欧洲', NOW(), NOW()),
  (10, '意大利', 'Italy', 'IT', '欧洲', NOW(), NOW()),
  (11, '俄罗斯', 'Russia', 'RU', '欧洲', NOW(), NOW()),
  (12, '瑞士', 'Switzerland', 'CH', '欧洲', NOW(), NOW()),
  (13, '瑞典', 'Sweden', 'SE', '欧洲', NOW(), NOW()),
  (14, '荷兰', 'Netherlands', 'NL', '欧洲', NOW(), NOW()),
  (15, '比利时', 'Belgium', 'BE', '欧洲', NOW(), NOW()),
  (16, '奥地利', 'Austria', 'AT', '欧洲', NOW(), NOW()),
  (17, '芬兰', 'Finland', 'FI', '欧洲', NOW(), NOW()),
  (18, '丹麦', 'Denmark', 'DK', '欧洲', NOW(), NOW()),
  (19, '挪威', 'Norway', 'NO', '欧洲', NOW(), NOW()),
  (20, '乌克兰', 'Ukraine', 'UA', '欧洲', NOW(), NOW()),

  -- 美洲
  (21, '美国', 'United States', 'US', '美洲', NOW(), NOW()),
  (22, '加拿大', 'Canada', 'CA', '美洲', NOW(), NOW()),
  (23, '巴西', 'Brazil', 'BR', '美洲', NOW(), NOW()),

  -- 大洋洲
  (24, '澳大利亚', 'Australia', 'AU', '大洋洲', NOW(), NOW()),
  (25, '新西兰', 'New Zealand', 'NZ', '大洋洲', NOW(), NOW());