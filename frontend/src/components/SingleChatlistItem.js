const SingleChatlistItem = ({ chat, userid, toggleChat, isOpen }) => {
  const handleToggle = () => {
    toggleChat(userid);
  };

  console.log(isOpen, userid, "OPEN");

  return (
    <div onClick={handleToggle}>
      <p>
        {chat.first_name} {chat.last_name}
      </p>
    </div>
  );
};

export default SingleChatlistItem;
