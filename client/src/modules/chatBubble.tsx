//-------------ALL OF THE CODE IN THIS FILE IS MY OWN------------------------
import React from "react";
import classes from "./chatBubble.module.css";


type props = {
  message: string,
  user: string,
}
const ChatBubble = ({ message, user }: props) => {
  return (
    <div className={classes.chatBubble}>
      <div className={classes.userName}> {user}</div>
      <div className={classes.message}>{message}</div>
    </div>
  );
};
export { ChatBubble };
