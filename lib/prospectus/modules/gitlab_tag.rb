require 'json'
require 'open-uri'

module LogCabin
  module Modules
    ##
    # Pull state from a Gitlab tag
    module GitlabTag
      include Prospectus.helpers.find(:regex)
      include Prospectus.helpers.find(:gitlab_api)

      def load!
        raise('No repo specified') unless @repo
        @state.value = regex_helper(tag)
      end

      private

      def tag
        @tag ||= gitlab_api.tags(gitlab_slug(@repo)).sort do |*points|
          dates = points.map { |x| Date.parse(x.commit.committed_date) }
          dates.last <=> dates.first
        end.first.name
      end
    end
  end
end
