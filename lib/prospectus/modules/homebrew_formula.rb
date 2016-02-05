module LogCabin
  module Modules
    ##
    # Pull state from a homebrew formula file
    module HomebrewFormula
      def load!
        raise('No name specified') unless @name
        cask_file = "Formula/#{@name}.rb"
        version_regex = /^\s+version ['"](.+)['"]$/
        Prospectus::State.from_block(@option, @state) do
          grep
          file cask_file
          regex version_regex
        end
      end

      private

      def name(value)
        @name = value
      end
    end
  end
end
