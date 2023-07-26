alter table order_products
    add is_deleted bool default false not null;