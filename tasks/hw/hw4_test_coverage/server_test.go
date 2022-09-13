package main

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_new_server(t *testing.T) {
	s, err := new_server("/", nil)
	assert.Error(t, err)
	assert.Nil(t, s)

	s, err = new_server(filename, nil)
	assert.NoError(t, err)
	assert.NotNil(t, s)

	tokens := []string{"lel", "kek", "xd"}
	s, err = new_server(filename, tokens)
	assert.NoError(t, err)
	assert.NotNil(t, s)
	m := map[string]any{}
	for _, token := range tokens {
		m[token] = nil
	}
	assert.Equal(t, m, s.tokens)
}

func Test_server_SearchRequest(t *testing.T) {
	type fields struct {
		tokens map[string]any
		r      *root
		users  []*User
	}
	type args struct {
		query string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantRv  *SearchRequest
		wantErr bool
	}{
		{"00", fields{}, args{query: "&&&;;;??======"}, nil, true},
		{"01", fields{}, args{query: "limit=x"}, nil, true},
		{"02", fields{}, args{query: "limit=-1"}, nil, true},
		{"03", fields{}, args{query: "limit=26"}, nil, true},
		{"10", fields{}, args{query: "limit=1&"}, nil, true},
		{"11", fields{}, args{query: "limit=1&offset=x"}, nil, true},
		{"12", fields{}, args{query: "limit=1&offset=-1"}, nil, true},
		{"13", fields{}, args{query: "limit=1&offset=26"}, nil, true},
		{"14", fields{users: []*User{nil, nil}}, args{query: "limit=2&offset=1"}, nil, true},
		{"20", fields{}, args{query: "limit=1&offset=0&query=x"}, nil, true},
		{"30", fields{}, args{query: "limit=1&offset=0&query=x&order_field"}, nil, true},
		{"31", fields{}, args{query: "limit=1&offset=0&query=x&order_field="}, nil, true},
		{"32", fields{}, args{query: "limit=1&offset=0&query=x&order_field=hehe"}, nil, true},
		{"40", fields{}, args{query: "limit=1&offset=0&query=x&order_field=&order_by="}, nil, true},
		{"41", fields{}, args{query: "limit=1&offset=0&query=x&order_field=&order_by=2"}, nil, true},
		{
			"50",
			fields{users: []*User{nil, nil}},
			args{query: "limit=1&offset=0&query=Kek&order_field=&order_by=1"},
			&SearchRequest{Limit: 1, Offset: 0, Query: "Kek", OrderField: "", OrderBy: 1},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &server{
				tokens: tt.fields.tokens,
				r:      tt.fields.r,
				users:  tt.fields.users,
			}
			gotRv, err := s.SearchRequest(tt.args.query)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)

			}
			assert.Equal(t, gotRv, tt.wantRv)
		})
	}
}

