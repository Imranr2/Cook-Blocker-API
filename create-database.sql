SET FOREIGN_KEY_CHECKS = 0;
DROP TABLE IF EXISTS user_tab;
CREATE TABLE user_tab (
    id INT NOT NULL AUTO_INCREMENT,
    username VARCHAR(30) NOT NULL,
    password VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    role TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX (username),
    PRIMARY KEY (`id`)
    -- FOREIGN KEY (role) REFERENCES role_tab(id)
);

DROP TABLE IF EXISTS role_tab;
CREATE TABLE role_tab (
    id INT NOT NULL AUTO_INCREMENT,
    role VARCHAR(255) NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`, `role`)
);


DROP TABLE IF EXISTS menu_item_tab;
CREATE TABLE menu_item_tab (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    price FLOAT NOT NULL,
    created_by INT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX (name),
    PRIMARY KEY (`id`, `name`),
    FOREIGN KEY (created_by) REFERENCES user_tab(id)
);

DROP TABLE IF EXISTS ingredient_tab;
CREATE TABLE ingredient_tab (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX(name),
    PRIMARY KEY (`id`, `name`)
);

DROP TABLE IF EXISTS menu_item_ingredient_tab;
CREATE TABLE menu_item_ingredient_tab (
    id INT NOT NULL AUTO_INCREMENT,
    menu_item_id INT NOT NULL,
    ingredient_id INT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    FOREIGN KEY (menu_item_id) REFERENCES menu_item_tab(id),
    FOREIGN KEY (ingredient_id) REFERENCES ingredient_tab(id)
);

DROP TABLE IF EXISTS item_image_tab;
CREATE TABLE item_image_tab (
    id INT NOT NULL AUTO_INCREMENT,
    menu_item_id INT NOT NULL,
    image_link TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX(menu_item_id),
    PRIMARY KEY (`id`),
    FOREIGN KEY (menu_item_id) REFERENCES menu_item_tab(id)
);

DROP TABLE IF EXISTS order_tab;
CREATE TABLE order_tab (
    id INT NOT NULL AUTO_INCREMENT,
    table_id INT NOT NULL,
    user_id INT NOT NULL,
    is_completed BIT NOT NULL,
    price FLOAT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    FOREIGN KEY (user_id) REFERENCES user_tab(id),
    FOREIGN KEY (table_id) REFERENCES table_tab(id)
);

DROP TABLE IF EXISTS order_item_tab;
CREATE TABLE order_item_tab (
    id INT NOT NULL AUTO_INCREMENT,
    quantity INT NOT NULL,
    order_id INT NOT NULL,
    menu_item_id INT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    FOREIGN KEY (order_id) REFERENCES order_tab(id),
    FOREIGN KEY (menu_item_id) REFERENCES menu_item_tab(id)
);

DROP TABLE IF EXISTS table_tab;
CREATE TABLE table_tab (
    id INT NOT NULL AUTO_INCREMENT,
    pax INT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
);

DROP TABLE IF EXISTS reservation_tab;
CREATE TABLE reservation_tab (
    id INT NOT NULL AUTO_INCREMENT,
    customer_name VARCHAR(255) NOT NULL,
    pax INT NOT NULL,
    table_id INT NOT NULL,
    reservation_datetime DATETIME NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    FOREIGN KEY (table_id) REFERENCES table_tab(id)
);