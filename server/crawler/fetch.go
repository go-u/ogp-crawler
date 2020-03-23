package crawler

import (
	"bytes"
	"errors"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/otiai10/opengraph"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	domain "server/domain/model"
	"strings"
	"time"
)

func (c *Crawler) fetchOgp(tweet twitter.Tweet) (*domain.Ogp, error) {
	Link := tweet.Entities.Urls[0].ExpandedURL
	buf := bytes.NewBuffer(nil)
	req, err := http.NewRequest("GET", Link, buf)
	if err != nil {
		return nil, err
	}

	const ua = "twitterbot"
	req.Header.Add("User-Agent", ua)

	//https://medium.com/@nate510/don-t-use-go-s-default-http-client-4804cb19f779
	var netTransport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 10 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 10 * time.Second,
	}

	var client = &http.Client{
		Transport: netTransport,
		Timeout:   time.Second * 10,
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	// copy body for later use as below parse method destruct res.body
	body_copy := copyBodyNonDestructive(res)
	ogpInfo := opengraph.New(Link)
	if err := ogpInfo.Parse(res.Body); err != nil {
		return nil, err
	}

	// ogp画像のチェック
	if len(ogpInfo.Image) == 0 {
		return nil, errors.New("no image")
	}
	if len(ogpInfo.Image) > 0 {
		isSecure := strings.HasPrefix(ogpInfo.Image[0].URL, "https")
		if !isSecure {
			return nil, errors.New("image url is not https ")
		}
	}

	// ogp情報のデコード
	decoder, err := getDecoder(body_copy, res.Header)
	if err != nil {
		return nil, err
	}
	titleDecoded, _ := decoder.String(ogpInfo.Title)
	descriptionDecoded, _ := decoder.String(ogpInfo.Description)

	// 引数のURLではなくリダイレクト先のURLを記録
	requestUrl, err := url.Parse(res.Request.URL.String())
	if err != nil || requestUrl == nil {
		return nil, err
	}

	ogp := &domain.Ogp{
		HostName:    c.HostName,
		Date:        time.Now(),
		FQDN:        requestUrl.Host,
		URL:         Link, // tweetに記載されたURL
		Title:       titleDecoded,
		Description: descriptionDecoded,
		Image:       ogpInfo.Image[0].URL,
		Type:        ogpInfo.Type,
		Lang:        tweet.Lang,
		TweetID:     tweet.ID,
	}
	return ogp, err
}

func getDecoder(body io.ReadCloser, header http.Header) (*encoding.Decoder, error) {
	b, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}
	enc, _, _ := charset.DetermineEncoding(b, header.Get("content-type"))
	return enc.NewDecoder(), nil
}

func copyBodyNonDestructive(res *http.Response) io.ReadCloser {
	buf, _ := ioutil.ReadAll(res.Body)
	res.Body = ioutil.NopCloser(bytes.NewBuffer(buf))
	body_copy := ioutil.NopCloser(bytes.NewBuffer(buf))
	return body_copy
}
