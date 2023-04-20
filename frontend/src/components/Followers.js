import List from "../components/List.js";

//insert correct URL
const FOLLOWERS_URL = "http://localhost:8000/profile";

const Followers = () => {
  const mapFollowers = (follower) => (
    <li key={follower.id}>
      {follower.firstName} {follower.lastName}
    </li>
  );

  return <List url={FOLLOWERS_URL} mapFunction={mapFollowers} />;
};

export default Followers;
