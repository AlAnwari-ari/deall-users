CREATE TABLE users
(
    user_id serial NOT NULL,
    role_id integer NOT NULL,
    username character varying(16) NOT NULL,
    email character varying(256) NOT NULL,
    fullname character varying(120) NOT NULL,
    password character varying(120) NOT NULL,
    created_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp without time zone,
    PRIMARY KEY (user_id),
    CONSTRAINT uniq_username UNIQUE (username),
    CONSTRAINT fk_role_user FOREIGN KEY (role_id)
        REFERENCES roles (role_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID
);

INSERT INTO users(role_id, username, email, fullname, password)
VALUES (1,'admin', 'admin@mail.com', 'Admin', '$2a$04$VmylnGWHTnYkepWSUYeb7uEwOHj1G/HGzDz17drQe64HbNsYyZzrm');