package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func TestRootDescribeTags(t *testing.T) {
	expected_body := `<DescribeTagsResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/"/">
   <requestId>7a62c49f-347e-4fc4-9331-6e8e</requestId>
   <tagSet>
      <item>
         <resourceId>i-asdfasdf</resourceId>
         <resourceType>instance</resourceType>
         <key>BLAH</key>
         <value>asdf</value>
      </item>
      <item>
         <resourceId>i-asdfasdf</resourceId>
         <resourceType>instance</resourceType>
         <key>aaaa</key>
         <value>bbbb</value>
      </item>
   </tagSet>
</DescribeTagsResponse>`

	form := url.Values{}
	form.Add("Action", "DescribeTags")
	form.Add("Version", "2016-11-15")
	form.Add("Filter.1.Name", "resource-id")
	form.Add("Filter.1.Value.1", "i-asdfasdf")
	form.Add("Filter.2.Name", "resource-type")
	form.Add("Filter.2.Value.1", "instance")

	client := &http.Client{}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/", testServer.URL), strings.NewReader(form.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	// Auth not actually checked by this API mock
	req.Header.Add("Authorization", "AWS4-HMAC-SHA256 ...")
	req.Header.Add("Accept-Encoding", "identity")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		t.Fatal(err)
	}
	if string(body) != expected_body {
		t.Errorf("Expected\n\n%s\n\ngot\n\n%s", expected_body, string(body))
	}
}
