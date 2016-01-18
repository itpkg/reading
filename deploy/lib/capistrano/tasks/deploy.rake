namespace :deploy do
  after 'deploy:updated', :build_reading do
    # on roles(:web), in: :groups, limit: 3, wait: 10 do
    #   within "#{release_path}/front-react" do
    #     execute :npm, :install, '--silent'
    #     execute :npm, :run, :build
    #   end
    # end

    on roles(:api), in: :groups, limit: 3, wait: 10 do
      within "#{release_path}/api" do
        execute :go, :get, '-u', 'github.com/itpkg/reading/api'
        execute :echo, '$GO_PATH'
        # execute :npm, :run, :build
      end
    end
  end

end