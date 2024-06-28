package auth

import (
	"github.com/ua-parser/uap-go/uaparser"
)

func GetSessionName(userAgent string) (string, error) {

	var sessionName string

	parser, err := uaparser.New("uaparser.yaml")
	if err != nil {
		return sessionName, err
	}

	ua := parser.Parse(userAgent)

	return ua.Os.ToString() + " - " + ua.UserAgent.ToString(), nil

}
