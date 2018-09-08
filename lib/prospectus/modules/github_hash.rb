module LogCabin
  module Modules
    ##
    # Pull state from the latest GitHub commit
    module GithubHash
      include Prospectus.helpers.find(:github_api)

      def load
        raise('No repo specified') unless @repo
        @branch ||= 'master'
        @long ? hash : hash.slice(0, 7)
      end

      private

      def hash
        @hash ||= github_api.branch(@repo, @branch).commit.sha
      end

      def branch(value)
        @branch = value
      end

      def long
        @long = true
      end
    end
  end
end
