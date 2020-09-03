CREATE TABLE `reservation` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `lesson_id` int,
  `charge_id` varchar(255) NOT NULL,
  `paid_price` int NOT NULL,
  `first_name` varchar(255) NOT NULL,
  `last_name` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
