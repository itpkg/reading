# Be sure to restart your server when you modify this file.

# Version of your assets, change this if you want to expire all your assets.
Rails.application.config.assets.version = '1.0'

# Add additional assets to the asset load path
# Rails.application.config.assets.paths << Emoji.images_path

# Precompile additional assets.
# application.js, application.css, and all non-JS/CSS in app/assets folder are already added.
# Rails.application.config.assets.precompile += %w( search.js )

# node_modules
Rails.application.root.join('node_modules').to_s.tap do |npm_path|
  Rails.application.config.sass.load_paths << npm_path
  Rails.application.config.assets.paths << npm_path
end
# Precompile Bootstrap fonts
Rails.application.config.assets.precompile << %r(bootstrap-sass/assets/fonts/bootstrap/[\w-]+\.(?:eot|svg|ttf|woff2?)$)
# Minimum Sass number precision required by bootstrap-sass
::Sass::Script::Value::Number.precision = [8, ::Sass::Script::Value::Number.precision].max

Rails.application.config.assets.precompile << 'marked/lib/marked.js'