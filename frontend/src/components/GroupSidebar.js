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

const GroupSidebar = () => {
  const [searchResults, setSearchResults] = useState([]);

  const sidebarItems = (
    <>
      
      <SearchBar setSearchResults={setSearchResults} />
      <SearchResults searchResults={searchResults} />
      <ul className="group-sidebar">
        <h1>Groups</h1>
        <li>
          <GenericGroupList url={USER_GROUPS_URL} />
        </li>
        <h1>My groups</h1>
        <i class="add-circle">
          <CreateGroup />
        </i>
        <li>
          <GenericGroupList url={USER_CREATED_GROUPS_URL} />
        </li>
        <h1>Events</h1>
        <li>
          <GenericEventList url={ACCEPTED_EVENTS_URL} />
        </li>
      </ul>
    </>
  );

  return <div className="group-extra">{sidebarItems} </div>;
};

export default GroupSidebar;
