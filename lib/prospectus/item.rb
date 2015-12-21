module Prospectus
  ##
  # Define item objects that defined expected vs actual state
  class Item
    attr_accessor :name, :expected, :actual

    def initialize(params = {})
      @options = params
    end
  end

  ##
  # DSL for wrapping eval of item files
  class ItemDSL
    def initialize(item)
      @item = item
    end

    def name(value)
      @item.name = value
    end

    def expected(&block)
      @item.expected = state(&block)
    end

    def actual(&block)
      @item.actual = state(&block)
    end

    private

    def state(&block)
      state = Prospectus::State.new
      dsl = Prosepectus::StateDSL.new(state)
      dsl.instance_eval(&block)
      state.version
    end
  end
end
