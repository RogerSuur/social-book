import React from "react";
import { useState, useEffect } from "react";
import FeedPosts from "../components/FeedPosts";
import CreatePost from "../components/CreatePost";
import axios from "axios";


const Posts = () => {
  const [posts, setPosts] = useState([]);
  const [loading, setLoading] = useState(false);
  let offset = 0

  const loader = () => {
    setLoading((prevLoading) => !prevLoading);
  };

  useEffect(() => {
    const loadPosts = async () => {
      await axios
        .get(`http://localhost:8000/feedposts/${offset}`, {
          withCredentials: true,
        })
        .then((response) =>  {{
          console.log(response.data);
         setPosts(response.data)}}
        );
    };

    loadPosts();
  }, [loading, offset]);

  return (
    <>
      <CreatePost handler={loader} />
      <div className="content-area">
       <FeedPosts offset={offset}/>
      </div>
    </>
  );
};

export default Posts;
