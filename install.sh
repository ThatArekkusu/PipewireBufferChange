#!/bin/bash

# Build the Go program
go build -o PipewireBufferChange main.go

# Move the binary to /usr/local/bin for global access
sudo mv PipewireBufferChange /usr/local/bin/PipewireBufferChange

# Create the bufferchange script
echo '#!/bin/bash' | sudo tee /usr/local/bin/bufferchange
echo 'if [ "$1" == "run" ]; then' | sudo tee -a /usr/local/bin/bufferchange
echo '  sudo /usr/local/bin/PipewireBufferChange' | sudo tee -a /usr/local/bin/bufferchange
echo 'else' | sudo tee -a /usr/local/bin/bufferchange
echo '  echo "Usage: bufferchange run"' | sudo tee -a /usr/local/bin/bufferchange
echo 'fi' | sudo tee -a /usr/local/bin/bufferchange

# Make the bufferchange script executable
sudo chmod +x /usr/local/bin/bufferchange

echo "Installation complete. You can now use the 'bufferchange run' command."