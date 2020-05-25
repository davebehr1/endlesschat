import React, { useEffect, useState } from "react";
import "./App.css";
import { connect, sendMsg } from "./api";
import { Header } from "./api/Header";
import { ChatHistory } from "./api/ChatHistory";
import ChatApp from "./ChatApp";
import classes from "./app.module.css";
import Modal from "./modules/modal";

export let socket;
function App() {
  const [chatHistory, setChatHistory] = useState([]);
  const [message, setMessage] = useState("");
  const [open, setOpen] = useState(true);
  const [username, setUsername] = useState("");

  useEffect(() => {
    console.log("hello");
    if (username) {
      console.log("setting socket", username);
      socket = new WebSocket(`ws://localhost:5000/chat/${username}`);
    }
    // if (socket) {
    //   connect((msg) => {
    //     console.log("New Message:", msg.data);
    //     console.log(typeof msg.data);
    //     var mes = JSON.parse(msg.data);
    //     if (typeof mes === "string") {
    //       mes = JSON.parse(mes);
    //     }
    //     setChatHistory([...chatHistory, mes]);
    //     console.log(chatHistory);
    //   });
    // }
  }, [username]);

  // function send() {
  //   sendMsg(message);
  //   console.log(message);
  //   console.log(chatHistory);
  // }
  return (
    <>
      {open && (
        <Modal>
          <div className={classes.modelContent}>
            <h2>Enter your alias</h2>
            <input
              type="text"
              onBlur={(username) => setUsername(username.currentTarget.value)}
            />
            <button
              onClick={() => {
                setTimeout(() => setOpen(false), 200);
              }}
            >
              send
            </button>
          </div>
        </Modal>
      )}
      {socket && <ChatApp socket={socket} />}
    </>
  );
}

export default App;
