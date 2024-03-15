CREATE TABLE tasks (
       id INTEGER,
       user_id INTEGER,
       title VARCHAR(255),
       description TEXT DEFAULT  NULL,
       status VARCHAR(50) DEFAULT 'pending',
       created_at TIMESTAMP WITH TIME ZONE NOT NULL,
       updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
       PRIMARY KEY (id)
);
