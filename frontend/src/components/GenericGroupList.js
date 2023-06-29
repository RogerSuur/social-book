import List from "./List.js";
import { Link } from "react-router-dom";

const GenericGroupList = ({ url }) => {
  const mapGenericGroupList = (group, index) => (
    <li key={index}>
      <Link to={`/groups/${group.id}`}>{group.name}</Link>
    </li>
  );

  return <List url={url} mapFunction={mapGenericGroupList} />;
};

export default GenericGroupList;
