'use strict'
const axios = require('axios');
const {fetchTableData,linkData} = require('./helper');

async function articleAndSPH(baseURL, authToken, articleTableID, sphCustomerTableID) {
    const sphRaw = await fetchTableData(baseURL, authToken, sphCustomerTableID);
    const sphMap = {};

    sphRaw.forEach(row => {
        const name = String(row["No SPH Customer"] || '').trim().toLowerCase();
        sphMap[name] = String(row["Id"]);
    });

    const articleRaw = await fetchTableData(baseURL, authToken, articleTableID);

    for (const article of articleRaw) {
        const sphNo = String(article["No SPH Customer"] || '').trim().toLowerCase();
        const articleID = String(article["Id"]);

        const sphID = sphMap[sphNo];

        if (sphID) {
            try {
                await linkData(baseURL, authToken, articleTableID, { Id: articleID, sph_customer_id: sphID });
                console.log(`✅ [Article - SPH] Linked ${article["Nama Article"]} -> SPH: ${sphNo}`);
            } catch (err) {
                console.log(`❌ [Article - SPH] Failed link ${article["Nama Article"]} -> SPH: ${sphNo}: ${err.message}`);
            }
        } else {
            console.log(`⚠️  No match for ${article["Nama Article"]} -> SPH: ${sphNo}`);
        }
    }
}

module.exports = articleAndSPH;