import express, { Express } from "express"
import bodyParser from "body-parser"
import { RequestIdGenerator } from "./middleware"
import { _config } from "./config"
import cors from "cors"

const __config = _config()
const app: Express = express()

app.use(cors())
app.use(bodyParser.urlencoded({ extended: false }))
app.use(bodyParser.json())
app.use(RequestIdGenerator())

app.get('/api/HealthCheck/Ping', (_: any, res: any) => res.status(200).send({
  statusCode: 200,
  message: "Ok - Project : " + __config.forProject
}))

app.use((error: Error, _req: any, res: any, _next: any) => res.status(500).send({
  statusCode: 500,
  message: error.message
}))

app.listen(__config.port, () => console.log(`[server]: Coming Soon API for project "${__config.forProject}" is running at: ${__config.protocol}://${__config.host}:${__config.port}`))
