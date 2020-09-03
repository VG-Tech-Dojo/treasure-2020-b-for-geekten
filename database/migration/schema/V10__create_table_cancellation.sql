CREATE TABLE `cancellation` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `reservation_id` int,
  `repaid_price` int NOT NULL,
  `canceled_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;