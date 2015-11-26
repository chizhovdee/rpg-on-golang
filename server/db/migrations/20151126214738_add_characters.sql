-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE `characters` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `basic_money` bigint(20) DEFAULT NULL,
  `vip_money` int(11) DEFAULT NULL,
  `level` int(11) DEFAULT '1',
  `experience` bigint(20) DEFAULT '0',
  `points` int(11) DEFAULT NULL,
  `health` int(11) DEFAULT NULL,
  `energy` int(11) DEFAULT NULL,
  `hp` int(11) DEFAULT '100',
  `ep` int(11) DEFAULT '10',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `index_characters_on_level` (`level`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back

drop table characters;