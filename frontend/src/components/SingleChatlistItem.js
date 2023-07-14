import ImageHandler from "../utils/imageHandler";

const SingleChatlistItem = ({ chat, toggleChat }) => {
  const id = chat?.user_id > 0 ? [chat.user_id, 0] : [0, chat.group_id];

  console.log(chat);

  const handleToggle = () => {
    toggleChat(id);
  };

  const image = () =>
    id[0] > 0
      ? ImageHandler(chat?.avatar_image, "defaultuser.jpg", "chatbox-img")
      : ImageHandler("", "defaultgroup.png", "chatbox-img");

  const listItem = (
    <p>
      {image()}
      {chat.name}
    </p>
  );

  return <div onClick={handleToggle}>{listItem}</div>;
};

export default SingleChatlistItem;
