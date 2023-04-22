alter table products
    add price decimal not null;

create index idx_products_price
    on products (price);