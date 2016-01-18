Deployment
---

### on server
    curl -sL https://deb.nodesource.com/setup_5.x | sudo -E bash -
    sudo apt-get install -y nodejs golang-go
    sudo sudo apt-get install redis-server postgresql nginx
    sudo useradd -s /bin/bash -m deploy
    sudo passwd -l deploy
    sudo su - deploy
    sudo echo 'deploy ALL=(ALL) NOPASSWD: ALL' > /etc/sudoers.d/deploy
    mkdir .ssh
    chmod 700 .ssh
    cat /tmp/id_rsa.pub >> .ssh/authorized_keys


### on local
    cd deploy
    vi config/deploy/production.rb
    cap production deploy:check
    cap production deploy
