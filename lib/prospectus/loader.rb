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

    def load
      return run_file(@options, @file) if File.exist? @file
      raise("No #{@file}/#{@dir} found") unless Dir.exist? @dir
      files = Dir.glob(@dir + '/*')
      raise('No files in ' + @dir) if files.empty?
      files.map { |x| run_file(@options, x, true) }.flatten
    end

    private

    def run_file(params, file, suffix_file = false)
      options = { file: file, suffix_file: suffix_file }.merge(params)
      Prospectus.load_from_file(options).check
    rescue RuntimeError
      puts "Failed parsing #{Dir.pwd}/#{file}"
      raise
    end
  end
end
