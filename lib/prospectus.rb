# frozen_string_literal: true

require 'logcabin'

##
# Tool and DSL for checking expected vs actual state
module Prospectus
  DEFAULT_FILE = './.prospectus'.freeze

  class << self
    ##
    # Insert a helper .new() method for creating a new Cache object
    def new(*args)
      self::List.new(*args)
    end

    ##
    # Method for loading list from DSL
    def load_from_file(params = {})
      file = params[:file] || DEFAULT_FILE
      list = List.new(params)
      dsl = ListDSL.new(list, params)
      dsl.instance_eval(File.read(file), file)
      list
    end

    def modules
      @modules ||= LogCabin.new(load_path: load_path(:modules))
    end

    def helpers
      @helpers ||= LogCabin.new(load_path: load_path(:helpers))
    end

    def extra_dep(name, dep)
      require dep
    rescue LoadError
      raise("The #{name} module requires the #{dep} gem")
    end

    private

    def gem_dir
      Gem::Specification.find_by_name('prospectus').gem_dir
    end

    def load_path(type)
      File.join(gem_dir, 'lib', 'prospectus', type.to_s)
    end
  end
end

require 'prospectus/version'
require 'prospectus/list'
require 'prospectus/item'
require 'prospectus/state'
