// Package kripa implements utility functions to generate Korean IPA
// transcriptions using the web http://pronunciation.cs.pusan.ac.kr/.
package kripa

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/transform"
)

/*
	EscapeEucKr returns a string with the escaped represetation in
	extended Unicode of the input Korean string.
*/
func EscapeEucKr(s string) string {
	var bufs bytes.Buffer
	wr := transform.NewWriter(&bufs, korean.EUCKR.NewEncoder())
	defer wr.Close()

	wr.Write([]byte(s))
	eucKrVals := bufs.String()

	escapedString := ""
	for i := 0; i < len(eucKrVals); i++ {
		escapedString += fmt.Sprintf("%%%02X", eucKrVals[i])
	}

	return escapedString
}

/*
	GetTranscriptIPA returns a string with the API transcription of the
	input Korean sentence escraped from
	http://pronunciation.cs.pusan.ac.kr.
*/
func GetTranscriptIPA(s string) (string, error) {
	url := fmt.Sprintf(
		"http://pronunciation.cs.pusan.ac.kr/pronunc2.asp?text1=%s&submit1=%s",
		EscapeEucKr(s),
		EscapeEucKr("확인하기"))

	client := &http.Client{}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	request.Header.Set("User-Agent",
		"Go-KrIPABot https://github.com/alejandrosame/go-kripa - Scraper bot to assist Korean to IPA dataset creation")

	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	utfBody := transform.NewReader(response.Body, korean.EUCKR.NewDecoder())

	document, err := goquery.NewDocumentFromReader(utfBody)
	if err != nil {
		return "", err
	}

	var transcription string
	document.Find("table").Each(func(i int, s *goquery.Selection) {
		tds := s.Find("td")
		title := tds.First()
		fmt.Println(fmt.Sprintf("%+v", title))
		if title.Text() == "IPA 결과" {
			transcription = strings.TrimSpace(title.Next().Text())
			fmt.Println(transcription)
		}
	})

	return transcription, nil
}
