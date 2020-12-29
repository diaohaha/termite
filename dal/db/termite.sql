DROP TABLE IF EXISTS termite_task_config;
CREATE TABLE `termite_task_config` (
  `id` bigint(24) NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `create_time` TIMESTAMP not null DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` TIMESTAMP not null DEFAULT CURRENT_TIMESTAMP on UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `work_key` varchar(32) NOT NULL DEFAULT '' COMMENT '任务ID',
  `work_name` varchar(64) NOT NULL DEFAULT '' COMMENT '任务名称',
  `work_desc` varchar(2048) NOT NULL DEFAULT '' COMMENT '任务描述',
  `work_config` varchar(1024) NOT NULL DEFAULT '' COMMENT '任务配置',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_work_key` (`work_key`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='Termite工作配置表';

DROP TABLE IF EXISTS termite_flow_config;
CREATE TABLE `termite_flow_config` (
  `id` bigint(24) NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `create_time` TIMESTAMP not null DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` TIMESTAMP not null DEFAULT CURRENT_TIMESTAMP on UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `flow_key` varchar(32) NOT NULL DEFAULT '' COMMENT '任务流ID',
  `flow_name` varchar(64) NOT NULL DEFAULT '' COMMENT '任务流名称',
  `flow_desc` varchar(2048) NOT NULL DEFAULT '' COMMENT '任务流描述',
  `flow_config` varchar(1024) NOT NULL DEFAULT '' COMMENT '任务流配置',
  `env` varchar(1024) NOT NULL default '' COMMENT '环境变量',
   PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_flow_key` (`flow_key`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='Termite工作流配置表';