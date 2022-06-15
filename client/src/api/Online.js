//-------------ALL OF THE CODE IN THIS FILE IS MY OWN------------------------
import React from "react";
import { OnlineBubble } from "../modules/onlineBubble";
import classes from "./online.module.css";



const online = ["ann","david"]


function Online() {
  return (
    <div className={classes.online}>
      <h2 style={{ margin: "10px" }}>Online</h2>
      <div style={{ marginLeft: "10px" }}>
      {online.map((name, _) =>
          <OnlineBubble  user={name} />
      )}
      </div>
    </div>
  );
}

export { Online };
