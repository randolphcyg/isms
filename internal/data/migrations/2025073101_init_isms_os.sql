INSERT INTO `isms_os` (
  `name`, `version`, `architecture`, `manufacturer`, `release_year`, `description`, `created_at`, `updated_at`
) VALUES
  -- Windows 桌面系统
  ('Microsoft Windows', '7', 'x86', 'Microsoft', 2009, '包含家庭版、专业版等所有细分版本（SP1及后续更新）', NOW(), NOW()),
  ('Microsoft Windows', '7', 'x64', 'Microsoft', 2009, '包含家庭版、专业版等所有细分版本（SP1及后续更新）', NOW(), NOW()),
  ('Microsoft Windows', '8', 'x86', 'Microsoft', 2012, '包含Windows 8及8.1的所有细分版本', NOW(), NOW()),
  ('Microsoft Windows', '8', 'x64', 'Microsoft', 2012, '包含Windows 8及8.1的所有细分版本', NOW(), NOW()),
  ('Microsoft Windows', '10', 'x86', 'Microsoft', 2015, '包含家庭版、专业版、LTSC等所有细分版本', NOW(), NOW()),
  ('Microsoft Windows', '10', 'x64', 'Microsoft', 2015, '包含家庭版、专业版、LTSC等所有细分版本', NOW(), NOW()),
  ('Microsoft Windows', '11', 'x64', 'Microsoft', 2021, '仅支持64位，包含所有细分版本', NOW(), NOW()),

  -- Windows Server 系统
  ('Microsoft Windows Server', '2016', 'x64', 'Microsoft', 2016, '包含Standard、Datacenter等版本', NOW(), NOW()),
  ('Microsoft Windows Server', '2019', 'x64', 'Microsoft', 2019, '包含Standard、Datacenter等版本', NOW(), NOW()),

  -- Ubuntu 系统
  ('Ubuntu', '20.04 LTS', 'x64', 'Canonical', 2020, '包含Desktop、Server等官方版本', NOW(), NOW()),
  ('Ubuntu', '20.04 LTS', 'arm64', 'Canonical', 2020, '适用于ARM64架构设备', NOW(), NOW()),
  ('Ubuntu', '22.04 LTS', 'x64', 'Canonical', 2022, '包含Desktop、Server等官方版本', NOW(), NOW()),

  -- CentOS 系统（开发商为Red Hat）
  ('CentOS', '7', 'x64', 'Red Hat', 2014, '企业级Linux发行版', NOW(), NOW()),
  ('CentOS Stream', '9', 'x64', 'Red Hat', 2021, '滚动更新版本，适用于开发测试', NOW(), NOW()),

  -- 工业软件常见系统（如ARM架构嵌入式）
  ('Linux', '4.19', 'arm32', 'Various', 2018, '嵌入式Linux内核，适用于32位ARM设备', NOW(), NOW()),
  ('SUSE Linux Enterprise', '15', 'ppc64le', 'SUSE', 2018, '企业级Linux，适用于PowerPC架构服务器', NOW(), NOW());