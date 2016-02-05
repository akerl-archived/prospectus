module LogCabin
  module Modules
    ##
    # Pull state from a homebrew cask file
    module HomebrewCask
      def load!
        raise('No name specified') unless @name
        cask_file = "Casks/#{@name}.rb"
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
