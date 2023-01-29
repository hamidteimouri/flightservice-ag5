up:
	curl  https://edu.postgrespro.com/demo-small-en.zip -o init-postgres/db.zip
	unzip init-postgres/db.zip -d init-postgres
	rm init-postgres/db.zip
	mv init-postgres/demo-small-en-20170815.sql init-postgres/2.sql
	docker-compose up -d