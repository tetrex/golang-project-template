CREATE TABLE Admin (
  id int PRIMARY KEY NOT NULL,
  email varchar DEFAULT null,
  password text DEFAULT null,
  isdelete ENUM ('1', '0') NOT NULL DEFAULT '0'
);
