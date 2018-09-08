module LogCabin
  module Modules
    ##
    # Simple text class, uses "set 'value'" to declare value
    module Static
      def load
        raise('Must use `set` to provide a value') unless @value
        @value
      end

      private

      def set(value)
        @value = value
      end
    end
  end
end
