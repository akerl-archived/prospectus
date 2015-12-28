Prospectus.extra_dep('url_xpath', 'nokogiri')

require 'open-uri'

module LogCabin
  module Modules
    ##
    # Pull state from a GitHub tag
    module UrlXpath
      include Prospectus.helpers.find(:regex)

      def load!
        fail('No url provided') unless @url
        fail('No xpath provided') unless @xpath
        text = parse_page
        @state.value = regex_helper(text)
      end

      private

      def parse_page
        page = Nokogiri::HTML(open(@url)) { |config| config.strict.nonet }
        page.xpath(@xpath).text
      end

      def url(value)
        @url = value
      end

      def xpath(value)
        @xpath = value
      end
    end
  end
end
