START TRANSACTION;

CREATE TABLE IF NOT EXISTS logs (
    id int NOT NULL, 
    user int,
    event varchar(255),
    info varchar(255),
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, 
    PRIMARY KEY (id)
);

COMMIT;