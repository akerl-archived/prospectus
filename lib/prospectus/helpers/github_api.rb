Prospectus.extra_dep('github_api', 'octoauth')

module LogCabin
  module Modules
    ##
    # Provide an api method for modules to query GitHub
    module GithubApi
      def github_api
        @github_api ||= Octokit::Client.new(
          access_token: auth.token,
          auto_paginate: true
        )
      end

      private

      def auth
        @auth ||= Octoauth.new(
          note: 'Prospectus',
          file: :default,
          autosave: true
        )
      end

      def repo(value)
        @repo = value
      end
    end
  end
end
