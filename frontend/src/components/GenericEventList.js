import List from "./List.js";
import { Link } from "react-router-dom";

const GenericEventList = ({ url }) => {
  const mapGenericEventList = (event, index) =>
    console.log(event)(
      <li key={index}>
        <Link to={`/event/${event.id}`}>{event.title}</Link>
      </li>
    );

  return <List url={url} mapFunction={mapGenericEventList} />;
};

export default GenericEventList;
