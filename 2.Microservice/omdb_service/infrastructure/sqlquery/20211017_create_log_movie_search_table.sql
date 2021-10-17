CREATE TABLE log_movie_search
(
    id          serial PRIMARY KEY,
    url         VARCHAR(100) not null,
    status_code int          not null,
    response    text,
    created_at  timestamp    not null
);