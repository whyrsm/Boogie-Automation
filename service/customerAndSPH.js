'use strict'
const axios = require('axios');
const {fetchTableData,linkData} = require('./helper');

async function customerAndSPH(baseURL, authToken, customerTableID, sphCustomerTableID) {
    const customersRaw = await fetchTableData(baseURL, authToken, customerTableID);
    const customerMap = {};

    customersRaw.forEach(row => {
        const name = String(row["Nama"] || '').trim().toLowerCase();
        customerMap[name] = String(row["Id"]);
    });

    const sphRaw = await fetchTableData(baseURL, authToken, sphCustomerTableID);

    for (const sph of sphRaw) {
        const custName = String(sph["Nama Account"] || '').trim().toLowerCase();
        const sphID = String(sph["Id"]);

        const custID = customerMap[custName];

        if (custID) {
            try {
                await linkData(baseURL, authToken, sphCustomerTableID, { Id: sphID, customer_id: custID });
                console.log(`✅ [Customer - SPH] Linked ${custName} (${custID}) -> SPH ${sphID}`);
            } catch (err) {
                console.log(`❌ [Customer - SPH] Failed link ${custName} (${custID}): ${err.message}`);
            }
        } else {
            console.log(`⚠️  No match for ${custName}`);
        }
    }
}

module.exports = customerAndSPH;