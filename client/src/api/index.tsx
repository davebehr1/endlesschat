import { client } from "../App";
// import { Message } from "../ChatApp"
let connect = (cb: (msg: string) => void) => {
  client!.onConnect = () => {
    console.log("Connected!!")
    client!.subscribe('/topic/greetings', function (msg) {
      console.log(msg);
      if (msg.body) {
        var jsonBody = JSON.parse(msg.body);
        console.log(jsonBody);
        if (jsonBody.name) {
          cb(jsonBody.name);
        }
      }
    });
  };

  client!.onDisconnect = (event) => {
    console.log("Socket Closed Connection: ", event);
  };

  client!.onStompError = (error) => {
    console.log("Socket Error: ", error);
  };

  client!.activate();
};

let sendMsg = (msg: string) => {
  var message = { "name": msg }
  console.log("sending msg: ", message);
  client!.publish({ destination: "/app/chat", body: JSON.stringify(message) });

};

export { connect, sendMsg };
