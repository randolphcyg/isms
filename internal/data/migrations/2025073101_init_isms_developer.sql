INSERT INTO `isms_developer` (
    `id`, `name_zh`, `name_en`, `website`, `description`, `country_id`, `created_at`, `updated_at`
) VALUES
  -- 未知开发商（ID=1，关联未知国家ID=1）
  (1, '未知开发商', 'Unknown Developer', '', '未明确的开发商', 1, NOW(), NOW()),

  -- 美国厂商（ID从2开始，关联美国ID=26）
  (2, '欧特克', 'Autodesk, Inc.', 'www.autodesk.com', '全球领先的设计软件公司', 26, NOW(), NOW()),
  (3, 'ANSYS', 'ANSYS, Inc.', 'www.ansys.com', '工程仿真软件领军企业', 26, NOW(), NOW()),
  (4, '楷登电子', 'Cadence Design Systems, Inc.', 'www.cadence.com', 'EDA领域龙头', 26, NOW(), NOW()),
  (5, '达索SolidWorks', 'Dassault Systèmes SolidWorks', 'www.solidworks.com', '3D CAD设计软件', 26, NOW(), NOW()),
  (6, '美国参数技术公司', 'PTC Inc.', 'www.ptc.com', '产品生命周期管理软件', 26, NOW(), NOW()),
  (7, 'CNC软件公司', 'CNC Software, Inc.', 'www.mastercam.com', '知名CAM软件开发商', 26, NOW(), NOW()),
  (8, 'Cambrio公司', 'Cambrio, Inc.', 'www.cambrio.com', '模具制造软件', 26, NOW(), NOW()),
  (9, '国家仪器', 'National Instruments (NI)', 'www.ni.com', '测试测量软件', 26, NOW(), NOW()),
  (10, 'DownStream科技', 'DownStream Technologies', 'www.downstreamtech.com', 'PCB制造软件', 26, NOW(), NOW()),
  (11, 'LightBurn软件公司', 'LightBurn Software', 'www.lightburnsoftware.com', '激光控制软件', 26, NOW(), NOW()),

  -- 中国厂商（ID从12开始，关联中国ID=2）
  (12, '中望软件', 'ZWSOFT Co., Ltd.', 'www.zwsoft.com', '中国领先CAD开发商', 2, NOW(), NOW()),
  (13, '北京华大九天科技股份有限公司', 'Beijing Huada Jiutian Technology Co., Ltd.', 'www.empyrean.com.cn', '集成电路EDA软件', 2, NOW(), NOW()),
  (14, '华为云计算技术有限公司', 'Huawei Cloud Computing Technology Co., Ltd.', 'www.huaweicloud.com', '企业级云ERP服务', 2, NOW(), NOW()),
  (15, '卡奥斯物联科技股份有限公司', 'Kaos IoT Technology Co., Ltd.', 'www.cosmoplat.com', '工业互联网平台', 2, NOW(), NOW()),
  (16, '广联达科技股份有限公司', 'Glodon Company Limited', 'www.glodon.com', '建筑工程信息化软件', 2, NOW(), NOW()),
  (17, '上海柏楚电子科技股份有限公司', 'Shanghai Bodor Laser Technology Co., Ltd.', 'www.bodor.com.cn', '激光切割控制软件', 2, NOW(), NOW()),
  (18, '用友网络科技股份有限公司', 'Yonyou Network Technology Co., Ltd.', 'www.yonyou.com', '中国ERP软件', 2, NOW(), NOW()),
  (19, '国电南瑞科技股份有限公司', 'NARI Technology Co., Ltd.', 'www.nari.com.cn', '中国电网自动化', 2, NOW(), NOW()),
  (20, '上海宝信软件股份有限公司', 'Baosight Software Co., Ltd.', 'www.baosight.com', '中国钢铁信息化', 2, NOW(), NOW()),
  (21, '金蝶软件（中国）有限公司', 'Kingdee Software (China) Co., Ltd.', 'www.kingdee.com', '中国企业管理软件', 2, NOW(), NOW()),

  -- 德国厂商（ID从22开始，关联德国ID=8）
  (22, '西门子', 'Siemens AG', 'www.siemens.com/plm', '工业软件巨头', 8, NOW(), NOW()),
  (23, 'EPLAN公司', 'EPLAN Software & Service', 'eplan.com', '电气设计软件领军者', 8, NOW(), NOW()),
  (24, 'Fritzing公司', 'Fritzing', 'https://fritzing.org/', '交互设计实验室项目', 8, NOW(), NOW()),
  (25, 'Softing公司', 'Softing AG', 'www.softing.com', '德国工业通信软件', 8, NOW(), NOW()),
  (26, 'CFTurbo公司', 'CFTurbo GmbH', 'www.cfturbo.com', '德国涡轮设计软件', 8, NOW(), NOW()),

  -- 法国厂商（ID从27开始，关联法国ID=9）
  (27, '达索系统', 'Dassault Systèmes', 'www.3ds.com', 'PLM与3D设计巨头', 9, NOW(), NOW()),
  (28, '施耐德电气', 'Schneider Electric', 'www.schneider-electric.com', '能源管理与自动化', 9, NOW(), NOW()),

  -- 其他国家厂商（ID从29开始）
  (29, '三菱电机', 'Mitsubishi Electric Corporation', 'www.mitsubishielectric.com', '日本PLC编程软件', 3, NOW(), NOW()), -- 日本ID=3
  (30, 'SolidCAM公司', 'SolidCAM Ltd.', 'www.solidcam.com', '以色列CAD/CAM方案', 5, NOW(), NOW()), -- 以色列ID=5
  (31, 'KISSsoft公司', 'KISSsoft AG', 'www.kisssoft.com', '瑞士机械传动设计', 11, NOW(), NOW()), -- 瑞士ID=11
  (32, '意法半导体', 'STMicroelectronics', 'www.st.com', '瑞士嵌入式设计软件', 11, NOW(), NOW()), -- 瑞士ID=11
  (33, 'Labcenter电子', 'Labcenter Electronics', 'www.labcenter.com', '英国Proteus设计软件', 10, NOW(), NOW()), -- 英国ID=10
  (34, '奥腾电子', 'Altium Limited', 'www.altium.com', '澳大利亚PCB设计', 30, NOW(), NOW()), -- 澳大利亚ID=30
  (35, 'ASCON集团', 'ASCON Group', 'www.ascon.ru', '俄罗斯CAD软件', 21, NOW(), NOW()), -- 俄罗斯ID=21
  (36, 'Model Studio CS', 'Model Studio CS', 'https://modelstudiocs.ru/', '俄罗斯工程设计软件，提供电气设计、管道设计等工业设施设计解决方案', 21, NOW(), NOW()), -- 俄罗斯ID=21
  (37, 'Alibre公司', 'Alibre, Inc.', 'www.alibre.com', '专业的3D CAD/CAM软件开发商', 26, NOW(), NOW()),
  (38, 'ntop', 'ntop', 'ntop.com', '开源网络监控工具，用于实时监控网络流量和使用情况', 26, NOW(), NOW()),
  (39, 'Ascon', 'Ascon', 'https://ascon.net/', '俄罗斯最大的工业软件厂商,是工业和建筑数字化的技术促进者，专注于设计和制造流程自动化领域，其产品用于 3D 建模、生产流程、管理工程数据、管理产品周期等，在汽车、重型机械、航空航天、国防等多个行业都有应用，全球有超过 14000 家客户。', 21, NOW(), NOW()),
  (40, 'Metalix公司', 'Metalix', 'www.metalix.com', '专业的钣金设计和制造软件', 26, NOW(), NOW());