# LanCtrl

HTTP-based shortcut execution and remote control for your PC. 

`lanctrl` spins up a lightweight web server that listens for incoming requests and simulates keyboard events on the host machine. This allows you to control your computer (e.g., switch tabs, pause media, close windows) from any device on your local network, such as a phone or tablet.

## Installation

### From Releases

*   Download the latest pre-compiled binary for your platform from the [GitHub Releases](https://github.com/dhr412/lanctrl/releases) page.
*   Run the executable (e.g., `./lanctrl` on Linux/macOS or `lanctrl.exe` on Windows).

### Build from Source

#### Prerequisites
*   [Go](https://go.dev/) (1.20 or later recommended)

1.  Clone the repository:
    ```bash
    git clone https://github.com/dhr412/lanctrl.git
    cd lanctrl
    ```

2.  Install dependencies:
    ```bash
    go mod tidy
    ```

3.  Build the binary:
    ```bash
    go build -o lanctrl src/main.go
    ```

    > Note: If you are compiling on Linux, you may need additional libraries for the `keybd_event` package (e.g., `libx11-dev`).*

## How It Works

The application initializes a local HTTP server (defaulting to port `8075`). Each defined route corresponds to a specific key combination using the [keybd_event](github.com/micmonay/keybd_event) library. When a route is hit, the server simulates the key press on the host OS.

It automatically attempts to detect your local IP address on the "Wi-Fi" interface to make connection easier.

### Usage

Run the application:
```bash
./lanctrl
```

Or specify a custom port:
```bash
./lanctrl -port 9090
```

Send a request to trigger an action (e.g., from a browser or using `curl`):
```bash
curl http://<YOUR_PC_IP>:8075/<endpoint>
```
or
Visit `http://<YOUR_PC_IP>:8075/<endpoint>` in your browser

### Available Endpoints

| Endpoint       | Action Description          |
| :------------- | :-------------------------- |
| `/alt-tab`     | Switch window (`Alt+Tab`)   |
| `/ctrl-tab`    | Switch tab (`Ctrl+Tab`)     |
| `/alt-esc`     | Switch window (`Alt+Esc`)   |
| `/alt-f4`      | Close window (`Alt+F4`)     |
| `/ctrl-w`      | Close tab (`Ctrl+W`)        |
| `/f`           | Press `F` (Fullscreen)      |
| `/space`       | Press `Space`               |
| `/k`           | Press `K` (Youtube Pause)   |
| `/play-pause`  | Media Play/Pause            |
| `/`            | Health Check                |
