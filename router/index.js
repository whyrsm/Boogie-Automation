'use strict'
const triggerCtrl = require('../controller/trigger')

module.exports = app =>{
	app.get('/api/health', (req, res) => {
		res.status(200).send({
			message: 'api is up and running',
		})
	})
	app.get('/api/sync', triggerCtrl.triggerSync)
}
