module LogCabin
  module Modules
    ##
    # Pull state from a git tag
    module GitTag
      include Prospectus.helpers.find(:regex)

      def load
        tag = `git describe --tags --abbrev=0 2>/dev/null`.chomp
        raise('No tags found') if tag.empty?
        regex_helper(tag)
      end
    end
  end
end
