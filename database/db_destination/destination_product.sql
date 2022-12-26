create table public.destination_product
(
    id            bigint    default nextval('table_name_id_seq'::regclass) not null
        constraint table_name_pk
            primary key,
    product_name  varchar,
    qty           integer,
    selling_price integer,
    promo_price   integer,
    created_at    timestamp default now()                                  not null,
    updated_at    timestamp default now()                                  not null
);