CREATE TABLE tasks (
    id serial PRIMARY KEY,
    create_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    update_date TIMESTAMP,
    owner INTEGER NOT NULL,
    name VARCHAR(200) NOT NULL,
    is_complete BOOLEAN DEFAULT FALSE,
    complete_date TIMESTAMP,

    FOREIGN KEY (owner) REFERENCES users (id) ON DELETE CASCADE
);

CREATE FUNCTION set_complete_date() RETURNS TRIGGER AS $set_complete_date$
    BEGIN
        IF NEW.is_complete IS TRUE AND NEW.is_complete <> OLD.is_complete THEN
            NEW.complete_date := CURRENT_TIMESTAMP;
        END IF;

        RETURN NEW;
    END;
$set_complete_date$ LANGUAGE plpgsql;

CREATE TRIGGER set_complete_date
    BEFORE UPDATE OF is_complete ON tasks
    FOR EACH ROW  
    EXECUTE FUNCTION set_complete_date();

CREATE FUNCTION set_update_date() RETURNS TRIGGER AS $set_update_date$
    BEGIN
        IF NEW.owner <> OLD.owner OR NEW.name <> OLD.name THEN 
            NEW.update_date := CURRENT_TIMESTAMP;
        END IF;

        RETURN NEW;
    END;
$set_update_date$ LANGUAGE plpgsql;

CREATE TRIGGER update_date
    BEFORE UPDATE ON tasks
    FOR EACH ROW
    EXECUTE FUNCTION set_update_date();

CREATE FUNCTION read_only_date_update() RETURNS TRIGGER AS $read_only_date_update$
    BEGIN
        IF NEW.create_date <> OLD.create_date THEN
            RAISE EXCEPTION 'cannot UPDATE create_date';
        END IF;

        IF NEW.update_date <> OLD.update_date THEN
            RAISE EXCEPTION 'cannot UPDATE update_date';
        END IF;

        IF NEW.complete_date <> OLD.complete_date THEN
            RAISE EXCEPTION 'cannot UPDATE complete_date';
        END IF;

        RETURN NEW;
    END;
$read_only_date_update$ LANGUAGE plpgsql;

CREATE TRIGGER read_only_date_update
    BEFORE UPDATE ON tasks
    FOR EACH ROW
    EXECUTE FUNCTION read_only_date_update();

-- todo -> read only trigger for insert