'use strict'
const axios = require('axios')


async function fetchTableData(baseURL, authToken, tableID) {
    const pageSize = 1000;
    const headers = { 'xc-token': authToken };
    const allRecords = [];

    // 1. Get total count
    const countURL = `${baseURL}/api/v2/tables/${tableID}/records/count`;
    const countResp = await axios.get(countURL, { headers });
    const count = countResp.data.count;
    console.log(`📊 Total records for table ${tableID}: ${count}`);

    // 2. Fetch All Data, Paginate with offset & limit
    for (let offset = 0; offset < count; offset += pageSize) {
        const pageURL = `${baseURL}/api/v2/tables/${tableID}/records?limit=${pageSize}&offset=${offset}`;
        const resp = await axios.get(pageURL, { headers });

        const records = resp.data.list;
        allRecords.push(...records);
        console.log(`📥 Fetched ${records.length} records (offset: ${offset})`);
    }

    return allRecords;
}


async function linkData(baseURL, authToken, tableID, payload) {
    const url = `${baseURL}/api/v2/tables/${tableID}/records`;

    try {
        const res = await axios.patch(url, payload, {
            headers: {
                'xc-token': authToken,
                'Content-Type': 'application/json'
            }
        });
        return true;
    } catch (err) {
        const status = err.response?.status;
        const data = err.response?.data;

        const errorMsg = `❌ Failed to link data | Status: ${status} | Response: ${JSON.stringify(data, null, 2)}`.trim();

        throw new Error(errorMsg);
    }
}

module.exports = {
    fetchTableData,
    linkData
};