##
# Tool and DSL for checking expected vs actual state
module Prospectus
  DEFAULT_FILE = './.prospectus'

  class << self
    ##
    # Insert a helper .new() method for creating a new Cache object
    def new(*args)
      self::List.new(*args)
    end

    ##
    # Method for loading list from DSL
    def load_from_file(params = {})
      params[:file] ||= DEFAULT_FILE
      list = List.new(params)
      dsl = ListDSL.new(list, params)
      dsl.instance_eval(File.read(params[:file]), params[:file])
      list
    end
  end
end

require 'prospectus/version'
require 'prospectus/module'
require 'prospectus/list'
require 'prospectus/item'
require 'prospectus/state'
