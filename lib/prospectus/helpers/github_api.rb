Prospectus.extra_dep('github_api', 'octoauth')

module LogCabin
  module Modules
    ##
    # Provide an api method for modules to query GitHub
    module GithubApi
      def github_api(endpoint = nil)
        cached_clients[endpoint] ||= Octokit::Client.new(octokit_args(endpoint))
      end

      private

      def octokit_args(endpoint)
        args = {
          access_token: auth(endpoint).token,
          auto_paginate: true
        }
        args[:api_endpoint] = endpoint if endpoint
        args
      end

      private

      def cached_clients
        @cached_clients ||= {}
      end

      def auth(endpoint)
        @auth ||= Octoauth.new(octoauth_args)
      end

      def octoauth_args(endpoint)
        args = {
          note: 'Prospectus',
          file: :default,
          autosave: true
        }
        args[:api_endpoint] = endpoint if endpoint
        args
      end

      def repo(value)
        @repo = value
      end
    end
  end
end
