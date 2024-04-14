CREATE TABLE `baby.user` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `uuid` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT 'UUID',
  `display_name` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '表示名',
  `created_at` datetime(6) NOT NULL COMMENT '作成日時',
  `updated_at` datetime(6) NOT NULL COMMENT '更新日時',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='ユーザー'
