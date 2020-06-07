//-------------ALL OF THE CODE IN THIS FILE IS MY OWN------------------------
import React, { useEffect, useState } from "react";
import "./App.css";
import ChatApp from "./ChatApp";
import classes from "./app.module.css";
import Modal from "./modules/modal";

export let socket;
function App() {
  const [open, setOpen] = useState(true);
  const [tempUsername, setTempUsername] = useState();
  const [username, setUsername] = useState();
  const [error, setError] = useState();

  useEffect(() => {
    console.log("hello");
    if (username) {
      console.log("setting socket", username);
      socket = new WebSocket(`wss://localhost:5000/chat/${username}`);
    }
  }, [username]);

  return (
    <>
      {open && (
        <Modal>
          <div className={classes.modelContent}>
            <h2>ENTER A CHAT NAME</h2>
            <input
              type="text"
              placeholder={"mr.robot92"}
              className={classes.aliasInput}
              onBlur={(username) =>
                setTempUsername(username.currentTarget.value)
              }
            />
            {error && <div className={classes.error}>{error}</div>}
            <button
              className={classes.sendButton}
              onClick={() => {
                if (tempUsername) {
                  setUsername(tempUsername);
                  setTimeout(() => setOpen(false), 200);
                } else {
                  setError("a chat name is required *");
                }
              }}
            >
              start chatting
            </button>
          </div>
        </Modal>
      )}
      {socket && <ChatApp socket={socket} />}
    </>
  );
}

export default App;
