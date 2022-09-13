package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type (
	row struct {
		ID            uint64 `xml:"id"`
		GUID          string `xml:"guid"`
		IsActive      bool   `xml:"isActive"`
		Balance       string `xml:"balance"`
		Picture       string `xml:"picture"`
		Age           uint8  `xml:"age"`
		EyeColor      string `xml:"eyeColor"`
		FirstName     string `xml:"first_name"`
		LastName      string `xml:"last_name"`
		Gender        string `xml:"gender"`
		Company       string `xml:"company"`
		Email         string `xml:"email"`
		Phone         string `xml:"phone"`
		Address       string `xml:"address"`
		About         string `xml:"about"`
		Registered    string `xml:"registered"`
		FavoriteFruit string `xml:"favoriteFruit"`
	}

	root struct {
		Rows []row `xml:"row"`
	}

	server struct {
		tokens map[string]any

		r     *root
		users []*User
	}
)

var (
	filename = "dataset.xml"

	ErrInvalidAccessToken = fmt.Errorf("invalid access token")
	ErrMissedField        = fmt.Errorf("missed field")
	ErrSubZeroLimit       = fmt.Errorf("limit must be >= 0")
	ErrSubZeroOffset      = fmt.Errorf("offset must be > 0")
	ErrOffsetTooBig       = fmt.Errorf("offset is too big")
	ErrInvalidOrderField  = fmt.Errorf("ErrorBadOrderField")
	ErrInvalidOrderBy     = fmt.Errorf("order_by should be in {-1,0,1}")

	ValidOrderFields = map[string]func(u1, u2 *User) int{
		"":     nil,
		"Id":   compareUserId,
		"Age":  compareUserAge,
		"Name": compareUserName,
	}
)

func (r *row) User() (rv *User) {
	rv = &User{
		Id:     int(r.ID),
		Name:   strings.Join([]string{r.FirstName, r.LastName}, " "),
		Age:    int(r.Age),
		About:  r.About,
		Gender: r.Gender,
	}
	return
}

func Load(filename string) (rv *root, err error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		err = errors.Wrap(err, "os.ReadFile(\""+filename+"\") failure:")
		return
	}

	rv = &root{}
	err = xml.Unmarshal(data, rv)
	err = errors.Wrap(err, "xml.Unmarshal failure:")

	return
}

func new_server(filename string, tokens []string) (rv *server, err error) {
	s := &server{
		tokens: map[string]any{},
	}

	s.r, err = Load(filename)
	if err != nil {
		err = errors.Wrap(err, "failed to create server")
		return
	}

	s.users = make([]*User, len(s.r.Rows))
	for i := range s.r.Rows {
		s.users[i] = s.r.Rows[i].User()
	}

	for _, token := range tokens {
		s.tokens[token] = nil
	}

	rv = s
	return
}

