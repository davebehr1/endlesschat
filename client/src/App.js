//-------------ALL OF THE CODE IN THIS FILE IS MY OWN------------------------
import React, { useEffect, useState } from "react";
import "./App.css";
import ChatApp from "./ChatApp";
import classes from "./app.module.css";
import Modal from "./modules/modal";
import { baseUrl,wsUrl } from "./const";

export let socket = null;
function App() {
  const [open, setOpen] = useState(true);
  const [tempUsername, setTempUsername] = useState();
  const [username, setUsername] = useState();
  const [error, setError] = useState(null);

  useEffect(() => {
    console.log(process.env.NODE_ENV);
    console.log(socket);
    if (username) {
      fetch(`${baseUrl}/v1/username/${username}`)
        .then(function (response) {
          return response.json();
        })
        .then((resp) => {
          if (resp.taken) {
            setError(resp.message);
          } else {
            console.log("setting socket", username);
            socket = new WebSocket(`${wsUrl}/v1/ws`);
            console.log(socket);
            setTimeout(() => setOpen(false), 200);
          }
        });
        // fetch('http://localhost:5002/signin',{
        //   method:"POST",
        //   body: JSON.stringify({
        //     username:username,
        //     password:"password1"
        //   }),
        //   headers:{
        //     'Content-type': 'application/json; charset=UTF-8',
        //   }
        // }).then((resp) => {
        //   return resp.json()
        // }).then((resp) => {
        //       console.log(resp)
        //       if (resp.taken) {
        //         setError(resp.message);
        //       } else {
        //         console.log("setting socket", username);
        //         socket = new WebSocket(`ws://localhost:5002/ws`);
        //         setTimeout(() => setOpen(false), 200);
        //       }
        //     }
        // )
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
              onClick={(value) => {
                if (value) {
                  setUsername(tempUsername);
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
      {console.log(socket)}
      {socket && <ChatApp socket={socket} />}
    </>
  );
}

export default App;
