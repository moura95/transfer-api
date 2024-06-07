SET TIME ZONE 'America/Sao_Paulo';
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE receivers (
     uuid UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
     nome VARCHAR(255) NOT NULL,
     cpfcnpj VARCHAR(14) NOT NULL,
     pix_key_type VARCHAR(20) NOT NULL CHECK (pix_key_type IN ('CPF', 'CNPJ', 'EMAIL', 'TELEFONE', 'CHAVE_ALEATORIA')),
     pix_key VARCHAR(140) NOT NULL,
     email VARCHAR(250) NULL,
     status VARCHAR(20) NOT NULL CHECK (status IN ('Validado', 'Rascunho'))
);

CREATE INDEX idx_status ON receivers(status);
CREATE INDEX idx_pix_key_type ON receivers(pix_key_type);
CREATE INDEX idx_pix_key ON receivers(pix_key);
