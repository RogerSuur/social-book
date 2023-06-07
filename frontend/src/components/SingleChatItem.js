import Chatbox from "../components/Chatbox";

const SingleChatItem = ({ chat, index, toggleChat, isOpen }) => {
  const handleToggle = () => {
    toggleChat(index);
  };

  console.log(isOpen, index, "OPEN");

  return (
    <>
      <div onClick={handleToggle}>
        <p>
          {chat.first_name} {chat.last_name}
        </p>
      </div>
      {isOpen && <Chatbox toggleChat={toggleChat} chat={chat} />}
    </>
  );
};

export default SingleChatItem;
