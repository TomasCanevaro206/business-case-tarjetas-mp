CREATE TABLE `cards` (
  `card_id` int NOT NULL AUTO_INCREMENT,
  `card_number` int NOT NULL,
  `card_type` varchar(255) NOT NULL,
  `expiration_date` varchar(255) NOT NULL,
  `card_state` varchar(255) NOT NULL,
  `timestamp_creation` varchar(255) NOT NULL,
  `timestamp_modificaction` varchar(255) NOT NULL,
  PRIMARY KEY (`card_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;