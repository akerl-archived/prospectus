require 'json'
require 'open-uri'

module LogCabin
  module Modules
    ##
    # Pull state from a GitHub tag
    module GithubTag
      include Prospectus.helpers.find(:regex)

      def load!
        tag = latest_tag
        @state.value = regex_helper(tag)
      end

      private

      def latest_tag
        JSON.load(open(url)).first['name']
      end

      def url
        fail('No repo specified') unless @repo
        @url ||= "https://api.github.com/repos/#{@repo}/tags"
      end

      def repo(value)
        @repo = value
      end
    end
  end
end
