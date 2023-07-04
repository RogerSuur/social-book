import { useEffect, useState, setError } from "react";
import { makeRequest } from "../services/makeRequest.js";

const SearchBar = ({ setSearchResults }) => {
  //   const [searchResults, setSearchResults] = useState([]);
  const [searchString, setSearchString] = useState("");

  //   useEffect(() => {
  //     const abortController = new AbortController();
  //     const loadPosts = async () => {
  //       try {
  //         console.log(searchString);
  //         // const response = await makeRequest(`search/${searchString}`, {
  //         //   signal: abortController.signal,
  //         // });
  //         fetch("https://jsonplaceholder.typicode.com/users")
  //           .then((response) => response.json())
  //           .then((response) => {
  //             console.log(response);
  //             setSearchResults(response);
  //           });
  //       } catch (error) {
  //         setError(error.message);
  //       }
  //     };
  //     loadPosts();

  //     console.log(searchResults.userName);

  //     return () => {
  //       abortController.abort();
  //     };
  //   }, [searchString]);

  const fetchData = () => {
    try {
      fetch("https://jsonplaceholder.typicode.com/users")
        .then((response) => response.json())
        .then((response) => {
          console.log(response);
          setSearchResults(response);
        });
    } catch (error) {
      setError(error.message);
    }
  };

  const handleChange = (e) => {
    if (!e.target.value) {
      setSearchResults([]);
    } else {
      setSearchString(e.target.value);
      fetchData(searchString);
    }
  };

  const handleSubmit = (e) => {
    e.preventDefault();
  };

  return (
    <>
      <form id="form" onSubmit={handleSubmit} onChange={handleChange}>
        <input
          className="search"
          type="text"
          id="search"
          placeholder="Search.."
          onChange={handleChange}
        />
        {/* <button className="search-button">Search</button> */}
      </form>
      {/* {searchResults.map((result) => (
        <div key={result.id}>{result.name}</div>
      ))} */}
    </>
  );
};

export default SearchBar;
