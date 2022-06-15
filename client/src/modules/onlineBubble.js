//-------------ALL OF THE CODE IN THIS FILE IS MY OWN------------------------
import React from "react";
import classes from "./onlineBubble.module.css";
const OnlineBubble = ({ user }) => {
  return (
    <div className={classes.onlineBubble}>
      <div className={classes.userName}> {user}</div>
      <div className={classes.online}/>
    </div>
  );
};
export { OnlineBubble };
