'use strict'
const axios = require('axios');
const {fetchTableData,linkData} = require('./helper');

async function customerAndPO(baseURL, authToken, customerTableID, POCustomerTableID) {
    const customersRaw = await fetchTableData(baseURL, authToken, customerTableID);
    const customerMap = {};

    customersRaw.forEach(row => {
        const name = String(row["Nama"] || '').trim().toLowerCase();
        customerMap[name] = String(row["Id"]);
    });

    const poRaw = await fetchTableData(baseURL, authToken, POCustomerTableID);

    for (const po of poRaw) {
        const custName = String(po["Nama Account"] || '').trim().toLowerCase();
        const poID = String(po["Id"]);

        const custID = customerMap[custName];

        if (custID) {
            try {
                await linkData(baseURL, authToken, POCustomerTableID, { Id: poID, customer_id: custID });
                console.log(`✅ [Customer - PO] Linked ${custName} (${custID}) -> PO ${poID}`);
            } catch (err) {
                console.log(`❌ [Customer - PO] Failed link ${custName} (${custID}): ${err.message}`);
            }
        } else {
            console.log(`⚠️  No match for ${custName}`);
        }
    }
}

module.exports = customerAndPO;