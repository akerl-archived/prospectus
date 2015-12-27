module Prospectus
  ##
  # Define item objects that defined expected vs actual state
  class Item
    def initialize(params = {})
      @options = params
    end

    def name
      @name ||= File.basename Dir.pwd
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

    private

    def state(name, &block)
      state = State.new(@options)
      dsl = StateDSL.new(state, @options)
      dsl.instance_eval(&block)
      dsl.load!
      @item.instance_variable_set(name, state.value)
    end
  end
end
