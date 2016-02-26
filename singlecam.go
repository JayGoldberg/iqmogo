package main

import (
	"net/url"
	"net/http"
	"fmt"
	"time"
	"mime"
	"mime/multipart"
	"errors"
	"strings"
	"os"
)

type Motionstreamer struct {
	freq			string
	waittime		time.Duration
	running			bool
	url			string
}

func NewMotionstreamer() *Motionstreamer {
	p := &Motionstreamer{
		waittime: time.Second * 1,
	}
	return p
}

func (m *Motionstreamer) getresponse(request *http.Request) (*http.Response, error) {
	tr := &http.Transport{DisableKeepAlives: true}
	client := &http.Client{Transport: tr}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		response.Body.Close()
		errs := "Got invalid response status: " + response.Status
		return nil, errors.New(errs)
	}
	return response, nil
}

func (m *Motionstreamer) openstream() {
	var boundary string
	var response *http.Response
	var mpread *multipart.Reader

	// if &noimage, timing is not respected and the camera will deliver as fast as possible
	m.url = "http://" + m.url + "/now.jpg?snap=spush" + m.freq + "&pragma=motion&noimage"
	var img *multipart.Part

	m.running = true

	for m.running {
		request, err := http.NewRequest("GET", m.url, nil)
		if err != nil {
			fmt.Println(m.url, err)
		}

		response, err = m.getresponse(request)
		if err != nil {
			fmt.Println(m.url, err)
                }
		m.running = true

		boundary, err = m.getboundary(response)
		if err != nil {
			fmt.Println(m.url, err)
			m.running = false
			response.Body.Close()
			time.Sleep(m.waittime)
		}

		mpread = multipart.NewReader(response.Body, boundary)

		for m.running {
			img, err = mpread.NextPart()
			if err != nil {
				fmt.Println(err)
				m.running = false
				break
			}

			// https://groups.google.com/forum/#!topic/golang-nuts/sdsA9U0Wpnk
			if strings.Contains(img.Header["Pragma"][0], "motion") {
				fmt.Println(img.Header["Pragma"][0])
			}

			img.Close()
		}
		response.Body.Close()
	}

	//time.Sleep(m.waittime) // or comment out for as fast as camera can deliver
}


func (m *Motionstreamer) getboundary(response *http.Response) (string, error) {
	header := response.Header.Get("Content-Type")

	if header == "" {
		return "", errors.New("Content-Type isn't specified!")
	}
	ct, params, err := mime.ParseMediaType(header)
	if err != nil {
		return "", err
	}
	if ct != "multipart/x-mixed-replace" {
		errs := "Wrong Content-Type: expected multipart/x-mixed-replace, got " + ct
		return "", errors.New(errs)
	}
	boundary, ok := params["boundary"]
	if !ok {
		return "", errors.New("No multipart boundary param in Content-Type header!")
	}
	// Some IP-cameras screw up boundary strings so we
	// have to remove excessive "--" characters manually.
	boundary = strings.Replace(boundary, "--", "", -1)

	return boundary, nil
}


func main() {
	camurl := os.Args[1]

	udecomp, err := url.Parse(camurl)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Reading motion data from http://" + udecomp.Host + "/serverpush.html?ds=4")
	motionHandler := NewMotionstreamer()
	motionHandler.url = udecomp.Host
	motionHandler.freq = "0.5" // 1 = 1 per sec, 0.1 = 10 per sec
	motionHandler.openstream()
}
