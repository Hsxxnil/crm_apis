alter table account_contacts
drop column is_deleted;

alter table accounts
drop column is_deleted;

alter table campaigns
drop column is_deleted;

alter table contacts
drop column is_deleted;

alter table contracts
drop column is_deleted;

alter table event_contacts
drop column is_deleted;

alter table event_user_attendees
drop column is_deleted;

alter table event_user_mains
drop column is_deleted;

alter table events
drop column is_deleted;

alter table leads
drop column is_deleted;

alter table opportunities
drop column is_deleted;

alter table opportunity_campaigns
drop column is_deleted;

alter table order_products
drop column is_deleted;

alter table orders
drop column is_deleted;

alter table products
drop column is_deleted;

alter table quote_products
drop column is_deleted;

alter table quotes
drop column is_deleted;

alter table roles
drop column is_deleted;

alter table users
drop column is_deleted;