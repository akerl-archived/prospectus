module Prospectus
  ##
  # Pull state from a git tag
  module GitTag
    def load!
      tag = `git describe --tags`.chomp
      if @find
        fail('Tag does not match regex') unless tag.match(@find)
        tag.sub!(@find, @replace)
      end
      @state.value = tag
    end

    private

    def regex(find, replace = '\1')
      @find = find
      @replace = replace
    end
  end
end
