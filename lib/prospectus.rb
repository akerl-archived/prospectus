##
# Tool and DSL for checking expected vs actual state
module Prospectus
  class << self
    ##
    # Insert a helper .new() method for creating a new Cache object

    def new(*args)
      self::List.new(*args)
    end
  end
end

require 'prospectus/version'
require 'prospectus/list'
require 'prospectus/item'
