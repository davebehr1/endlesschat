import { socket } from "../App";
import { Message } from "../ChatApp"
let connect = (cb: (msg: Message) => void) => {
  console.log("Attempting Connection...");

  socket!.onopen = () => {
    console.log("Successfully Connected");
  };

  socket!.onmessage = (msg) => {
    console.log(msg);
    cb(msg);
  };

  socket!.onclose = (event) => {
    console.log("Socket Closed Connection: ", event);
  };

  socket!.onerror = (error) => {
    console.log("Socket Error: ", error);
  };
};

let sendMsg = (msg: string) => {
  var message = { "message": msg, "to": "ann" }
  console.log("sending msg: ", message);
  socket!.send(JSON.stringify(message));
};

export { connect, sendMsg };
