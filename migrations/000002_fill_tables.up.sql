INSERT INTO users (name, money_amount)
SELECT * FROM (VALUES ('George', 25000),
                      ('Kate', 23000),
                      ('Eric', 37000)) AS foo
WHERE NOT EXISTS (SELECT 1 FROM users);
