alter table account_contacts
drop column deleted_at;

alter table accounts
drop column deleted_at;

alter table campaigns
drop column deleted_at;

alter table contacts
drop column deleted_at;

alter table contracts
drop column deleted_at;

alter table event_contacts
drop column deleted_at;

alter table event_user_attendees
drop column deleted_at;

alter table event_user_mains
drop column deleted_at;

alter table events
drop column deleted_at;

alter table leads
drop column deleted_at;

alter table opportunities
drop column deleted_at;

alter table opportunity_campaigns
drop column deleted_at;

alter table order_products
drop column deleted_at;

alter table orders
drop column deleted_at;

alter table products
drop column deleted_at;

alter table quote_products
drop column deleted_at;

alter table quotes
drop column deleted_at;

alter table roles
drop column deleted_at;

alter table users
drop column deleted_at;