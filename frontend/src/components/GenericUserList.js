import List from "../components/List.js";
import { Link } from "react-router-dom";
import ImageHandler from "../utils/imageHandler.js";

const image = (user) =>
  ImageHandler(user?.avatarImage, "defaultuser.jpg", "profile-image");

const GenericUserList = ({ url }) => {
  const mapUsers = (user, index) => {
    return (
      <li className="follower pepe" key={index}>
        {image(user)}
        <Link to={`/profile/${user.id}`}>
          {user?.username
            ? `${user.username}`
            : `${user.firstName} ${user.lastName}`}
        </Link>
      </li>
    );
  };

  return <List url={url} mapFunction={mapUsers} />;
};

export default GenericUserList;
