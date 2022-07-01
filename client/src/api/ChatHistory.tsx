//-------------ALL OF THE CODE IN THIS FILE IS MY OWN------------------------
import React from "react";
import { ChatBubble } from "../modules/chatBubble";
import { InfoBubble } from "../modules/infoBubble";
import classes from "./chatHistory.module.css";
import { ExtractedMessage } from "../ChatApp"




type Props = {
  chatHistory: ExtractedMessage[]
}

const ChatHistory = ({ chatHistory }: Props) => {
  return (
    <div className={classes.historyWrapper}>
      {chatHistory.map((msg, _) =>
        "User" in msg && msg["User"] !== "" ? (
          <ChatBubble message={msg.body} user={msg.User} />
        ) : (
          (console.log("no user"), (<InfoBubble message={msg.body} />))
        )
      )}
    </div>
  );
};
export { ChatHistory };
