Prospectus.extra_dep('gitlab_api', 'keylime')
Prospectus.extra_dep('gitlab_api', 'gitlab')

module LogCabin
  module Modules
    ##
    # Provide an api method for modules to query GitLab
    module GitlabApi
      def gitlab_api
        @gitlab_api ||= Gitlab.client(
          endpoint: gitlab_endpoint,
          private_token: gitlab_token
        )
      end

      def gitlab_slug(repo)
        repo.sub('/', '%2F')
      end

      private

      def gitlab_token
        @gitlab_token ||= Keylime.new(
          server: gitlab_endpoint,
          account: 'prospectus'
        ).get!('GitLab API token (https://gitlab.com/profile/account)').password
      end

      def gitlab_endpoint
        @gitlab_endpoint ||= 'https://gitlab.com/api/v3'
      end
    end
  end
end
