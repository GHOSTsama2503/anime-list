CREATE TABLE IF NOT EXISTS users (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  username TEXT UNIQUE NOT NULL,
  password TEXT NOT NULL,
  is_admin INTEGER NOT NULL
);


CREATE TABLE IF NOT EXISTS animes (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  id_al INTEGER NOT NULL,
  title_romaji TEXT UNIQUE NOT NULL,
  title_native TEXT,
  title_english TEXT,
  format TEXT NOT NULL,
  status TEXT NOT NULL,
  description TEXT NOT NULL,
  start_date TEXT NOT NULL,
  end_date TEXT NOT NULL,
  season TEXT NOT NULL,
  season_year INT,
  episodes INTEGER NOT NULL,
  duration INTEGER NOT NULL,
  banner_image TEXT,
  st_image TEXT NOT NULL,
  group_position INTEGER
);


CREATE TABLE IF NOT EXISTS genres (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT UNIQUE NOT NULL
);


CREATE TABLE IF NOT EXISTS anime_genres (
  anime_id INTEGER NOT NULL,
  genre_id INTEGER NOT NULL,

  PRIMARY KEY (anime_id, genre_id),

  CONSTRAINT anime_genres_anime_id FOREIGN KEY (anime_id) REFERENCES animes(id) ON DELETE CASCADE,
  CONSTRAINT anime_genres_genre_id FOREIGN KEY (genre_id) REFERENCES genres(id) ON DELETE CASCADE
);


CREATE TABLE IF NOT EXISTS studios (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT UNIQUE NOT NULL
);


CREATE TABLE IF NOT EXISTS anime_studios (
  anime_id INTEGER NOT NULL,
  studio_id INTEGER NOT NULL,

  PRIMARY KEY (anime_id, studio_id),

  CONSTRAINT anime_studios_anime_id FOREIGN KEY (anime_id) REFERENCES animes(id) ON DELETE CASCADE,
  CONSTRAINT anime_studios_studio_id FOREIGN KEY (studio_id) REFERENCES studios(id) ON DELETE CASCADE
);


CREATE TABLE IF NOT EXISTS synonyms (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  anime_id INTEGER NOT NULL,
  name TEXT NOT NULL,

  CONSTRAINT fk_synonyms_anime_id FOREIGN KEY (anime_id) REFERENCES animes(id) ON DELETE CASCADE
);


CREATE TABLE IF NOT EXISTS cover_images (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  anime_id INTEGER NOT NULL,
  extra_large TEXT,
  large TEXT,
  medium TEXT,
  color TEXT,

  CONSTRAINT fk_cover_image_anime_id FOREIGN KEY (anime_id) REFERENCES animes(id) ON DELETE CASCADE
);


CREATE TABLE IF NOT EXISTS groups (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS anime_groups (
  anime_id INTEGER NOT NULL,
  group_id INTEGER NOT NULL,

  PRIMARY KEY (anime_id, group_id),

  CONSTRAINT fk_anime_groups_anime_id FOREIGN KEY (anime_id) REFERENCES animes(id) ON DELETE CASCADE,
  CONSTRAINT fk_anime_groups_group_id FOREIGN KEY (group_id) REFERENCES groups(id) ON DELETE CASCADE
);
