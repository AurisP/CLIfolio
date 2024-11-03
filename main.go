package main

import (
	"errors"
	"net"
	"os"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/logging"
)

const (
	host = "localhost"
	port = "23234"
)

func main() {
	// Initialize progress bar model
	m := loadingModel{
		progress: progress.New(progress.WithScaledGradient("#f0f2f2", "#08b9ff")),
	}

	srv, err := wish.NewServer(
		wish.WithAddress(net.JoinHostPort(host, port)),
		wish.WithHostKeyPath(".ssh/id_ed25519"),
		wish.WithMiddleware(
			func(next ssh.Handler) ssh.Handler {
				return func(sess ssh.Session) {
					if _, err := tea.NewProgram(m, tea.WithInput(sess), tea.WithOutput(sess)).Run(); err != nil {
						wish.Println(sess, "Error starting the application: "+err.Error())
					}
					next(sess)
				}
			},
			logging.Middleware(),
		),
	)

	if err != nil {
		log.Error("Could not start server", "error", err)
		os.Exit(1)
	}

	log.Info("Starting SSH server", "host", host, "port", port)
	if err = srv.ListenAndServe(); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
		log.Error("Could not start server", "error", err)
		os.Exit(1)
	}
}
