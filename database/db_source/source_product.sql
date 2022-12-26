create table if not exists public.source_product
(
    id            bigint    default nextval('db_source_id_seq'::regclass) not null
    constraint db_source_pk
    primary key,
    product_name  varchar,
    qty           integer,
    selling_price integer,
    promo_price   integer,
    created_at    timestamp default now()                                 not null,
    updated_at    timestamp default now()                                 not null
);