SET TIME ZONE 'America/Sao_Paulo';
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE receivers (
     uuid UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
     name VARCHAR(255) NOT NULL,
     cpf_cnpj VARCHAR(14) UNIQUE NOT NULL,
     pix_key_type VARCHAR(20) NOT NULL CHECK (pix_key_type IN ('CPF', 'CNPJ', 'EMAIL', 'TELEFONE', 'CHAVE_ALEATORIA')),
     pix_key VARCHAR(140) NOT NULL,
     email VARCHAR(250) UNIQUE NOT NULL,
     status VARCHAR(20) NOT NULL CHECK (status IN ('Validado', 'Rascunho')),
     created_at    TIMESTAMP    NOT NULL DEFAULT NOW(),
     update_at    TIMESTAMP    NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_status ON receivers(status);
CREATE INDEX idx_pix_key_type ON receivers(pix_key_type);
CREATE INDEX idx_pix_key ON receivers(pix_key);
