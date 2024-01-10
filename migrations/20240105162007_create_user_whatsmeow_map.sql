-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_whatsmeow_map (
    user_id VARCHAR(255) NOT NULL,
    jid VARCHAR(255) NOT NULL,
    phone VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, jid)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table user_whatsmeow_map
-- +goose StatementEnd
