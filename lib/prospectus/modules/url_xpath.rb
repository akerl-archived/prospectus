begin
  require 'nokogiri'
rescue LoadError
  raise('The url_xpath module requires the nokogiri gem')
end

require 'open-uri'

module LogCabin
  module Modules
    ##
    # Pull state from a GitHub tag
    module UrlXpath
      def load!
        fail('No url provided') unless @url
        fail('No xpath provided') unless @xpath
        text = parse_page
        if @find
          fail('Text does not match regex') unless text.match(@find)
          text.sub!(@find, @replace)
        end
        @state.value = text
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

      def regex(find, replace = '\1')
        @find = find
        @replace = replace
      end
    end
  end
end
