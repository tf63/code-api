# postgres周り

### ログイン
コンテナの立ち上げ
```
    docker compose up -d
```

postgresコンテナをシェルにアタッチ
```
    docker compose exec postgres /bin/bash
```

postgresへ接続
```
    root@146370d17ba4:/# psql -U postgres postgres
```

### CSVファイルからインポート

ホスト側のCSVファイルからシーディングする
```
    \COPY language FROM /data/language.csv DELIMITER ',' CSV HEADER;
```

あるいはホスト側で
```
    psql -d postgres -U postgres -c "\COPY language FROM /data/language.csv DELIMITER ',' CSV HEADER;"
```

https://www.ashisuto.co.jp/db_blog/article/how-import-and-export-data-using-csv-files-postgresql.html

