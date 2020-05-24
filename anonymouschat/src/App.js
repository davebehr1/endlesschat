import React, { useEffect, useState } from "react";
import logo from "./logo.svg";
import "./App.css";

import { connect, sendMsg } from "./api";
import { Header } from "./api/Header";
import { ChatHistory } from "./api/ChatHistory";

function App() {
  const [chatHistory, setChatHistory] = useState([]);
  useEffect(() => {
    connect((msg) => {
      console.log("New Message:", msg);

      setChatHistory([...chatHistory, msg]);
      console.log(chatHistory);
    });
  });
  function send() {
    sendMsg("hello");
    console.log(chatHistory);
  }
  return (
    <div className="App">
      <Header />
      <ChatHistory chatHistory={chatHistory} />
      <button onClick={() => send()}>Hit</button>
    </div>
  );
}

export default App;
