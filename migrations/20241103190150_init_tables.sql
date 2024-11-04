-- +goose Up
-- +goose StatementBegin
CREATE TABLE operators
(
    id       serial,
    login    VARCHAR(50)  NOT NULL UNIQUE,
    password VARCHAR(256) NOT NULL,
    name     VARCHAR(100)
);

CREATE TABLE dialogs
(
    id         serial,
    creator_by BIGINT      not null,
    created_at BIGINT      not null,
    status     VARCHAR(64) NOT NULL
);

CREATE TABLE messages
(
    id         serial,
    dialog_id  BIGINT REFERENCES dialogs (id),
    created_at BIGINT      not null,
    created_by BIGINT      not null,
    text       TEXT        NOT NULL,
    type       VARCHAR(64) NOT NULL
);

CREATE TABLE user_dialogs
(
    id         serial,
    support_id BIGINT REFERENCES operators(id),
    dialog_id  BIGINT REFERENCES dialogs (id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
