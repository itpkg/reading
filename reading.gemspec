$:.push File.expand_path('../lib', __FILE__)

# Maintain your gem's version:
require 'reading/version'

# Describe your gem and declare its dependencies:
Gem::Specification.new do |s|
  s.name        = 'reading'
  s.version     = Reading::VERSION
  s.authors     = ['Jitang Zheng']
  s.email       = ['jitang.zheng@gmail.com']
  s.homepage    = 'https://github.com/itpkg/reading'
  s.summary     = 'reading engine'
  s.description = 'read epub books'
  s.license     = 'MIT'

  s.files = Dir['{app,config,db,lib}/**/*', 'MIT-LICENSE', 'Rakefile', 'README.md']

  s.add_dependency 'rails', '~> 5.0.0', '>= 5.0.0.1'
  s.add_dependency 'epub-parser'

  s.add_development_dependency 'sqlite3'
end
