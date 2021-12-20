SELECT *, UNIX_TIMESTAMP(updated_at) AS unix_ts_in_secs
FROM messages
WHERE UNIX_TIMESTAMP(updated_at) > :sql_last_value AND updated_at < NOW()
ORDER BY updated_at ASC