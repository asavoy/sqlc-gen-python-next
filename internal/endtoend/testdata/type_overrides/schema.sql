CREATE TABLE articles (
    id         BIGSERIAL PRIMARY KEY,
    metadata   JSONB NOT NULL,
    settings   JSONB,
    author_id  UUID NOT NULL
);
