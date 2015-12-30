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
      @name << "::#{File.basename @options[:file]}" if @options[:file]
      @name
    end

    def expected
      @expected || fail("No expected state was loaded for #{name}")
    end

    def actual
      @actual || fail("No actual state was loaded for #{name}")
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

    def expected(&block)
      state(:@expected, &block)
    end

    def actual(&block)
      state(:@actual, &block)
    end

    def deps(&block)
      dsl = ListDSL.new(@item.list, @options)
      dsl.instance_eval(&block)
      @item.list.items.each do |dep|
        dep.instance_variable_set(:@name, "#{@item.name}::#{dep.name}")
      end
    end

    private

    def state(name, &block)
      state = Prospectus::State.from_block(@options, &block)
      @item.instance_variable_set(name, state.value)
    end
  end
end
