-- init_db.sql

CREATE DATABASE IF NOT EXISTS test_bookstore;
GRANT ALL PRIVILEGES ON test_bookstore.* TO 'mysql_user'@'%';
FLUSH PRIVILEGES;