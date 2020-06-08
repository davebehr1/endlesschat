//-------------ALL OF THE CODE IN THIS FILE IS MY OWN------------------------
import React, { useEffect, useState } from "react";
import "./App.css";
import { connect, sendMsg } from "./api";
import { Header } from "./api/Header";
import { ChatHistory } from "./api/ChatHistory";
import classes from "./app.module.css";

export let socket;
function ChatApp({ socket }) {
  const [chatHistory, setChatHistory] = useState([]);
  const [message, setMessage] = useState("");

  useEffect(() => {
    if (socket) {
      connect((msg) => {
        console.log("New Message:", msg.data);
        console.log(typeof msg.data);
        var mes = JSON.parse(msg.data);
        if (typeof mes === "string") {
          mes = JSON.parse(mes);
        }
        setChatHistory([...chatHistory, mes]);
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
  );
}

export default ChatApp;
