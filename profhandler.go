package profhandler

import (
	"net/http"
	"fmt"
	"github.com/pkg/profile"
)

var prof interface{
	Stop()
}

func Start(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.Header().Add("X-Error-Desc", err.Error())
		w.WriteHeader(500)
	} else if prof != nil {
		fmt.Fprintln(w, "Profiling already started.")
	} else {
		profiles := []func(*profile.Profile){}
		mode := ""
		path := ""
		for k, v := range r.Form {
			key := string(k)
			switch key {
			case "cpu":
				if mode == "" {
					profiles = append(profiles, profile.CPUProfile)
					mode = "cpu"
				}
			case "mem":
				if mode == "" {
					profiles = append(profiles, profile.MemProfile)
					mode = "mem"
				}
			case "block":
				if mode == "" {
					profiles = append(profiles, profile.BlockProfile)
					mode = "block"
				}
			}
			if key == "path" {
				if path == "" {
					path = string(v[0])
					profiles = append(profiles, profile.ProfilePath(path))
				}
			}
		}
		prof = profile.Start(profiles...)
		if mode == "" {
			mode = "cpu"
		}
		if path == "" {
			path = "temporary path"
		}
		fmt.Fprintln(w, "Profiling started. Mode:", mode, "ProfilePath:", path)
	}
}

func Stop(w http.ResponseWriter, r *http.Request) {
	if prof != nil {
		prof.Stop()
		prof = nil
		fmt.Fprintln(w, "Profiling stopped")
	} else {
		fmt.Fprintln(w, "Error: Profiling already stopped.")
	}

}
