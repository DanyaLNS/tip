CREATE TABLE artists (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    genre VARCHAR(255) NOT NULL,
    country VARCHAR(255) NOT NULL
);

CREATE TABLE albums (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    artist_id INT REFERENCES artists(id) ON DELETE CASCADE,
    year INT NOT NULL,
    genre VARCHAR(255) NOT NULL
);

CREATE TABLE songs (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    artist_id INT REFERENCES artists(id) ON DELETE CASCADE,
    album_id INT REFERENCES albums(id) ON DELETE CASCADE,
    genre VARCHAR(255) NOT NULL,
    duration VARCHAR(10) NOT NULL
);


