-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';                                            
CREATE TABLE users (                                              
    id SERIAL PRIMARY KEY,                                        
    email VARCHAR(255) NOT NULL UNIQUE,                           
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP 
);                                                                
                                                                  
CREATE INDEX idx_users_email ON users(email);                     
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
