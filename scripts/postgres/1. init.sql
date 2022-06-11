CREATE USER akaddr
    WITH PASSWORD 'aswqas';

CREATE DATABASE base10
    WITH  OWNER = akaddr
    ENCODING = 'UTF8'
    LC_COLLATE = 'en_US.utf8'
    LC_CTYPE = 'en_US.utf8'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1;