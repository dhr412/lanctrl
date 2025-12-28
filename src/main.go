package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
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
	kb.Press()
	kb.Release()
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
	kb.Press()
	kb.Release()
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
	kb.Press()
	kb.Release()
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
	kb.Press()
	kb.Release()
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
	kb.Press()
	kb.Release()
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
	kb.Press()
	kb.Release()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("F pressed"))
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
	kb.Press()
	kb.Release()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Space pressed"))
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
	kb.Press()
	kb.Release()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("K pressed"))
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
	kb.Press()
	kb.Release()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Play/Paused"))
}

func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[%s] %s %s from %s\n", time.Now().Format(time.RFC1123), r.Method, r.URL.Path, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}

func main() {
	portFlag := flag.Int("port", 8075, "Port to run the server on")
	flag.Parse()

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
	http.Handle("/k", logMiddleware(http.HandlerFunc(pressK)))
	http.Handle("/play-pause", logMiddleware(http.HandlerFunc(playPause)))

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

	fmt.Print(
		"Endpoints:\n" +
			"Alt+Tab: /alt-tab\n" +
			"Ctrl+Tab: /ctrl-tab\n" +
			"Alt+Esc: /alt-esc\n" +
			"Alt+F4: /alt-f4\n" +
			"Ctrl+W: /ctrl-w\n" +
			"F: /f\n" +
			"Space: /Space\n" +
			"K: /k\n" +
			"Play/Pause (Windows only): /play-pause\n" +
			"\n",
	)

	if ip != "" {
		fmt.Printf("Running on http://%s:%d\n\n", ip, *portFlag)
	} else {
		fmt.Printf("Running on port: %d\n\n", *portFlag)
	}

	http.ListenAndServe(fmt.Sprintf(":%d", *portFlag), nil)
}
