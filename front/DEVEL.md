DEVELOPMENT
---

### Atom-Editor
    git clone https://github.com/atom/atom.git
    git fetch -p
    git checkout $(git describe --tags `git rev-list --tags --max-count=1`)

    cd atom
    sudo pacman -S --needed gconf base-devel git nodejs npm libgnome-keyring python2
    export PYTHON=/usr/bin/python2
    script/build
    ls -l $TMPDIR/atom-build/Atom
    sudo script/grunt install
    
### Env
    npm install -g ember-cli
    npm install -g phantomjs
    npm install
    bower install
    ember server --proxy http://127.0.0.1:3000

### Deploy
    ember build --environment="production"
