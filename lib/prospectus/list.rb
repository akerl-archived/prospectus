module Prospectus
  ##
  # Define list object that contains items
  class List
    def initialize(params = {})
      @options = params
    end

    def items
      @items ||= []
    end

    def check
      items
    end
  end

  ##
  # DSL for wrapping eval of list files
  class ListDSL
    def initialize(list, params)
      @list = list
      @options = params
    end

    def item(&block)
      item = Item.new(@options)
      dsl = ItemDSL.new(item, @options)
      dsl.instance_eval(&block)
      @list.items << item
    end
  end
end
