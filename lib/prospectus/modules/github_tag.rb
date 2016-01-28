require 'json'
require 'open-uri'

module LogCabin
  module Modules
    ##
    # Pull state from a GitHub tag
    module GithubTag
      include Prospectus.helpers.find(:regex)
      include Prospectus.helpers.find(:github_api)

      def load!
        fail('No repo specified') unless @repo
        @state.value = regex_helper(tag)
      end

      private

      def tag
        @tag ||= github_api.tags(@repo).first.name
      end
    end
  end
end
