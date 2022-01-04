CREATE TABLE clients(
    id serial NOT NULL,
    name VARCHAR(255) NOT NULL,
    PRIMARY KEY(id)
);

CREATE TABLE accounts(
    id SERIAL NOT NULL,
    client_id INT NOT NULL,
    currency VARCHAR(255) NOT NULL,
    amount DECIMAL NOT NULL DEFAULT 0,
    FOREIGN KEY(client_id) REFERENCES clients(id),
    PRIMARY KEY(id)
);

CREATE TABLE transactions(
    id serial NOT NULL,
    sender_account_id INT,
    receiver_account_id INT,
    type VARCHAR(255) NOT NULL,
    amount DECIMAL(9,2) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY(sender_account_id) REFERENCES accounts(id),
    FOREIGN KEY(receiver_account_id) REFERENCES accounts(id),
    PRIMARY KEY(id)
);