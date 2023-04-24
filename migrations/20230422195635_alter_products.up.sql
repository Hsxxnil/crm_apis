alter table products
    add price numeric not null;

create index idx_products_price
    on products (price);