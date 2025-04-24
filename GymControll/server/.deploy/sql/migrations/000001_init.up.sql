-- Criação de um usuário administrador
CREATE USER admin WITH PASSWORD 'adminpassword';

-- Concede permissões de superusuário ao usuário criado
ALTER USER admin WITH SUPERUSER;