import React from "react";
import { useState, useEffect } from "react";
import Select from "react-select";
import axios from "axios";

const CreatePost = (props) => {
  const [formData, setFormData] = useState({
    title: "",
    body: "",
    categories: [],
  });
  const [categories, setCategories] = useState([]);
  const [errMsg, setErrMsg] = useState("");

  useEffect(() => {
    const loadCategories = async () => {
      await axios
        .get("http://localhost:8000/categories", {
          withCredentials: true,
        })
        .then((response) =>
          setCategories(
            response.data.map((category) => {
              return { label: category.title, value: category.category_id };
            })
          )
        );
    };

    loadCategories();
  }, []);

  const handleChange = (event) => {
    const { name, value, type, checked } = event.target;

    setFormData((prevFormData) => {
      return {
        ...prevFormData,
        [name]: type === "checkbox" ? checked : value,
      };
    });
  };

  const handleOptions = (selected) => {
    setFormData((prevFormData) => {
      return {
        ...prevFormData,
        categories: selected,
      };
    });
  };

  const handleSubmit = async (event) => {
    event.preventDefault();

    try {
      const response = await axios.post(
        "http://localhost:8000/createpost",
        JSON.stringify(formData),
        { withCredentials: true },
        {
          headers: { "Content-Type": "application/json" },
        }
      );

      setErrMsg(response.data?.message);
      props.handler();
    } catch (err) {
      if (!err?.response) {
        setErrMsg("No Server Response");
      } else {
        setErrMsg("Internal Server Error");
      }
    }

    setFormData({
      title: "",
      body: "",
      categories: [],
    });
  };

  return (
    <>
      <div className="content-area">
        {errMsg && <h2>{errMsg}</h2>}
        <form onSubmit={handleSubmit}>
          <input
            type="text"
            placeholder="Title"
            onChange={handleChange}
            name="title"
            value={formData.title}
            required
          />
          <textarea
            placeholder="Write something..."
            onChange={handleChange}
            value={formData.body}
            name="body"
            required
          />
          <Select
            options={categories}
            isMulti
            name="categories"
            onChange={handleOptions}
            value={formData.categories}
            required
          />
          <button>Post</button>
        </form>
      </div>
    </>
  );
};

export default CreatePost;
