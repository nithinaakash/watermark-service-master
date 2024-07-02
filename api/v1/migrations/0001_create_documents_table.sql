CREATE TABLE documents (
    id SERIAL PRIMARY KEY,
    content VARCHAR(255) NOT NULL,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    topic VARCHAR(255) NOT NULL,
    watermark VARCHAR(255)
);


psql -U postgres -d postgres -f migrations/0001_create_documents_table.sql
