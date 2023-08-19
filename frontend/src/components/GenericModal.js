import React, { useState } from "react";
import { Modal, Button } from "react-bootstrap";

const GenericModal = ({ buttonText, headerText, children }) => {
  const [show, setShow] = useState(false);

  const handleClose = () => setShow(false);
  const handleShow = () => setShow(true);

  return (
    <>
      <Button className="w-100" variant="primary" onClick={handleShow}>
        {buttonText}
      </Button>

      <Modal centered show={show} onHide={handleClose}>
        <Modal.Header closeButton>
          <h3>{headerText ? headerText : buttonText}</h3>
        </Modal.Header>
        <Modal.Body>
          {React.Children.map(children, (child) =>
            React.cloneElement(child, { handleClose })
          )}
        </Modal.Body>
      </Modal>
    </>
  );
};

export default GenericModal;
