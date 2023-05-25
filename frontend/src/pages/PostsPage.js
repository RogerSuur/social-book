import React from "react";
import { useState, useEffect } from "react";
import FeedPosts from "../components/FeedPosts";
import CreatePost from "../components/CreatePost";
import { makeRequest } from "../services/makeRequest";

const Posts = () => {
  const [posts, setPosts] = useState([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);
  let offset = 0;

  const loader = () => {
    setLoading(true);
  };

  useEffect(() => {
    const loadPosts = async () => {
      try {
        const response = await makeRequest(`feedposts/${offset}`);
        setPosts(response);
        setLoading(false);
      } catch (error) {
        setError(error.message);
        setLoading(false);
      }

      // try {
      //   const response = await axios.get(`http://localhost:8000/feedposts/${offset}`, {
      //     withCredentials: true,
      //   });
      //   setPosts(response.data);
      // } catch (error) {
      //   setError(error.message);
      // }
    };

    loadPosts();
  }, [loading, offset]);

  return (
    <>
      <CreatePost handler={loader} />
      {error ? (
        <div className="error">{error}</div>
      ) : (
        <div className="content-area">
          <FeedPosts posts={posts} />
        </div>
      )}
    </>
  );
};

export default Posts;
