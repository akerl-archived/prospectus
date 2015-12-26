module Prospectus
  ##
  # Handles loading of dynamic modules for state checking
  module Module
    BUILTIN_MODULE_PATH = File.join(
      Gem::Specification.find_by_name('prospectus').gem_dir,
      'lib',
      'modules'
    )
    MODULE_LOAD_PATH = [
      BUILTIN_MODULE_PATH
    ]

    class << self
      ##
      # Method for finding modules to load
      def find(method)
        file = find_module_file(method)
        require file
        class_name = parse_class_name(method)
        Prospectus.const_get(class_name)
      end

      private

      ##
      # Convert file name to class name
      # Borrowed with love from Homebrew: http://git.io/vEoDg
      def parse_class_name(method)
        class_name = method.to_s.capitalize
        class_name.gsub!(/[-_.\s]([a-zA-Z0-9])/) { Regexp.last_match[1].upcase }
        class_name.tr!('+', 'x')
        class_name
      end

      ##
      # Check module path for module
      def find_module_file(method)
        MODULE_LOAD_PATH.each do |dir|
          file_path = File.join(dir, "#{method}.rb")
          return file_path if File.exist? file_path
        end
        fail("Module #{method} not found")
      end
    end
  end
end
