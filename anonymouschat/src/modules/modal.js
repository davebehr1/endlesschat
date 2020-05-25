import styles from "./modal.module.css";
import React, { useEffect } from "react";

const Modal = ({ children }) => {
  const ref = React.createRef();
  //   useEffect(() => {
  //     if (ref.current) disableBodyScroll(ref.current);
  //   });
  //   useEffect(() => {
  //     return () => {
  //       clearAllBodyScrollLocks();
  //     };
  //   }, []);
  return (
    <div className={styles.modal} ref={ref}>
      {children}
    </div>
  );
};

export default Modal;
