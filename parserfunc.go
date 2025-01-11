package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func validateInt(value string) (interface{}, error) {
	return strconv.Atoi(value)
}

func validateGroup(value string) (interface{}, error) {
	group := strings.TrimSpace(value)
	if len([]rune(group)) > 24 {
		err := fmt.Errorf("Group name is too long (%d/24)", len([]rune(group)))
		return "", err
	}
	return group, nil
}

func validateShort(value string) (interface{}, error) {
	short := strings.TrimSpace(value)
	if len([]byte(short)) > 255 {
		err := fmt.Errorf("Group name is too long (%d/255)", len([]byte(short)))
		return "", err
	}
	return short, nil
}

func validateString(value string) (interface{}, error) {
	return strings.TrimSpace(value), nil
}

func validateDate(value string) (interface{}, error) {
	re := regexp.MustCompile(`(\d{4})-{0,1}(\d{2})-{0,1}(\d{2})`)
	result := re.FindAllStringSubmatch(value, -1)
	if len(result) == 0 {
		err := fmt.Errorf("Invalid date %s", value)
		return time.Now, err
	}
	if len(result[0]) != 4 {
		err := fmt.Errorf("Invalid date %s", value)
		return time.Now, err
	}

	dateStr := fmt.Sprintf("%s-%s-%s", result[0][1], result[0][2], result[0][3])
	return time.ParseInLocation(time.DateOnly, dateStr, time.Local)
}
