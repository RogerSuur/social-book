import List from "./List.js";
import { Link } from "react-router-dom";

const GenericEventList = ({ url }) => {
  const mapGenericEventList = (event, index) => (
    <li key={index}>
      <Link to={`/event/${event.id}`}>{event.name}</Link>
    </li>
  );

  return <List url={url} mapFunction={mapGenericEventList} />;
};

export default GenericEventList;
