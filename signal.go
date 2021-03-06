package profhandler

import (
	"fmt"
	"github.com/pkg/profile"
	"os"
	"os/signal"
	"syscall"
	"strings"
	"io/ioutil"
)

func NewSignalHandler() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGUSR1, syscall.SIGUSR2)
	go func() {
		for {
			s := <- sig
			switch s {
			case syscall.SIGUSR1:
				signalStart()
			case syscall.SIGUSR2:
				signalStop()
			}
		}
	}()
}

func signalStart() {
	if prof != nil {
		fmt.Println("Error: Profiling already started.")
		return
	}
	mode := ""
	mode_file := fmt.Sprintf("%s/data/PROFILING_MODE", os.Getenv("SRC_ROOT"))
	f, err := os.Open(mode_file)
	if err == nil {
		buf, err := ioutil.ReadAll(f)
		if err == nil {
			mode = strings.TrimSpace(string(buf))
		}
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
