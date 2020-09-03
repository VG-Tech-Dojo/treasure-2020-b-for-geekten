CREATE TABLE `zoom` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `owner_id` int,
  `access_token` varchar(1024) NOT NULL,
  `refresh_token` varchar(1024) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;