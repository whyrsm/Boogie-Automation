'use strict'
const axios = require('axios');
const {fetchTableData,linkData} = require('./helper');

async function poAndSPH(baseURL, authToken, poTableID, sphCustomerTableID) {
    const sphRaw = await fetchTableData(baseURL, authToken, sphCustomerTableID);
    const sphMap = {};

    sphRaw.forEach(row => {
        const name = String(row["No SPH Customer"] || '').trim().toLowerCase();
        sphMap[name] = String(row["Id"]);
    });

    const poRaw = await fetchTableData(baseURL, authToken, poTableID);

    for (const po of poRaw) {
        const sphNO = String(po["No SPH Customer"] || '').trim().toLowerCase();
        const poID = String(po["Id"]);

        const sphID = sphMap[sphNO];

        if (sphID) {
            try {
                await linkData(baseURL, authToken, poTableID, { Id: poID, sph_customer_id: sphID });
                console.log(`✅ [SPH - PO] Linked ${sphNO}`);
            } catch (err) {
                console.log(`❌ [SPH - PO] Failed link ${sphNO}: ${err.message}`);
            }
        } else {
            console.log(`⚠️  No match for SPH: ${sphNO}`);
        }
    }
}

module.exports = poAndSPH;