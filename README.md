# reading
I love reading.

## Development


### Install ruby
    sudo apt-get install -y git build-essential make libssl-dev libreadline-dev

    git clone https://github.com/rbenv/rbenv.git ~/.rbenv
    git clone https://github.com/rbenv/ruby-build.git ~/.rbenv/plugins/ruby-build
    git clone https://github.com/rbenv/rbenv-vars.git ~/.rbenv/plugins/rbenv-vars

    # Modify your ~/.zshrc file instead of ~/.bash_profile
    echo 'export PATH="$HOME/.rbenv/bin:$PATH"\neval "$(rbenv init -)"' >> ~/.bash_profile 
    
    rbenv install -l    
    CONFIGURE_OPTS="--disable-install-doc" rbenv install 2.3.1
    rbenv local 2.3.1
    gem install bundler
    
### Install nodejs
    curl -sL https://deb.nodesource.com/setup_6.x | sudo -E bash -
    sudo apt-get install -y nodejs
    
### Starting
    git clone https://github.com/itpkg/reading.git
    cd reading
    bundler install
    npm install
    
    
    
## Deployment
    cd reading
    vi config/deploy.rb
    cap production deploy
    cap production puma:config
    cap production puma:nginx_config
    

## Documents
* [Material-UI](http://www.material-ui.com)
* [Material-UI Icons](https://design.google.com/icons/)
* [Redux](http://redux.js.org/index.html)
