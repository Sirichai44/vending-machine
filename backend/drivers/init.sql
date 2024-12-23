CREATE DATABASE IF NOT EXISTS vending_machine;

USE vending_machine;

CREATE TABLE IF NOT EXISTS products (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    image_url VARCHAR(255),
    price INT NOT NULL,
    stock INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS money (
    value INT PRIMARY KEY,
    quantity INT NOT NULL
);

CREATE TABLE IF NOT EXISTS transactions (
    id INT AUTO_INCREMENT PRIMARY KEY,
    product_id INT NOT NULL,
    amount INT NOT NULL,
    `change` INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO products (name, image_url, price, stock) VALUES
    ('MEGA SPACE MOLLY 400% Sanrio Characters Series', 'https://prod-thailand-res.popmart.com/default/20241024_102631_678853____list_pic_____800x800.png', 200, 10),
    ('KAIJU Sitting in a Row Series Figures', 'https://prod-eurasian-res.popmart.com/default/20240117_173553_344734__1200x1200.jpg', 270, 10),
    ('Ultraman Shooting Studio Series Figures', 'https://prod-eurasian-res.popmart.com/default/20231208_161319_877211__1200x1200.jpg', 350, 10),
    ('SKULLPANDA Winter Symphony', 'https://prod-thailand-res.popmart.com/default/20241204_111332_857297____list_pic_____800x800.png', 300, 10),
    ('HIRONO x CLOT Series Figures', 'https://prod-thailand-res.popmart.com/default/20241211_153448_174631____list_pic_____800x800.png', 250, 10),
    ('Hirono × Le Petit Prince Series Figures', 'https://prod-thailand-res.popmart.com/default/20241023_142026_927920____10_____1200x1200.jpg', 250, 10);

INSERT INTO money (value, quantity) VALUES
    (1000, 10),
    (500, 10),
    (100, 30),
    (50, 30),
    (20, 30),
    (10, 50),
    (5, 50),
    (2, 50),
    (1, 50);
