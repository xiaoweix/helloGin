# helloGin
###运行开始前请修改数据库配置
###1、Mysql 数据库新建一个表
```
CREATE TABLE `t_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID ',
  `username` varchar(64) NOT NULL COMMENT '用户名称 ',
  `age` int(11) NOT NULL COMMENT 'ID ',
  `mobile` varchar(64) NOT NULL COMMENT '手机号码 ',
  `sex` varchar(32) DEFAULT NULL COMMENT '性别 ',
  `address` varchar(64) DEFAULT NULL COMMENT '地址 ',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户表';
```
###2、去database下修改mysql的连接配置
修改const的相关信息

###3、拉去需要的包

###4、运行即可
