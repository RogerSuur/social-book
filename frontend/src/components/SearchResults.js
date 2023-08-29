import React from "react";
import { LinkContainer } from "react-router-bootstrap";
import { Container, ListGroup, ListGroupItem } from "react-bootstrap";

const SearchResults = ({ searchResults, setSearchResults }) => {
  const searchResultsMap = searchResults.map((result, index) => (
    <LinkContainer
      to={
        result.userId === 0
          ? `/groups/${result.groupId}`
          : `/profile/${result.userId}`
      }
    >
      <ListGroupItem
        action
        onClick={() => {
          setSearchResults([]);
        }}
        key={index}
      >
        <>{result.name}</>
      </ListGroupItem>
    </LinkContainer>
  ));

  const hasResults = searchResultsMap.length > 0;

  return <ListGroup className="position-fixed">{searchResultsMap}</ListGroup>;
};

export default SearchResults;
