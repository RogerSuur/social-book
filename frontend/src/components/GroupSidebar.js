import React, { useState } from "react";
import GenericGroupList from "../components/GenericGroupList";
import GenericEventList from "../components/GenericEventList";
import SearchBar from "../components/SearchBar";
import {
  USER_CREATED_GROUPS_URL,
  USER_GROUPS_URL,
  ACCEPTED_EVENTS_URL,
} from "../utils/routes";
import { SearchResults } from "../components/SearchResults";
import CreateGroup from "../components/CreateGroup";
import { Container, ListGroup, Row, Col } from "react-bootstrap";

const GroupSidebar = () => {
  const [searchResults, setSearchResults] = useState([]);
  const [loadNewGroups, setLoadNewGroups] = useState(0);

  const handleGroupUpdate = () => {
    setLoadNewGroups((prevCount) => prevCount + 1);
  };

  const sidebarItems = (
    <Container>
      <SearchBar setSearchResults={setSearchResults} />
      <SearchResults searchResults={searchResults} />
      <h1>Groups</h1>
      <ListGroup variant="flush">
        <GenericGroupList url={USER_GROUPS_URL} />
      </ListGroup>
      <Row>
        <Col>
          <h1>My groups</h1>
        </Col>
        <Col xs="3" className="m-auto">
          <CreateGroup onGroupCreated={handleGroupUpdate} />
        </Col>
      </Row>

      <ListGroup variant="flush">
        <GenericGroupList
          url={USER_CREATED_GROUPS_URL}
          loadNewGroups={loadNewGroups}
        />
      </ListGroup>

      <h1>Events</h1>
      <ListGroup variant="flush">
        <GenericEventList url={ACCEPTED_EVENTS_URL} />
      </ListGroup>
    </Container>
  );

  return <>{sidebarItems}</>;
};

export default GroupSidebar;
