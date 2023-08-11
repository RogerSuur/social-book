import React, { useState } from "react";
import GenericGroupList from "../components/GenericGroupList";
import GenericEventList from "../components/GenericEventList";
import SearchBar from "./SearchBar";
import {
  USER_CREATED_GROUPS_URL,
  USER_GROUPS_URL,
  ACCEPTED_EVENTS_URL,
} from "../utils/routes";
import { SearchResults } from "./SearchResults";
import CreateGroup from "./CreateGroup";
import Container from "react-bootstrap/Container";

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
      <ul className="group-sidebar">
        <h1>Groups</h1>
        <li className="pepe">
          <GenericGroupList url={USER_GROUPS_URL} />
        </li>
        <h1>My groups</h1>
        <CreateGroup onGroupCreated={handleGroupUpdate} />
        <li className="pepe">
          <GenericGroupList
            url={USER_CREATED_GROUPS_URL}
            loadNewGroups={loadNewGroups}
          />
        </li>
        <h1>Events</h1>
        <li className="pepe">
          <GenericEventList url={ACCEPTED_EVENTS_URL} />
        </li>
      </ul>
    </Container>
  );

  return <>{sidebarItems}</>;
};

export default GroupSidebar;
