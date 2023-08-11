alter table account_contacts
    add deleted_at timestamp;

alter table accounts
    add deleted_at timestamp;

alter table campaigns
    add deleted_at timestamp;

alter table contacts
    add deleted_at timestamp;

alter table contracts
    add deleted_at timestamp;

alter table event_contacts
    add deleted_at timestamp;

alter table event_user_attendees
    add deleted_at timestamp;

alter table event_user_mains
    add deleted_at timestamp;

alter table events
    add deleted_at timestamp;

alter table leads
    add deleted_at timestamp;

alter table opportunities
    add deleted_at timestamp;

alter table opportunity_campaigns
    add deleted_at timestamp;

alter table order_products
    add deleted_at timestamp;

alter table orders
    add deleted_at timestamp;

alter table products
    add deleted_at timestamp;

alter table quote_products
    add deleted_at timestamp;

alter table quotes
    add deleted_at timestamp;

alter table roles
    add deleted_at timestamp;

alter table users
    add deleted_at timestamp;



