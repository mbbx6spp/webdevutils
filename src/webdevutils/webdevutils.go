package webdevutils

import (
	"crypto/tls"
	"errors"
	"net/http"
	"os"
  "os/signal"
)

type trap func(<-chan os.Signal)

// StaticServerTLS binds to ipAddr using certificate file, crtFile, and PEM-
// encoded key file, keyFile, to serve static content from directory, rootDir.
// It checks that rootDir is a directory, both crtFile and keyFile exist and
// are a valid X509 keypair together.
// Returns an error from any of those check or if the HTTP server encountered
// and error while starting up.
func StaticServerTLS(ipAddr, crtFile, keyFile, rootDir string) error {
	fi, err := os.Stat(rootDir)
	if err != nil {
		return err
	}
	if !fi.IsDir() {
		return errors.New(rootDir + " is not a directory")
	}

	err = ValidateX509KeyPair(crtFile, keyFile)
	if err != nil {
		return err
	}

	return http.ListenAndServeTLS(ipAddr, crtFile, keyFile,
		http.FileServer(http.Dir(rootDir)))
}

// StaticServer binds to ipAddr and serves static content from rootDir.
// Checks rootDir is a valid directory.
// It returns an error if rootDir is not a directory, does not exist, or
// if there was a problem starting the HTTP server to serve static content.
func StaticServer(ipAddr, rootDir string) error {
	fi, err := os.Stat(rootDir)
	if err != nil {
		return err
	}
	if !fi.IsDir() {
		return errors.New(rootDir + " is not a directory")
	}

	return http.ListenAndServe(ipAddr,
		http.FileServer(http.Dir(rootDir)))
}

// ValidateX509KeyPair validates a certificate file, crtFile, and a PEM-
// encoded key file, keyFile, are a valid X509 key pair.
// It returns an error if either certFile or keyFile do not exist or
// when they are not a valid X509 key pair.
// It returns nil in all successful cases.
func ValidateX509KeyPair(crtFile, keyFile string) error {
	_, err := os.Stat(crtFile)
	if err != nil {
		return err
	}

	_, err = os.Stat(keyFile)
	if err != nil {
		return err
	}

	_, err = tls.LoadX509KeyPair(crtFile, keyFile)
	if err != nil {
		return err
	}

	return nil
}

// RegisterSignalTrap will register a signal trap goroutine, signalTrap, for
// a list of OS signals, signals.
// Returns channel that received OS signals so that client can
// UnregisterSignalTrap if or when needed.
func RegisterSignalTrap(sTrap trap, sigs ...os.Signal) chan<- os.Signal {
  c := make(chan os.Signal, 1)

  for _, s := range sigs {
    signal.Notify(c, s)
  }
  go sTrap(c)

  return c
}

// UnregisterSignalTrap will stop signals being sent to a channel already
// registered for specific OS signals. It will send the specified OS signal
// to the channel after it is unregistered to receive signals from the OS.
// This could be nil and should be handled appropriately by the goroutine
// listening for os.Signals on the given channel passed to RegisterSignalTrap
// originally.
func UnregisterSignalTrap(c chan<- os.Signal, sig os.Signal) {
  signal.Stop(c)
  c<- sig
}
