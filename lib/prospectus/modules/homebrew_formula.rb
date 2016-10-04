require 'json'

module LogCabin
  module Modules
    ##
    # Pull state from a homebrew formula file
    module HomebrewFormula
      def load!
        raise('No name specified') unless @name
        cask_file = "Formula/#{@name}.rb"
        output = `brew info --json=v1 #{cask_file}`
        @state.value = JSON.parse(output).first.dig('versions', 'stable')
      end

      private

      def name(value)
        @name = value
      end
    end
  end
end
