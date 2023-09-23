CREATE OR REPLACE FUNCTION get_total_row_count(
    _table_name VARCHAR,
    _col_name VARCHAR,
    _clause VARCHAR,
	OUT count INTEGER
)
AS $$
BEGIN
    EXECUTE format('SELECT COUNT(%s) FROM %s WHERE %s', _col_name, _table_name, _clause)
    INTO count;
END;
$$
LANGUAGE plpgsql;


-- CALLS
SELECT * FROM get_total_row_count('users', 'id', 'is_deleted = FALSE');

