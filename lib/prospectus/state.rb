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
    def initialize(state, params)
      @state = state
      @options = params
    end

    def method_missing(method, *args, &block)
      return super if @module
      @module = Prospectus::Module.find(method)
      return super unless @module
      extend @module
    end
  end
end
