CREATE TABLE users (
  id INT NOT NULL AUTO_INCREMENT,
  name VARCHAR(100),
  mail VARCHAR(100),
  level INT,
  pass VARCHAR(1000),
  created_at DATETIME,
  updated_at DATETIME,
  primary key(id)
);
CREATE TABLE management (
  id INT NOT NULL AUTO_INCREMENT,
  user_id INT,
  sake_name VARCHAR(500),
  amount INT,
  date DATE,
  created_at DATETIME,
  updated_at DATETIME,
  primary key(id),
  FOREIGN KEY key_user_id(user_id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE RESTRICT
);
CREATE TABLE recipe (
  id INT NOT NULL AUTO_INCREMENT,
  user_id INT,
  title VARCHAR(100),
  text VARCHAR(5000),
  sumbnail VARCHAR(10000),
  created_at DATETIME,
  updated_at DATETIME,
  primary key(id),
  FOREIGN KEY key_user_id(user_id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE RESTRICT
);
LOAD DATA INFILE "/docker-entrypoint-initdb.d/users.csv" INTO TABLE users FIELDS TERMINATED BY ',' OPTIONALLY ENCLOSED BY '"' (`name`, `mail`, `level`, `pass`, `created_at`, `updated_at`);
LOAD DATA INFILE "/docker-entrypoint-initdb.d/management.csv" INTO TABLE management FIELDS TERMINATED BY ',' OPTIONALLY ENCLOSED BY '"' (`user_id`, `sake_name`, `amount`, `date`, `created_at`, `updated_at`);
LOAD DATA INFILE "/docker-entrypoint-initdb.d/recipe.csv" INTO TABLE recipe FIELDS TERMINATED BY ',' OPTIONALLY ENCLOSED BY '"' (`user_id`, `title`, `text`, `sumbnail`, `created_at`, `updated_at`);