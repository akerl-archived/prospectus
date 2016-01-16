module LogCabin
  module Modules
    ##
    # Change directory before running module
    module Chdir
      def chdir_helper
        @dir ||= '.'
        Dir.chdir(@dir) { yield }
      end

      private

      def dir(value)
        @dir = value
      end
    end
  end
end
