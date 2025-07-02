'use strict'
require('dotenv').config()
const express = require('express')

const app = express()
const cors = require('cors')
const Routes = require('./router')

app.use(cors())
app.use(express.urlencoded({ extended: true }))
Routes(app)

const server = require('http').createServer(app) 
const PORT = process.env.PORT || process.env.APP_PORT || 4000
if (!module.parent) {
	server.listen(PORT, () => {
		console.log('Express Server Now Running. port:'+PORT)
	})
}
module.exports = app