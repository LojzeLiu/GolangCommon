package Common

import (
	"crypto/md5"
	"errors"
	"fmt"
	"regexp"
	"sort"
	"strings"
)

type HttpRequestSecurity struct {
	TokenUrl string
}

func RequestLegalCheck(req string) bool {
	regexpReq := regexp.MustCompile(`(^[a-zA-Z0-9_]*=[^&=]*&{0,1})([a-zA-Z0-9_]*=[^=&]*&)*([a-zA-Z0-9_]*=[^=&]*){0,1}([^&]$)`)

	return regexpReq.MatchString(req)
}

func (h *HttpRequestSecurity) SignCheck(req string) error {
	//检查请求是否合法
	requestInfo := req
	reqs := strings.Split(req, "?")
	if len(reqs) == 2 {
		if !RequestLegalCheck(reqs[1]) {
			return errors.New("The request is not legal.")
		}
		requestInfo = reqs[1]
	}

	//转换为Map
	MapReques := make(map[string]string)
	var Ssign string
	var Skey []string
	requestSplit := strings.Split(requestInfo, "&")
	reqLen := len(requestSplit)
	for i := 0; i < reqLen; i++ {
		currReq := requestSplit[i]
		currRequest := strings.Split(currReq, "=")
		if len(currRequest) != 2 {
			continue
		}
		if strings.Compare(string(currRequest[0]), "sign") == 0 {
			Ssign = string(currRequest[1])
			continue
		}
		MapReques[currRequest[0]] = currRequest[1]
		Skey = append(Skey, currRequest[0])
	}

	//排序
	var strReq string
	sort.Strings(Skey)
	for _, key := range Skey {
		strReq += key + MapReques[key]
	}

	//生成MD5串
	DEBUG("Request: ", strReq, ";")
	MD5Sum := fmt.Sprintf("%x", md5.Sum([]byte(strReq)))
	DEBUG("MD5:", MD5Sum)

	//对比
	if strings.Compare(MD5Sum, Ssign) != 0 {
		return errors.New("The Sign is invalid.")
	}

	return nil
}
