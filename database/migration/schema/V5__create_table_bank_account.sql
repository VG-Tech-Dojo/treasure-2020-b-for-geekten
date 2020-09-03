CREATE TABLE `bank_account` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `owner_id` int,
  `bank_account_number` varchar(255) NOT NULL,
  `bank_branch_code` varchar(255) NOT NULL,
  `bank_code` varchar(255) NOT NULL,
  `bank_account_holder_name` varchar(255) NOT NULL,
  `bank_account_type` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;;