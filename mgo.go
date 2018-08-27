package util

import (
	"gopkg.in/mgo.v2/bson"
	"errors"
	"net/http"
	"strings"
)

func ValidateObjectId(s string) error {
	if !bson.IsObjectIdHex(s) {
		return errors.New("object id is not valid")
	}
	return nil
}

func RequestQueryString(r *http.Request, name string, q map[string]interface{}) {
	value := r.URL.Query().Get(name)
	if value!="" {
		q[name]=value
	}
}

func RequestQueryArrayString(r *http.Request, name string, q map[string]interface{}) {
	value := r.URL.Query().Get(name)
	if value!="" {
		arr := strings.Split(value,",")
		q[name]=bson.M{"$in":arr}
	}
}

func RequestQueryDBRefString(r *http.Request, name string, q map[string]interface{}) error {
	value := r.URL.Query().Get(name)
	if value!="" {
		if !bson.IsObjectIdHex(value) {
			return errors.New("object id not valid")
		}
		q[name+".$id"]=bson.ObjectIdHex(value)
	}
	return nil
}

func RequestQueryInteger(r *http.Request, name string, q map[string]interface{}) {
	value, _ := strconv.Atoi(r.URL.Query().Get(name))
	if value!=0 {
		q[name]=value
	}
}

func RequestQueryBoolean(r *http.Request, name string, q map[string]interface{}) {
	value := r.URL.Query().Get(name)
	if value=="false" {
		q[name]=false
	}else {
		q[name]=true
	}
}

func RequestPagingAndSorting(r *http.Request) (int,int,string,error) {
	//Paging
	skip := RequestParamInt(r,"skip")
	limit := RequestParamInt(r,"limit")

	if skip==0 {
		skip=1
	}
	if limit==0 {
		limit=20
	}

	//Sorting
	sort := RequestParamString(r,"sort")
	if sort!="" {
		s:= strings.Split(sort,",")
		if len(s)!=2 {
			return 0,0,"", errors.New("wrong sort format")
		}
		if s[1]=="desc" {
			sort="-"+s[0]
		}else {
			sort=s[0]
		}
	}else {
		sort="createdAt"
	}
	return skip,limit,sort,nil
}