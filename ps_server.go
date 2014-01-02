package main

import (
    "io"
    "net/http"
    "os/exec"
)

func ps() ([]byte, error) {
	cmd := exec.Command("ps")
	return cmd.Output()
}

func psServer(w http.ResponseWriter, req *http.Request) {
	output, err := ps()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

//  Convert the raw uptime output to an UptimeInfo object.
//  ui, err := parseUptimeInfo(output)
//  if err != nil {
// 	w.WriteHeader(http.StatusInternalServerError)
//		return
//	}

	// Create the JSON representation of the system uptime.
//	data, err := json.MarshalIndent(ui, " ", "")
//	if err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//		return
//	}

	// Write the HTTP response headers and body.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, string(output))
}
