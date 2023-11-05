-- Master
INSERT INTO language (language_id, name)
    VALUES (1, 'python'), (2, 'go'), (3, 'typescript'), (4, 'javascript')
;

INSERT INTO framework (framework_id, name)
    VALUES (1, 'something'), 
            (2, 'fastapi'), (3, 'flask'), (4, 'django'), 
            (5, 'react'), (6, 'nextjs'), (7, 'vue'), (8, 'nuxtjs'), (9, 'nodejs'), (10, 'express')
;

INSERT INTO tool (language_id, framework_id, tool_id)
    VALUES (1, 1, 1), (1, 2, 2), (1, 3, 3),
            (2, 1, 4), 
            (3, 1, 5), (3, 5, 6), (3, 6, 7), (3, 7, 8), (3, 8, 9), (3, 9, 10), (3, 10, 11),
            (4, 1, 12), (4, 5, 13), (4, 6, 14), (4, 7, 15), (4, 8, 16), (4, 9, 17), (4, 10, 18)
;

INSERT INTO algorithm (algorithm_id, name)
    VALUES (1, 'dfs'), (2, 'bfs'), (3, 'dp'), (4, 'binary search')
;

INSERT INTO pattern (pattern_id, name)
    VALUES (1, 'abstract factory pattern'), (2, 'builder pattern'), (3, 'factory method pattern'), (4, 'prototype pattern'), (5, 'singleton pattern')
;

-- not Master ?
INSERT INTO program_code (content, nrow, tool_id)
    VALUES ('this is python code', 1, 1),
            ('this is fastapi code', 1, 2),
            ('this is go code', 1, 4)
;

INSERT INTO algorithm_code (content, nrow, algorithm_id, language_id)
    VALUES ('this is python dfs code', 1, 1, 1),
            ('this is go dfs code', 1, 1, 2)
;

INSERT INTO pattern_code (content, nrow, pattern_id, language_id)
    VALUES ('this is python abstract factory pattern code', 1, 1, 1),
        ('this is typescript abstract factory pattern code', 1, 1, 3)
;
