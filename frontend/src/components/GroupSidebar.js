import React from "react";
import GenericGroupList from "../components/GenericGroupList";
import GenericEventList from "../components/GenericEventList";
import SearchBar from "./SearchBar";
import {
  USER_CREATED_GROUPS_URL,
  USER_GROUPS_URL,
  ACCEPTED_EVENTS_URL,
} from "../utils/routes";

const GroupSidebar = () => {
  const sidebarItems = (
    <>
      <h1>Search</h1>
      <SearchBar />
      <ul className="group-sidebar">
        <h1>Groups</h1>
        <li>
          <GenericGroupList url={USER_GROUPS_URL} />
        </li>
        <h1>My groups</h1>
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
