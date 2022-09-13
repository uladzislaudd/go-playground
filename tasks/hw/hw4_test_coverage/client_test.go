package main

// код писать тут
import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type (
	stub struct {
		w time.Duration
		s int
		d []byte
	}
)

var (
	first25 = make([]User, 25)
)

func init() {
	s, err := new_server(filename, []string{"kek"})
	if err != nil {
		panic(err)
	}

	for i := 0; i < 25; i++ {
		first25[i] = *s.users[i]
	}
}

func (s *stub) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	time.Sleep(s.w)
	rw.WriteHeader(s.s)
	rw.Write(s.d)
}

func TestSearchClient_FindUsers1(t *testing.T) {
	s, err := new_server(filename, []string{"kek"})
	assert.NotNil(t, s)
	assert.NoError(t, err)

	s1 := httptest.NewServer(s)
	assert.NotNil(t, s1)

	s2 := httptest.NewServer(&stub{w: 1050 * time.Millisecond})
	assert.NotNil(t, s2)

	s3 := httptest.NewServer(&stub{s: http.StatusInternalServerError})
	assert.NotNil(t, s3)

	s4 := httptest.NewServer(&stub{s: http.StatusBadRequest, d: []byte("hehexe")})
	assert.NotNil(t, s4)

	s5 := httptest.NewServer(&stub{s: http.StatusOK, d: []byte("hehexe")})
	assert.NotNil(t, s5)

	srv := &SearchClient{}

	tests := []struct {
		name string
		at   string
		url  string
		req  SearchRequest
		err  error
	}{
		{"00", "", "", SearchRequest{Limit: -1}, fmt.Errorf("limit must be > 0")},
		{"01", "", "", SearchRequest{Offset: -1}, fmt.Errorf("offset must be > 0")},
		{"02", "", "", SearchRequest{}, fmt.Errorf("unknown error Get \"?limit=1&offset=0&order_by=0&order_field=&query=\": unsupported protocol scheme \"\"")},
		{"03", "", s2.URL, SearchRequest{}, fmt.Errorf("timeout for limit=1&offset=0&order_by=0&order_field=&query=")},
		{"04", "", s3.URL, SearchRequest{}, fmt.Errorf("SearchServer fatal error")},
		{"05", "", s4.URL, SearchRequest{}, fmt.Errorf("cant unpack error json: invalid character 'h' looking for beginning of value")},
		{"06", "", s5.URL, SearchRequest{}, fmt.Errorf("cant unpack result json: invalid character 'h' looking for beginning of value")},
		{"10", "xd", s1.URL, SearchRequest{}, fmt.Errorf("Bad AccessToken")},
		{"11", "kek", s1.URL, SearchRequest{OrderField: "XD"}, fmt.Errorf("OrderFeld XD invalid")},
		{"12", "kek", s1.URL, SearchRequest{OrderBy: 2}, fmt.Errorf("unknown bad request error: order_by should be in {-1,0,1}")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv.AccessToken = tt.at
			srv.URL = tt.url
			_, err := srv.FindUsers(tt.req)
			assert.Equal(t, tt.err, err)
		})
	}
}

func TestSearchClient_FindUsers2(t *testing.T) {
	s, err := new_server(filename, []string{"kek"})
	assert.NotNil(t, s)
	assert.NoError(t, err)

	s1 := httptest.NewServer(s)
	assert.NotNil(t, s1)

	srv := &SearchClient{
		AccessToken: "kek",
		URL:         s1.URL,
	}
	tests := []struct {
		name string
		req  SearchRequest
		want *SearchResponse
	}{
		{"00", SearchRequest{Limit: 26, Query: "Boyd Wolf"}, &SearchResponse{Users: []User{*s.users[0]}}},
		{"01", SearchRequest{Limit: 24, Query: ""}, &SearchResponse{NextPage: true, Users: first25[:24]}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := srv.FindUsers(tt.req)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
