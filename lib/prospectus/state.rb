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

    def =~(other)
      return super unless other.is_a? Prospectus::State
      ov = other.value
      return ov.include?(@value) if ov.is_a? Enumerable
      return @value =~ ov if ov.is_a? Regexp
      @value == ov
    end

    def to_s
      @value.to_s
    end
  end

  ##
  # DSL for wrapping eval of states
  class StateDSL
    def initialize(state, params)
      @state = state
      @options = params
    end

    def respond_to_missing?(method, _ = false)
      return super if @module
      Prospectus.modules.find(method)
      true
    rescue RuntimeError
      super
    end

    def method_missing(method, *args, &block)
      return super if @module
      @module = Prospectus.modules.find(method)
      return super unless @module
      extend @module
    end
  end
end
