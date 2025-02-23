import * as dotenv from "dotenv"

dotenv.config()

export type Config = {
  protocol: string
  host: string
  port: number
}

export const _config = (): Config => {
  const port = process.env.COMINGSOON_PORT || "3000"
  console.log("port : ", port)
  console.log("typeof port : ", typeof port)
  return {
    protocol: process.env.COMINGSOON_PROTOCOL || "http",
    host: process.env.COMINGSOON_HOST || "localhost",
    port: parseInt(port, 10),
  }
}