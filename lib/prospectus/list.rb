module Prospectus
  ##
  # Define list object that contains items
  class List
    DEFAULT_PATH = './.prospectus'

    def initialize(params = {})
      @options = params
    end

    def items
      @items ||= []
    end

    def check
      items
    end

    def self.load(params = {})
      params[:directory] ||= DEFAULT_PATH
      list = new
      dsl = Prospectus::ListDSL.new(list)
      dsl.instance_eval(File.read(params[:directory]), path)
      list
    end
  end

  ##
  # DSL for wrapping eval of list files
  class ListDSL
    def initialize(list)
      @list = list
    end

    def item(&block)
      item = Prospectus::Item.new
      dsl = Prospectus::ItemDSL.new(item)
      dsl.instance_eval(&block)
      @list.items << item
    end
  end
end
