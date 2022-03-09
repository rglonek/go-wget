package wget

import (
	"bytes"
	"log"
	"testing"
	"time"
)

func TestGet(t *testing.T) {
	w := new(bytes.Buffer)
	output, err := Get(&GetInput{
		Url:    "http://212.183.159.230/5MB.zip",
		Writer: w,
	})
	if err != nil {
		t.Logf("Get Error: %s", err)
		t.FailNow()
	}
	t.Logf("transferred:%s statusCode:%d status:%s total:%s", SizeToString(output.NumBytes), output.ResponseCode, output.Response, SizeToString(output.TotalBytes))
}

func TestGetWithProgress(t *testing.T) {
	w := new(bytes.Buffer)
	output, err := GetWithProgress(&GetInput{
		Url:               "http://212.183.159.230/20MB.zip",
		Writer:            w,
		CallbackFrequency: time.Second,
		CallbackFunc:      callback,
	})
	if err != nil {
		t.Logf("Get Error: %s", err)
		t.FailNow()
	}
	t.Logf("transferred:%s statusCode:%d status:%s total:%s", SizeToString(output.NumBytes), output.ResponseCode, output.Response, SizeToString(output.TotalBytes))
}

func callback(p *Progress) {
	log.Printf("%d%% complete @ %s / second (%s elapsed)", p.PctComplete, SizeToString(p.BytesPerSecond), p.TimeElapsed.Round(time.Second))
}
