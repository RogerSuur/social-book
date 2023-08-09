import ImageHandler from "../utils/imageHandler";
import Container from "react-bootstrap/Container";

const SingleChatlistItem = ({ chat }) => {
  const image =
    chat?.user_id > 0
      ? ImageHandler(chat?.avatar_image, "defaultuser.jpg", "chatbox-img")
      : ImageHandler("", "defaultgroup.png", "chatbox-img");

  const listItem = (
    <>
      {image} {chat.name}
    </>
  );

  return listItem;
};

export default SingleChatlistItem;
