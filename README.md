reading
---
I love reading.

## Resources
 * [https://developers.google.com/youtube/v3/getting-started](Youtube API)
 * [https://developers.google.com/identity/protocols/OAuth2](Google Oauth2)
 * [https://addons.mozilla.org/en-US/firefox/addon/react-devtools/](React devtools for Firefox)
 * [https://chrome.google.com/webstore/detail/react-developer-tools/fmkadmapgofadopljbjfkapdkoienihi](React devtools for Chrome)

## Development

### Install packages
    curl -sL https://deb.nodesource.com/setup_5.x | sudo -E bash -
    sudo apt-get install -y nodejs golang-go
    sudo sudo apt-get install redis-server postgresql nginx

### Install go
    bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
    gvm install go1.4
    GOROOT_BOOTSTRAP=/home/flamen/.gvm/gos/go1.4 gvm install go1.5.3
    gvm use go1.5.3 --default
    go get github.com/itpkg/reading/api


### Setup database (as postgres)

    psql
    CREATE USER reading WITH PASSWORD 'changeme';
    CREATE DATABASE itpkg_reading_prod WITH ENCODING='UTF8';
    GRANT ALL PRIVILEGES ON DATABASE itpkg_reading_prod TO reading;
    \q

#### test database connection


    psql -U reading -d itpkg_reading_prod    

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
    source .env
    ls config # edit config files
    ./itpkg -h
    ./itpkg db migrate
    ./itpkg db seed
    ./itpkg nginx
    ./itpkg build
    ./itpkg status
    ./itpkg server
