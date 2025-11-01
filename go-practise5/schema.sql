CREATE TABLE jobs (
                      id SERIAL PRIMARY KEY,
                      title TEXT NOT NULL,
                      company TEXT NOT NULL,
                      salary INTEGER NOT NULL,
                      created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

INSERT INTO jobs (title, company, salary) VALUES
                                              ('Go Developer', 'Kolesa', 615000),
                                              ('Frontend Developer', 'Kolesa', 550000),
                                              ('Backend Developer', 'Chocofamily', 600000),
                                              ('DevOps Engineer', 'One Tech', 700000);