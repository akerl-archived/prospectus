Prospectus.extra_dep('url_xpath', 'nokogiri')

require 'open-uri'

module LogCabin
  module Modules
    ##
    # Pull state from a GitHub tag
    module UrlXpath
      include Prospectus.helpers.find(:regex)

      def load
        raise('No url provided') unless @url
        raise('No xpath provided') unless @xpath
        text = parse_page
        regex_helper(text)
      end

      private

      def parse_page
        page = open(@url) # rubocop:disable Security/Open
        html = Nokogiri::HTML(page) { |config| config.strict.nonet }
        html.xpath(@xpath).text.strip
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
