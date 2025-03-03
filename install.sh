#!/bin/bash

# Build the Go programs
go build -o PipewireBufferChange main.go
go build -o PipewireBufferStatus src/status/status.go
go build -o PipewireBufferDirectory src/directory/directory.go  

# Move the binaries to /usr/local/bin for global access
sudo mv PipewireBufferChange /usr/local/bin/PipewireBufferChange
sudo mv PipewireBufferStatus /usr/local/bin/PipewireBufferStatus
sudo mv PipewireBufferDirectory /usr/local/bin/PipewireBufferDirectory

# Create the bufferchange script
echo '#!/bin/bash' | sudo tee /usr/local/bin/bufferchange
echo 'if [ "$1" == "run" ]; then' | sudo tee -a /usr/local/bin/bufferchange
echo '  sudo /usr/local/bin/PipewireBufferChange' | sudo tee -a /usr/local/bin/bufferchange
echo 'elif [ "$1" == "status" ]; then' | sudo tee -a /usr/local/bin/bufferchange
echo '  sudo /usr/local/bin/PipewireBufferStatus' | sudo tee -a /usr/local/bin/bufferchange 
echo 'elif [ "$1" == "directory" ]; then' | sudo tee -a /usr/local/bin/bufferchange
echo '  sudo /usr/local/bin/PipewireBufferDirectory' | sudo tee -a /usr/local/bin/bufferchange 
echo 'else' | sudo tee -a /usr/local/bin/bufferchange
echo '  echo "Usage: bufferchange {run|status}"' | sudo tee -a /usr/local/bin/bufferchange
echo 'fi' | sudo tee -a /usr/local/bin/bufferchange

# Make the bufferchange script executable
sudo chmod +x /usr/local/bin/bufferchange
sudo chmod +x /usr/local/bin/PipewireBufferStatus 
sudo chmod +x /usr/local/bin/PipewireBufferDirectory
sudo chmod +x ./audio.sh

echo "Installation complete. You can now use the 'bufferchange run', 'bufferchange status' or bufferchange directory commands."
