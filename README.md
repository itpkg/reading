reading
---
I love reading.

## Documents
 * [https://developers.google.com/youtube/v3/getting-started](https://developers.google.com/youtube/v3/getting-started)

## Development



### Install
    bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
    gvm install go1.4
    GOROOT_BOOTSTRAP=/home/flamen/.gvm/gos/go1.4 gvm install go1.5.2
    gvm use go1.5.2 --default
    go get github.com/itpkg/reading/api


### Setup database (as postgres)

    psql
    CREATE USER reading WITH PASSWORD 'changeme';
    CREATE DATABASE reading_production WITH ENCODING='UTF8';
    GRANT ALL PRIVILEGES ON DATABASE reading_production TO reading;
    \q

#### test database connection


    psql -U reading -d reading_production    

* if report 'FATAL:  Peer authentication failed for user "reading"', open file "/etc/postgresql/9.3/main/pg_hba.conf" change line "local   all             all                                     peer" to "local   all             all                                     md5" and then run:

    service postgresql restart

### Start
    cd $GOPATH/src/github.com/itpkg/reading/api
    go run app.go server

    cd $GOPATH/src/github.com/itpkg/reading/front
    npm run start

## Deployment

### Build
    make
    ls release

### Run
    export ITPKG_ENV=production
    ls config # edit config files
    ./itpkg -h
    ./itpkg db migrate
    ./itpkg db seed
    ./itpkg nginx
    ./itpkg build
    ./itpkg status
    ./itpkg server
