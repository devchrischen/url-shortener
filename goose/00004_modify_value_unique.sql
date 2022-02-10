-- +goose Up
-- +goose StatementBegin
ALTER TABLE `Url_Shortener`.`Hash`
  ADD UNIQUE KEY (`Value`);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE `Url_Shortener`.`Hash`
  DROP KEY `Value`;
-- +goose StatementEnd
