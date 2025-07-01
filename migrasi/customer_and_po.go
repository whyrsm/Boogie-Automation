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

func linkCustomerToPO(baseURL, authToken, tableID string, sphID, customerID interface{}) error {
	url := fmt.Sprintf("%s/api/v2/tables/%s/records", baseURL, tableID)
	payload := map[string]interface{}{
		"customer_id": customerID,
		"Id":          sphID,
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

// CustomerAndPO ...
func CustomerAndPO(baseUrL, authToken, customerTableID, POCustomerTableID string) {

	customersRaw, err := fetchTableData(baseUrL, authToken, customerTableID)
	if err != nil {
		panic("customersRaw: " + err.Error())
	}

	customerMap := make(map[string]string)
	for _, row := range customersRaw {
		name := strings.ToLower(strings.TrimSpace(fmt.Sprint(row["Nama"])))
		customerMap[name] = fmt.Sprint(row["Id"])
	}

	poRaw, err := fetchTableData(baseUrL, authToken, POCustomerTableID)
	if err != nil {
		panic("poRaw Error: " + err.Error())
	}

	for _, po := range poRaw {
		custName := strings.ToLower(strings.TrimSpace(fmt.Sprint(po["Nama Account"])))
		sphID := fmt.Sprint(po["Id"])
		if custID, ok := customerMap[custName]; ok {
			err := linkCustomerToSPH(baseUrL, authToken, POCustomerTableID, sphID, custID)
			if err != nil {
				fmt.Printf("❌ [PO] Failed link %s (%s): %v\n", custName, custID, err)
			} else {
				fmt.Printf("✅ [PO] Linked %s (%s) -> PO %s\n", custName, custID, sphID)
			}
		} else {
			fmt.Printf("⚠️  No match for %s (%s)\n", custName, custID)
		}
	}
}
