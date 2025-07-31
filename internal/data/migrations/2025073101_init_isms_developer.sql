-- 开发商表初始化（关联国家表ID）
INSERT INTO `isms_developer` (
    `id`, `name_zh`, `name_en`, `website`, `description`, `country_id`, `created_at`, `updated_at`
) VALUES
-- 美国厂商
(1, '欧特克', 'Autodesk, Inc.', 'www.autodesk.com', '全球领先的设计软件公司，产品涵盖CAD、BIM、工程仿真等', 21, NOW(), NOW()),
(2, 'ANSYS', 'ANSYS, Inc.', 'www.ansys.com', '全球领先的工程仿真软件公司，涵盖多物理场仿真', 21, NOW(), NOW()),
(3, '参数技术公司', 'PTC Inc.', 'www.ptc.com', '主打Creo和Windchill，专注产品生命周期管理', 21, NOW(), NOW()),
(4, 'CNC软件公司', 'CNC Software, Inc.', 'www.mastercam.com', '知名CAM软件开发商，Mastercam系列产品', 21, NOW(), NOW()),
(5, 'Cambrio公司', 'Cambrio, Inc.', 'www.cambrio.com', '专注模具制造与数控加工软件', 21, NOW(), NOW()),
(6, '达索SolidWorks', 'Dassault Systèmes SolidWorks', 'www.solidworks.com', '3D CAD设计软件，应用于机械设计领域', 21, NOW(), NOW()),
(7, '国家仪器', 'National Instruments (NI)', 'www.ni.com', '测试测量与自动化软件提供商', 21, NOW(), NOW()),
(8, '楷登电子', 'Cadence Design Systems, Inc.', 'www.cadence.com', 'EDA领域领军企业，提供芯片与PCB设计软件', 21, NOW(), NOW()),
(9, 'DownStream科技', 'DownStream Technologies', 'www.downstreamtech.com', 'PCB制造与DFM软件开发商', 21, NOW(), NOW()),
(10, 'LightBurn软件公司', 'LightBurn Software', 'www.lightburnsoftware.com', '激光切割与雕刻控制软件开发商', 21, NOW(), NOW()),

-- 德国厂商
(11, '西门子', 'Siemens AG', 'www.siemens.com/plm', '工业软件巨头，产品涵盖PLM与工业自动化', 7, NOW(), NOW()),
(12, 'EPLAN公司', 'EPLAN Software & Service', 'eplan.com', '专注电气设计软件，服务工业自动化领域', 7, NOW(), NOW()),
(13, 'Softing公司', 'Softing AG', 'www.softing.com', '工业通信软件与工具开发商', 7, NOW(), NOW()),
(14, 'CFTurbo公司', 'CFTurbo GmbH', 'www.cfturbo.com', '涡轮机械设计软件研发商', 7, NOW(), NOW()),
(66, 'Fritzing公司', 'Fritzing', 'https://fritzing.org/', '德国波茨坦应用科学大学（University of Applied Sciences Potsdam）的交互设计实验室（Interaction Design Lab）', 7, NOW(), NOW()),

-- 法国厂商
(15, '达索系统', 'Dassault Systèmes', 'www.3ds.com', 'PLM与3D设计软件公司，产品涵盖CATIA等', 8, NOW(), NOW()),
(16, '施耐德电气', 'Schneider Electric', 'www.schneider - electric.com', '能源管理与自动化巨头，工业软件开发商', 8, NOW(), NOW()),

-- 瑞士厂商
(17, 'KISSsoft公司', 'KISSsoft AG', 'www.kisssoft.com', '机械传动设计与计算软件提供商', 12, NOW(), NOW()),
(18, '意法半导体', 'STMicroelectronics', 'www.st.com', '半导体公司，开发嵌入式系统设计软件', 12, NOW(), NOW()),

-- 以色列厂商
(19, 'SolidCAM公司', 'SolidCAM Ltd.', 'www.solidcam.com', 'CAD/CAM集成解决方案提供商', 4, NOW(), NOW()),

