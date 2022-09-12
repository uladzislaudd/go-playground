package fast

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

type (
	//easyjson:json
	User struct {
		Browsers []string `json:"browsers,omitempty,nocopy"`
		Company  string   `json:"company,omitempty,nocopy"`
		Country  string   `json:"country,omitempty,nocopy"`
		Email    string   `json:"email,omitempty,nocopy"`
		Job      string   `json:"job,omitempty,nocopy"`
		Name     string   `json:"name,omitempty,nocopy"`
		Phone    string   `json:"phone,omitempty,nocopy"`
	}
)

func FastSearch(out io.Writer, data []byte) {
	seenBrowsers := map[string]interface{}{}

	user := User{}
	fmt.Fprintln(out, "found users:")
	for i, line := range bytes.Split(data, []byte("\n")) {
		// fmt.Printf("%v %v\n", err, line)
		err := user.UnmarshalJSON(line)
		if err != nil {
			panic(err)
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
		j := strings.IndexRune(user.Email, '@')
		fmt.Fprintf(out, "[%d] %s <%s [at] %s>\n", i, user.Name, user.Email[:j], user.Email[j+1:])
	}

	fmt.Fprintln(out, "\nTotal unique browsers", len(seenBrowsers))
}
