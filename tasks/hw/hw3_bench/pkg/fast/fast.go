package fast

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	jsoniter "github.com/json-iterator/go"
)

type (
	//easyjson:json
	User struct {
		Browsers []string `json:"browsers,omitempty,nocopy"`
		// Company  string   `json:"company,omitempty,nocopy"`
		// Country  string   `json:"country,omitempty,nocopy"`
		Email string `json:"email,omitempty,nocopy"`
		// Job      string   `json:"job,omitempty,nocopy"`
		Name string `json:"name,omitempty,nocopy"`
		// Phone    string   `json:"phone,omitempty,nocopy"`
	}

	Unmarshaler interface {
		Unmarshal([]byte, *User) error
	}

	um1 struct{}
	um2 struct{}
	um3 struct{}
)

var (
	json2 = jsoniter.ConfigCompatibleWithStandardLibrary
	um1o  = um1{}
	um2o  = um2{}
	um3o  = um3{}
)

func (um1) Unmarshal(data []byte, u *User) error {
	return u.UnmarshalJSON(data)
}

func (um2) Unmarshal(data []byte, u *User) error {
	return json2.Unmarshal(data, u)
}

func (um3) Unmarshal(data []byte, u *User) error {
	return json.Unmarshal(data, u)
}

func FastSearch(out io.Writer, data []byte) {
	FastSearch1(out, data)
}

func FastSearch1(out io.Writer, data []byte) {
	fastSearch(out, data, um1o)
}

func FastSearch2(out io.Writer, data []byte) {
	fastSearch(out, data, um2o)
}

func FastSearch3(out io.Writer, data []byte) {
	fastSearch(out, data, um3o)
}

func fastSearch(out io.Writer, data []byte, um Unmarshaler) {
	seenBrowsers := map[string]interface{}{}

	user := User{}
	fmt.Fprintln(out, "found users:")
	for i := 0; data != nil; i++ {
		j := bytes.IndexRune(data, '\n')
		if j == -1 {
			j = len(data)
		}

		// fmt.Printf("%v %v\n", err, line)
		if err := um.Unmarshal(data[:j], &user); err != nil {
			panic(err)
		}

		if j == len(data) {
			data = nil
		} else {
			data = data[j+1:]
		}

		isAndroid := false
		isMSIE := false

		for _, browser := range user.Browsers {
			if strings.Contains(browser, "Android") {
				seenBrowsers[browser] = nil
				isAndroid = true
				continue
			}

			if strings.Contains(browser, "MSIE") {
				seenBrowsers[browser] = nil
				isMSIE = true
				continue
			}
		}

		if !(isAndroid && isMSIE) {
			continue
		}

		// log.Println("Android and MSIE user:", user["name"], user["email"])
		h := strings.IndexRune(user.Email, '@')
		fmt.Fprintf(out, "[%d] %s <%s [at] %s>\n", i, user.Name, user.Email[:h], user.Email[h+1:])
	}

	fmt.Fprintln(out, "\nTotal unique browsers", len(seenBrowsers))
}
