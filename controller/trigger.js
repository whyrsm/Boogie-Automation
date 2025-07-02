'use strict'

const customerAndPO = require('../service/customerAndPO')
const customerAndSPH = require('../service/customerAndSPH')
const articleAndSPH = require('../service/articleAndSPH')
const poAndSPH = require('../service/poAndSPH')

const triggerSync = async (req, res, next) => {
    try {
        const { type } = req.query;

        // kalo ga ada type, maka jalankan semua
        if (!type) {
            await customerAndPO(process.env.NOCO_URL, process.env.NOCO_TOKEN, process.env.CUSTOMER_TABLE_ID, process.env.PO_CUSTOMER_TABLE_ID)
            await customerAndSPH(process.env.NOCO_URL, process.env.NOCO_TOKEN, process.env.CUSTOMER_TABLE_ID, process.env.SPH_CUSTOMER_TABLE_ID)
            await articleAndSPH(process.env.NOCO_URL, process.env.NOCO_TOKEN, process.env.ARTICLE_TABLE_ID, process.env.SPH_CUSTOMER_TABLE_ID)
            await poAndSPH(process.env.NOCO_URL, process.env.NOCO_TOKEN, process.env.PO_CUSTOMER_TABLE_ID, process.env.SPH_CUSTOMER_TABLE_ID)
        }

        switch (type) {
            case 'customer-po':
                await customerAndPO(process.env.NOCO_URL, process.env.NOCO_TOKEN, process.env.CUSTOMER_TABLE_ID, process.env.PO_CUSTOMER_TABLE_ID)
                break;
            case 'customer-sph':
                await customerAndSPH(process.env.NOCO_URL, process.env.NOCO_TOKEN, process.env.CUSTOMER_TABLE_ID, process.env.SPH_CUSTOMER_TABLE_ID)
                break;
            case 'article-sph':
                await articleAndSPH(process.env.NOCO_URL, process.env.NOCO_TOKEN, process.env.ARTICLE_TABLE_ID, process.env.SPH_CUSTOMER_TABLE_ID)
                break;
            case 'po-sph':
                await poAndSPH(process.env.NOCO_URL, process.env.NOCO_TOKEN, process.env.PO_CUSTOMER_TABLE_ID, process.env.SPH_CUSTOMER_TABLE_ID)
                break;
            default:
                break;
        }

        return res.status(200).json({ message: "done" });
    } catch (err) {
        next(err); // Biarkan error handler middleware yang proses
    }
}

module.exports = {
    triggerSync
}