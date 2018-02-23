package main

import (
	"io/ioutil"
	"regexp"
)

// return []User include all user info
func GetUsers(ssusersPath, sstrafficPath string) []User {
	var multPort []string
	var multPassword []string
	var multLimit []string
	var multUsed []string
	var multRemaining []string
	var users []User
	// load ssusers (ssadmin showpw > ssusers)
	ssusers, _ := ioutil.ReadFile(ssusersPath)
	re := regexp.MustCompile("(?m)^(\\d+)\\s+(.*?)\\s+")
	multUser := re.FindAll(ssusers, -1)
	for _, u := range multUser {
		uString := string(u)
		multPort = append(multPort, re.ReplaceAllString(uString, "$1"))
		multPassword = append(multPassword, re.ReplaceAllString(uString, "$2"))
	}
	// append "total" port which is not recorded in ssusers file
	multPort = append(multPort, "total")
	multPassword = append(multPassword, "nil")
	// load sstraffic (ssadmin show > ssusers)
	sstraffic, _ := ioutil.ReadFile(sstrafficPath)
	re = regexp.MustCompile("(?m)^([^#\\s]+)\\s+\\d+\\((.*?)\\)\\s+\\d+\\((.*?)\\)\\s+\\d+\\((.*?)\\)\\s+")
	multTraffic := re.FindAll(sstraffic, -1)
	for _, t := range multTraffic {
		tString := string(t)
		multLimit = append(multLimit, re.ReplaceAllString(tString, "$2"))
		multUsed = append(multUsed, re.ReplaceAllString(tString, "$3"))
		multRemaining = append(multRemaining, re.ReplaceAllString(tString, "$4"))
	}
	// build users slice include User structs
	for i := 0; i < len(multPort); i++ {
		users = append(users, User{
			Port:      multPort[i],
			Password:  multPassword[i],
			Limit:     multLimit[i],
			Used:      multUsed[i],
			Remaining: multRemaining[i]})
	}
	return users
}

func checkUsernameAndPassword(username, password string) bool {
	if username != USERNAME || password != PASSWORD {
		return false
	}
	return true
}

func CheckLogin(username, password string) string {
	var loginPrompt string
	if checkUsernameAndPassword(username, password) {
		loginPrompt = "Login Success"
	} else if username == "" || password == "" {
		loginPrompt = "Invalid username or password"
	} else {
		loginPrompt = "Incorrect username or password"
	}
	return loginPrompt
}
