package main

import (
	"log"
	"net/http"
	"os/exec"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)

	http.HandleFunc("/callvlc", func(w http.ResponseWriter, r *http.Request) {
		url := r.FormValue("url")
		cmd := exec.Command("am", "start", "-n", "org.videolan.vlc/org.videolan.vlc.gui.video.VideoPlayerActivity", "-a", "android.intent.action.VIEW", "-d", url)
		err := cmd.Run()

		if err != nil {
			log.Fatal(err)
		}
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
