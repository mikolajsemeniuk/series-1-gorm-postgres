-- SELECT EXISTS(SELECT datname FROM pg_catalog.pg_database WHERE datname = 'orders');

-- IF SELECT count(datname) FROM pg_catalog.pg_database WHERE datname = 'orders') > 0 
-- THEN

-- ELSE
-- --     REVOKE CONNECT ON DATABASE orders FROM public;
-- --     SELECT pg_terminate_backend(pg_stat_activity.pid)
-- --     FROM pg_stat_activity
-- --     WHERE pg_stat_activity.datname = 'orders';
-- --     DROP DATABASE IF EXISTS orders;
-- END IF;
-- create database orders;

-- REVOKE CONNECT ON DATABASE orders FROM public;
-- SELECT pg_terminate_backend(pg_stat_activity.pid)
-- FROM pg_stat_activity
-- WHERE pg_stat_activity.datname = 'orders';
-- DROP DATABASE orders;

-- DO $$
-- BEGIN
--     IF EXISTS (SELECT 1 FROM pg_catalog.pg_database WHERE datname = 'orders') THEN
--         REVOKE CONNECT ON DATABASE orders FROM public;
--         SELECT pg_terminate_backend(pg_stat_activity.pid)
--         FROM pg_stat_activity
--         WHERE pg_stat_activity.datname = 'orders';
--         DROP DATABASE orders;
--     END IF;
-- END; 
-- $$ 
-- DO $$
-- BEGIN
--   IF SELECT EXISTS(SELECT datname FROM pg_catalog.pg_database WHERE datname = 'orders') THEN
--      --some logic
--   END IF;
-- END; 
-- $$ 
-- DO $$
-- BEGIN
--     IF EXISTS (SELECT 1 FROM people p WHERE p.person_id = my_person_id) THEN
    
--     END IF
-- END 
-- $$;