\COPY language FROM /data/language.csv DELIMITER ',' CSV HEADER QUOTE '"' ESCAPE '\';
\COPY pattern FROM /data/pattern.csv DELIMITER ',' CSV HEADER QUOTE '"' ESCAPE '\';
\COPY framework FROM /data/framework.csv DELIMITER ',' CSV HEADER QUOTE '"' ESCAPE '\';
\COPY algorithm FROM /data/algorithm.csv DELIMITER ',' CSV HEADER QUOTE '"' ESCAPE '\';
\COPY tool FROM /data/tool.csv DELIMITER ',' CSV HEADER QUOTE '"' ESCAPE '\';
\COPY language_code ("nrow", "language_id", "content") FROM /data/language_code.csv DELIMITER ',' CSV HEADER QUOTE '"' ESCAPE '\';
\COPY framework_code ("nrow", "tool_id", "content") FROM /data/framework_code.csv DELIMITER ',' CSV HEADER QUOTE '"' ESCAPE '\';
\COPY algorithm_code ("nrow", "language_id", "algorithm_id", "content") FROM /data/algorithm_code.csv DELIMITER ',' CSV HEADER QUOTE '"' ESCAPE '\';
\COPY pattern_code ("nrow", "language_id", "pattern_id", "content") FROM /data/pattern_code.csv DELIMITER ',' CSV HEADER QUOTE '"' ESCAPE '\';