module Prospectus
  ##
  # Define a state object that supports modular checks
  class State
    attr_accessor :value

    def initialize(params = {})
      @options = params
    end
  end

  ##
  # DSL for wrapping eval of states
  class StateDSL
    def initialize(state)
      @state = state
      @extended = false
    end

    def method_missing(method, *args, &block)
      return super if @extended
      module_obj = Prospectus::Module.find(method)
      return super unless module_obj
      extend module_obj
    end
  end
end
