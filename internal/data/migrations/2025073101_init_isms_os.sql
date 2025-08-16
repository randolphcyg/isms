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
  ('SUSE Linux Enterprise', '15', 'ppc64le', 'SUSE', 2018, '企业级Linux，适用于PowerPC架构服务器', NOW(), NOW()),

  -- MS Access 相关记录（注：实际应区分应用软件与操作系统）
  ('Microsoft Access', '2010', 'x86', 'Microsoft', 2010, '32位版本，Microsoft Office套件中的关系型数据库管理系统，支持Jet数据库引擎，适用于小型数据库应用开发', NOW(), NOW()),
  ('Microsoft Access', '2010', 'x64', 'Microsoft', 2010, '64位版本，需运行在64位Windows系统，支持更大内存寻址，提升对大型数据库的处理能力', NOW(), NOW()),
  ('Microsoft Access', '2013', 'x86', 'Microsoft', 2012, '32位版本，新增应用程序部件、改进的数据宏功能，集成Office 365服务', NOW(), NOW()),
  ('Microsoft Access', '2013', 'x64', 'Microsoft', 2012, '64位版本，优化对大文件和复杂查询的处理性能，兼容64位Office组件', NOW(), NOW()),
  ('Microsoft Access', '2016', 'x86', 'Microsoft', 2015, '32位版本，支持SharePoint集成、数据损失 prevention (DLP) 功能，增强安全性', NOW(), NOW()),
  ('Microsoft Access', '2016', 'x64', 'Microsoft', 2015, '64位版本，提升与大型Excel文件的交互效率，支持更多并发用户访问', NOW(), NOW()),
  ('Microsoft Access', '2019', 'x86', 'Microsoft', 2018, '32位版本，包含在Office 2019套件中，新增现代图表、改进的SQL视图，优化触摸设备支持', NOW(), NOW()),
  ('Microsoft Access', '2019', 'x64', 'Microsoft', 2018, '64位版本，强化对Access数据库引擎的64位支持，提升复杂报表生成速度', NOW(), NOW())

  ('Microsoft Windows Server', '2012', 'x64', 'Microsoft', 2012, '包含Standard、Datacenter等版本', NOW(), NOW()),
  ('Microsoft Windows Server', '2022', 'x64', 'Microsoft', 2022, '包含Standard、Datacenter等版本', NOW(), NOW()),

  -- 补充Ubuntu系统
  ('Ubuntu', '18.04 LTS', 'x64', 'Canonical', 2018, '包含Desktop、Server等官方版本', NOW(), NOW()),
  
  -- Red Hat Enterprise Workstation/Server 系统
  ('Red Hat Enterprise Linux', '8.6', 'x64', 'Red Hat', 2022, '企业级Linux发行版', NOW(), NOW()),
  ('Red Hat Enterprise Linux', '8.7', 'x64', 'Red Hat', 2023, '企业级Linux发行版', NOW(), NOW()),
  ('Red Hat Enterprise Linux', '9.0', 'x64', 'Red Hat', 2022, '企业级Linux发行版', NOW(), NOW()),
  ('Red Hat Enterprise Linux', '9.1', 'x64', 'Red Hat', 2023, '企业级Linux发行版', NOW(), NOW()),
  
  -- SUSE Linux Enterprise Server 系统
  ('SUSE Linux Enterprise Server', '12 SP5', 'x64', 'SUSE', 2020, '企业级Linux发行版', NOW(), NOW()),
  ('SUSE Linux Enterprise Server', '15 SP3', 'x64', 'SUSE', 2021, '企业级Linux发行版', NOW(), NOW()),
  ('SUSE Linux Enterprise Server', '15 SP4', 'x64', 'SUSE', 2022, '企业级Linux发行版', NOW(), NOW()),

  -- 补充Windows历史系统
  ('Microsoft Windows', 'Vista', 'x86', 'Microsoft', 2007, '包含家庭版、商业版、旗舰版等所有细分版本', NOW(), NOW()),
  ('Microsoft Windows', 'Vista', 'x64', 'Microsoft', 2007, '64位版本，包含家庭版、商业版、旗舰版等所有细分版本', NOW(), NOW()),
  ('Microsoft Windows', 'XP', 'x86', 'Microsoft', 2001, '包含家庭版、专业版等所有细分版本（SP1及后续更新）', NOW(), NOW()),
  ('Microsoft Windows', '2000', 'x86', 'Microsoft', 2000, '包含专业版、服务器版等所有细分版本', NOW(), NOW()),
  ('Microsoft Windows', 'NT', 'x86', 'Microsoft', 1999, 'Windows NT 4.0及其后续更新版本', NOW(), NOW()),
  ('Microsoft Windows', 'ME', 'x86', 'Microsoft', 2000, 'Windows Millennium Edition', NOW(), NOW()),
  ('Microsoft Windows', '98', 'x86', 'Microsoft', 1998, '包含第二版(SE)在内的所有更新版本', NOW(), NOW()),
  ('Microsoft Windows', '95', 'x86', 'Microsoft', 1995, '包含OSR2在内的所有更新版本', NOW(), NOW()),
  ('Microsoft Windows', 'XP', 'x64', 'Microsoft', 2005, '64位版本，基于Windows Server 2003代码', NOW(), NOW()),

  -- 补充 Fedora 系统
  ('Fedora', '37', 'x64', 'Red Hat', 2022, '由Fedora Project社群开发、Red Hat赞助的Linux发行版', NOW(), NOW()),
  ('Fedora', '37', 'arm64', 'Red Hat', 2022, '适用于ARM64架构设备', NOW(), NOW()),

  -- 补充 macOS 系统
  ('macOS', '13 Ventura', 'x64', 'Apple', 2022, 'Apple开发的Mac操作系统，支持Apple Silicon和Intel处理器', NOW(), NOW()),
  ('macOS', '14 Sonoma', 'x64', 'Apple', 2023, 'Apple开发的Mac操作系统，引入桌面小组件、Safari浏览器配置文件等功能', NOW(), NOW()),

  -- Android 系统
  ('Android', '8.0', 'arm64', 'Google', 2017, 'Android Oreo，代号奥利奥', NOW(), NOW()),
  ('Android', '8.1', 'arm64', 'Google', 2017, 'Android Oreo，小版本更新', NOW(), NOW()),
  ('Android', '9.0', 'arm64', 'Google', 2018, 'Android Pie，代号派', NOW(), NOW()),
  ('Android', '10.0', 'arm64', 'Google', 2019, 'Android 10，首次不用甜品命名', NOW(), NOW()),
  ('Android', '11.0', 'arm64', 'Google', 2020, 'Android 11，代号R', NOW(), NOW()),
  ('Android', '12.0', 'arm64', 'Google', 2021, 'Android 12，代号S', NOW(), NOW()),
  ('Android', '13.0', 'arm64', 'Google', 2022, 'Android 13，代号T', NOW(), NOW()),
  ('Android', '14.0', 'arm64', 'Google', 2023, 'Android 14，代号UpsideDownCake', NOW(), NOW()),
  ('Android', '15.0', 'arm64', 'Google', 2024, 'Android 15，代号Vanilla Ice Cream', NOW(), NOW());