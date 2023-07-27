FROM busybox
# アプリ（バッチ）
COPY bin/go1/main /app/main
# CRON
COPY build/go1/test.crontab /var/spool/cron/crontabs/root
# コンテナ起動時に実行するコマンド
CMD crond -f