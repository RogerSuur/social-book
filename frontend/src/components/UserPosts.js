import List from "../components/List.js";

//insert correct URL
const USER_POSTS_URL = "http://localhost:8000/profile";

const UserPosts = () => {
  const mapUserPosts = (post) => <li key={post.id}>{post.title}</li>;

  return <List url={USER_POSTS_URL} mapFunction={mapUserPosts} />;
};

export default UserPosts;
