import { useEffect, useState } from "react";
import "./App.css";
import ChatApp from "./ChatApp";
import classes from "./app.module.css";
import Modal from "./modules/modal";
import { httpbaseUrl, wsbaseUrl } from "./consts"
import { Client } from '@stomp/stompjs';



const SOCKET_URL = `${wsbaseUrl}/chat-websocket`;


export let client: Client | null = null;
function App() {
  const [open, setOpen] = useState(true);
  const [tempUsername, setTempUsername] = useState<String>();
  const [username, setUsername] = useState<String>();
  const [error, setError] = useState<String | null>(null);

  useEffect(() => {
    if (username) {
      fetch(`${httpbaseUrl}/username/${username}`)
        .then(function (response) {
          return response.json();
        })
        .then((resp) => {
          if (resp.taken) {
            setError(resp.message);
          } else {
            client = new Client({
              brokerURL: SOCKET_URL,
              reconnectDelay: 5000,
              heartbeatIncoming: 4000,
              heartbeatOutgoing: 4000
            });


            setTimeout(() => setOpen(false), 200);
          }
        });

      fetch(`/v2/`)
        .then(function (response) {
          console.log(response.body);
        });
      // fetch('http://localhost:5001/signin',{
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
      //         socket = new WebSocket(`ws://localhost:5001/ws`);
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
      {client && <ChatApp socket={client} />}
    </>
  );
}

export default App;
