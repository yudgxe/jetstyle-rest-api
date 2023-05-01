DROP TRIGGER set_complete_date ON tasks;
DROP TRIGGER update_date ON tasks;
DROP TRIGGER read_only_date_update ON tasks;


DROP FUNCTION set_complete_date();
DROP FUNCTION set_update_date();
DROP FUNCTION read_only_date_update();

DROP TABLE tasks;