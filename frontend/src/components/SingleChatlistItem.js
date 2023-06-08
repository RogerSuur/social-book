const SingleChatlistItem = ({ chat, userid, toggleChat }) => {
  const handleToggle = () => {
    toggleChat(userid);
  };

  let listItem;

  if (chat?.username) {
    listItem = <p>{chat.username}</p>;
  } else {
    listItem = (
      <p>
        {chat.first_name} {chat.last_name}
      </p>
    );
  }

  return <div onClick={handleToggle}>{listItem}</div>;
};

export default SingleChatlistItem;
