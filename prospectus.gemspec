$:.unshift File.expand_path('../lib/', __FILE__)
require 'prospectus/version'

Gem::Specification.new do |s|
  s.name        = 'prospectus'
  s.version     = Prospectus::VERSION
  s.date        = Time.now.strftime('%Y-%m-%d')

  s.summary     = 'Tool and DSL for checking expected vs actual state'
  s.description = "Tool and DSL for checking expected vs actual state"
  s.authors     = ['Les Aker']
  s.email       = 'me@lesaker.org'
  s.homepage    = 'https://github.com/akerl/prospectus'
  s.license     = 'MIT'

  s.files       = `git ls-files`.split
  s.test_files  = `git ls-files spec/*`.split
  s.executables = ['prospectus']

  s.add_dependency 'mercenary', '~> 0.3.4'

  s.add_development_dependency 'rubocop', '~> 0.35.0'
  s.add_development_dependency 'rake', '~> 10.4.0'
  s.add_development_dependency 'codecov', '~> 0.1.1'
  s.add_development_dependency 'rspec', '~> 3.4.0'
  s.add_development_dependency 'fuubar', '~> 2.0.0'
end