func (s *server) SearchRequest(query string) (rv *SearchRequest, err error) {
	sr := &SearchRequest{}

	vs, err := url.ParseQuery(query)
	if err != nil {
		err = errors.Wrap(err, "failed to ParseQuery(\""+query+"\")")
		return
	}

	if v, ok := vs["limit"]; ok && len(v) > 0 {
		sr.Limit, err = strconv.Atoi(v[0])
		if err != nil {
			err = errors.Wrap(err, "failed to convert "+v[0]+" to int")
			return
		}
	} else {
		err = errors.Wrap(ErrMissedField, "limit")
		return
	}
	if sr.Limit < 0 {
		err = errors.Wrap(ErrSubZeroLimit, "invalid search request")
		return
	}
	if sr.Limit > 25 {
		sr.Limit = 25
	}

	if v, ok := vs["offset"]; ok && len(v) > 0 {
		sr.Offset, err = strconv.Atoi(v[0])
		if err != nil {
			err = errors.Wrap(err, "failed to convert "+v[0]+" to int")
			return
		}
	} else {
		err = errors.Wrap(ErrMissedField, "offset")
		return
	}
	if sr.Offset < 0 {
		err = errors.Wrap(ErrSubZeroOffset, "invalid search request")
		return
	}
	if sr.Offset > len(s.users) {
		err = errors.Wrap(ErrOffsetTooBig, "invalid search request")
		return
	}
	if len(s.users)-sr.Offset < sr.Limit {
		sr.Limit = len(s.users) - sr.Offset
	}

	if v, ok := vs["query"]; ok && len(v) > 0 {
		sr.Query = v[0]
	} else {
		err = errors.Wrap(ErrMissedField, "query")
		return
	}

	if v, ok := vs["order_field"]; ok && len(v) > 0 {
		sr.OrderField = v[0]
	} else {
		err = errors.Wrap(ErrMissedField, "order_field")
		return
	}
	for _, field := range strings.Split(sr.OrderField, ",") {
		if _, ok := ValidOrderFields[field]; !ok {
			err = errors.Wrap(ErrInvalidOrderField, sr.OrderField)
			return
		}
	}

	if v, ok := vs["order_by"]; ok && len(v) > 0 {
		sr.OrderBy, err = strconv.Atoi(v[0])
		if err != nil {
			err = errors.Wrap(err, "failed to convert "+v[0]+" to int")
			return
		}
	} else {
		err = errors.Wrap(ErrMissedField, "order_by")
		return
	}
	if sr.OrderBy != OrderByAsc && sr.OrderBy != OrderByAsIs && sr.OrderBy != OrderByDesc {
		err = errors.Wrap(ErrInvalidOrderBy, fmt.Sprint(sr.OrderBy))
		return
	}

	return sr, nil
}

func compareUserId(u1, u2 *User) int {
	return 1
}

func compareUserAge(u1, u2 *User) int {
	switch {
	case u1.Age < u2.Age:
		return -1
	case u1.Age > u2.Age:
		return 1
	}
	return 0
}

func compareUserName(u1, u2 *User) int {
	switch {
	case u1.Name < u2.Name:
		return -1
	case u1.Name > u2.Name:
		return 1
	}
	return 0
}

func (s *server) SearchServer(sr *SearchRequest) (rv []*User) {
	count := 0
	for i := sr.Offset; (i < len(s.users)) && (count < sr.Limit); i++ {
		user := s.users[i]
		switch {
		case sr.Query == "":
			fallthrough
		case strings.Contains(user.Name, sr.Query):
			fallthrough
		case strings.Contains(user.About, sr.Query):
			rv = append(rv, user)
			count++
		}
	}

	if sr.OrderBy == 0 || sr.OrderField == "" {
		return
	}

	comparers := []func(u1, u2 *User) int{}
	for _, field := range strings.Split(sr.OrderField, ",") {
		if v, ok := ValidOrderFields[field]; ok {
			comparers = append(comparers, v)
		}
	}

	asc := sr.OrderBy == OrderByAsc
	less := func(i, j int) bool {
		for _, l := range comparers {
			switch l(rv[i], rv[j]) {
			case -1:
				return asc
			case 1:
				return !asc
			}
		}
		return !asc
	}

	sort.SliceStable(rv, less)

	return
}

func (s *server) serveHTTP(token string, query string) (int, []byte) {
	if _, ok := s.tokens[token]; !ok {
		log.Printf("failed to serve request \"%s\", invalid access token: \"%s\"", query, token)
		return http.StatusUnauthorized, nil
	}

	sr, err := s.SearchRequest(query)
	if err != nil {
		log.Printf("failed to serve request \"%s\", error: \"%v\"", query, err)
		data, _ := json.Marshal(SearchErrorResponse{Error: errors.Cause(err).Error()})
		return http.StatusBadRequest, data
	}

	users := []User{}
	for _, user := range s.SearchServer(sr) {
		users = append(users, *user)
	}
	data, _ := json.Marshal(users)
	return http.StatusOK, data
}

func (s *server) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	status, data := s.serveHTTP(r.Header.Get("AccessToken"), r.URL.RawQuery)
	rw.WriteHeader(status)
	rw.Write(data)
}
