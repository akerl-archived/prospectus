Prospectus.extra_dep('gemnasium', 'netrc')
require 'net/http'
require 'json'

module LogCabin
  module Modules
    ##
    # Pull state from Gemnasium
    module Gemnasium
      def load!
        raise('No slug provided') unless @slug
        @state.value = parse_api
      end

      private

      def parse_api
        %w(red yellow).each do |color|
          return color if colors.include? color
        end
        'green'
      end

      def colors
        @colors ||= api_data.map { |x| x['first_level'] && x['color'] }
      end

      def api_data
        return @api_data if @api_data
        resp = Net::HTTP.start(uri.host, uri.port, use_ssl: true) do |http|
          request = Net::HTTP::Get.new uri.request_uri
          request.basic_auth(*creds)
          response = http.request(request)
          JSON.parse(response.body)
        end
        @api_data = validate_response(resp)
      end

      def validate_response(resp)
        return resp if resp.is_a? Array
        raise("API lookup on gemnasium failed: #{resp['message']}")
      end

      def creds
        netrc[site] || prompt_for_creds
      end

      def prompt_for_creds
        puts 'Please enter your API key from https://gemnasium.com/settings >> '
        resp = gets.chomp
        netrc[site] = 'X', resp
        netrc.save
        netrc[site]
      end

      def netrc
        @netrc ||= Netrc.read
      end

      def uri
        @uri ||= URI("https://api.gemnasium.com/projects/#{@slug}/dependencies")
      end

      def site
        'gemnasium.com'
      end

      def slug(value)
        @slug = value
      end
    end
  end
end
