package tests

import (
	"gopkg.in/dutchcoders/goftp.v1"
	"sync"
	"testing"
)

func TestConcurrency(t *testing.T) {
	s := getServer(false)
	defer s.Stop()

	nbClients := 100

	waitGroup := sync.WaitGroup{}
	waitGroup.Add(nbClients)

	for i := 0; i < nbClients; i++ {
		go func() {
			var err error
			var ftp *goftp.FTP

			if ftp, err = goftp.Connect(s.Listener.Addr().String()); err != nil {
				panic(err)
			}
			defer ftp.Close()

			if err = ftp.Login("test", "test"); err != nil {
				t.Fatal("Failed to login:", err)
			}

			waitGroup.Done()
		}()
	}

	waitGroup.Wait()
}
