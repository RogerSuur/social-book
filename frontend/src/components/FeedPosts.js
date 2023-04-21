import List from "./List";


const FEEDPOSTS_URL = "http://localhost:8000/feedposts";

const FeedPosts = () => {
    const mapFeedPosts = (post) => (
        <div className="content-area" key={post.id}>
          <div className="row">{post.content}</div>
          <div className="row">
            {new Date(post.createdAt).toLocaleString("et-EE")}
          </div>
        </div>
        )


  return <List url={FEEDPOSTS_URL} mapFunction={mapFeedPosts} />;
}


export default FeedPosts;
