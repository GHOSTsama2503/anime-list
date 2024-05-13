CREATE TABLE IF NOT EXISTS animes (
  id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  id_al INT NOT NULL,
  title_romaji VARCHAR(255) NOT NULL,
  title_native VARCHAR(255),
  title_english VARCHAR(255),
  format VARCHAR(10) NOT NULL,
  status VARCHAR(10) NOT NULL,
  description VARCHAR(2000) NOT NULL,
  start_date DATE NOT NULL,
  end_date DATE NOT NULL,
  season VARCHAR(6) NOT NULL,
  season_year SMALLINT,
  episodes SMALLINT NOT NULL,
  duration SMALLINT NOT NULL,
  banner_image VARCHAR(255),
  st_image VARCHAR(255) NOT NULL
);


CREATE TABLE IF NOT EXISTS genres (
  id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  name VARCHAR(32) UNIQUE NOT NULL
);


CREATE TABLE IF NOT EXISTS anime_genres (
  anime_id INT,
  genre_id INT,

  PRIMARY KEY (anime_id, genre_id),

  CONSTRAINT anime_genres_anime_id FOREIGN KEY (anime_id) REFERENCES animes(id),
  CONSTRAINT anime_genres_genre_id FOREIGN KEY (genre_id) REFERENCES genres(id)
);


CREATE TABLE IF NOT EXISTS studios (
  id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  name VARCHAR(32) UNIQUE NOT NULL
);


CREATE TABLE IF NOT EXISTS anime_studios (
  anime_id INT,
  studio_id INT,

  PRIMARY KEY (anime_id, studio_id),

  CONSTRAINT anime_studios_anime_id FOREIGN KEY (anime_id) REFERENCES animes(id),
  CONSTRAINT anime_studios_studio_id FOREIGN KEY (studio_id) REFERENCES studios(id)
);


CREATE TABLE IF NOT EXISTS synonyms (
  id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  anime_id INT NOT NULL,
  name VARCHAR(255) NOT NULL,

  CONSTRAINT fk_synonyms_anime_id FOREIGN KEY (anime_id) REFERENCES animes(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS cover_images (
  id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  anime_id INT NOT NULL,
  extra_large VARCHAR(255),
  large VARCHAR(255),
  medium VARCHAR(255),
  color VARCHAR(255),

  CONSTRAINT fk_cover_image_anime_id FOREIGN KEY (anime_id) REFERENCES animes(id) ON DELETE CASCADE
);
