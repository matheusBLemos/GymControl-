CREATE TYPE gender_enum AS ENUM ('M', 'F');
CREATE TYPE account_type_enum AS ENUM ('ALUNO', 'PERSONAL', 'NUTRICIONISTA', 'ACADEMIA');
CREATE TYPE status_enum AS ENUM ('0', '1');

CREATE TABLE users (
    id UUID PRIMARY KEY,
    name VARCHAR(255),
    password VARCHAR(255),
    email VARCHAR(255),
    birthday TIMESTAMP,
    gender gender_enum,
    account_type account_type_enum,
    status status_enum,
    created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated TIMESTAMP
);

CREATE TABLE exercice (
    id UUID PRIMARY KEY,
    name VARCHAR(255),
    description VARCHAR(255),
    body_part VARCHAR(255),
    fk_created_user UUID,  
    created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    status status_enum,
    FOREIGN KEY (fk_created_user) REFERENCES users(id)
);

CREATE TABLE training (
    id UUID PRIMARY KEY,
    fk_user UUID,
    fk_create_user UUID,
    muscle VARCHAR(255),
    description TEXT,
    status status_enum,
    created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (fk_user) REFERENCES users(id),
    FOREIGN KEY (fk_create_user) REFERENCES users(id)
);

CREATE TABLE exercice_training (
    id UUID PRIMARY KEY,
    fk_exercice UUID,
    fk_training UUID,
    sets INT,
    reps INT,
    durations_seconds INT,
    total_durations_seconds INT,
    status status_enum,
    created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE user_training (
    id UUID PRIMARY KEY,
    fk_user UUID,
    fk_training UUID,
    fk_create_user UUID,
    status status_enum,
    created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (fk_user) REFERENCES users(id),
    FOREIGN KEY (fk_training) REFERENCES training(id)
);

CREATE OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_users_updated_at
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();

CREATE TRIGGER update_exercice_updated_at
BEFORE UPDATE ON exercice
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();

CREATE TRIGGER update_training_updated_at
BEFORE UPDATE ON training
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();

CREATE TRIGGER update_exercice_training_updated_at
BEFORE UPDATE ON exercice_training
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();

CREATE TRIGGER update_user_training_updated_at
BEFORE UPDATE ON user_training
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();
