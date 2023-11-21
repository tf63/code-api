----------------------------------------------------------------
-- レコード数 1000000
----------------------------------------------------------------
-- Master
INSERT INTO language (language_id, name)
    SELECT i, format('language %s', i)
        FROM generate_series(1, 1000000) as i
;

INSERT INTO framework (framework_id, name)
    SELECT i, format('framework %s', i)
        FROM generate_series(1, 1000000) as i
;

INSERT INTO tool (language_id, framework_id, tool_id)
    SELECT 1, 1, i
        FROM generate_series(1, 1000000) as i
;

INSERT INTO algorithm (algorithm_id, name)
    SELECT i, format('algorithm %s', i)
        FROM generate_series(1, 1000000) as i
;

INSERT INTO pattern (pattern_id, name)
    SELECT i, format('pattern %s', i)
        FROM generate_series(1, 1000000) as i
;

-- not Master ?
INSERT INTO language_code (language_code_id, content, nrow, language_id)
    SELECT i, format('code %s', i), mod(i, 30)+1, mod(i, 30)+1
        FROM generate_series(1, 1000000) as i
;

INSERT INTO algorithm_code (algorithm_code_id, content, nrow, language_id, algorithm_id)
    SELECT i, format('code %s', i), mod(i, 30)+1, mod(i, 30)+1, mod(i, 30)+1
        FROM generate_series(1, 1000000) as i
;

----------------------------------------------------------------
-- レコード数 50000
----------------------------------------------------------------
-- -- Master
-- INSERT INTO language (language_id, name)
--     SELECT i, format('language %s', i)
--         FROM generate_series(1, 2000) as i
-- ;

-- INSERT INTO framework (framework_id, name)
--     SELECT i, format('framework %s', i)
--         FROM generate_series(1, 2000) as i
-- ;

-- INSERT INTO tool (language_id, framework_id, tool_id)
--     SELECT 1, 1, i
--         FROM generate_series(1, 2000) as i
-- ;

-- INSERT INTO algorithm (algorithm_id, name)
--     SELECT i, format('algorithm %s', i)
--         FROM generate_series(1, 2000) as i
-- ;

-- INSERT INTO pattern (pattern_id, name)
--     SELECT i, format('pattern %s', i)
--         FROM generate_series(1, 2000) as i
-- ;

-- -- not Master ?
-- INSERT INTO program_code (program_code_id, content, nrow, tool_id)
--     SELECT i, format('code %s', i), mod(i, 30), i/30 + 1
--         FROM generate_series(1, 50000) as i
-- ;

-- INSERT INTO algorithm_code (algorithm_code_id, content, nrow, language_id, algorithm_id)
--     SELECT i, format('code %s', i), mod(i, 30), i/30 + 1, i/30 + 1
--         FROM generate_series(1, 50000) as i
-- ;
