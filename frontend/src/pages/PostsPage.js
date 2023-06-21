import React from "react";
import { useState, useEffect } from "react";
import FeedPosts from "../components/FeedPosts";
import CreatePost from "../components/CreatePost";
import { makeRequest } from "../services/makeRequest";
import GroupSidebar from "../components/GroupSidebar";

const Posts = () => {
  const [posts, setPosts] = useState([]);
  const [isPostsLoading, setPostsLoading] = useState(false);
  const [error, setError] = useState(null);
  let offset = 0;

  const handlePostUpdate = () => {
    setPostsLoading(!isPostsLoading);
  };

  useEffect(() => {
    const loadPosts = async () => {
      try {
        const response = await makeRequest(`feedposts/${offset}`);
        setPosts(response);
      } catch (error) {
        setError(error.message);
        // setLoading(false);
      }

      // try {
      //   const response = await axios.get(`http://localhost:8000/feedposts/${offset}`, {
      //     withCredentials: true,
      //   });
      //   setPosts(response.data);
      // } catch (error) {
      //   setError(error.message);
      // }
      // setLoading(false);
    };

    loadPosts();
  }, [isPostsLoading, offset]);

  return (
    <>
      <GroupSidebar />
      <CreatePost onPostsUpdate={handlePostUpdate} />
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
