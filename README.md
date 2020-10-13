# DB_MONTEIRO_PO1

## What is this API ?
It's a student project allowing to manipulate sql.

**The given database in migrations folder was poorly designed. It will be the subject of a future optimization work.**

## Set up project
Set up variable environnement

Create a .env file and add following vars with your data

```
USERDB=root
PASSDB=root
IPDB=db:3306
NAMEDB=classicmodels
```

```sh
$ docker-compose up --build
```

Replace "root" (it's username) and "root" (it's password) by your own login

```sh
$ docker exec -i db sh -c 'exec mysql -uroot -p"root"' < ./migrations/classicmodels.sql
```

Check API host 

```sh
$ docker ps
```

## API DOC
https://app.swaggerhub.com/apis/wyllisMonteiro/swagger-db_monteiro_p_01/1.0.0

## Unit tests
Coming soon