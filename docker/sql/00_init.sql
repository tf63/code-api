CREATE TABLE IF NOT EXISTS todos (
    id SERIAL PRIMARY KEY,
    title varchar(50) NOT NULL,
    done BOOLEAN DEFAULT false
);