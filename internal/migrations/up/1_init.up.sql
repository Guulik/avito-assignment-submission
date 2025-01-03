CREATE TABLE IF NOT EXISTS "banner"
(
    banner_id  serial    not null primary key,
    feature_id bigint    not null,
    tag_ids    bigint[]  not null,
    content    JSON      not null,
    is_active  boolean   not null default true,
    created_at timestamp not null,
    updated_at timestamp not null
);

CREATE TABLE IF NOT EXISTS "banner_definition"
(
    banner_id  BIGINT NOT NULL REFERENCES "banner" (banner_id) ON DELETE CASCADE,
    feature_id BIGINT NOT NULL,
    tag_id     BIGINT NOT NULL,
    PRIMARY KEY (feature_id, tag_id)
);