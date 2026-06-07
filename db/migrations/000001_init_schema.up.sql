CREATE TABLE IF NOT EXISTS traits (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    category VARCHAR(100),
    generated_mood VARCHAR(100)
);

CREATE TABLE IF NOT EXISTS careers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    branch VARCHAR(100),
    base_salary INT,
    ideal_mood VARCHAR(100)
);

CREATE TABLE IF NOT EXISTS career_recommendations (
    career_id INT REFERENCES careers(id) ON DELETE CASCADE,
    trait_id INT REFERENCES traits(id) ON DELETE CASCADE,
    compatibility_score INT CHECK (compatibility_score BETWEEN 1 AND 5),
    reason TEXT,
    PRIMARY KEY (career_id, trait_id)
);