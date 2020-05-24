import React, { useEffect } from "react";

const ChatHistory = ({ chatHistory }) => {
  return (
    <div className="ChatHistory">
      <h2>Chat History</h2>
      {chatHistory.map((msg, index) => (
        <p key={index}>{msg.data}</p>
      ))}
    </div>
  );
};
export { ChatHistory };