-- 中国厂商
(20, '中望软件', 'ZWSOFT Co., Ltd.', 'www.zwsoft.com', '中国领先的CAD软件开发商', 1, NOW(), NOW()),
(21, '卡奥斯物联科技股份有限公司', 'Kaos IoT Technology Co., Ltd.', 'www.cosmoplat.com', '工业互联网平台开发商，提供MOM、SIM等解决方案', 1, NOW(), NOW()),
(22, '用友网络科技股份有限公司', 'Yonyou Network Technology Co., Ltd.', 'www.yonyou.com', '提供ERP、CRM、SCM等软件及云服务', 1, NOW(), NOW()),
(23, '国电南瑞科技股份有限公司', 'NARI Technology Co., Ltd.', 'www.nari.com.cn', '专注电网自动化及工业控制软件', 1, NOW(), NOW()),
(24, '上海宝信软件股份有限公司', 'Baosight Software Co., Ltd.', 'www.baosight.com', '钢铁信息化软件开发商', 1, NOW(), NOW()),
(25, '金蝶软件（中国）有限公司', 'Kingdee Software (China) Co., Ltd.', 'www.kingdee.com', '提供金蝶云EBC等企业管理软件', 1, NOW(), NOW()),
(26, '广联达科技股份有限公司', 'Glodon Company Limited', 'www.glodon.com', '专注建筑工程信息化软件', 1, NOW(), NOW()),
(27, '北京华大九天科技股份有限公司', 'Beijing Huada Jiutian Technology Co., Ltd.', 'www.empyrean.com.cn', '集成电路EDA软件开发商', 1, NOW(), NOW()),
(28, '浙江中控技术股份有限公司', 'Zhejiang SUPCON Technology Co., Ltd.', 'www.supcontech.com', '提供DCS、MES、PLC等工业自动化软件', 1, NOW(), NOW()),
(29, '上海柏楚电子科技股份有限公司', 'Shanghai Bodor Laser Technology Co., Ltd.', 'www.bodor.com.cn', '激光切割控制软件开发商', 1, NOW(), NOW()),
(30, '固高科技股份有限公司', 'Googol Technology Co., Ltd.', 'www.googol.com.cn', '提供开放式二次开发平台相关软件', 1, NOW(), NOW()),
(31, '华为云计算技术有限公司', 'Huawei Cloud Computing Technology Co., Ltd.', 'www.huaweicloud.com', '提供企业级云ERP等软件服务', 1, NOW(), NOW()),
(32, '北京神舟航天软件技术股份有限公司', 'Beijing Shenzhou Aerospace Software Technology Co., Ltd.', 'www.hty.com.cn', '专注军工航天领域工业软件', 1, NOW(), NOW()),
(33, '广州赛意信息科技股份有限公司', 'Saiyi Information Technology Co., Ltd.', 'www.chinasie.com', '工业管理软件开发商', 1, NOW(), NOW()),
(34, '南京科远智慧科技集团股份有限公司', 'Nanjing Keyuan Intelligence Technology Group Co., Ltd.', 'www.keyuanauto.com', '提供DCS、PLC等工业自动化软件', 1, NOW(), NOW()),
(35, '安世亚太科技股份有限公司', 'PERA Global Technology Co., Ltd.', 'www.peraglobal.com', 'CAE领军企业', 1, NOW(), NOW()),
(36, '北京机械工业自动化研究所有限公司', 'Beijing Institute of Mechanical Industry Automation Co., Ltd.', 'www.bimai.com', '提供ERP等工业管理软件', 1, NOW(), NOW()),
(37, '北京盈建科软件股份有限公司', 'Beijing YJK Software Co., Ltd.', 'www.yjk.cn', '专注建筑结构设计软件', 1, NOW(), NOW()),
(38, '依柯力信息科技（上海）股份有限公司', 'Inkelink Information Technology (Shanghai) Co., Ltd.', 'www.inkelink.com', '提供MOM、IQM、MES等软件及数字孪生解决方案', 1, NOW(), NOW()),
(39, '安徽容知日新科技股份有限公司', 'Richtech Rockontrol Co., Ltd.', 'www.uzertech.com', '设备管理软件开发商', 1, NOW(), NOW()),
(40, '中车信息技术有限公司', 'CRRC Information Technology Co., Ltd.', 'www.crrcit.com', 'PLM软件开发商，服务于轨道交通行业', 1, NOW(), NOW()),
(41, '北京数码大方科技股份有限公司', 'Digital大方 Technology Co., Ltd.', 'www.digital - nc.com', '提供CAD电子图板等软件', 1, NOW(), NOW()),
(42, '索为技术股份有限公司', 'Sowell Technology Co., Ltd.', 'www.sowelltech.com', '产品需求管理系统软件开发商', 1, NOW(), NOW()),
(43, '苏州浩辰软件股份有限公司', 'Suzhou Gosunsoft Co., Ltd.', 'www.gstarcad.com', 'CAD软件与云方案提供商', 1, NOW(), NOW()),
(44, '中仿智能科技（上海）股份有限公司', 'Shanghai CF Tech Co., Ltd.', 'www.cftooldesign.com', '提供模拟器、虚拟仿真系统软件', 1, NOW(), NOW()),
(45, '上海哥瑞利软件股份有限公司', 'Shanghai Gorilly Software Co., Ltd.', 'www.gorillysoft.com', '半导体MES、CIM软件开发商', 1, NOW(), NOW()),
(46, '北京亚控科技发展有限公司', 'Beijing Kingview Technology Co., Ltd.', 'www.kingview.com', '工业APP组态平台开发商', 1, NOW(), NOW()),
(47, '苏州同元软控信息技术有限公司', 'Suzhou Tongyuan Softcontrol Information Technology Co., Ltd.', 'www.mworks - sim.com', '提供系统设计与仿真验证平台MWorks', 1, NOW(), NOW()),
(48, '北京安怀信科技股份有限公司', 'Beijing Anhuai Xin Technology Co., Ltd.', 'www.perasim.com', 'CAE软件开发商', 1, NOW(), NOW()),
(49, '武汉开目信息技术股份有限公司', 'Wuhan Kaimu Information Technology Co., Ltd.', 'www.kmsoft.com.cn', '提供CAPP、3DCAPP、CAD、PLM等软件', 1, NOW(), NOW()),
(50, '南京国睿信维软件有限公司', 'Nanjing Goriway Software Co., Ltd.', 'www.goriway.com', 'PLM软件开发商', 1, NOW(), NOW()),

