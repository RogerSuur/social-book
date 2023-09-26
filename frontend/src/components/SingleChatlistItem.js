import ImageHandler from "../utils/imageHandler";

const SingleChatlistItem = ({ chat }) => {
  const image =
    chat?.user_id > 0
      ? ImageHandler(chat?.avatar_image, "defaultuser.jpg", "chatbox-img")
      : ImageHandler(chat?.avatar_image, "defaultgroup.png", "chatbox-img");

  const styles = chat?.user_id > 0 ? "" : "group-chatlist-item";

  const listItem = (
    <span className={`me-1 ${styles}`}>
      {image} {chat.name}
    </span>
  );

  return listItem;
};

export default SingleChatlistItem;
