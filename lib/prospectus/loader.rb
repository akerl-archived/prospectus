module Prospectus
  DEFAULT_FILE = './.prospectus'.freeze

  ##
  # Helper for loading prospectus from the current directory
  class Loader
    def initialize(params = {})
      @options = params
      @file = params[:file] || DEFAULT_FILE
      @dir = @file + '.d'
    end

    def list
      return @list if @list
      items = load_file_or_dir
      params = { items: items }.merge(@options)
      @list = Prospectus::List.new(params)
    end

    def check
      list.check
    end

    private

    def load_file_or_dir
      return parse_file(@file) if File.exist? @file
      raise("No #{@file}/#{@dir} found") unless Dir.exist? @dir
      files = Dir.glob(@dir + '/*')
      raise('No files in ' + @dir) if files.empty?
      files.map { |x| parse_file(x, true) }.flatten
    end

    def parse_file(file, suffix_file = false)
      params = { file: file, suffix_file: suffix_file }.merge(@options)
      Prospectus::List.from_file(params).items
    rescue RuntimeError
      puts "Failed parsing #{Dir.pwd}/#{file}"
      raise
    end
  end
end
