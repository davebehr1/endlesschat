//-------------ALL OF THE CODE IN THIS FILE IS MY OWN------------------------
import React from "react";
import classes from "./chatBubble.module.css";


type Props = {
  message: string,
}
const InfoBubble = ({ message }: Props) => {
  console.log("hey");
  return (
    <div className={classes.infoBubble}>
      <div className={classes.message}>{message}</div>
    </div>
  );
};
export { InfoBubble };
