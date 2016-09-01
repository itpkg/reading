namespace :deploy do
  after :compile_assets, :npm_build do
    on roles(:web) do |host|
      info "Install npm packages on #{host}"
      within release_path do
        execute :npm, 'run', 'build'
      end
    end
  end
end