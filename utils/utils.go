package utils

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
)

func RootDir() *string {
	_, b, _, _ := runtime.Caller(0)
	rootDir := filepath.Join(filepath.Dir(b), "..")
	return &rootDir
}

func GoDotEnvVariable(key string) *string {
	envFileName := filepath.Join(*RootDir(), ".env")
	err := godotenv.Load(envFileName)
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	value := os.Getenv(key)
	return &value
}

type PostgresBackup struct {
	Name     string
	Password string
	Host     string
	Database string
	Port     string
}

func Setup() (*PostgresBackup, error) {
	var DbHost string
	if runtime.GOOS != "windows" {
		DbHost = *GoDotEnvVariable("DB_HOST")
	} else {
		DbHost = "localhost"
	}
	return &PostgresBackup{
		Name:     *GoDotEnvVariable("DB_USER"),
		Password: *GoDotEnvVariable("DB_PASSWORD"),
		Host:     DbHost,
		Port:     *GoDotEnvVariable("DB_PORT"),
		Database: *GoDotEnvVariable("DB_NAME"),
	}, nil
}

func (s *PostgresBackup) CreateDump() error {
	var args []string
	if s.Host != "" {
		args = append(args, "-h", s.Host)
	} else {
		args = append(args, "-h", "127.0.0.1")
	}
	if s.Port != "" {
		args = append(args, "-p", s.Port)
	}
	args = append(args, "-Fc")
	if s.Name != "" {
		args = append(args, "-U", s.Name)
	}
	if s.Password != "" {
		_ = os.Setenv("PGPASSWORD", s.Password)
	}
	if s.Database == "" {
		return errors.New("Missing database name")
	}
	args = append(args, s.Database)
	out, err := os.OpenFile(path.Join(*RootDir(), "my_backup.sql"), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	cmd := exec.Command("pg_dump", args...)
	cmd.Stdout = out
	fmt.Println(args)
	return cmd.Run()
}

func (s *PostgresBackup) ApplyDump() error {
	var args []string
	if s.Host != "" {
		args = append(args, "-h", s.Host)
	} else {
		args = append(args, "-h", "127.0.0.1")
	}
	if s.Port != "" {
		args = append(args, "-p", s.Port)
	}
	args = append(args, "-Fc", "--clean")
	if s.Name != "" {
		args = append(args, "-U", s.Name)
	}
	if s.Password != "" {
		_ = os.Setenv("PGPASSWORD", s.Password)
	}
	if s.Database == "" {
		return errors.New("Missing database name")
	}
	args = append(args, "-d", s.Database)
	args = append(args, filepath.Join(*RootDir(), "my_backup.sql"))
	cmd := exec.Command("pg_restore", args...)
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return nil
}
