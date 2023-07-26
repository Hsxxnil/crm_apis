alter table orders
    add is_deleted bool default false not null;