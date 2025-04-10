# XCard

XCard is a Go-based service for the MiSTer FPGA that manages and launches games stored in the XCard storage format. The XCard format is designed to organize a set of related files (e.g., ROMs, assets) and their metadata within a folder or volume, described by a manifest file. This service allows MiSTer users to launch games from XCard folders, leveraging metadata for enhanced presentation and configuration, such as menu names, marquee bitmaps, controller settings, and screen filters.

## Features
*XCard Format Support*: Reads a manifest file in an XCard folder to identify and manage game files and metadata.

*Game Launching*: Launches games on MiSTer FPGA by interfacing with the system's core-loading mechanism.

*Metadata Handling*: Supports optional metadata like:
* Display names for menus or screens.
* Marquee definitions for bitmap display.
* Controller configurations.
* Preferred screen filters.
* Systemd Service: Runs as a background service on MiSTer, enabling automatic or triggered game launches.
* Extensible: Easily adaptable for additional metadata or MiSTer-specific features.

## XCard Format
The XCard format is a folder or volume containing:

Game Files: ROMs, assets, or other files required by an FPGA core (e.g., .nes, .sfc).

Manifest File: A file (e.g., manifest.json) describing the contents and metadata.

Manifest Example:
```json
{
  "name": "Panzer Dragoon Saga",
  "core": "SATURN",
  "files": [
    {
      "path": "PanzerDragoonSaga.saturn",
      "type": "rom"
    },
    {
      "path": "cover.png",
      "type": "asset"
    }
  ],
  "metadata": {
    "display_name": "Panzer Dragoon Saga",
    "marquee": {
      "bitmap_path": "marquee.png",
      "width": 320,
      "height": 32
    },
    "controller": {
      "type": "sega-saturn",
      "mapping": {
        "a": "button1",
        "b": "button2"
      }
    },
    "screen_filter": "scanlines"
  }
}
```

### Prerequisites
* Go: Version 1.16 or higher for development and compilation.
* MiSTer FPGA: Running a Linux-based OS with systemd (standard MiSTer setup).
* Storage: SD card with XCard folders at /media/fat/xcards/ (configurable).
* Access: SSH or physical access to MiSTer for deployment.

## Installation
#### Clone the Repository:
```bash
git clone https://github.com/yourusername/xcard.git
cd xcard
```
#### Build the Binary:

Cross-compile for MiSTer’s ARM architecture:
```bash
GOOS=linux GOARCH=arm GOARM=7 go build -o xcard
```
#### Deploy to MiSTer:

Copy the binary to MiSTer:
```bash
scp xcard root@<mister-ip>:/usr/local/bin/
```

#### Set Up XCard Folders:
Place XCard folders (each with a manifest.json and game files) in /media/fat/xcards/. 

Example:
```
/media/fat/xcards/SuperMario/
  ├── manifest.json
  ├── SuperMario.nes
  ├── cover.png
  └── marquee.png
```
#### Configure systemd Service:

Create /etc/systemd/system/xcard.service on MiSTer:
```ini
[Unit]
Description=XCard Game Launcher Service for MiSTer FPGA
After=network.target

[Service]
ExecStart=/usr/local/bin/xcard -xcard /media/fat/xcards/SuperMario
Restart=always
User=root
Group=root

[Install]
WantedBy=multi-user.target
```
* Replace /media/fat/xcards/SuperMario with your XCard folder path.

#### Enable and start the service:
``` bash
systemctl daemon-reload
systemctl enable xcard
systemctl start xcard
```

