CREATE TABLE meows (
  id SERIAL PRIMARY KEY,
  body TEXT NOT NULL,
  created_on TIMESTAMP WITH TIME ZONE NOT NULL
);