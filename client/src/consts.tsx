export const wsbaseUrl = process.env.NODE_ENV === "development" ? "ws://localhost:8080" : "ws://endlesschat.net/v2"
export const httpbaseUrl = process.env.NODE_ENV === "development" ? "http://localhost:5003" : "https://endlesschat.net/api"