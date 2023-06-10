const SingleChatlistItem = ({ chat, toggleChat }) => {
  const id = chat?.userid > 0 ? [chat.userid, 0] : [0, chat.group_id];

  const handleToggle = () => {
    toggleChat(id);
  };

  const listItem = (
    <p>
      {chat.name}
      {/* <img
        style={{
          width: "10px",
          height: "10px",
          objectFit: "cover",
          objectPosition: "0% 100%",
        }}
        src={`images/${id}/${chat.avatarImage}`}
      ></img> */}
    </p>
  );

  return <div onClick={handleToggle}>{listItem}</div>;
};

export default SingleChatlistItem;
