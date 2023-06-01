import List from "../components/List.js";
import { Link } from "react-router-dom";

const FOLLOWERS_URL = "http://localhost:8000/followers";

const Followers = () => {
  const mapFollowers = (follower) => (
    <li key={follower.id}>
      <Link to={`/profile/${follower.id}`}>
        {follower.firstName} {follower.lastName}
      </Link>
    </li>
  );

  return <List url={FOLLOWERS_URL} mapFunction={mapFollowers} />;
};

export default Followers;
