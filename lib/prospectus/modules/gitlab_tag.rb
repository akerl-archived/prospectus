require 'json'
require 'open-uri'

Prospectus.extra_dep('gitlab_tag', 'version_sorter')

module LogCabin
  module Modules
    ##
    # Pull state from a Gitlab tag
    module GitlabTag
      include Prospectus.helpers.find(:regex)
      include Prospectus.helpers.find(:gitlab_api)

      def load!
        fail('No repo specified') unless @repo
        @state.value = regex_helper(tag)
      end

      private

      def tag
        VersionSorter.rsort(tags).first
      end

      def tags
        @tags ||= gitlab_api.tags(gitlab_slug(@repo), per_page: 1).map(&:name)
      end

      def repo(value)
        @repo = value
      end
    end
  end
end