func Test_server_SearchServer(t *testing.T) {
	s, err := new_server(filename, nil)
	assert.NotNil(t, s)
	assert.NoError(t, err)
	s.users = append(s.users, s.users[len(s.users)-1])

	tests := []struct {
		name   string
		sr     *SearchRequest
		wantRv []*User
	}{
		{"00", &SearchRequest{Limit: 1}, []*User{s.users[0]}},
		{"01", &SearchRequest{Limit: 2}, []*User{s.users[0], s.users[1]}},
		{"02", &SearchRequest{Limit: 3}, []*User{s.users[0], s.users[1], s.users[2]}},
		{"03", &SearchRequest{Limit: 4}, []*User{s.users[0], s.users[1], s.users[2], s.users[3]}},
		{"10", &SearchRequest{Limit: 1, Offset: 1}, []*User{s.users[1]}},
		{"11", &SearchRequest{Limit: 2, Offset: 2}, []*User{s.users[2], s.users[3]}},
		{"12", &SearchRequest{Limit: 3, Offset: 3}, []*User{s.users[3], s.users[4], s.users[5]}},
		{"13", &SearchRequest{Limit: 4, Offset: 4}, []*User{s.users[4], s.users[5], s.users[6], s.users[7]}},
		{"20", &SearchRequest{Limit: 4, OrderBy: OrderByAsc, OrderField: ""}, []*User{s.users[0], s.users[1], s.users[2], s.users[3]}},
		{"21", &SearchRequest{Limit: 4, OrderBy: OrderByAsc, OrderField: "Id"}, []*User{s.users[0], s.users[1], s.users[2], s.users[3]}},
		{"22", &SearchRequest{Limit: 4, OrderBy: OrderByAsc, OrderField: "Age"}, []*User{s.users[1], s.users[0], s.users[2], s.users[3]}},
		{"23", &SearchRequest{Limit: 4, OrderBy: OrderByAsc, OrderField: "Name"}, []*User{s.users[0], s.users[2], s.users[3], s.users[1]}},
		{"30", &SearchRequest{Limit: 4, OrderBy: orderDesc, OrderField: ""}, []*User{s.users[0], s.users[1], s.users[2], s.users[3]}},
		{"31", &SearchRequest{Limit: 4, OrderBy: orderDesc, OrderField: "Id"}, []*User{s.users[3], s.users[2], s.users[1], s.users[0]}},
		{"32", &SearchRequest{Limit: 4, OrderBy: orderDesc, OrderField: "Age"}, []*User{s.users[3], s.users[2], s.users[0], s.users[1]}},
		{"33", &SearchRequest{Limit: 4, OrderBy: orderDesc, OrderField: "Name"}, []*User{s.users[1], s.users[3], s.users[2], s.users[0]}},
		{"44", &SearchRequest{Limit: 4, Offset: 4, OrderBy: OrderByAsc, OrderField: ""}, []*User{s.users[4], s.users[5], s.users[6], s.users[7]}},
		{"45", &SearchRequest{Limit: 4, Offset: 4, OrderBy: OrderByAsc, OrderField: "Id"}, []*User{s.users[4], s.users[5], s.users[6], s.users[7]}},
		{"46", &SearchRequest{Limit: 4, Offset: 4, OrderBy: OrderByAsc, OrderField: "Age"}, []*User{s.users[4], s.users[5], s.users[7], s.users[6]}},
		{"47", &SearchRequest{Limit: 4, Offset: 4, OrderBy: OrderByAsc, OrderField: "Name"}, []*User{s.users[5], s.users[6], s.users[7], s.users[4]}},
		{"50", &SearchRequest{Limit: 2, Offset: 34, OrderBy: OrderByAsc, OrderField: "Name"}, []*User{s.users[34], s.users[35]}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRv := s.SearchServer(tt.sr)
			assert.Equal(t, gotRv, tt.wantRv)
		})
	}
}

func Test_server_serveHTTP(t *testing.T) {
	s, err := new_server(filename, []string{"xd"})
	assert.NotNil(t, s)
	assert.NoError(t, err)

	type args struct {
		token string
		query string
	}
	tests := []struct {
		name   string
		args   args
		status int
		data   any
	}{
		{"00", args{token: "xD"}, http.StatusUnauthorized, ""},

		{"10", args{token: "xd", query: ""}, http.StatusBadRequest, "missed field"},
		{"11", args{token: "xd", query: "limit=1"}, http.StatusBadRequest, "missed field"},
		{"12", args{token: "xd", query: "limit=1&offset=0"}, http.StatusBadRequest, "missed field"},
		{"13", args{token: "xd", query: "limit=1&offset=0&query=x"}, http.StatusBadRequest, "missed field"},
		{"14", args{token: "xd", query: "limit=1&offset=0&query=x&order_field=Name"}, http.StatusBadRequest, "missed field"},

		{"20", args{token: "xd", query: "limit=1&offset=0&query=x&order_field=Name&order_by=0"}, http.StatusOK, []User{*s.users[0]}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			status, data := s.serveHTTP(tt.args.token, tt.args.query)
			assert.Equal(t, tt.status, status)
			if data == nil {
				return
			}
			if status == http.StatusBadRequest {
				er := &SearchErrorResponse{}
				err := json.Unmarshal(data, &er)
				assert.NoError(t, err)
				assert.Equal(t, tt.data, er.Error)
			}
			if status == http.StatusOK {
				users := []User{}
				err := json.Unmarshal(data, &users)
				assert.NoError(t, err)
				assert.Equal(t, tt.data, users)
			}
		})
	}
}
