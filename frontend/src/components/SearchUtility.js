import React, { useState } from "react";
import SearchBar from "../components/SearchBar";
import SearchResults from "../components/SearchResults";
import { Col, Stack, Row } from "react-bootstrap";

const SearchUtility = () => {
  const [searchResults, setSearchResults] = useState([]);

  return (
    <div>
      <SearchBar setSearchResults={setSearchResults} />
      {searchResults.length > 0 && (
        <SearchResults
          searchResults={searchResults}
          setSearchResults={setSearchResults}
        />
      )}
    </div>
  );
};

export default SearchUtility;
