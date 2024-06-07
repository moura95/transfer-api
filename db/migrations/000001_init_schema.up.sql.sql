SET TIME ZONE 'America/Sao_Paulo';
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS receivers (
                                       uuid          UUID PRIMARY KEY      DEFAULT uuid_generate_v4(),
                                       name VARCHAR(255) NOT NULL,
                                       email VARCHAR(255) NOT NULL UNIQUE,
                                       tax_id varchar(255) NOT NULL UNIQUE,
                                       created_at    TIMESTAMP    NOT NULL DEFAULT NOW(),
                                       update_at    TIMESTAMP    NOT NULL DEFAULT NOW()
);
CREATE INDEX idx_receivers_email ON receivers (email);

CREATE INDEX idx_receivers_tax_id ON receivers (tax_id);

