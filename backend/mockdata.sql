CREATE TABLE IF NOT EXISTS species (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    location VARCHAR(255)
);

INSERT INTO species (name, description, location) VALUES
('Emperor Penguin', 'The largest penguin species', 'Antarctica'),
('King Penguin', 'Second largest species', 'Sub-Antarctic islands')
ON CONFLICT DO NOTHING;
