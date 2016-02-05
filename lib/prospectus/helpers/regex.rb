module LogCabin
  module Modules
    ##
    # Use regex to adjust state value
    module Regex
      def regex_helper(value)
        return value unless @find
        raise("Value does not match regex: #{value}") unless value.match(@find)
        value.sub(@find, @replace)
      end

      private

      def regex(find, replace = '\1')
        @find = find
        @replace = replace
      end
    end
  end
end
