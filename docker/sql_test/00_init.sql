CREATE TABLE IF NOT EXISTS algorithm (
    algorithm_id BIGINT PRIMARY KEY,
    name varchar(50) NOT NULL
);
COMMENT ON TABLE algorithm IS 'algorithmを管理するテーブル';

CREATE TABLE IF NOT EXISTS pattern (
    pattern_id BIGINT PRIMARY KEY,
    name varchar(50) NOT NULL
);
COMMENT ON TABLE pattern IS 'designpatternを管理するテーブル';

CREATE TABLE IF NOT EXISTS language (
    language_id BIGINT PRIMARY KEY,
    name varchar(50) NOT NULL
);
COMMENT ON TABLE language IS 'languageを管理するテーブル';

CREATE TABLE IF NOT EXISTS framework (
    framework_id BIGINT PRIMARY KEY,
    name varchar(50) NOT NULL
);
COMMENT ON TABLE framework IS 'frameworkを管理するテーブル';

CREATE TABLE IF NOT EXISTS tool (
    tool_id BIGINT PRIMARY KEY,
    language_id BIGINT NOT NULL,
    framework_id BIGINT DEFAULT 1 NOT NULL,
    FOREIGN KEY (language_id) REFERENCES public.language(language_id) ON DELETE RESTRICT,
    FOREIGN KEY (framework_id) REFERENCES public.framework(framework_id) ON DELETE SET DEFAULT
);
COMMENT ON TABLE tool IS 'frameworkとlanguageをまとめるテーブル';

CREATE TABLE IF NOT EXISTS language_code (
    language_code_id SERIAL PRIMARY KEY,
    content TEXT NOT NULL,
    nrow BIGINT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    language_id BIGINT NOT NULL,
    FOREIGN KEY (language_id) REFERENCES public.language(language_id) ON DELETE RESTRICT
);
CREATE INDEX ON language_code (language_id, nrow);
COMMENT ON TABLE language_code IS 'プログラミング言語のソースコードを管理するテーブル';

CREATE TABLE IF NOT EXISTS framework_code (
    framework_code_id SERIAL PRIMARY KEY,
    content TEXT NOT NULL,
    nrow BIGINT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    tool_id BIGINT NOT NULL,
    FOREIGN KEY (tool_id) REFERENCES public.tool(tool_id) ON DELETE RESTRICT
);
CREATE INDEX ON framework_code (tool_id, nrow);
COMMENT ON TABLE framework_code IS 'フレームワークやライブラリのソースコードを管理するテーブル';

CREATE TABLE IF NOT EXISTS algorithm_code (
    algorithm_code_id SERIAL PRIMARY KEY,
    content TEXT NOT NULL,
    nrow BIGINT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    language_id BIGINT NOT NULL,
    algorithm_id BIGINT NOT NULL,
    FOREIGN KEY (language_id) REFERENCES public.language(language_id) ON DELETE RESTRICT,
    FOREIGN KEY (algorithm_id) REFERENCES public.algorithm(algorithm_id) ON DELETE RESTRICT
);
CREATE INDEX ON algorithm_code (language_id, algorithm_id, nrow);
COMMENT ON TABLE algorithm_code IS '様々なプログラミング言語で書かれたアルゴリズムのソースコードを管理するテーブル';

CREATE TABLE IF NOT EXISTS pattern_code (
    pattern_code_id SERIAL PRIMARY KEY,
    content TEXT NOT NULL,
    nrow BIGINT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    language_id BIGINT NOT NULL,
    pattern_id BIGINT NOT NULL,
    FOREIGN KEY (language_id) REFERENCES public.language(language_id) ON DELETE RESTRICT,
    FOREIGN KEY (pattern_id) REFERENCES public.pattern(pattern_id) ON DELETE RESTRICT
);
CREATE INDEX ON pattern_code (language_id, pattern_id, nrow);
COMMENT ON TABLE pattern_code IS '様々なプログラミング言語で書かれたデザインパターンのソースコードを管理するテーブル';
