const SingleChatItem = ({ chat, index, toggleChat, isOpen }) => {
  const handleToggle = () => {
    console.log(index);
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
      {isOpen && <div>Chat</div>}
    </>
  );
};

export default SingleChatItem;
