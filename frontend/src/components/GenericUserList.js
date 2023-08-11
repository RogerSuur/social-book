import List from "../components/List.js";
import ImageHandler from "../utils/imageHandler.js";
import { ListGroup } from "react-bootstrap";
import { LinkContainer } from "react-router-bootstrap";

const image = (user) =>
  ImageHandler(user?.avatarImage, "defaultuser.jpg", "userlist-img");

const GenericUserList = ({ url }) => {
  const mapUsers = (user, index) => {
    return (
      <ListGroup.Item action key={index}>
        <LinkContainer to={`/profile/${user.id}`}>
          <div>
            {image(user)}
            {user?.nickname
              ? `${user.nickname}`
              : `${user.firstName} ${user.lastName}`}
          </div>
        </LinkContainer>
      </ListGroup.Item>
    );
  };

  return <List url={url} mapFunction={mapUsers} />;
};

export default GenericUserList;
