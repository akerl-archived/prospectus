module Prospectus
  ##
  # Simple test class, uses "set 'value'" to declare value
  module StaticTest
    def set(value)
      @state.value = value
    end
  end
end
