CREATE TABLE IF NOT EXISTS songs (
    id_song SERIAL PRIMARY KEY, 
    group_name VARCHAR(255) NOT NULL,          
    song VARCHAR(255) NOT NULL,             
    release_date VARCHAR(255),                      
    link TEXT                               
);

-- DROP TABLE IF EXISTS songs;
