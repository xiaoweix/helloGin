CREATE TABLE `t_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID ',
  `username` varchar(64) NOT NULL COMMENT '用户名称 ',
  `age` int(11) NOT NULL COMMENT 'ID ',
  `mobile` varchar(64) NOT NULL COMMENT '手机号码 ',
  `sex` varchar(32) DEFAULT NULL COMMENT '性别 ',
  `address` varchar(64) DEFAULT NULL COMMENT '地址 ',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户表';