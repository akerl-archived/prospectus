require 'json'
require 'open-uri'

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
        @tag ||= gitlab_api.tags(gitlab_slug(@repo), per_page: 1).first.name
      end

      def repo(value)
        @repo = value
      end
    end
  end
end
