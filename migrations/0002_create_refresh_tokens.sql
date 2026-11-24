-- +goose Up
CREATE TABLE refresh_token (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    chain_id UUID NOT NULL DEFAULT gen_random_uuid(),
    account_id BIGINT NOT NULL REFERENCES account (id),
    selector VARCHAR NOT NULL UNIQUE,
    verifier_hash VARCHAR NOT NULL,
    issued_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    revoked_at TIMESTAMPTZ NULL,
    expires_at TIMESTAMPTZ NOT NULL
);

CREATE INDEX token_chain_idx ON refresh_token(chain_id);

-- +goose Down
DROP TABLE IF EXISTS refresh_token;