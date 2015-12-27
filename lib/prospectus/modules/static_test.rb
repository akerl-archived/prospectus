module LogCabin
  module Modules
    ##
    # Simple test class, uses "set 'value'" to declare value
    module StaticTest
      def load!
        @state.value = @value
      end

      private

      def set(value)
        @value = value
      end
    end
  end
end
