module Prospectus
  ##
  # Define a state object that supports modular checks
  class State
    def initialize(params = {})
      @options = params
    end

    def version
      '0.0.2'
    end
  end

  ##
  # DSL for wrapping eval of states
  class StateDSL
    def initialize(state)
      @state = state
    end
  end
end
