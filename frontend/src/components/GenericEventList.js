import List from "./List.js";
import { ListGroup } from "react-bootstrap";
import { LinkContainer } from "react-router-bootstrap";

const GenericEventList = ({ url }) => {
  const mapGenericEventList = (event, index) => (
    <LinkContainer action to={`/event/${event.id}`}>
      <ListGroup.Item key={index}>
        <>{event.title}</>
      </ListGroup.Item>
    </LinkContainer>
  );

  return <List url={url} mapFunction={mapGenericEventList} />;
};

export default GenericEventList;
