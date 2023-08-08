import ImageHandler from "../utils/imageHandler";

const SingleChatlistItem = ({ chat, toggleChat }) => {
  const handleToggle = () => {
    toggleChat(chat);
  };

  const image = () =>
    chat?.user_id > 0
      ? ImageHandler(chat?.avatar_image, "defaultuser.jpg", "chatbox-img")
      : ImageHandler("", "defaultgroup.png", "chatbox-img");

  const listItem = (
    <p>
      {image()} {chat.name}{" "}
    </p>
  );

  return <div onClick={handleToggle}>{listItem}</div>;
};

export default SingleChatlistItem;
