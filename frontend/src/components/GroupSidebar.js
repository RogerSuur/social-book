import React from "react";
import GenericGroupList from "../components/GenericGroupList";
import GenericEventList from "../components/GenericEventList";
import {
  USER_CREATED_GROUPS_URL,
  USER_GROUPS_URL,
  ACCEPTED_EVENTS_URL,
} from "../utils/routes";

const GroupSidebar = () => {
  const sidebarItems = "sidebar";
  // const sidebarItems = (
  //   <>
  //     <li>
  //       <GenericGroupList url={USER_GROUPS_URL} />
  //     </li>
  //     <li>
  //       <GenericGroupList url={USER_CREATED_GROUPS_URL} />
  //     </li>
  //     <li>
  //       <GenericEventList url={ACCEPTED_EVENTS_URL} />
  //     </li>
  //   </>
  // );

  return <ul className="group-sidebar">{sidebarItems}</ul>;
};

export default GroupSidebar;
