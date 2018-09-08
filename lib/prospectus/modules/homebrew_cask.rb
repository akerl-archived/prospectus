module LogCabin
  module Modules
    ##
    # Pull state from a homebrew cask file
    module HomebrewCask
      def load
        raise('No name specified') unless @name
        cask_file = "Casks/#{@name}.rb"
        output = `brew cask _stanza version #{cask_file}`
        output.strip
      end

      private

      def name(value)
        @name = value
      end
    end
  end
end
