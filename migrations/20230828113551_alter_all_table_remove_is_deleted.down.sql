alter table account_contacts
    add is_deleted timestamp;

alter table accounts
    add is_deleted timestamp;

alter table campaigns
    add is_deleted timestamp;

alter table contacts
    add is_deleted timestamp;

alter table contracts
    add is_deleted timestamp;

alter table event_contacts
    add is_deleted timestamp;

alter table event_user_attendees
    add is_deleted timestamp;

alter table event_user_mains
    add is_deleted timestamp;

alter table events
    add is_deleted timestamp;

alter table leads
    add is_deleted timestamp;

alter table opportunities
    add is_deleted timestamp;

alter table opportunity_campaigns
    add is_deleted timestamp;

alter table order_products
    add is_deleted timestamp;

alter table orders
    add is_deleted timestamp;

alter table products
    add is_deleted timestamp;

alter table quote_products
    add is_deleted timestamp;

alter table quotes
    add is_deleted timestamp;

alter table roles
    add is_deleted timestamp;

alter table users
    add is_deleted timestamp;



