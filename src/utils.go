package richpresence

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type ApplicationInfo struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Icon        string `json:"icon"`
	Description string `json:"description"`
}

func LookupApplication(clientID string) (*ApplicationInfo, error) {
	url := fmt.Sprintf("https://discord.com/api/v10/applications/%s/rpc", clientID)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("discord API returned status %d", resp.StatusCode)
	}

	var info ApplicationInfo
	if err := json.NewDecoder(resp.Body).Decode(&info); err != nil {
		return nil, err
	}

	return &info, nil
}

func GetApplicationIconURL(appID, iconHash string) string {
	if iconHash == "" {
		return ""
	}
	return fmt.Sprintf("https://cdn.discordapp.com/app-icons/%s/%s.png", appID, iconHash)
}

func GetAssetURL(appID, assetKey string) string {
	if assetKey == "" {
		return ""
	}
	if len(assetKey) > 3 && assetKey[:3] == "mp:" {
		return ""
	}
	return fmt.Sprintf("https://cdn.discordapp.com/app-assets/%s/%s.png", appID, assetKey)
}
