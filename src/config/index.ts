import * as dotenv from "dotenv"

dotenv.config()

export type Config = {
  protocol: string
  host: string
  port: number
  forProject: string
}

export const _config = (): Config => ({
  protocol: process.env.COMINGSOON_PROTOCOL || "http",
  host: process.env.COMINGSOON_HOST || "localhost",
  port: parseInt(process.env.COMINGSOON_PORT || "3000", 10),
  forProject: process.env.COMINGSOON_FOR_PROJECT || "unknown",
})