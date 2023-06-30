const SingleChatlistItem = ({ chat, toggleChat }) => {
  const id = chat?.user_id > 0 ? [chat.user_id, 0] : [0, chat.group_id];

  const handleToggle = () => {
    toggleChat(id);
  };

  const defaultImage = () =>
    id[0] > 0 ? "defaultuser.jpg" : "defaultgroup.png";

  const imageHandler = () => {
    const source = chat?.avatarImage
      ? `images/${chat.id}/${chat.avatarImage}`
      : defaultImage();

    const image = (
      <img
        style={{
          width: "20px",
          height: "20px",
        }}
        src={source}
      ></img>
    );
    return image;
  };

  const listItem = (
    <p>
      {chat.name}
      {imageHandler()}
    </p>
  );

  return <div onClick={handleToggle}>{listItem}</div>;
};

export default SingleChatlistItem;
