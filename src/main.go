package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"syscall"
	"time"

	"github.com/micmonay/keybd_event"
)

func altTab(w http.ResponseWriter, r *http.Request) {
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(fmt.Appendf(nil, "%v", err))
		return
	}
	kb.HasALT(true)
	kb.SetKeys(keybd_event.VK_TAB)
	err = kb.Launching()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Alt Tabbed"))
}

func ctrlTab(w http.ResponseWriter, r *http.Request) {
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(fmt.Appendf(nil, "%v", err))
		return
	}
	kb.HasCTRL(true)
	kb.SetKeys(keybd_event.VK_TAB)
	err = kb.Launching()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Ctrl Tabbed"))
}

func altEsc(w http.ResponseWriter, r *http.Request) {
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(fmt.Appendf(nil, "%v", err))
		return
	}
	kb.HasALT(true)
	kb.SetKeys(keybd_event.VK_ESC)
	err = kb.Launching()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Alt Escaped"))
}

func altf4(w http.ResponseWriter, r *http.Request) {
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(fmt.Appendf(nil, "%v", err))
		return
	}
	kb.HasALT(true)
	kb.SetKeys(keybd_event.VK_F4)
	err = kb.Launching()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Alt F4-ed"))
}

func ctrlw(w http.ResponseWriter, r *http.Request) {
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(fmt.Appendf(nil, "%v", err))
		return
	}
	kb.HasCTRL(true)
	kb.SetKeys(keybd_event.VK_W)
	err = kb.Launching()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Ctrl W-ed"))
}

func pressF(w http.ResponseWriter, r *http.Request) {
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(fmt.Appendf(nil, "%v", err))
		return
	}
	kb.SetKeys(keybd_event.VK_F)
	err = kb.Launching()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("F-ed"))
}

func pressSpace(w http.ResponseWriter, r *http.Request) {
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(fmt.Appendf(nil, "%v", err))
		return
	}
	kb.SetKeys(keybd_event.VK_SPACE)
	err = kb.Launching()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Space-ed"))
}

func pressLeft(w http.ResponseWriter, r *http.Request) {
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(fmt.Appendf(nil, "%v", err))
		return
	}
	kb.SetKeys(37) // VK_LEFT
	err = kb.Launching()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Left arrowed"))
}

func pressUp(w http.ResponseWriter, r *http.Request) {
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(fmt.Appendf(nil, "%v", err))
		return
	}
	kb.SetKeys(38) // VK_UP
	err = kb.Launching()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Up arrowed"))
}

func pressDown(w http.ResponseWriter, r *http.Request) {
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(fmt.Appendf(nil, "%v", err))
		return
	}
	kb.SetKeys(40) // VK_DOWN
	err = kb.Launching()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Down arrowed"))
}

func pressRight(w http.ResponseWriter, r *http.Request) {
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(fmt.Appendf(nil, "%v", err))
		return
	}
	kb.SetKeys(39) // VK_RIGHT
	err = kb.Launching()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Right arrowed"))
}

func pressK(w http.ResponseWriter, r *http.Request) {
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(fmt.Appendf(nil, "%v", err))
		return
	}
	kb.SetKeys(keybd_event.VK_K)
	err = kb.Launching()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("K-ed"))
}

func playPause(w http.ResponseWriter, r *http.Request) {
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(fmt.Appendf(nil, "%v", err))
		return
	}
	kb.SetKeys(179) // VK_MEDIA_PLAY_PAUSE
	err = kb.Launching()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Played/Paused"))
}

func showDesktop(w http.ResponseWriter, r *http.Request) {
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(fmt.Appendf(nil, "%v", err))
		return
	}
	kb.SetKeys(keybd_event.VK_D)
	kb.HasSuper(true)
	err = kb.Launching()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Desktoped"))
}

func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[%s] %s %s from %s\n", time.Now().Format(time.RFC1123), r.Method, r.URL.Path, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}

func main() {
	portFlag := flag.Int("port", 8075, "Port to run the server on")
	debugFlag := flag.Bool("debug", false, "Enable debug stats")
	flag.Parse()

	if *debugFlag {
		go func() {
			for range time.Tick(16 * time.Second) {
				runtime.GC()
				var m runtime.MemStats
				runtime.ReadMemStats(&m)
				fmt.Printf("Heap: %d MB  GC pauses: last=%d ns\n",
					m.HeapAlloc>>20, m.PauseTotalNs)
			}
		}()
	}

	isWin := strings.Contains(strings.ToLower(runtime.GOOS), "windows")

	http.Handle("/", logMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Running"))
	})))
	http.Handle("/alt-tab", logMiddleware(http.HandlerFunc(altTab)))
	http.Handle("/ctrl-tab", logMiddleware(http.HandlerFunc(ctrlTab)))
	http.Handle("/alt-esc", logMiddleware(http.HandlerFunc(altEsc)))
	http.Handle("/alt-f4", logMiddleware(http.HandlerFunc(altf4)))
	http.Handle("/ctrl-w", logMiddleware(http.HandlerFunc(ctrlw)))
	http.Handle("/f", logMiddleware(http.HandlerFunc(pressF)))
	http.Handle("/space", logMiddleware(http.HandlerFunc(pressSpace)))
	http.Handle("/left", logMiddleware(http.HandlerFunc(pressLeft)))
	http.Handle("/up", logMiddleware(http.HandlerFunc(pressUp)))
	http.Handle("/down", logMiddleware(http.HandlerFunc(pressDown)))
	http.Handle("/right", logMiddleware(http.HandlerFunc(pressRight)))
	http.Handle("/k", logMiddleware(http.HandlerFunc(pressK)))
	http.Handle("/desktop", logMiddleware(http.HandlerFunc(showDesktop)))
	if isWin {
		http.Handle("/play-pause", logMiddleware(http.HandlerFunc(playPause)))
	}

	var ip string
	iface, err := net.InterfaceByName("Wi-Fi")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	}
	addrs, err := iface.Addrs()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && ipnet.IP.To4() != nil {
			ip = ipnet.IP.String()
		}
	}

	endpts := "Endpoints:\n" +
		"Alt+Tab: /alt-tab\n" +
		"Ctrl+Tab: /ctrl-tab\n" +
		"Alt+Esc: /alt-esc\n" +
		"Alt+F4: /alt-f4\n" +
		"Ctrl+W: /ctrl-w\n" +
		"F: /f\n" +
		"Space: /space\n" +
		"Left arrow: /left\n" +
		"Up arrow: /up\n" +
		"Down arrow: /down\n" +
		"Right arrow: /right\n" +
		"K: /k\n" +
		"Desktop: /desktop\n"
	if isWin {
		endpts += "Play/Pause: /play-pause\n"
	}

	fmt.Print(endpts + "\n")

	if ip != "" {
		fmt.Printf("Running on http://%s:%d\n\n", ip, *portFlag)
	} else {
		fmt.Printf("Running on port: %d\n\n", *portFlag)
	}

	runtime.GC()

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", *portFlag),
		Handler: nil,
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Fprintf(os.Stderr, "Server error: %v\n", err)
		}
	}()

	<-sigChan
	fmt.Println("\nShutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "Server shutdown error: %v\n", err)
		os.Exit(1)
	}
}
