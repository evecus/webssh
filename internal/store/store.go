package store

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sync"
	"time"
)

const DataDir = "data"

var mu sync.RWMutex

// AuthData stores hashed credentials
type AuthData struct {
	Username     string `json:"username"`
	PasswordHash string `json:"password_hash"`
}

// SSHProfile stores an SSH connection profile
type SSHProfile struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Host       string `json:"host"`
	Port       int    `json:"port"`
	Username   string `json:"username"`
	Password   string `json:"password,omitempty"`
	PrivateKey string `json:"private_key,omitempty"`
	Passphrase string `json:"passphrase,omitempty"`
	AuthType   string `json:"auth_type"` // "password" or "key"
	CreatedAt  string `json:"created_at"`
}

// Settings stores UI preferences
type Settings struct {
	Theme    string `json:"theme"`
	UIFont   string `json:"ui_font"`
	TermFont string `json:"term_font"`
	Lang     string `json:"lang"`
}

func EnsureDataDir() error {
	return os.MkdirAll(DataDir, 0750)
}

func filePath(name string) string {
	return filepath.Join(DataDir, name)
}

func readJSON(name string, v interface{}) error {
	mu.RLock()
	defer mu.RUnlock()
	data, err := os.ReadFile(filePath(name))
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}

func writeJSON(name string, v interface{}) error {
	mu.Lock()
	defer mu.Unlock()
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filePath(name), data, 0640)
}

// ---- Auth ----

func LoadAuth() (*AuthData, error) {
	var a AuthData
	if err := readJSON("auth.json", &a); err != nil {
		return nil, err
	}
	return &a, nil
}

func SaveAuth(a *AuthData) error {
	return writeJSON("auth.json", a)
}

func AuthExists() bool {
	_, err := os.Stat(filePath("auth.json"))
	return err == nil
}

// ---- Settings ----

func LoadSettings() (*Settings, error) {
	var s Settings
	if err := readJSON("settings.json", &s); err != nil {
		return &Settings{
			Theme:    "purple-pink",
			UIFont:   "'Outfit','Noto Sans SC',sans-serif",
			TermFont: "'JetBrains Mono',monospace",
			Lang:     "zh",
		}, nil
	}
	return &s, nil
}

func SaveSettings(s *Settings) error {
	return writeJSON("settings.json", s)
}

// ---- SSH Profiles ----

func LoadSSHProfiles() ([]SSHProfile, error) {
	var profiles []SSHProfile
	if err := readJSON("ssh.json", &profiles); err != nil {
		return []SSHProfile{}, nil
	}
	return profiles, nil
}

func SaveSSHProfile(p SSHProfile) ([]SSHProfile, error) {
	profiles, _ := LoadSSHProfiles()
	// Check if updating existing
	for i, existing := range profiles {
		if existing.ID == p.ID {
			profiles[i] = p
			return profiles, writeJSON("ssh.json", profiles)
		}
	}
	if p.ID == "" {
		p.ID = generateID()
	}
	p.CreatedAt = time.Now().Format(time.RFC3339)
	profiles = append(profiles, p)
	return profiles, writeJSON("ssh.json", profiles)
}

func DeleteSSHProfile(id string) ([]SSHProfile, error) {
	profiles, _ := LoadSSHProfiles()
	var updated []SSHProfile
	for _, p := range profiles {
		if p.ID != id {
			updated = append(updated, p)
		}
	}
	if updated == nil {
		updated = []SSHProfile{}
	}
	return updated, writeJSON("ssh.json", updated)
}

func generateID() string {
	return time.Now().Format("20060102150405.000")
}
