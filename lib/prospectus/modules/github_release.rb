require 'json'
require 'open-uri'

module LogCabin
  module Modules
    ##
    # Pull state from a GitHub release
    module GithubRelease
      include Prospectus.helpers.find(:regex)
      include Prospectus.helpers.find(:github_api)
      include Prospectus.helpers.find(:filter)

      def load!
        raise('No repo specified') unless @repo
        @state.value = regex_helper(release)
      end

      private

      def release
        return @release if @release
        releases = github_api.releases(@repo).map(&:tag_name)
        @release = filter_helper(releases).first
      end
    end
  end
end
