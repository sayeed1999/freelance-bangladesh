DO $$
BEGIN
    IF NOT EXISTS (SELECT FROM pg_database WHERE datname = 'freelance-db') THEN
        EXECUTE 'CREATE DATABASE freelance-db';
    END IF;
END $$;
