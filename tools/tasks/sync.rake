desc 'sync files to server'
task :sync do
  puts `rsync -a dst ssh://deploy@#{ENV['DOMAIN']}/var/www/#{ENV['DOMAIN']}/current/public/assets`
end