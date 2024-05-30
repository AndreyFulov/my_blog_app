const {createServer} = require('node:http')

const server = createServer((req,res) => {
    res.writeHead(200, {"Content-Type":"text/plain"});
    res.end()
})

server.listen(3000, "127.0.0.1",() => {
    console.log("listening server!")
})