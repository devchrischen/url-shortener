-- +goose Up
-- +goose StatementBegin
-- CREATE DATABASE IF NOT EXISTS `Shortener` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci ;
-- USE `Shortener`;

CREATE TABLE IF NOT EXISTS `Url_Shortener`.`Hash` (
  `ID` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `Value` VARCHAR(16) NOT NULL COMMENT 'hash value for short URL',
  `CreatedAt` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `UpdatedAt` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`ID`))
ENGINE = InnoDB;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `Url_Shortener`.`Hash`;
-- +goose StatementEnd
