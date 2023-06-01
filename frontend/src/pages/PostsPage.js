import React from "react";
import { useState, useEffect } from "react";
import FeedPosts from "../components/FeedPosts";
import CreatePost from "../components/CreatePost";
import { makeRequest } from "../services/makeRequest";

const Posts = () => {
  const [posts, setPosts] = useState([]);
  const [error, setError] = useState(null);
  const [offset, setOffset] = useState(0);
  const [hasMore, setHasMore] = useState(false);

  const handlePostUpdate = () => {
    setOffset(0);
    setPosts([]);
  };

  const handlePageChange = () => {
    setOffset((prevOffset) => prevOffset + 1);
  };

  useEffect(() => {
    const abortController = new AbortController();
    const loadPosts = async () => {
      try {
        const response = await makeRequest(`feedposts/${offset}`, {
          signal: abortController.signal,
        });
        setPosts((prevPosts) => {
          return [...prevPosts, ...response];
        });
        setHasMore(response.length > 0);
      } catch (error) {
        setError(error.message);
      }

      // try {
      //   const response = await axios.get(
      //     `http://localhost:8000/feedposts/${offset}`,
      //     {
      //       withCredentials: true,
      //       signal: abortController.signal,
      //     }
      //   );
      //   setPosts((prevPosts) => {
      //     const newPosts = response.data.filter(
      //       (post) => !prevPosts.some((prevPost) => prevPost.id === post.id)
      //     );
      //     return [...prevPosts, ...newPosts];
      //   });
      // } catch (error) {
      //   setError(error.message);
      // }
      // setLoading(false);
    };
    loadPosts();

    return () => {
      abortController.abort();
    };
  }, [offset]);

  return (
    <>
      <CreatePost onPostsUpdate={handlePostUpdate} />
      {error ? (
        <div className="error">{error}</div>
      ) : (
        <div className="content-area">
          <FeedPosts
            posts={posts}
            hasMore={hasMore}
            onLoadMore={handlePageChange}
          />
        </div>
      )}
    </>
  );
};

export default Posts;
