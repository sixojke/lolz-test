CREATE TABLE genre(
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE book(
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    author VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    genre_id UUID REFERENCES genre(id)
);
