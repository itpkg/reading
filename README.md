# reading
I love reading.

## Development

### Install ruby
    git clone https://github.com/rbenv/rbenv.git ~/.rbenv
    git clone https://github.com/rbenv/ruby-build.git ~/.rbenv/plugins/ruby-build
    git clone https://github.com/rbenv/rbenv-vars.git ~/.rbenv/plugins/rbenv-vars

    # Modify your ~/.zshrc file instead of ~/.bash_profile
    echo 'export PATH="$HOME/.rbenv/bin:$PATH"\neval "$(rbenv init -)"' >> ~/.bash_profile 
    
    rbenv install -l
    rbenv install 2.3.1
    rbenv local 2.3.1
    gem install bundler
    npm install
    
### Starting
    git clone https://github.com/itpkg/reading.git
    cd reading
    bundler install
    
    
    
## Deployment
    cd reading
    vi config/deploy.rb
    cap production deploy
    

## Documents
* [Material-UI](http://www.material-ui.com)
* [Material-UI Icons](https://design.google.com/icons/)
* [Redux](http://redux.js.org/index.html)
