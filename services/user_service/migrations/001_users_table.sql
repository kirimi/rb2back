-- +goose Up
CREATE TABLE IF NOT EXISTS public.users
(
    uid text COLLATE pg_catalog."default" NOT NULL,
    is_enabled boolean NOT NULL,
    locale text COLLATE pg_catalog."default",
    is_anonymous boolean NOT NULL,
    premium_type integer NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    launch_count bigint NOT NULL,
    app_version text COLLATE pg_catalog."default",
    build_version text COLLATE pg_catalog."default",
    device_info text COLLATE pg_catalog."default",
    name text COLLATE pg_catalog."default",
    email text COLLATE pg_catalog."default",
    picture_url text COLLATE pg_catalog."default",
    fb_created_at timestamp without time zone,
    CONSTRAINT users_pkey PRIMARY KEY (uid)
    )

    TABLESPACE pg_default;

CREATE UNIQUE INDEX IF NOT EXISTS __user_uid
    ON public.users USING btree
    (uid COLLATE pg_catalog."default" ASC NULLS LAST)
    TABLESPACE pg_default;

-- +goose Down
DROP TABLE IF EXISTS public.users;
DROP INDEX IF EXISTS public.__user_uid;

