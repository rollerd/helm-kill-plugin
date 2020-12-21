package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"syscall"
)

func main() {
	// crude argument parsing
	if os.Args[1] == "--help" {
		fmt.Println("Delete chart from http repo\nRequires environment variables 'USER=<username>' and 'HELM_HTTP_URL=<chart repo url>' be set\n\nUsage: helm kill <chartname> <chart_version>\n")
		os.Exit(0)
	} else if len(os.Args) < 3 {
		fmt.Println("Error: Missing required number of args: 'charts' 'version'")
		os.Exit(1)
	}

	// create required string variables
	user := getEnvVar("USER")
	base_url := getEnvVar("HELM_HTTP_URL")
	password := getPassword(fmt.Sprintf("Enter password for %s: \n", user))
	user_pwd_string := fmt.Sprintf("%s:%s", user, password)
	b64_encoded_user_pwd := b64Encode(user_pwd_string)

	// call the http endpoint to delete the chart
	deleteChart(b64_encoded_user_pwd, base_url, os.Args[1], os.Args[2])
}

// get user from environment
func getEnvVar(envvar string) string {
	value := os.Getenv(envvar)
	if value == "" {
		fmt.Printf("Error: Could not find %s environment variable, please export %s=<value>\n", envvar, envvar)
		os.Exit(1)
	}

	return value
}

// create URL and auth strings and call endpoint to delete chart
// https://stackoverflow.com/a/37091538
func deleteChart(b64_user_pwd, base_url, chart, version string) {
	url := fmt.Sprintf("%s/%s/%s", base_url, chart, version)
	fmt.Printf("Removing chart at: %s\n", string(url))
	auth_string := fmt.Sprintf("Basic %s", b64_user_pwd)
	req, _ := http.NewRequest(http.MethodDelete, url, nil)
	// add basic auth header to request
	req.Header.Add("authorization", auth_string)

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	respStatus := resp.Status
	fmt.Println(respStatus)
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
}

// base64 encode the username:password to avoid special char issues when calling URL
func b64Encode(string_to_encode string) string {
	data := []byte(string_to_encode)
	str := base64.StdEncoding.EncodeToString(data)

	return str
}

// request user enter password without echoing input
func getPassword(prompt string) string {
	fmt.Print(prompt)

	// common settings and variables for both stty calls
	attrs := syscall.ProcAttr{
		Dir:   "",
		Env:   []string{},
		Files: []uintptr{os.Stdin.Fd(), os.Stdout.Fd(), os.Stderr.Fd()},
		Sys:   nil}

	var ws syscall.WaitStatus

	// disable echoing
	pid, err := syscall.ForkExec(
		"/bin/stty",
		[]string{"stty", "-echo"},
		&attrs)
	if err != nil {
		panic(err)
	}

	// wait for the stty process to complete
	_, err = syscall.Wait4(pid, &ws, 0, nil)
	if err != nil {
		panic(err)
	}

	// echo is disabled, now grab the data
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	// re-enable echo
	pid, err = syscall.ForkExec(
		"/bin/stty",
		[]string{"stty", "echo"},
		&attrs)
	if err != nil {
		panic(err)
	}

	// wait for the stty process to complete
	_, err = syscall.Wait4(pid, &ws, 0, nil)
	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(text)
}
