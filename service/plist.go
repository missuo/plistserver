package service

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PlistQuery struct {
	Bundleid string `form:"bundleid"`
	Name     string `form:"name"`
	Version  string `form:"version"`
	Fetchurl string `form:"fetchurl"`
}

const plistTemplate = `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>items</key>
    <array>
        <dict>
            <key>assets</key>
            <array>
                <dict>
                    <key>kind</key>
                    <string>software-package</string>
                    <key>url</key>
                    <string>%s</string>
                </dict>
            </array>
            <key>metadata</key>
            <dict>
                <key>bundle-identifier</key>
                <string>%s</string>
                <key>bundle-version</key>
                <string>%s</string>
                <key>kind</key>
                <string>software</string>
                <key>title</key>
                <string>%s</string>
            </dict>
        </dict>
    </array>
</dict>
</plist>`

func GeneratePlist(c *gin.Context) {
	var query PlistQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		c.String(http.StatusBadRequest, "Invalid query parameters")
		return
	}

	plistXML := fmt.Sprintf(
		plistTemplate,
		query.Fetchurl,
		query.Bundleid,
		query.Version,
		query.Name,
	)

	c.Header("Content-Type", "application/octet-stream")
	c.String(http.StatusOK, plistXML)
}
