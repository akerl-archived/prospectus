##
# Tool and DSL for checking expected vs actual state
module Prospectus
  DEFAULT_PATH = './.prospectus'

  class << self
    ##
    # Insert a helper .new() method for creating a new Cache object

    def new(*args)
      self::List.new(*args)
    end

    def load_from_file(params = {})
      path = params[:directory] || DEFAULT_PATH
      list = List.new
      dsl = ListDSL.new(list)
      dsl.instance_eval(File.read(path), path)
      list
    end
  end
end

require 'prospectus/version'
require 'prospectus/list'
require 'prospectus/item'
require 'prospectus/state'
