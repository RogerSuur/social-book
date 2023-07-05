import { useEffect, useState } from "react";
import { makeRequest } from "../services/makeRequest.js";

const SearchBar = ({ setSearchResults }) => {
  const [searchString, setSearchString] = useState("");
  const [error, setError] = useState(null);

  const fetchData = async (value) => {
    try {
      const response = await makeRequest(`search/${value}`, {});
      setSearchResults(response);
    } catch (error) {
      setError(error.message);
    }
  };

  const handleChange = (e) => {
    const value = e.target.value;
    if (!value) {
      setSearchResults([]);
    } else {
      fetchData(value);
      setSearchString(value);
    }
  };

  const handleSubmit = (e) => {
    e.preventDefault();
  };

  return (
    <>
      {error && <div>Error: {error}</div>}
      <form id="form" onSubmit={handleSubmit} onChange={handleChange}>
        <input
          className="search-it" type="text" placeholder="Search here"
          onChange={handleChange}
        />
      </form>
    </>
  );
};

export default SearchBar;
