import React, { useState, useEffect } from "react";
import { useParams, useNavigate, Link } from "react-router-dom";
import { EVENT_URL, EVENT_ATTENDANCE_URL } from "../utils/routes";
import axios from "axios";
import ImageHandler from "../utils/imageHandler";
import { Container, Row, Col, Button, Stack, ListGroup } from "react-bootstrap";
import { LinkContainer } from "react-router-bootstrap";
import GenericModal from "../components/GenericModal";
import { ShortDatetime } from "../utils/datetimeConverters";

const EventPage = () => {
  const [event, setEvent] = useState({});
  const [error, setError] = useState("");
  const [response, setResponse] = useState(false);
  const navigate = useNavigate();
  const { id } = useParams();

  useEffect(() => {
    const loadEvent = async () => {
      try {
        await axios
          .get(EVENT_URL + id, {
            withCredentials: true,
          })
          .then((response) => {
            console.log("RESP: ", response.data);
            setEvent(response.data);
          });
      } catch (err) {
        if (!err?.response) {
          setError("No Server Response");
        } else if (err.response?.status === 404) {
          navigate("404", { replace: true });
        } else if (err.response?.status > 200) {
          setError("Internal Server Error");
        }
      }
    };
    loadEvent();
  }, [id, response]);

  const image = (user) =>
    ImageHandler(user?.imagePath, "defaultuser.jpg", "userlist-img");

  const userList = (attendance) => {
    const users = event?.members?.filter(
      (member) => member.isAttending === attendance
    );

    return users?.map((member, index) => (
      <ListGroup.Item action key={index}>
        <LinkContainer to={`/profile/${member.id}`}>
          <div>
            {image(member)}
            {member?.nickname
              ? `${member.nickname}`
              : `${member.firstName} ${member.lastName}`}
          </div>
        </LinkContainer>
      </ListGroup.Item>
    ));
  };

  const handleResponse = async (isAttending) => {
    console.log("IS ATTENDING: ", isAttending);
    const data = { eventId: +id, isAttending };
    try {
      await axios.post(
        EVENT_ATTENDANCE_URL,
        JSON.stringify(data),
        { withCredentials: true },
        {
          headers: { "Content-Type": "application/json" },
        }
      );

      setResponse(!response);
    } catch (err) {
      if (!err?.response) {
        setError("No Server Response");
      } else if (err.response?.status > 200) {
        setError("Internal Server Error");
      }
    }
  };

  const countUsers = (attending) => {
    const userArray = event?.members?.map((member) => member.isAttending);

    return userArray?.reduce(
      (count, obj) => (obj === attending ? count + 1 : count),
      0
    );
  };

  const renderedEvent = (
    <Container fluid>
      <Row>
        <Col className="m-auto text-center">
          <h1>{event?.title}</h1>
        </Col>
        <Col md="3" className="m-auto">
          <Stack gap={2}>
            <Button onClick={() => handleResponse(true)}>Attend</Button>
            <Button onClick={() => handleResponse(false)}>Skip</Button>
          </Stack>
        </Col>
      </Row>
      <p>{event?.description}</p>
      <p>Start: {ShortDatetime(event?.eventTime)}</p>
      <p>End: {ShortDatetime(event?.eventEndTime)}</p>
      <Row className="gap-2">
        <Col xs="12" md>
          <GenericModal
            buttonText={`Going ${countUsers(true) > 0 ? countUsers(true) : ""}`}
            headerText="Going"
          >
            {userList(true)}
          </GenericModal>
        </Col>
        <Col xs="12" md>
          <GenericModal
            buttonText={`Not going ${
              countUsers(false) > 0 ? countUsers(false) : ""
            }`}
            headerText="Not Going"
          >
            {userList(false)}
          </GenericModal>
        </Col>
      </Row>
    </Container>
  );

  return <>{error ? <div>{error}</div> : <div>{renderedEvent}</div>}</>;
};

export default EventPage;
