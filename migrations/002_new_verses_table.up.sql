CREATE TABLE IF NOT EXISTS verses (
    id_verses SERIAL PRIMARY KEY, 
    text_verses TEXT NOT NULL,                        
    song_id INT NOT NULL,                      
    verses_num INT,                            
    FOREIGN KEY (song_id) REFERENCES songs(id_song) ON DELETE CASCADE
);

-- DROP TABLE IF EXISTS verses;
