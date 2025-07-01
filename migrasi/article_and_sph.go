// Package migrasi ...
package migrasi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func linkArticleAndSPH(baseUrL, authToken, articleTableID string, articleID, sphID interface{}) error {
	url := fmt.Sprintf("%s/api/v2/tables/%s/records", baseUrL, articleTableID)
	payload := map[string]interface{}{
		"sph_customer_id": sphID,
		"Id":              articleID,
	}
	data, _ := json.Marshal(payload)

	req, _ := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(data))
	req.Header.Set("xc-token", authToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))

	if resp.StatusCode >= 300 || resp.StatusCode < 200 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to link record: %s", string(body))
	}

	return nil
}

// ArticleAndSPH ...
func ArticleAndSPH(baseUrL, authToken, articleTableID, sphCustomerTableID string) {

	sphRaw, err := fetchTableData(baseUrL, authToken, sphCustomerTableID)
	if err != nil {
		panic("customersRaw: " + err.Error())
	}

	sphMap := make(map[string]string)
	for _, row := range sphRaw {
		name := strings.ToLower(strings.TrimSpace(fmt.Sprint(row["No SPH Customer"])))
		sphMap[name] = fmt.Sprint(row["Id"])
	}

	articleRaw, err := fetchTableData(baseUrL, authToken, articleTableID)
	if err != nil {
		panic("articleRaw Error: " + err.Error())
	}

	b, _ := json.Marshal(sphMap)
	fmt.Println(string(b))

	for _, sph := range articleRaw {
		sphNO := strings.ToLower(strings.TrimSpace(fmt.Sprint(sph["No SPH Customer"])))
		articleID := fmt.Sprint(sph["Id"])

		if sphID, ok := sphMap[sphNO]; ok {
			err := linkArticleAndSPH(baseUrL, authToken, articleTableID, articleID, sphID)
			if err != nil {
				fmt.Printf("❌ [Article] Failed link %s : %v\n", sphNO, err)
			} else {
				fmt.Printf("✅ [Article] Linked SPH %s\n", sphNO)
			}
		} else {
			fmt.Printf("⚠️  No match for (%s)\n", sphNO)
		}
	}
}
