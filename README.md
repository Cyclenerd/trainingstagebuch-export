# Trainingstagebuch.org Export

Create a local copy of your [Trainingstagebuch.org](https://trainingstagebuch.org/) workout data.

![Trainingstagebuch.org Logo](https://trainingstagebuch.org/static/images/apple-touch-icon.png)

This command line application creates an export of your recorded activities at Trainingstagebuch.org in both GPX and CSV formats.

## Setup

1. Create a directory for your Trainingstagebuch.org export
1. [Download](https://github.com/Cyclenerd/trainingstagebuch-export/releases/latest) the trainingstagebuch-export application for your operating system and CPU architecture. If you are unsure, usually `x86_64` will work.
    * <img src="https://www.microsoft.com/favicon.ico" width=16 height=16/> **Windows**
        * [x86_64](https://github.com/Cyclenerd/trainingstagebuch-export/releases/latest/download/trainingstagebuch-export-windows-x86_64.exe) Intel or AMD 64-Bit CPU
        * [arm64](https://github.com/Cyclenerd/trainingstagebuch-export/releases/latest/download/trainingstagebuch-export-windows-arm64.exe) Arm-based 64-Bit CPU
    * <img src="https://developer.apple.com/favicon.ico" width=16 height=16/> **macOS**
        * [x86_64](https://github.com/Cyclenerd/trainingstagebuch-export/releases/latest/download/trainingstagebuch-export-macos-x86_64) Intel 64-bit CPU
        * [arm64](https://github.com/Cyclenerd/trainingstagebuch-export/releases/latest/download/trainingstagebuch-export-macos-arm64) Apple silicon 64-bit CPU (M1, M2...)
    * <img src="https://www.kernel.org/theme/images/logos/favicon.png" width=16 height=16/> **Linux**
        * [x86_64](https://github.com/Cyclenerd/trainingstagebuch-export/releases/latest/download/trainingstagebuch-export-linux-x86_64) Intel or AMD 64-Bit CPU
        * [arm64](https://github.com/Cyclenerd/trainingstagebuch-export/releases/latest/download/trainingstagebuch-export-linux-x86_64) Arm-based 64-Bit CPU (i.e. Raspberry Pi)
1. Rename it to:
    * `trainingstagebuch-export` (macOS, Linux)
    * `trainingstagebuch-export.exe` (Windows)
1. Go to <https://trainingstagebuch.org/login/sso> to get your private session key
    * Copy the key between `<session>` and `</session>`
1. Run the application - it will prompt you to enter your session key

## Usage

This project require you to use a command-line interface.
Don't worry, it's easier than it looks!
Here's how to open one:

## Windows (PowerShell)

1. Press the <kbd>Windows key</kbd> + <kbd>X</kbd>.
1. Choose "Windows PowerShell".
1. Type `cd` followed by the path to your folder (e.g., `cd C:\Users\YourName\Documents\Trainingstagebuch`) and press Enter.
1. Type `trainingstagebuch-export.exe` and press Enter.

## macOS (Terminal)

1. Press <kbd>Command</kbd> + <kbd>Space</kbd>.
1. Type "Terminal" and press Enter.
1. Type `cd` followed by the path to your folder (e.g., `cd /Users/YourName/Documents/Trainingstagebuch`) and press Enter.
1. Type `chmod +x trainingstagebuch-export` and press Enter. (This makes the tool work).
1. Type `./trainingstagebuch-export` and press Enter.

## License

This project is licensed under the [GNU General Public License v3](./LICENSE).