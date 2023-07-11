import ImageHandler from "../utils/imageHandler";

const SingleChatlistItem = ({ chat, toggleChat }) => {
  const id = chat?.user_id > 0 ? [chat.user_id, 0] : [0, chat.group_id];

  const handleToggle = () => {
    toggleChat(id);
  };

  const defaultImage = () =>
    id[0] > 0 ? "defaultuser.jpg" : "defaultgroup.png";

  const image = () =>
    ImageHandler(chat?.avatarImage, defaultImage(), "chatbox-img");

  const listItem = (
    <p>
      {image()}
      {chat.name}
    </p>
  );

  return <div onClick={handleToggle}>{listItem}</div>;
};

export default SingleChatlistItem;
