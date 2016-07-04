module LogCabin
  module Modules
    ##
    # Use regex to filter out from a list of matches
    module Filter
      def filter_helper(list)
        return list unless @filter_regex
        new_list = list.select { |x| x =~ @filter_regex }
        return new_list unless new_list.empty?
        raise("No matches found in list: #{@filter_regex} / #{list}")
      end

      private

      def filter(regex)
        @filter_regex = regex
      end
    end
  end
end
