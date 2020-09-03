CREATE TABLE `website` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `owner_id` int,
  `title` varchar(255),
  `profile` varchar(255),
  `theme` varchar(255),
  `content` text,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;