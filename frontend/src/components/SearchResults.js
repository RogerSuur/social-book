import React from "react";
import { LinkContainer } from "react-router-bootstrap";
import { ListGroup, ListGroupItem, Container } from "react-bootstrap";
import { Scrollbars } from "react-custom-scrollbars-2";

const SearchResults = ({ searchResults, setSearchResults }) => {
  const searchResultsMap = searchResults?.map((result, index) => (
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

  return (
    <ListGroup className="search-results position-fixed w-25 h-50">
      <Scrollbars autoHide>{searchResultsMap}</Scrollbars>
    </ListGroup>
  );
};

export default SearchResults;
