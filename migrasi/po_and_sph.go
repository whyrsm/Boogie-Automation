package migrasi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func linkPOAndSPH(baseUrL, authToken, poTableID string, poID, sphID interface{}) error {
	url := fmt.Sprintf("%s/api/v2/tables/%s/records", baseUrL, poTableID)
	payload := map[string]interface{}{
		"po_customer_id": poID,
		"Id":             sphID,
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

// POAndSPH ...
func POAndSPH(baseUrL, authToken, poTableID, sphCustomerTableID string) {

	sphRaw, err := fetchTableData(baseUrL, authToken, sphCustomerTableID)
	if err != nil {
		panic("customersRaw: " + err.Error())
	}

	sphMap := make(map[string]string)
	for _, row := range sphRaw {
		name := strings.ToLower(strings.TrimSpace(fmt.Sprint(row["No SPH Customer"])))
		sphMap[name] = fmt.Sprint(row["Id"])
	}

	poRaw, err := fetchTableData(baseUrL, authToken, poTableID)
	if err != nil {
		panic("articleRaw Error: " + err.Error())
	}

	for _, po := range poRaw {
		sphNO := strings.ToLower(strings.TrimSpace(fmt.Sprint(po["No SPH Customer"])))
		poID := fmt.Sprint(po["Id"])

		if sphID, ok := sphMap[sphNO]; ok {
			err := linkArticleAndSPH(baseUrL, authToken, poTableID, poID, sphID)
			if err != nil {
				fmt.Printf("❌ [Article] PO SPH link %s : %v\n", sphNO, err)
			} else {
				fmt.Printf("✅ [Article] PO SPH linked %s\n", sphNO)
			}
		} else {
			fmt.Printf("⚠️  No match for (%s)\n", sphNO)
		}
	}
}
