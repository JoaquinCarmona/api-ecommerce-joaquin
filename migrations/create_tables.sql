
CREATE EXTENSION "uuid-ossp";

CREATE TABLE public.users (
          id uuid NOT NULL DEFAULT uuid_generate_v4(),
          "name" varchar(60) NOT NULL,
          email varchar(60) NOT NULL,
          phone varchar(10) NULL,
          "password" varchar NULL,
          master_api_key varchar NULL,
          created_at timestamp NULL,
          updated_at timestamp NULL,
          deleted_at timestamp NULL,
          CONSTRAINT users_pk PRIMARY KEY (id)
);


CREATE TABLE public.products (
          id uuid NOT NULL DEFAULT uuid_generate_v4(),
          "name" varchar(60) NOT NULL,
          description varchar(255) NOT NULL,
          sku varchar(10) NULL,
          stock int not NULL DEFAULT 0,
          image_url varchar(255) NULL,
          price numeric(11,2) NOT NULL default 0.00,
          price_in_cents bigint NOT NULL default 0,
          currency varchar(5) NOT NULL default 'USD',
          created_at timestamp NULL,
          updated_at timestamp NULL,
          deleted_at timestamp NULL,
          CONSTRAINT products_pk PRIMARY KEY (id)
);

CREATE TABLE public.carts (
          id uuid NOT NULL DEFAULT uuid_generate_v4(),
          total_price_in_cents bigint NOT NULL default 0,
          currency varchar(5) NOT NULL default 'USD',
          status varchar(10) NULL,
          created_at timestamp NULL,
          updated_at timestamp NULL,
          CONSTRAINT carts_pk PRIMARY KEY (id)
);


CREATE TABLE public.cart_product (
          cart_id uuid NOT NULL,
          product_id uuid NOT NULL,
          qty bigint NOT NULL,
          price_in_cents bigint NOT NULL,
          added_at timestamp NULL,
          CONSTRAINT fk_cart_product_cart FOREIGN KEY (cart_id) REFERENCES public.carts(id),
          CONSTRAINT fk_cart_product_product FOREIGN KEY (product_id) REFERENCES public.products(id),
          CONSTRAINT pk_cart_product PRIMARY KEY (cart_id, product_id)
);