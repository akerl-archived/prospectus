require 'json'
require 'open-uri'

module LogCabin
  module Modules
    ##
    # Pull state from a GitHub tag
    module GithubTag
      include Prospectus.helpers.find(:regex)
      include Prospectus.helpers.find(:github_api)
      include Prospectus.helpers.find(:filter)

      def load
        raise('No repo specified') unless @repo
        regex_helper(tag)
      end

      private

      def tag
        return @tag if @tag
        @tags = filter_helper(github_api.tags(@repo).map { |x| x[:name] }).first
      end
    end
  end
end
