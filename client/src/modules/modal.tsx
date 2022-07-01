import styles from "./modal.module.css";
import React from "react";

type props = {
  children: React.ReactNode
}

const Modal = ({ children }: props) => {
  const ref = React.createRef<HTMLDivElement>();
  return (
    <div className={styles.modal} ref={ref}>
      {children}
    </div>
  );
};

export default Modal;
