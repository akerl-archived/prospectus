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
      all, good_only = @options.values_at(:all, :good_only)
      items.select do |x|
        match = x.expected == x.actual
        true if all || (!match ^ good_only)
      end
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
      item.list.items.each do |x|
        x.prefix item.name
        @list.items << x
      end
    end
  end
end
