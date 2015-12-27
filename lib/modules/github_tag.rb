require 'json'
require 'open-uri'

module Prospectus
  ##
  # Pull state from a GitHub tag
  module GithubTag
    def load!
      tag = latest_tag
      if @find
        fail('Tag does not match regex') unless tag.match(@find)
        tag.sub!(@find, @replace)
      end
      @state.value = tag
    end

    private

    def latest_tag
      JSON.load(open(url)).first['name']
    end

    def url
      fail('No repo specified') unless @repo
      @url ||= "https://api.github.com/repos/#{@repo}/tags"
    end

    def repo(value)
      @repo = value
    end

    def regex(find, replace = '\1')
      @find = find
      @replace = replace
    end
  end
end
