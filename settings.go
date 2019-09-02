package godns

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Domain struct
type Domain struct {
	DomainName string   `json:"domain_name"`
	SubDomains []string `json:"sub_domains"`
}

// Notify struct for SMTP notification
type Notify struct {
	Enabled      bool   `json:"enabled"`
	SMTPServer   string `json:"smtp_server"`
	SMTPUsername string `json:"smtp_username"`
	SMTPPassword string `json:"smtp_password"`
	SMTPPort     int    `json:"smtp_port"`
	SendTo       string `json:"send_to"`
}

// Settings struct
type Settings struct {
	Provider    string   `json:"provider"`
	Email       string   `json:"email"`
	Password    string   `json:"password"`
	LoginToken  string   `json:"login_token"`
	Domains     []Domain `json:"domains"`
	IPUrl       string   `json:"ip_url"`
	Interval    int      `json:"interval"`
	UserAgent   string   `json:"user_agent,omitempty"`
	LogPath     string   `json:"log_path"`
	Socks5Proxy string   `json:"socks5_proxy"`
	Notify      Notify   `json:"notify"`
	IPInterface string   `json:"ip_interface"`
	//the code is not ready to update AAAA record
	//IPType      string   `json:"ip_type"`
}

// LoadSettings -- Load settings from config file
func LoadSettings(configPath string, settings *Settings) error {
	// LoadSettings from config file
	file, err := ioutil.ReadFile(configPath)
	if err != nil {
		fmt.Println("Error occurs while reading config file, please make sure config file exists!")
		return err
	}

	err = json.Unmarshal(file, settings)
	if err != nil {
		fmt.Println("Error occurs while unmarshal config file, please make sure config file correct!")
		return err
	}

	if settings.Interval == 0 {
		// set default interval as 5 minutes if interval is 0
		settings.Interval = 5 * 60
	}

	return nil
}
