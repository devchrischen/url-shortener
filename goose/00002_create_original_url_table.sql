-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `Url_Shortener`.`OriginalUrl` (
  `ID` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `HashID` INT UNSIGNED NOT NULL,
  `Url` VARCHAR(256) NOT NULL COMMENT 'original long URL',
  `CreatedAt` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `UpdatedAt` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`ID`),
  CONSTRAINT `FK_OriginalUrl_HashID`
  FOREIGN KEY (`HashID`)
  REFERENCES `Url_Shortener`.`Hash` (`ID`)
  ON DELETE NO ACTION
  ON UPDATE NO ACTION)
ENGINE = InnoDB;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `Url_Shortener`.`OriginalUrl`;
-- +goose StatementEnd
