CREATE TABLE IF NOT EXISTS orders (
    id BIGSERIAL PRIMARY KEY,
    status SMALLINT NOT NULL,
    user_id BIGINT NOT NULL,
    payment_type SMALLINT NOT NULL
);

CREATE TABLE IF NOT EXISTS items (
    item_id BIGSERIAL PRIMARY KEY,
    order_id BIGINT NOT NULL REFERENCES orders(id) ON DELETE CASCADE,
    amount BIGINT NOT NULL,
    discounted_amount BIGINT NOT NULL
    );

ALTER TABLE orders
    ADD CONSTRAINT status_check CHECK (status IN (0, 1, 2, 3)),
    ADD CONSTRAINT payment_type_check CHECK (payment_type IN (0, 1, 2));
