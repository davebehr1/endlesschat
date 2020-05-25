import React from "react";
import classes from "./chatBubble.module.css";
const ChatBubble = ({ message, user }) => {
  return (
    <div className={classes.chatBubble}>
      <div className={classes.userName}> {user}</div>
      <div className={classes.message}>{message}</div>
    </div>
  );
};
export { ChatBubble };
