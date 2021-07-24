package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strconv"
)

type checkServerResponse struct {
	Check string "json check"
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {

	response := checkServerResponse{Check: "Hello"}
	encoder := json.NewEncoder(w)
	encoder.Encode(&response)

}

func main() {
	http.HandleFunc("/", defaultHandler)

	err := http.ListenAndServe(":9090", nil)

	if err != nil {
		log.Fatal("ListenAndServe : ", err)
	} else {
		fmt.Println("listen at 9090")
	}
}

func checkOk(out string, err error) {
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))

}

func createPhotos(filename string, startTime string, term string) {
	videofile := filename + ".mp4"
	photofile := filename + "_%d.jpg"
	ffmpeg, err := exec.Command("ffmpeg", "-ss", startTime, "-t", term, "-i", videofile, photofile).Output()
	checkOk(string(ffmpeg), err)
}

func deletePhotos() {
	i := 1
	for {
		v := strconv.Itoa(i)
		file := "Waves_" + v + ".jpg"
		_, err := exec.Command("rm", file).Output()
		if err != nil {
			return
		}
		i += 1
	}
}

func checkServer() string {
	return "ok"
}
