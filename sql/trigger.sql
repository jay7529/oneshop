DROP TRIGGER IF EXISTS `after_shop_insert`;
CREATE TRIGGER `after_shop_insert` AFTER INSERT ON `shop` FOR EACH ROW BEGIN
INSERT INTO shop_detail (shop_id) VALUES (new.shop_id);
END