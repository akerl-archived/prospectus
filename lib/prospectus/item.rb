require 'json'

module Prospectus
  ##
  # Define item objects that defined expected vs actual state
  class Item
    attr_reader :list

    def initialize(params = {})
      @options = params
      @list = List.new(params)
    end

    def name
      return @name if @name
      @name = File.basename Dir.pwd
      @name << "::#{File.basename @options[:file]}" if @options[:suffix_file]
      @name
    end

    def prefix(value)
      raise('Name not set for sub-item') unless @name
      @name = value + '::' + @name
    end

    def noop
      @expected = 'noop'
      @actual = 'noop'
    end

    def expected
      @expected || raise("No expected state was loaded for #{name}")
    end

    def actual
      @actual || raise("No actual state was loaded for #{name}")
    end

    def to_json(_ = {})
      { name: name, expected: expected, actual: actual }.to_json
    end
  end

  ##
  # DSL for wrapping eval of item files
  class ItemDSL
    def initialize(item, params)
      @item = item
      @options = params
    end

    def name(value)
      @item.instance_variable_set(:@name, value)
    end

    def noop
      @item.noop
    end

    def expected(&block)
      state(:@expected, &block)
    end

    def actual(&block)
      state(:@actual, &block)
    end

    def deps(&block)
      dsl = ListDSL.new(@item.list, @options)
      dsl.instance_eval(&block)
    end

    private

    def state(name, &block)
      state = Prospectus::State.from_block(@options, &block)
      @item.instance_variable_set(name, state.value)
    rescue => e
      raise("Failed to set #{name} state for #{@item.name}: #{e.message}")
    end
  end
end
