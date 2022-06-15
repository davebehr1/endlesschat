//-------------ALL OF THE CODE IN THIS FILE IS MY OWN------------------------
import styles from "./modal.module.css";
import React from "react";

const Modal = ({ children }) => {
  const ref = React.createRef();
  return (
    <div className={styles.modal} ref={ref}>
      {children}
    </div>
  );
};

export default Modal;
