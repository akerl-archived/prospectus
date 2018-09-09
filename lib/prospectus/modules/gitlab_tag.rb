require 'json'
require 'open-uri'

module LogCabin
  module Modules
    ##
    # Pull state from a Gitlab tag
    module GitlabTag
      include Prospectus.helpers.find(:regex)
      include Prospectus.helpers.find(:gitlab_api)
      include Prospectus.helpers.find(:filter)

      def load!
        raise('No repo specified') unless @repo
        @state.value = regex_helper(tag)
      end

      private

      def tags
        @tags ||= gitlab_api.tags(@repo).sort do |*points|
          dates = points.map { |x| Date.parse(x.commit.committed_date) }
          dates.last <=> dates.first
        end.map(&:name)
      end

      def tag
        @tag = filter_helper(tags).first
      end
    end
  end
end
