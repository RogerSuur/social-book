import List from "./List";


const FEEDPOSTS_URL = `http://localhost:8000/feedposts/`;

const FeedPosts = (props) => {
  const { offset } = props;
  console.log(offset)
    const mapFeedPosts = (post) => {

      const numComments = 7

      return (
        <div className="content-area" key={post.id}>
          <div className="row">{post.userId}</div>
          <div className="row">{post.content}</div>
          <div className="row">
            {new Date(post.createdAt).toLocaleString("et-EE")}
          </div>
          {/* ADD A LINK TO comments if not empty*/}
          {/* {post.comments !== 0 && <div className="row"> 7 comments</div>} */}
          {post.comments !== 0 && (
        <div className="row">
          <a href={`/comments/${post.id}`}>{numComments} comments</a>
        </div>
      )}
        </div>
      )  
  }


  return <List url={`${FEEDPOSTS_URL}${offset}`} mapFunction={mapFeedPosts} />;
}


export default FeedPosts;
