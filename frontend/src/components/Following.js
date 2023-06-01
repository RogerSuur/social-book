import List from "../components/List.js";
import { Link } from "react-router-dom";

const FOLLOWING_URL = "http://localhost:8000/following";

const Following = () => {
  const mapFollowing = (following) => (
    <li key={following.id}>
      {following.firstName} {following.lastName}
    </li>
  );

  return <List url={FOLLOWING_URL} mapFunction={mapFollowing} />;
};

export default Following;
