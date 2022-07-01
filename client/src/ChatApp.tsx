import React, { useEffect, useState } from "react";
import "./App.css";
import { connect, sendMsg } from "./api";
import { Header } from "./api/Header";
import { Online } from "./api/Online";
import { ChatHistory } from "./api/ChatHistory";
import classes from "./app.module.css";
import { Client } from '@stomp/stompjs';

type props = {
  socket: Client | null;
}

export interface Message {
  name: string,

}

export interface ExtractedMessage {
  body: string,
  User: string
}

function ChatApp({ socket }: props) {
  const [chatHistory, setChatHistory] = useState<ExtractedMessage[]>([]);
  const [message, setMessage] = useState("");

  useEffect(() => {
    if (socket) {
      connect((msg: string) => {
        console.log("New Message:", msg);

        setChatHistory([...chatHistory, { body: msg, User: "David" }]);
        console.log(chatHistory);
      });
    }
  });

  function send() {
    sendMsg(message);
    console.log(message);
    console.log(chatHistory);
  }
  return (
    <div className="App">
      <Header />
      <div className={classes.bigWrapper}>
        <Online />
        <div className={classes.innerWrapper} >
          <ChatHistory chatHistory={chatHistory} />
          <div className={classes.messageContainer}>
            <input
              type="text"
              placeholder="new message"
              className={classes.textInput}
              onChange={(val) => setMessage(val.currentTarget.value)}
            />
            <button className={classes.sendButton} onClick={() => send()}>
              Send
            </button>
          </div>
        </div>
      </div>
    </div>
  );
}

export default ChatApp;
