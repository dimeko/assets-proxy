START TRANSACTION;

CREATE TABLE IF NOT EXISTS users (
    id int NOT NULL,
    username varchar(255), 
    password  varchar(255),
    website varchar(255),
    website_image_dir varchar(255),
    website_db_dir varchar(255),
    website_auth_username varchar(255),
    website_auth_password varchar(255),
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, 
    PRIMARY KEY (id)
);

COMMIT;