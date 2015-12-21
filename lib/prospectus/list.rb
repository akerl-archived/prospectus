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
    def initialize(list)
      @list = list
    end

    def item(&block)
      item = Item.new
      dsl = ItemDSL.new(item)
      dsl.instance_eval(&block)
      @list.items << item
    end
  end
end
