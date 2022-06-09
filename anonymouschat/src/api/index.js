//-------------THE CODE IN THIS FILE IS STANDARD FOR LISTENING FOR WEB SOCKET EVENTS---------------
//-------------THE SEND MESSAGE FUNCTION ON LINE 25 is My own-----------------
import { socket } from "../App";
let connect = (cb) => {
  console.log("Attempting Connection...");

  socket.onopen = () => {
    console.log("Successfully Connected");
  };

  socket.onmessage = (msg) => {
    console.log(msg);
    cb(msg);
  };

  socket.onclose = (event) => {
    console.log("Socket Closed Connection: ", event);
  };

  socket.onerror = (error) => {
    console.log("Socket Error: ", error);
  };
};

let sendMsg = (msg) => {
  var message = {"message":msg,"to":"ann"}
  console.log("sending msg: ", message);
  socket.send(JSON.stringify(message));
};

export { connect, sendMsg };
