DROP TABLE IF EXISTS `shop`;
CREATE TABLE `shop` (
  `shop_id` int NOT NULL AUTO_INCREMENT,
  `account` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `status` int DEFAULT 0,
  PRIMARY KEY (`shop_id`)
);

DROP TABLE IF EXISTS `shop_detail`;
CREATE TABLE `shop_detail` (
  `shop_id` int NOT NULL AUTO_INCREMENT COMMENT 'Primary Key',
  `shop_name` varchar(50) DEFAULT NULL,
  `shop_info` varchar(50) DEFAULT NULL,
  `shop_image` varchar(250) DEFAULT NULL,
  `corporation_name` varchar(50) DEFAULT NULL,
  `shop_location` varchar(100) DEFAULT NULL,
  `open_time` varchar(50) DEFAULT NULL,
  `dayoff` varchar(50) DEFAULT NULL,
  `phonenumber` varchar(50) DEFAULT NULL,
  `email` varchar(50) DEFAULT NULL,
  `shop_city` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  PRIMARY KEY (`shop_id`),
  CONSTRAINT `shop_detail_ibfk_1` FOREIGN KEY (`shop_id`) REFERENCES `shop` (`shop_id`)
)

DROP TABLE IF EXISTS `car`;
CREATE TABLE `car` (
  `car_id` int NOT NULL AUTO_INCREMENT COMMENT 'Primary Key',
  `shop_id` int NOT NULL,
  `car_name` varchar(255) DEFAULT NULL,
  `car_brand` varchar(255) DEFAULT NULL,
  `car_image` varchar(255) DEFAULT NULL,
  `car_price` int DEFAULT '0',
  `car_fee` int DEFAULT '0',
  `shelves` tinyint(1) DEFAULT '0',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `car_year` int DEFAULT NULL,
  PRIMARY KEY (`car_id`)
)

