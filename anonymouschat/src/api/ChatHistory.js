import React, { useEffect } from "react";
import { ChatBubble } from "../modules/chatBubble";
import { InfoBubble } from "../modules/infoBubble";
import classes from "./chatHistory.module.css";
const ChatHistory = ({ chatHistory }) => {
  return (
    <div className={classes.historyWrapper}>
      {chatHistory.map((msg, index) =>
        "User" in msg && msg["User"] != "" ? (
          <ChatBubble message={msg.body} user={msg.User} />
        ) : (
          (console.log("no user"), (<InfoBubble message={msg.body} />))
        )
      )}
    </div>
  );
};
export { ChatHistory };
