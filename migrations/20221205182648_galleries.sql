-- +goose Up
CREATE TABLE Galleries(
  id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  name text NOT NULL,
  description text,
  imageLocation text
);

-- +goose Down
DROP TABLE Galleries
