DROP TABLE IF EXISTS gym_controll_users;

CREATE DATABASE gym_controll_db;

\c gym_controll_db;

-- Cria o enum para o tipo de conta
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'account_type_enum') THEN
        CREATE TYPE account_type_enum AS ENUM ('default_user', 'personal', 'nutritionist', 'coach', 'gym_manager');
    END IF;
END$$;

CREATE OR REPLACE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TABLE gym_controll_users (
    id VARCHAR(255) PRIMARY KEY, 
    name VARCHAR(255) NOT NULL,
    password VARCHAR(255),
    email VARCHAR(255) UNIQUE, 
    birthday TIMESTAMPTZ,
    gender VARCHAR(55),
    account_type account_type_enum,
    activate INT DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE gym_controll_exercices (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    muscle VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    equipament VARCHAR(255) NOT NULL,
    video VARCHAR(255),
    activate INT DEFAULT 0,
    created_user VARCHAR(255),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_created_user FOREIGN KEY (created_user) REFERENCES gym_controll_users(id)
);

CREATE TABLE relationate_personal_custumer (
    id VARCHAR(255) PRIMARY KEY,
    personal_email VARCHAR(255) NOT NULL,
    custumer_email VARCHAR(255) NOT NULL,
    activate INT DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_personal_email FOREIGN KEY (personal_email) REFERENCES gym_controll_users(email),
    CONSTRAINT fk_custumer_email FOREIGN KEY (custumer_email) REFERENCES gym_controll_users(email)
);