import React, { useState } from "react";
import SearchBar from "../components/SearchBar";
import SearchResults from "../components/SearchResults";

const SearchUtility = () => {
  const [searchResults, setSearchResults] = useState([]);

  return (
    <div>
      <SearchBar setSearchResults={setSearchResults} />
      <SearchResults
        searchResults={searchResults}
        setSearchResults={setSearchResults}
      />
    </div>
  );
};

export default SearchUtility;
