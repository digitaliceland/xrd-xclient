package web

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Client struct {
	BaseURL   *url.URL
	XRoadClient string
	XRoadService string
	httpClient *http.Client
}


//Use: x:=flattenHeader(m)
//pre: m is a http.Header
//post: x is now a multi string with one line per header and a colon separating field name and value
func flattenHeader(m http.Header) string {
	s := ""
	for k, v := range m {
		for _, xx := range v {
			s = fmt.Sprintf("%s%s: %s\n", s, k, xx)
		}
	}
	return s
}


//Pre: Client object has been initialized with proper service, client and URL info.  path is relative url of service
//Post: the requested service has been called and resulting text string returned
func (c *Client) do(path string) (string, string, string, error) {
	service_path := fmt.Sprintf("/r1/%s/%s", c.XRoadService, path)
	rel := &url.URL{Path: service_path}
	u := c.BaseURL.ResolveReference(rel)
	fmt.Println("Accessing: ", u.String())
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return "","","", err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-Road-Client", c.XRoadClient)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "",flattenHeader(req.Header),"", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	res := string(body)
	return res, flattenHeader(req.Header), flattenHeader(resp.Header), err
}
