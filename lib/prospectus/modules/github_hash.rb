module LogCabin
  module Modules
    ##
    # Pull state from the latest GitHub commit
    module GithubHash
      def load!
        repo_url = _repo_url
        repo_xpath = _xpath
        Prospectus::State.from_block(@options, @state) do
          url_xpath
          url repo_url
          xpath repo_xpath
        end
      end

      private

      def repo(value)
        @repo = value
      end

      def _repo_url
        fail('No repo provided') unless @repo
        "https://github.com/#{@repo}"
      end

      def _xpath
        '//a[@class="commit-tease-sha"]'
      end
    end
  end
end
