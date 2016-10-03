package profhandler

import (
	"fmt"
	"github.com/pkg/profile"
	"os"
	"os/signal"
	"syscall"
	"strings"
)

func NewSignalHandler() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGUSR1, syscall.SIGUSR2)
	go func() {
		for {
			s := <- sig
			switch s {
			case syscall.SIGUSR1:
				mode := os.Getenv("PROFILING_MODE")
				signalStart(strings.TrimSpace(mode))
			case syscall.SIGUSR2:
				signalStop()
			}
		}
	}()
}

func signalStart(mode string) {
	if prof != nil {
		fmt.Println("Error: Profiling already started.")
		return
	}
	profiles := []func(*profile.Profile){}
	switch mode {
	case "cpu":
		profiles = append(profiles, profile.CPUProfile)
	case "mem":
		profiles = append(profiles, profile.MemProfile)
	case "block":
		profiles = append(profiles, profile.BlockProfile)
	default:
		profiles = append(profiles, profile.CPUProfile)
		mode = "cpu"
	}
	profiles = append(profiles, profile.ProfilePath("."))
	prof = profile.Start(profiles...)
	fmt.Println("Profiling started. Mode:", mode)

}

func signalStop() {
	if prof != nil {
		prof.Stop()
		prof = nil
		fmt.Println("Profiling stopped")
	} else {
		fmt.Println("Error: Profiling already stopped.")
	}

}
