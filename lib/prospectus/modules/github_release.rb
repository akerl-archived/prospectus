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

      def load
        raise('No repo specified') unless @repo
        regex_helper(release)
      end

      private

      def allow_prerelease
        @allow_prerelease = true
      end

      def release
        return @release if @release
        releases = github_api.releases(@repo)
        releases.reject!(&:draft)
        releases.reject!(&:prerelease) unless @allow_prerelease
        @release = filter_helper(releases.map(&:tag_name)).first
      end
    end
  end
end
