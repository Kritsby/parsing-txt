CREATE TABLE if not exists temp (data jsonb);
COPY temp (data)
FROM './migrations/output.json';
CREATE TABLE if not exists country (id SERIAL, code TEXT, ru text, en text);
INSERT INTO country
SELECT (data->>'code')::text,
	(data->>'ru')::text,
	(data->>'en')::text
FROM temp;
DROP TABLE if exists temp;