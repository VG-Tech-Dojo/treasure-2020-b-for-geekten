CREATE TABLE `lesson` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `owner_id` int,
  `start_time` datetime NOT NULL,
  `end_time` datetime NOT NULL,
  `meeting_id` varchar(255) NOT NULL,
  `price` int NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
