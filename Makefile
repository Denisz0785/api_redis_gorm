.PHONY: db redis

db:
	mysql -h 127.0.0.1 -P 3333 -u myuser -p
	SHOW DATABASES;
	USE goredisfiber;
	SHOW TABLES;
redis:
	redis-cli -h 127.0.0.1 -p 1111
	KEYS *
	GET key

