module LogCabin
  module Modules
    ##
    # Change directory before running module
    module Chdir
      def chdir_helper(&block)
        @dir ||= '.'
        Dir.chdir(@dir) { block.call }
      end

      private

      def dir(value)
        @dir = value
      end
    end
  end
end
