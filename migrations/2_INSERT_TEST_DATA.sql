INSERT INTO artists (name, genre, country)
VALUES
    ('Radiohead', 'Alternative Rock', 'UK');

INSERT INTO albums (title, artist_id, year, genre)
VALUES
    ('OK Computer', 1, 1997, 'Alternative Rock'),
    ('In Rainbows', 1, 2007, 'Alternative Rock');

INSERT INTO songs (title, artist_id, album_id, genre, duration)
VALUES
    ('Karma Police', 1, 1, 'Alternative Rock', '4:21'),
    ('No Surprises', 1, 1, 'Alternative Rock', '3:48');