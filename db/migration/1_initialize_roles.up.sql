CREATE TABLE roles
(
    role_id serial NOT NULL,
    role_name character varying(64) NOT NULL,
    can_add boolean NOT NULL DEFAULT false,
    can_update boolean NOT NULL DEFAULT false,
    can_read boolean NOT NULL DEFAULT false,
    can_delete boolean NOT NULL DEFAULT false,
    created_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp without time zone,
    PRIMARY KEY (role_id)
);

INSERT INTO roles(role_id, role_name, can_read, can_add, can_update, can_delete)
VALUES (1,'admin', true,true,true,true), (2,'user', true,false,false,false);