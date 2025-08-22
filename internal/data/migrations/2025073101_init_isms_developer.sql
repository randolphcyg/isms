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
  (40, 'Metalix公司', 'Metalix', 'www.metalix.com', '专业的钣金设计和制造软件', 26, NOW(), NOW()),
  (41, 'esko公司', 'ELKO EP', 'https://www.esko.com/', '一家在包装、标签和印刷行业提供软件和硬件解决方案的公司，专注于帮助企业优化包装设计、生产流程等', 23, NOW(), NOW()),
  (42, 'CAMWorks公司', 'CAMWorks', 'www.camworks.com', 'CAD/CAM集成解决方案提供商', 26, NOW(), NOW()),
  (43, 'Eficad公司', 'Eficad', 'swood.eficad.com', '核心产品为SWOOD，这是一款专门为木工行业打造的软件，于 2010 年推出，可用于木工项目管理，涵盖从设计到生产的全过程。', 9, NOW(), NOW()),
  (44, 'MagiCAD公司', 'MagiCAD Group', 'www.magicad.com', '一家全球软件解决方案和制造商服务提供商，专注于建筑行业软件和服务', 15, NOW(), NOW()),
  (45, 'Keysight公司', 'Keysight Technologies', 'www.keysight.com', '一家全球领先的电子测量技术公司，主要提供电子测试与测量仪器、系统及相关软件解决方案，应用于通信、航空航天、汽车电子、半导体等多个领域。', 26, NOW(), NOW()),
  (46, 'Intel公司', 'Intel Corporation', 'www.intel.com', '知名半导体芯片制造商', 26, NOW(), NOW()),
  (47, 'ProgeSOFT公司', 'ProgeSOFT S.A.', 'www.progesoft.com', '专注于开发低成本的 2D/3D DWG CAD 技术，以及工业自动化、工程和资源管理等领域的垂直应用程序。', 11, NOW(), NOW()),
  (48, 'Einkaufsplaner', 'Einkaufsplaner', 'www.einkaufsplaner.de', '德国采购计划软件', 8, NOW(), NOW()),
  (49, 'Wonderware', 'Wonderware', 'www.wonderware.com', '工业自动化软件解决方案提供商', 26, NOW(), NOW()),
  (50, 'Matrix TSL', 'Matrix TSL', 'www.matrixtsl.com', '工程教育领域的硬件和软件解决方案提供商', 10, NOW(), NOW()),
  (51, 'Abacom Technologies Inc.', 'Abacom Technologies Inc.', 'abacom-tech.com', '专注于以太网模块、数据协议分析仪和抓包工具等产品', 26, NOW(), NOW()),
  (52, 'Novarm Ltd.', 'Novarm Ltd.', 'https://www.novarm.com/', '专业的PCB设计软件开发商，提供DipTrace等产品', 21, NOW(), NOW());(48, '欧姆龙', 'Omron Corporation', 'www.omron.com', '日本著名的电子制造厂商，涉及工业自动化控制系统、电子控制设备元件、社会系统以及健康医疗设备等领域', 3, NOW(), NOW()),
  (53, '欧姆龙', 'Omron Corporation', 'www.omron.com', '日本著名的电子制造厂商，涉及工业自动化控制系统、电子控制设备元件、社会系统以及健康医疗设备等领域', 3, NOW(), NOW()),
  (54, '赛灵思', 'Xilinx, Inc.', 'www.xilinx.com', '可编程逻辑器件(FPGA)的发明者和领先供应商', 26, NOW(), NOW()),
  (55, '罗克韦尔自动化', 'Rockwell Automation, Inc.', 'www.rockwellautomation.com', '全球最大的工业自动化与信息技术公司之一，致力于帮助客户提高生产力并推动世界实现可持续发展', 26, NOW(), NOW()),
  (56, 'KNX协会', 'KNX Association', 'www.knx.org', 'KNX标准的制定者和维护者，提供ETS等工程工具软件', 8, NOW(), NOW()),
  (57, 'Mentor Graphics', 'Mentor Graphics Corporation', 'www.mentor.com', '电子设计自动化(EDA)技术的领导厂商，提供完整的软件和硬件设计解决方案', 26, NOW(), NOW()),
  (58, 'SIMetrix Technologies Ltd.', 'SIMetrix Technologies Ltd.', 'https://www.simplistechnologies.com/', '电路仿真软件开发商', 26, NOW(), NOW()),
  (59, '微芯科技', 'Microchip Technology', 'https://www.microchip.com/', '致力于智能、互联和安全的嵌入式控制与处理解决方案的领先供应商', 26, NOW(), NOW()),
  (60, 'Lambda Research Group', 'Lambda Research Group', 'https://lambdameet.com/', '科学会议组织机构，专门组织学术和企业界的会议', 10, NOW(), NOW()),
  (61, 'Aldec公司', 'Aldec, Inc.', 'https://www.aldec.com/', '致力于EDA工具研发，为全球IC设计师提供功能强大、易学易用的设计手段', 26, NOW(), NOW()),
  (62, 'MHJ-Software公司', 'MHJ-Software', 'https://www.mhj-tools.com/', '专注于自动化技术软件，提供PLC培训软件、数字孪生服务、机电一体化系统的3D模拟等', 8, NOW(), NOW()),
  (63, 'IAR Systems', 'IAR Systems AB', 'https://www.iar.com/', '瑞典著名嵌入式系统开发工具供应商，提供IAR Embedded Workbench等产品', 6, NOW(), NOW()),
  (64, 'LPKF激光电子股份公司', 'LPKF Laser & Electronics AG', 'https://www.lpkf.com/', '德国领先的激光技术解决方案提供商，专注于PCB原型制作、太阳能电池、塑料焊接、医疗技术等领域的精密激光加工', 8, NOW(), NOW());
  (65, '库卡', 'KUKA AG', 'www.kuka.com', '世界领先的工业机器人制造商之一，成立于1898年，总部位于德国奥格斯堡', 8, NOW(), NOW()),
  (66, 'TP-LINK', 'TP-LINK Corporation', 'www.tp-link.com', '全球领先的网络通讯设备供应商，致力于提供全面的ICT设备与解决方案', 2, NOW(), NOW()),
  (67, '中控技术股份有限公司', 'SUPCON Technology Co., Ltd.', 'www.supcon.com', '国内领先、全球化布局的智能制造整体解决方案供应商，致力于"AI+数据"核心能力的构建及落地应用，已累计服务海内外客户3万多家，覆盖化工、石化、油气、电力、制药等数十个重点行业。', 2, NOW(), NOW()),
  (68, 'Vero Software', 'Vero Software', 'www.verosoftware.com', '英国知名的CAD/CAM软件开发商，专注于为制造业提供专业的工程软件解决方案，其产品包括Alphacam等', 10, NOW(), NOW()),
  (69, '菲尼克斯电气', 'Phoenix Contact', 'www.phoenixcontact.com', '德国工业自动化和电气连接技术领域的领先企业', 8, NOW(), NOW()),
  (70, 'Spectrum软件公司', 'Spectrum', 'http://www.spectrum-digitizer.com/', '德国高精度及高速度数字化仪制造商，专注于电子信号的获取、产生和分析', 8, NOW(), NOW()),
  (71, '高云半导体', 'GOWIN Semiconductor', 'https://www.gowinsemi.com.cn/', '中国领先的FPGA供应商，专注于集成电路设计工具，提供FPGA集成开发环境(IDE)用于FPGA设计输入、代码合成、布局路由等', 2, NOW(), NOW()),
  (72, 'Synopsys', 'Synopsys, Inc.', 'www.synopsys.com', '半导体工艺和器件仿真软件开发商', 26, NOW(), NOW()),
  (73, 'Robert McNeel & Assoc', 'Robert McNeel & Associates', 'https://www.rhino3d.com/', '美国著名的3D建模软件开发商，其Rhino（犀牛）软件广泛应用于三维动画、工业制造、科研、机械设计等领域', 26, NOW(), NOW()),
  (74, '博思软件', 'Beiswenger and Associates', '', '建筑和工程领域创新技术解决方案提供商，Fuzor软件开发商', 26, NOW(), NOW()),
  (75, 'Graphisoft', 'Graphisoft', 'www.graphisoft.com', 'BIM软件开发商，Archicad软件开发商', 34, NOW(), NOW()),
  (76, 'Jan Adamec', 'Jan Adamec', 'www.roomarranger.com', '室内设计软件开发商，开发了Room Arranger软件', 26, NOW(), NOW()),
  (77, 'Vectorworks公司', 'Vectorworks, Inc.', 'https://www.vectorworks.net/', '全球设计和BIM软件开发商，服务于85个国家的AEC、景观和娱乐行业', 26, NOW(), NOW()),
  (78, 'ecru sc', 'ecru sc', 'www.ecru.pl', '室内设计软件开发商，开发了Pro100 7.08软件', 24, NOW(), NOW()),
  (79, '建设专家集团公司', 'Group of Companies "Stroy Expertiza"', 'basegroup.su', 'Base 10 + Фундамент 14 + Плита 6 软件开发商', 21, NOW(), NOW()),
  (80, 'Interior3D', 'Interior3D', 'interior3d.su', '室内设计软件开发商', 21, NOW(), NOW()),
  (81, 'Стройэкспертиза', 'Stroy Ekspertiza', 'https://basegroup.su', 'Фундамент 13.3 软件开发商', 21, NOW(), NOW()),
  (82, 'MicroCrowd', 'MicroCrowd', 'http://www.sweethome3d.com/', 'Sweet Home 3D 2.6便携式软件开发商，支持32/64位系统，多语言界面（含俄语）', 14, NOW(), NOW()),
  (83, 'IMSI Design', 'IMSI Design', 'www.imsidesign.com', '知名的CAD和家居设计软件开发商，主要产品包括TurboCAD等', 26, NOW(), NOW()),
  (84, 'AMS Software', 'AMS Software', 'https://amssoft.ru/interior/', '专业的室内设计软件开发商，提供Interior 3D Design等产品', 21, NOW(), NOW()),
  (85, 'SCADSoft', 'SCADSoft', 'scadsoft.com', 'SCAD Office结构分析与设计集成系统，包含多个工程计算程序', 21, NOW(), NOW()),
  (86, 'Sisoft Development', 'Sisoft Development', 'csdev.ru', 'SPDS图形软件开发商', 21, NOW(), NOW()),
  (87, 'Dlubal软件', 'Dlubal Software', 'http://www.dlubal.ru/rstab-8xx.aspx', '结构分析软件开发商', 8, NOW(), NOW()),
  (88, 'CSoft', 'CSoft', 'www.csoft.ru', '俄罗斯软件开发商', 21, NOW(), NOW()),
  (89, 'Module软件', 'Module-Soft', 'https://www.modul-company.com/', 'HouseCreator+软件开发商', 21, NOW(), NOW()),
  (90, 'Trimble公司', 'Trimble Inc.', 'https://www.trimble.com', '美国GPS技术领先创新者，提供先进的GPS组件以及通过其他定位技术增强的完整客户解决方案', 26, NOW(), NOW()),
  (91, 'Optimal Programs', 'Optimal Programs', 'http://optimalprograms.com/cutting-optimization/', '专业的切割优化软件开发商', 26, NOW(), NOW());
