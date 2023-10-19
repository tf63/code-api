CREATE TABLE IF NOT EXISTS todo (
    id SERIAL PRIMARY KEY,
    title varchar(50) NOT NULL,
    done BOOLEAN DEFAULT false
);
COMMENT ON TABLE todo IS 'todoを管理するテーブル (サンプル)';

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
    language_id BIGINT,
    framework_id BIGINT DEFAULT 1,
    FOREIGN KEY (language_id) REFERENCES public.language(language_id) ON DELETE RESTRICT,
    FOREIGN KEY (framework_id) REFERENCES public.framework(framework_id) ON DELETE SET DEFAULT
);
COMMENT ON TABLE tool IS 'frameworkとlanguageをまとめるテーブル';

CREATE TABLE IF NOT EXISTS program_code (
    program_code_id SERIAL PRIMARY KEY,
    content TEXT NOT NULL,
    nrow BIGINT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now(),
    tool_id BIGINT,
    FOREIGN KEY (tool_id) REFERENCES public.tool(tool_id) ON DELETE RESTRICT
);
COMMENT ON TABLE program_code IS 'プログラミング言語のソースコードを管理するテーブル';

CREATE TABLE IF NOT EXISTS algorithm_code (
    algorithm_code_id SERIAL PRIMARY KEY,
    content TEXT NOT NULL,
    nrow BIGINT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now(),
    language_id BIGINT,
    algorithm_id BIGINT,
    FOREIGN KEY (language_id) REFERENCES public.language(language_id) ON DELETE RESTRICT,
    FOREIGN KEY (algorithm_id) REFERENCES public.algorithm(algorithm_id) ON DELETE RESTRICT
);
COMMENT ON TABLE algorithm_code IS '様々なプログラミング言語で書かれたアルゴリズムのソースコードを管理するテーブル';

CREATE TABLE IF NOT EXISTS pattern_code (
    pattern_code_id SERIAL PRIMARY KEY,
    content TEXT NOT NULL,
    nrow BIGINT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now(),
    language_id BIGINT,
    pattern_id BIGINT,
    FOREIGN KEY (language_id) REFERENCES public.language(language_id) ON DELETE RESTRICT,
    FOREIGN KEY (pattern_id) REFERENCES public.pattern(pattern_id) ON DELETE RESTRICT
);
COMMENT ON TABLE pattern_code IS '様々なプログラミング言語で書かれたデザインパターンのソースコードを管理するテーブル';