-- 日本厂商
(51, '三菱电机', 'Mitsubishi Electric Corporation', 'www.mitsubishielectric.com', '工业自动化巨头，PLC编程软件开发商', 2, NOW(), NOW()),

-- 英国厂商
(52, 'Labcenter电子', 'Labcenter Electronics', 'www.labcenter.com', 'Proteus系列电子设计软件开发商', 9, NOW(), NOW()),
(53, 'Matrix TSL公司', 'Matrix TSL Limited', 'matrixtsl.com', '嵌入式系统设计软件开发商', 9, NOW(), NOW()),

-- 加拿大厂商
(54, 'EFICAD公司', 'EFICAD Inc.', 'swood.eficad.com', '家具设计软件SWOOD开发商', 22, NOW(), NOW()),

-- 芬兰厂商
(55, 'MagiCAD公司', 'MagiCAD Group', 'http://www.magicad.com', '建筑机电BIM软件提供商', 17, NOW(), NOW()),

-- 比利时厂商
(56, '艾司科', 'Esko Group', 'https://www.esko.com/', '包装印刷行业软件开发商', 15, NOW(), NOW()),

-- 澳大利亚厂商
(57, '奥腾电子', 'Altium Limited', 'www.altium.com', 'PCB设计软件Altium Designer开发商', 24, NOW(), NOW()),

-- 乌克兰厂商
(58, 'DipTrace公司', 'DipTrace LLC', 'https://diptrace.com/', 'PCB设计软件开发商', 20, NOW(), NOW()),

-- 俄罗斯厂商
(59, 'ASCON集团', 'ASCON Group', 'www.ascon.ru', '俄罗斯CAD软件开发商，代表产品КОМПАС - 3D', 11, NOW(), NOW()),
(60, 'C3D Labs公司', 'C3D Labs', 'www.c3dlabs.com', '开发几何内核C3D，用于工程软件创建', 11, NOW(), NOW()),
(61, 'STC“APM”公司', 'STC “APM”', '无公开官网', '俄罗斯工程软件开发商，Razvitie联盟成员', 11, NOW(), NOW()),
(62, 'ADEM公司', 'ADEM', '无公开官网', '专注CAD相关软件，Razvitie联盟成员', 11, NOW(), NOW()),
(63, 'TESIS公司', 'TESIS', '无公开官网', '涉及工程软件领域，Razvitie联盟成员', 11, NOW(), NOW()),
(64, 'EREMEX公司', 'EREMEX', 'www.eremex.ru', '基于C3D内核开发PCB设计等软件', 11, NOW(), NOW()),
(65, 'Sigma Technology公司', 'Sigma Technology', '无公开官网', 'Razvitie联盟成员，从事CAD相关开发', 11, NOW(), NOW());