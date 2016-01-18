server '192.168.1.109', user: 'deploy', roles: %w{db api web}
set :deploy_to, '/var/www/192.168.1.109'