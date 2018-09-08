module LogCabin
  module Modules
    ##
    # Pull state from a git hash
    module GitHash
      include Prospectus.helpers.find(:chdir)

      def load
        chdir_helper do
          short_arg = @long ? '' : '--short'
          hash = `git rev-parse #{short_arg} HEAD 2>/dev/null`.chomp
          raise('No hash found') if hash.empty?
          hash
        end
      end

      private

      def long
        @long = true
      end
    end
  end
end
