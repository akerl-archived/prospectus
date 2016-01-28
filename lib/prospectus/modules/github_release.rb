require 'json'
require 'open-uri'

module LogCabin
  module Modules
    ##
    # Pull state from a GitHub release
    module GithubRelease
      include Prospectus.helpers.find(:regex)
      include Prospectus.helpers.find(:github_api)

      def load!
        fail('No repo specified') unless @repo
        @state.value = regex_helper(release)
      end

      private

      def release
        @release ||= github_api.latest_release(@repo).tag_name
      end
    end
  end
end
