package config

import (
	"fmt"
	"os"
	"strings"
)

func GetGCPProjectID() string {
	return os.Getenv("GCP_PROJECT_ID")
}

func ConnectDBInfo() string {
	type opt struct {
		Key   string
		Value string
	}
	type opts []*opt

	options := opts{
		{Key: "parseTime", Value: "true"},
	}
	info := strings.Builder{}
	var base string
	if GetDBPort() != "" {
		base = fmt.Sprintf("%s:%s@%s(%s:%s)/%s", GetDBUserName(), GetDBPassword(), GetDBConnMethod(), GetDBHostName(), GetDBPort(), GetDatabaseName())
	} else {
		base = fmt.Sprintf("%s:%s@%s(%s)/%s", GetDBUserName(), GetDBPassword(), GetDBConnMethod(), GetDBHostName(), GetDatabaseName())
	}
	info.WriteString(base)

	for _, option := range options {
		info.WriteString(fmt.Sprintf("?%s=%s", option.Key, option.Value))
	}

	return info.String()
}

func GetDBUserName() string {
	defaultName := "root"
	if userName := os.Getenv("DATABASE_USER_NAME"); userName != "" {
		return userName
	} else {
		return defaultName
	}
}

func GetDBPassword() string {
	defaultPassword := "root"
	if password := os.Getenv("DATABASE_PASSWORD"); password != "" {
		return password
	} else {
		return defaultPassword
	}
}

func GetDBConnMethod() string {
	defaultConnMethod := "tcp"
	if ConnMethod := os.Getenv("DATABASE_CONN_METHOD"); ConnMethod != "" {
		return ConnMethod
	} else {
		return defaultConnMethod
	}
}

func GetDBHostName() string {
	defaultHostname := "rdb"
	if hostname := os.Getenv("DATABASE_HOSTNAME"); hostname != "" {
		return hostname
	} else {
		return defaultHostname
	}
}

func GetDBPort() string {
	defaultPort := "3306"
	if port := os.Getenv("DATABASE_PORT"); port != "" {
		return port
	} else {
		return defaultPort
	}
}

func GetDatabaseName() string {
	defaultDatabaseName := "ks-laboratory-backend"
	if databaseName := os.Getenv("DATABASE_NAME"); databaseName != "" {
		return databaseName
	} else {
		return defaultDatabaseName
	}
}
