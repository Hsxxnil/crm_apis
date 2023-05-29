drop index idx_quote_products_code;
drop index idx_quote_products_description;

alter table quote_products
drop column code;

alter table quote_products
drop column description;