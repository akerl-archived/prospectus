module LogCabin
  module Modules
    ##
    # Use regex to adjust state value
    module Regex
      def regex_helper(value)
        return value unless @find
        m = value.match(@find)
        raise("Value does not match regex: #{value}") unless m
        m.to_s.sub(@find, @replace)
      end

      private

      def regex(find, replace = '\1')
        @find = find
        @replace = replace
      end
    end
  end
end
