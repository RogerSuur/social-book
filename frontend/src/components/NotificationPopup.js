import React, { useState } from "react";
import Notification from "../components/Notification";
import { Toast, CloseButton, Row, Col, Container } from "react-bootstrap";

const NotificationPopup = ({ notification, onPopupClose }) => {
  const [show, setShow] = useState(true);

  return (
    <Toast
      className="d-none d-md-block position-absolute"
      bg="info"
      show={show}
      autohide
      onClose={() => {
        setShow(false);
        onPopupClose();
      }}
    >
      <Toast.Body>
        <Container>
          <Row>
            <Col>
              <Notification notification={notification} popup={true} />
            </Col>
            <Col md="1">
              <CloseButton
                onClick={() => {
                  setShow(false);
                  onPopupClose();
                }}
              />
            </Col>
          </Row>
        </Container>
      </Toast.Body>
    </Toast>
  );
};

export default NotificationPopup;
