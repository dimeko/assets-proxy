START TRANSACTION;

-- Create a new user in the 'users' table
INSERT INTO users (username, password, website, website_image_dir, website_db_dir, website_auth_username, website_auth_password)
VALUES ('user1', 'put_your_hashed_password', 'user1-website.com', '_', '_', 'username', 'password');

INSERT INTO users (username, password, website, website_image_dir, website_db_dir, website_auth_username, website_auth_password)
VALUES ('user2', 'put_your_hashed_password', 'user2-website.com', '_', '_', 'username', 'password');

COMMIT;
