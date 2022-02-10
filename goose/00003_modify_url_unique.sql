-- +goose Up
-- +goose StatementBegin
ALTER TABLE `Url_Shortener`.`OriginalUrl`
  ADD UNIQUE KEY (`Url`);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE `Url_Shortener`.`OriginalUrl`
  DROP KEY `Url`;
-- +goose StatementEnd
