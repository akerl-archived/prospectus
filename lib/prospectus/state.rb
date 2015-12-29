module Prospectus
  ##
  # Define a state object that supports modular checks
  class State
    attr_accessor :value

    def initialize(params = {})
      @options = params
    end

    def self.from_block(params = {}, state = nil, &block)
      state ||= State.new(params)
      dsl = StateDSL.new(state, params)
      dsl.instance_eval(&block)
      dsl.load!
      state
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
      @module = Prospectus.modules.find(method)
      return super unless @module
      extend @module
    end
  end
end
