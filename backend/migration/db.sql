CREATE DATABASE IF NOT EXISTS musicplayer DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE musicplayer;

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

DROP TABLE IF EXISTS categories;
CREATE TABLE categories  (
  id int(11) NOT NULL AUTO_INCREMENT,
  title varchar(255) NOT NULL,
  position int(11) NOT NULL,
  image_url varchar(255) NOT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);


DROP TABLE IF EXISTS roles;
CREATE TABLE roles (
    id        int(11) NOT NULL AUTO_INCREMENT,
    name varchar(255) NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

DROP TABLE IF EXISTS users;
CREATE TABLE users  (
                           id int(11) NOT NULL AUTO_INCREMENT,
                           roles_id int,
                           user_name varchar(255) NOT NULL,
                           pass_word varchar(255) NOT NULL,
                           avatar varchar(255) NOT NULL,
                           age decimal(10,2) NOT NULL,
                           signature varchar(255) NOT NULL,
                           created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                           updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                           PRIMARY KEY (id),
                           FOREIGN KEY (roles_id) REFERENCES roles(id)
);

DROP TABLE IF EXISTS sheets;
CREATE TABLE sheets  (
                             id int(11) NOT NULL AUTO_INCREMENT,
                             title varchar(255) NOT NULL,
                             position int(11) NOT NULL,
                             image_url varchar(255) NOT NULL,
                             created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                             updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                             PRIMARY KEY (id)
);


DROP TABLE IF EXISTS products;
CREATE TABLE products  (
                           id int(11) NOT NULL AUTO_INCREMENT,
                           category_id int,
                           title varchar(255) NOT NULL,
                           image_url varchar(255) NOT NULL,
                           price decimal(10,2) NOT NULL,
                           description text NOT NULL,
                           created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                           updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                           PRIMARY KEY (id),
                           FOREIGN KEY (category_id) REFERENCES categories(id)
);

DROP TABLE IF EXISTS songs;
CREATE TABLE songs  (
                           id int(11) NOT NULL AUTO_INCREMENT,
                           sheet_id int,
                           name varchar(255) NOT NULL,
                           singer varchar(255) NOT NULL,
                           image_url varchar(255) NOT NULL,
                           description text NOT NULL,
                           created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                           updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                           PRIMARY KEY (id),
                           FOREIGN KEY (sheet_id) REFERENCES sheets(id)
);

SET FOREIGN_KEY_CHECKS = 1;