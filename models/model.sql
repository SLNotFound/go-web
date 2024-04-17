CREATE TABLE `hrms_department` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `pid` int(10) unsigned NOT NULL,
  `name` varchar(100) NOT NULL COMMENT '部门名称',
  `code` varchar(100) NOT NULL COMMENT '部门code',
  `managerId` int(10) NOT NULL COMMENT '主管id',
  `managerName` varchar(100) NOT NULL COMMENT '主管名称',
  `introduce` varchar(100) NOT NULL COMMENT '介绍',
  `createTime` varchar(100) NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='部门管理';

CREATE TABLE `hrms_auth` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(10) NOT NULL COMMENT '员工名字',
  `mobile` varchar(11) NOT NULL COMMENT '员工手机号',
  `formOfEmployment` int(1) NOT NULL COMMENT '员工聘用形式',
  `workNumber` varchar(10) NOT NULL COMMENT '员工工号',
  `departmentId` int(10) NOT NULL COMMENT '员工部门id',
  `timeOfEntry` varchar(10) NOT NULL COMMENT '员工入职日期',
  `correctionTime` varchar(10) NOT NULL COMMENT '员工转正日期',
  `staffPhoto` varchar(10) COMMENT '头像信息',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='员工表';
