package a3rt

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSmallTalk(t *testing.T) {
	expected := &SmallTalkResult{
		Perplexity: 3.2993975249803382,
		Reply:      "poyopi!",
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/talk/v1/smalltalk" {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		b, err := json.Marshal(&SmallTalkResponse{
			Status:  0,
			Message: "ok",
			Results: []*SmallTalkResult{expected},
		})
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		if _, err := w.Write(b); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		return
	}))
	defer ts.Close()

	client := NewClient()
	client.SetBaseUri(ts.URL)
	client.SetApiKey("secret")

	res, err := client.SmallTalk(context.Background(), "poyopi?")
	if err != nil {
		t.Fatalf("should not be fail: %v", err)
	}
	if res.Perplexity != expected.Perplexity {
		t.Fatalf("want %f but %f", expected.Perplexity, res.Perplexity)
	}
	if res.Reply != expected.Reply {
		t.Fatalf("want %q but %q", expected.Reply, res.Reply)
	}
}
