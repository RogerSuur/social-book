import List from "../components/List.js";
import { Link } from "react-router-dom";

const GenericUserList = ({ url }) => {
  const mapUsers = (user, index) => (
    <li key={index}>
      <Link to={`/profile/${user.id}`}>
        {user?.username
          ? `${user.username}`
          : `${user.firstName} ${user.lastName}`}
      </Link>
    </li>
  );

  return <List url={url} mapFunction={mapUsers} />;
};

export default GenericUserList;
