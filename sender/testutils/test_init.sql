CREATE TABLE IF NOT EXISTS questions (
    Date TIME PRIMARY KEY NOT NULL,
    Question STRING NOT NULL
);

CREATE TABLE IF NOT EXISTS users_data (
    User_nickname STRING REFERENCES users_auth (Nickname),
    Name          STRING NOT NULL,
    Sex           STRING
);

CREATE TABLE IF NOT EXISTS users_auth (
    Nickname STRING UNIQUE PRIMARY KEY,
    Login STRING NOT NULL,
    Password STRING NOT NULL
);

CREATE TABLE IF NOT EXISTS users_questions (
    Question_id   TIME   REFERENCES questions (Date),
    User_nickname STRING REFERENCES users_auth (Nickname),
    Answer        STRING NOT NULL,
    Created_at    TIME   NOT NULL,
    Updated_at    TIME   NOT NULL
);