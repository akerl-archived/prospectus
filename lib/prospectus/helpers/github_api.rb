Prospectus.extra_dep('github_api', 'octoauth')

module LogCabin
  module Modules
    ##
    # Provide an api method for modules to query GitHub
    module GithubApi
      def github_api
        cached_clients[@endpoint] ||= Octokit::Client.new(octokit_args)
      end

      private

      def octokit_args
        args = {
          access_token: auth.token,
          auto_paginate: true
        }
        args[:api_endpoint] = @endpoint if @endpoint
        args
      end

      private

      def cached_clients
        @cached_clients ||= {}
      end

      def auth
        @auth ||= Octoauth.new(octoauth_args)
      end

      def octoauth_args
        args = {
          note: 'Prospectus',
          file: :default,
          autosave: true
        }
        args[:api_endpoint] = @endpoint if @endpoint
        args
      end

      def repo(value)
        @repo = value
      end

      def endpoint(value)
        @endpoint = value
      end
    end
  end
end
