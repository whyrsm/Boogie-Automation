package migrasi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func fetchTableData(baseURL, authToken, tableID string) ([]map[string]interface{}, error) {
	const pageSize = 1000
	var allRecords []map[string]interface{}

	// 1. Get total count
	countURL := fmt.Sprintf("%s/api/v2/tables/%s/records/count", baseURL, tableID)
	countReq, _ := http.NewRequest("GET", countURL, nil)
	countReq.Header.Set("xc-token", authToken)

	countResp, err := http.DefaultClient.Do(countReq)
	if err != nil {
		return nil, err
	}
	defer countResp.Body.Close()

	var countRes struct {
		Count int `json:"count"`
	}
	if err := json.NewDecoder(countResp.Body).Decode(&countRes); err != nil {
		return nil, err
	}
	fmt.Printf("📊 Total records for table %s: %d\n", tableID, countRes.Count)

	// 2. Loop with offset & limit
	for offset := 0; offset < countRes.Count; offset += pageSize {
		url := fmt.Sprintf("%s/api/v2/tables/%s/records?limit=%d&offset=%d", baseURL, tableID, pageSize, offset)
		req, _ := http.NewRequest(http.MethodGet, url, nil)
		req.Header.Set("xc-token", authToken)

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		var page struct {
			List []map[string]interface{} `json:"list"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&page); err != nil {
			return nil, err
		}

		// b, _ := json.Marshal(page)
		// fmt.Println("page:", string(b))

		allRecords = append(allRecords, page.List...)
		fmt.Printf("📥 Fetched %d records (offset: %d)\n", len(page.List), offset)
	}

	return allRecords, nil
}
