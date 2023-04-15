import React, { useState, useEffect } from "react";
import { useParams, Link } from "react-router-dom";
import axios from "axios";

const Category = () => {
  const [category, setCategory] = useState({
    posts: [],
  });
  const { id } = useParams();

  useEffect(() => {
    const loadCategory = async () => {
      await axios
        .get(`http://localhost:8000/categories/${id}`, {
          withCredentials: true,
        })
        .then((response) => setCategory(response.data));
    };

    loadCategory();
  }, []);

  return (
    <>
      <div>
        {category.posts.map((post) => (
          <div className="content-area" key={post.post_id}>
            <Link to={`/posts/${post.post_id}`}>
              {post.title}
              <br />
            </Link>
            <div className="row">{post.body}</div>
            <div className="row">
              {new Date(post.post_datetime).toLocaleString("et-EE")}
            </div>
          </div>
        ))}
      </div>
    </>
  );
};

export default Category;
