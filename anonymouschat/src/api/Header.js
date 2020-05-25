import React from "react";
import classes from "./header.module.css";

function Header() {
  return (
    <div className={classes.header}>
      <h2 style={{ marginLeft: "10px" }}>Anonymous Chat Group</h2>
    </div>
  );
}

export { Header };
