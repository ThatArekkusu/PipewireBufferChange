# PipewireBufferChange

## Overview
PipewireBufferChange is a project designed to automate buffer size changes in the pipewire directory.

## Installation
To install PipewireBufferChange, follow these steps:

1. Clone the repository:
    ```bash
    git clone https://github.com/ThatArekkusu/PipewireBufferChange.git
    ```
2. Navigate to the project directory:
    ```bash
    cd PipewireBufferChange
    ```
3. Make the install script executable:
    ```bash
    chmod +x install.sh
    ```

4. Run the installation script:
    ```bash
    ./install.sh
    ```

5. After setting buffer size:
    ```bash
    chmod +x audio.sh
    ```

6. Restart audio drivers (this will be automated in the future)
    ```bash
    ./audio.sh
    ```

## Usage
To use PipewireBufferChange, execute the following command:
```bash
bufferchange run
