CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE todos (
    id uuid DEFAULT uuid_generate_v4 (),
    title VARCHAR NOT NULL,
    PRIMARY KEY (id)
);

CREATE INDEX ON todos (id, title);