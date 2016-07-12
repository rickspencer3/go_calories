CREATE TABLE IF NOT EXISTS `foods` (
  `id` INTEGER PRIMARY KEY  NOT NULL,
  `calories` int(11) default NULL,
  `description` text NOT NULL,
  `short_description` varchar(255) NOT NULL default '',
  `user_id` int(11) default NULL,
  `created_at` datetime default NULL,
  `instructions` text,
  `recipe` tinyint(1) default NULL,
  `fat_grams` float default NULL,
  `prot_grams` float default NULL,
  `carbs_grams` float default NULL,
  `alco_grams` float default NULL,
  `sat_grams` float default NULL,
  `trans_grams` float default NULL,
  `chol_mgrams` float default NULL,
  `sod_mgrams` float default NULL,
  `fiber_grams` float default NULL,
  `sugar_grams` float default NULL,
  `fluid` tinyint(1) default '0'
) ;


