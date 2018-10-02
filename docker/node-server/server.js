const express = require('express')
const app = express()

const PORT = process.env.PORT || 3000

app.get('/', (_, res) => res.send('welcome to node server'))

app.listen(PORT, () => console.log(`node server running on port *${PORT}`))