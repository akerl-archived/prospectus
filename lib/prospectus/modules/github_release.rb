require 'json'
require 'open-uri'

module LogCabin
  module Modules
    ##
    # Pull state from a GitHub tag
    module GithubRelease
      include Prospectus.helpers.find(:regex)

      def load!
        release = latest_release
        @state.value = regex_helper(release)
      end

      private

      def latest_release
        JSON.load(open(url))['name']
      end

      def url
        fail('No repo specified') unless @repo
        @url ||= "https://api.github.com/repos/#{@repo}/releases/latest"
      end

      def repo(value)
        @repo = value
      end
    end
  end
end
