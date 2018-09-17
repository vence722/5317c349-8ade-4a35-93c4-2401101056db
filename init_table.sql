CREATE TABLE `order` (
  `id` varchar(128) NOT NULL DEFAULT '',
  `origin_latitude` double NOT NULL DEFAULT '0',
  `origin_longitude` double NOT NULL DEFAULT '0',
  `destination_latitude` double NOT NULL DEFAULT '0',
  `destination_longitude` double NOT NULL DEFAULT '0',
  `distance` double NOT NULL DEFAULT '0',
  `status` varchar(128) NOT NULL DEFAULT '',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;