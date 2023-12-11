package internal

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"net/netip"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"time"
)

const caddyfileTpl = `
{
    auto_https off
    log {
        format json
    }
}
http://localhost:0 {
    respond "Hello, world!"
}
`

type serviceInfo struct {
	Name string `json:"name"`
	// We do not want to omit empty values in the json response
	Version string `json:"version"`
	Health  bool   `json:"health"`
}

type versionsResponse struct {
	Versions []*serviceInfo `json:"versions"`
}

type gateway struct {
	test    *Test
	cmd     *exec.Cmd
	pidfile string
}

func (g *gateway) run() error {

	caddyfileTmpDir := os.TempDir()
	caddyfile := filepath.Join(caddyfileTmpDir, "Caddyfile")
	g.pidfile = filepath.Join(caddyfileTmpDir, "pid")

	if err := os.WriteFile(caddyfile, []byte(caddyfileTpl), 0o666); err != nil {
		return err
	}

	conn, err := net.Listen("tcp", ":0")
	if err != nil {
		return err
	}

	parsedAddr, err := netip.ParseAddrPort(conn.Addr().String())
	if err != nil {
		panic(err)
	}

	buf := bytes.NewBufferString("")
	g.cmd = exec.Command("go", "run", "main.go", "run",
		"-c", caddyfile,
		"--pidfile", g.pidfile,
		"--pingback", fmt.Sprintf("127.0.0.1:%d", parsedAddr.Port()),
		"--watch",
	)
	g.cmd.Stdout = io.MultiWriter(os.Stdout, buf)
	g.cmd.Stderr = io.MultiWriter(os.Stdout, buf)
	g.cmd.Dir = "../../../ee/gateway"
	fmt.Println("will run caddy", time.Now())
	if err := g.cmd.Start(); err != nil {
		panic(err)
	}

	client, err := conn.Accept()
	if err != nil {
		return err
	}

	if err := client.Close(); err != nil {
		return err
	}

	scanner := bufio.NewScanner(buf)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		type line struct {
			ActualAddress string `json:"actual_address"`
		}
		x := line{}
		if err := json.Unmarshal(scanner.Bytes(), &x); err != nil {
			return err
		}
		if x.ActualAddress != "" {
			fmt.Println("Caddy listening on", x.ActualAddress)
			break
		}
	}

	return nil
}

func (g *gateway) stop() error {
	data, err := os.ReadFile(g.pidfile)
	if err != nil {
		return err
	}

	pid, err := strconv.ParseInt(strings.TrimSuffix(string(data), "\n"), 10, 64)
	if err != nil {
		return err
	}

	if err := syscall.Kill(int(pid), syscall.SIGINT); err != nil {
		return err
	}
	if err := g.cmd.Wait(); err != nil {
		return errors.Wrap(err, "stopping gateway")
	}
	return nil
}

func (g gateway) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/versions" {
		res := versionsResponse{
			Versions: []*serviceInfo{
				{
					Name:    "ledger",
					Version: "v2.0.0",
				},
				// If needed, add other services version
			},
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			panic(err)
		}
		return
	}

	if !strings.HasPrefix(r.URL.Path, "/api/") {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	paths := strings.Split(r.URL.Path, "/")
	service := paths[2]
	routings, ok := g.test.servicesToRoute[service]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	found := false
	var port uint16
	for _, routing := range routings {
		if routing.routingFunc == nil {
			port = routing.port
			found = true
			break
		}

		ok = routing.routingFunc("/"+strings.Join(paths[3:], "/"), r.Method)
		if ok {
			port = routing.port
			found = true
			break
		}
	}
	if !found {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	url, _ := url.Parse(fmt.Sprintf("http://127.0.0.1:%d", port))
	proxy := httputil.NewSingleHostReverseProxy(url)

	http.StripPrefix("/api/"+service, proxy).ServeHTTP(w, r)
}

var _ http.Handler = &gateway{}

func newGateway(test *Test) *gateway {
	return &gateway{
		test: test,
	}
}
