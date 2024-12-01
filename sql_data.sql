-- Вставка данных в таблицу songs с возвратом ID песен
WITH inserted_songs AS (
    INSERT INTO songs (group_name, song, release_date, link)
    VALUES 
        ('Imagine Dragons', 'Believer', '2017-02-01', 'https://example.com/believer'),
        ('Queen', 'Bohemian Rhapsody', '1975-10-31', 'https://example.com/bohemian_rhapsody'),
        ('Coldplay', 'Yellow', '2000-06-26', 'https://example.com/yellow'),
        ('The Beatles', 'Hey Jude', '1968-08-26', 'https://example.com/hey_jude'),
        ('Linkin Park', 'Numb', '2003-03-25', 'https://example.com/numb'),
        ('Pink Floyd', 'Comfortably Numb', '1979-11-30', 'https://example.com/comfortably_numb'),
        ('Led Zeppelin', 'Stairway to Heaven', '1971-11-08', 'https://example.com/stairway_to_heaven'),
        ('Eminem', 'Lose Yourself', '2002-10-28', 'https://example.com/lose_yourself'),
        ('Adele', 'Someone Like You', '2011-01-24', 'https://example.com/someone_like_you'),
        ('Taylor Swift', 'Love Story', '2008-09-15', 'https://example.com/love_story'),
        ('Michael Jackson', 'Thriller', '1982-11-30', 'https://example.com/thriller'),
        ('U2', 'With or Without You', '1987-03-16', 'https://example.com/with_or_without_you'),
        ('Nirvana', 'Smells Like Teen Spirit', '1991-09-10', 'https://example.com/smells_like_teen_spirit'),
        ('Elton John', 'Your Song', '1970-10-26', 'https://example.com/your_song'),
        ('Beyoncé', 'Halo', '2008-01-20', 'https://example.com/halo'),
        ('Ed Sheeran', 'Shape of You', '2017-01-06', 'https://example.com/shape_of_you'),
        ('The Rolling Stones', 'Paint It Black', '1966-05-06', 'https://example.com/paint_it_black'),
        ('AC/DC', 'Back In Black', '1980-07-25', 'https://example.com/back_in_black'),
        ('Katy Perry', 'Firework', '2010-10-26', 'https://example.com/firework'),
        ('Billie Eilish', 'Bad Guy', '2019-03-29', 'https://example.com/bad_guy')
    RETURNING id_song, song
)
-- Привязка куплетов к песням
INSERT INTO verses (text_verses, song_id, verses_num)
VALUES 
    -- Куплеты для Believer
    ('First things first, I''ma say all the words inside my head', (SELECT id_song FROM inserted_songs WHERE song = 'Believer'), 1),
    ('Second things second, don''t you tell me what you think that I could be', (SELECT id_song FROM inserted_songs WHERE song = 'Believer'), 2),
    -- Куплеты для Bohemian Rhapsody
    ('Is this the real life? Is this just fantasy?', (SELECT id_song FROM inserted_songs WHERE song = 'Bohemian Rhapsody'), 1),
    ('Caught in a landslide, no escape from reality', (SELECT id_song FROM inserted_songs WHERE song = 'Bohemian Rhapsody'), 2),
    -- Куплеты для Yellow
    ('Look at the stars, look how they shine for you', (SELECT id_song FROM inserted_songs WHERE song = 'Yellow'), 1),
    ('And everything you do, yeah, they were all yellow', (SELECT id_song FROM inserted_songs WHERE song = 'Yellow'), 2),
    -- Куплеты для Hey Jude
    ('Hey Jude, don''t make it bad', (SELECT id_song FROM inserted_songs WHERE song = 'Hey Jude'), 1),
    ('Take a sad song and make it better', (SELECT id_song FROM inserted_songs WHERE song = 'Hey Jude'), 2),
    -- Куплеты для Numb
    ('I''ve become so numb, I can''t feel you there', (SELECT id_song FROM inserted_songs WHERE song = 'Numb'), 1),
    ('I''m tired of being what you want me to be', (SELECT id_song FROM inserted_songs WHERE song = 'Numb'), 2),
    -- Куплеты для Comfortably Numb
    ('Hello, is there anybody in there?', (SELECT id_song FROM inserted_songs WHERE song = 'Comfortably Numb'), 1),
    ('Just nod if you can hear me, is there anyone home?', (SELECT id_song FROM inserted_songs WHERE song = 'Comfortably Numb'), 2),
    -- Куплеты для Stairway to Heaven
    ('There''s a lady who''s sure all that glitters is gold', (SELECT id_song FROM inserted_songs WHERE song = 'Stairway to Heaven'), 1),
    ('And she''s buying a stairway to heaven', (SELECT id_song FROM inserted_songs WHERE song = 'Stairway to Heaven'), 2),
    -- Куплеты для Lose Yourself
    ('Look, if you had one shot, or one opportunity', (SELECT id_song FROM inserted_songs WHERE song = 'Lose Yourself'), 1),
    ('To seize everything you ever wanted, in one moment', (SELECT id_song FROM inserted_songs WHERE song = 'Lose Yourself'), 2),
    -- Куплеты для Someone Like You
    ('I heard that you''re settled down', (SELECT id_song FROM inserted_songs WHERE song = 'Someone Like You'), 1),
    ('That you found a girl and you''re married now', (SELECT id_song FROM inserted_songs WHERE song = 'Someone Like You'), 2),
    -- Куплеты для Love Story
    ('We were both young when I first saw you', (SELECT id_song FROM inserted_songs WHERE song = 'Love Story'), 1),
    ('I close my eyes and the flashback starts', (SELECT id_song FROM inserted_songs WHERE song = 'Love Story'), 2),
    -- Куплеты для Thriller
    ('It''s close to midnight, and something evil''s lurking in the dark', (SELECT id_song FROM inserted_songs WHERE song = 'Thriller'), 1),
    ('You try to scream, but terror takes the sound before you make it', (SELECT id_song FROM inserted_songs WHERE song = 'Thriller'), 2),
    -- Куплеты для With or Without You
    ('See the stone set in your eyes', (SELECT id_song FROM inserted_songs WHERE song = 'With or Without You'), 1),
    ('I can''t live with or without you', (SELECT id_song FROM inserted_songs WHERE song = 'With or Without You'), 2),
    -- Куплеты для Smells Like Teen Spirit
    ('Load up on guns, bring your friends', (SELECT id_song FROM inserted_songs WHERE song = 'Smells Like Teen Spirit'), 1),
    ('It''s fun to lose and to pretend', (SELECT id_song FROM inserted_songs WHERE song = 'Smells Like Teen Spirit'), 2),
    -- Куплеты для Your Song
    ('It''s a little bit funny, this feeling inside', (SELECT id_song FROM inserted_songs WHERE song = 'Your Song'), 1),
    ('I hope you don''t mind that I put down in words', (SELECT id_song FROM inserted_songs WHERE song = 'Your Song'), 2),
    -- Куплеты для Halo
    ('Remember those walls I built?', (SELECT id_song FROM inserted_songs WHERE song = 'Halo'), 1),
    ('Everywhere I''m looking now, I''m surrounded by your embrace', (SELECT id_song FROM inserted_songs WHERE song = 'Halo'), 2),
    -- Куплеты для Shape of You
    ('The club isn''t the best place to find a lover', (SELECT id_song FROM inserted_songs WHERE song = 'Shape of You'), 1),
    ('I''m in love with the shape of you', (SELECT id_song FROM inserted_songs WHERE song = 'Shape of You'), 2),
    -- Куплеты для Paint It Black
    ('I see a red door and I want it painted black', (SELECT id_song FROM inserted_songs WHERE song = 'Paint It Black'), 1),
    ('No colors anymore, I want them to turn black', (SELECT id_song FROM inserted_songs WHERE song = 'Paint It Black'), 2),
    -- Куплеты для Back In Black
    ('Back in black, I hit the sack', (SELECT id_song FROM inserted_songs WHERE song = 'Back In Black'), 1),
    ('Yes, I''m let loose from the noose', (SELECT id_song FROM inserted_songs WHERE song = 'Back In Black'), 2)
;
