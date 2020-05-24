import React, { useEffect, useState } from "react";
import "./App.css";
import { connect, sendMsg } from "./api";
import { Header } from "./api/Header";
import { ChatHistory } from "./api/ChatHistory";

function App() {
  const [chatHistory, setChatHistory] = useState([]);
  const [message, setMessage] = useState("");
  useEffect(() => {
    connect((msg) => {
      console.log("New Message:", msg);

      setChatHistory([...chatHistory, msg]);
      console.log(chatHistory);
    });
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

      <input
        type="text"
        onChange={(val) => setMessage(val.currentTarget.value)}
      />
      <button onClick={() => send()}>Hit</button>
    </div>
  );
}

export default App;
