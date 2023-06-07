import { WS_URL } from "../utils/routes";

const Chatbox = ({ toggleChat, chat }) => {
  const closeChat = () => {
    toggleChat(0);
  };

  const chatbox = (
    <div className="chatbox">
      <div className="chat-title">
        {chat.first_name} {chat.last_name}
        <button onClick={closeChat}>Close</button>
      </div>
      <div className="message-box">Messages</div>
    </div>
  );

  return <>{chatbox}</>;
};

export default Chatbox;
