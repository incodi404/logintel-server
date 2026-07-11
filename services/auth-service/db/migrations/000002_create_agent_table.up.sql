CREATE TABLE agents (
    id INT PRIMARY KEY,
    agent_name VARCHAR(100) NOT NULL,
    agent_ip VARCHAR(100) NOT NULL UNIQUE,
    agent_id VARCHAR(100) NOT NULL UNIQUE,
    description TEXT
)