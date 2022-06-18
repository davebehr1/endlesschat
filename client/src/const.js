export const baseUrl = process.env.NODE_ENV === 'development' ? 'http://localhost:5002' : `http://${process.env.CHAT_SERVER_HOST}:${process.env.CHAT_SERVER_PORT}`


export const wsUrl = process.env.NODE_ENV === 'development' ? 'ws://localhost:5002' : `ws://${process.env.CHAT_SERVER_HOST}:${process.env.CHAT_SERVER_PORT}`