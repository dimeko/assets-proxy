START TRANSACTION;

-- Create a new user in the 'users' table
INSERT INTO users (username, password, website, website_image_dir, website_db_dir, website_auth_username, website_auth_password)
VALUES ('userA', 'd476944f46865d03e83421d3b52e763f', 'user1-website.com', '_', '_', 'username', 'password');

INSERT INTO users (username, password, website, website_image_dir, website_db_dir, website_auth_username, website_auth_password)
VALUES ('userB', '74f7da5e7e1c5fe23ba76e4e6db7ebc6', 'user2-website.com', '_', '_', 'username', 'password');

COMMIT;
