CREATE TABLE IF NOT EXISTS clients (
    id          INTEGER      PRIMARY KEY UNIQUE,
    phone       INTEGER (11) NOT NULL,
    mobile_code INTEGER      NOT NULL,
    tag         STRING       NOT NULL
);

CREATE TABLE IF NOT EXISTS mails (
    id          INTEGER PRIMARY KEY UNIQUE,
    start_time  STRING  NOT NULL,
    end_time    STRING  NOT NULL,
    text        STRING  NOT NULL,
    mobile_code INTEGER NOT NULL,
    tag         STRING  NOT NULL
);

CREATE TABLE IF NOT EXISTS message (
    id        INTEGER PRIMARY KEY UNIQUE,
    send_time STRING  NOT NULL,
    mail_id   INTEGER REFERENCES mail (id) NOT NULL,
    client_id INTEGER REFERENCES client (id) NOT NULL
);