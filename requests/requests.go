package requests

import (
	"bytes"
	"context"
	"crypto/tls"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/un4gi/mBot/env"
)

var Token = ""
var Urls = []string{
	// urls[0] = all unregistered targets
	"https://platform.synack.com/api/targets?filter%5Bprimary%5D=unregistered&filter%5Bsecondary%5D=all&filter%5Bcategory%5D=all&filter%5Bindustry%5D=all&sorting%5Bfield%5D=dateUpdated&sorting%5Bdirection%5D=desc",
	// urls[1] = available missions sorted by price
	// "https://platform.synack.com/api/tasks/v1/tasks?sortBy=price-sort-desc&withHasBeenViewedInfo=true&status=PUBLISHED&page=0&pageSize=20",
	// urls[1] = available missions sorted by price (v2)
	"https://platform.synack.com/api/tasks/v2/tasks?perPage=20&viewed=true&page=1&status=PUBLISHED&sort=AMOUNT&sortDir=desc",
	// urls[2] = QR window
	"https://platform.synack.com/api/targets?filter%5Bprimary%5D=all&filter%5Bsecondary%5D%5B%5D=a&filter%5Bsecondary%5D%5B%5D=l&filter%5Bsecondary%5D%5B%5D=l&filter%5Bsecondary%5D%5B%5D=quality_period&filter%5Bcategory%5D=all&filter%5Bindustry%5D=all&sorting%5Bfield%5D=dateUpdated&sorting%5Bdirection%5D=desc",
	// urls[3] = claimed missions
	"https://platform.synack.com/api/tasks/v2/tasks?perPage=20&viewed=true&page=1&status=CLAIMED",
	// urls[4] = beginning of URL to edit missions
	"https://platform.synack.com/api/tasks/v2/tasks/",
	// urls[5] = authenticate URL
	"https://login.synack.com/api/authenticate",
}

func SetHeaders(req *http.Request) {
	req.Header.Set("User-Agent", "mBot (https://github.com/un4gi/mBot)")
	req.Header.Set("Authorization", "Bearer "+Token)
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Referer", "https://platform.synack.com/tasks/user/available")
	req.Header.Set("X-CSRF-Token", "xxxx")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Connection", "close")
}

func DoGetRequest(target string) (int, io.ReadCloser) {
	ctx, cancel := context.WithTimeout(context.Background(), 180000*time.Millisecond)
	defer cancel()

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	req, err := http.NewRequest("GET", target, nil)
	if err != nil {
		log.Println(err)
		return 0, nil
	}
	SetHeaders(req)

	resp, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		log.Println(err)
		return 0, nil
	}
	return resp.StatusCode, resp.Body
}

func DoPostRequest(target string, jsonStr []byte) (int, io.ReadCloser) {
	ctx, cancel := context.WithTimeout(context.Background(), 180000*time.Millisecond)
	defer cancel()

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	req, err := http.NewRequest("POST", target, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Println(err)
	}

	SetHeaders(req)

	resp, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		log.Println(err)
	}
	return resp.StatusCode, resp.Body
}

func SetLoginHeaders(req *http.Request, token string, cookie string) {
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:95.0) Gecko/20100101 Firefox/95.0")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Charset", "utf-8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("Referer", "https://login.synack.com")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-CSRF-Token", token)
	req.Header.Set("Origin", "https://login.synack.com")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cookie", cookie)
}

func DoLoginGetRequest(target string) (int, io.ReadCloser, http.Header) {
	ctx, cancel := context.WithTimeout(context.Background(), 180000*time.Millisecond)
	defer cancel()

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	req, err := http.NewRequest("GET", target, nil)
	if err != nil {
		log.Println(err)
		return 0, nil, nil
	}
	SetHeaders(req)

	resp, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		log.Println(err)
		return 0, nil, nil
	}
	return resp.StatusCode, resp.Body, resp.Header
}

func DoLoginPostRequest(target string, jsonStr []byte, token string, cookie string) (int, io.ReadCloser) {
	ctx, cancel := context.WithTimeout(context.Background(), 180000*time.Millisecond)
	defer cancel()

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	req, err := http.NewRequest("POST", target, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Println(err)
	}

	SetLoginHeaders(req, token, cookie)

	resp, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		log.Println(err)
	}
	return resp.StatusCode, resp.Body
}

func SetGrantTokenHeaders(req *http.Request) {
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:95.0) Gecko/20100101 Firefox/95.0")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Charset", "utf-8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("Referer", "https://login.synack.com")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-CSRF-Token", "xxxx")
	req.Header.Set("Origin", "https://login.synack.com")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
}

func DoGrantTokenRequest(target string) (int, io.ReadCloser) {
	ctx, cancel := context.WithTimeout(context.Background(), 180000*time.Millisecond)
	defer cancel()

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	req, err := http.NewRequest("GET", target, nil)
	if err != nil {
		log.Println(err)
		return 0, nil
	}
	SetGrantTokenHeaders(req)

	resp, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		if err == context.DeadlineExceeded {
			log.Printf(env.ErrorColor, err)
		} else {
			log.Printf(env.ErrorColor, err)
			return 0, nil
		}
	}
	return resp.StatusCode, resp.Body
}
