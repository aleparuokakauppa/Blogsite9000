DROP TABLE IF EXISTS comments;
CREATE TABLE comments (
  ID         INT AUTO_INCREMENT NOT NULL,
  Author     VARCHAR(512) NOT NULL,
  Text       TEXT NOT NULL,
  Target     INT,
  Time       TEXT NOT NULL,
  PRIMARY KEY (`id`)
);
