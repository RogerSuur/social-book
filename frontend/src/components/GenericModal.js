import React, { useState } from "react";
import { Modal, Button } from "react-bootstrap";

const GenericModal = ({
  img,
  variant,
  linkText,
  buttonText,
  headerText,
  headerButton,
  children,
}) => {
  const [show, setShow] = useState(false);

  const handleClose = () => setShow(false);
  const handleShow = () => setShow(true);

  return (
    <>
      {linkText ? (
        <p onClick={handleShow}>{linkText}</p>
      ) : (
        <Button
          className="w-100"
          variant={variant ? variant : "primary"}
          onClick={handleShow}
        >
          {img ? img : buttonText}
        </Button>
      )}

      <Modal
        centered
        scrollable
        show={show}
        fullscreen="md-down"
        onHide={handleClose}
      >
        <Modal.Header closeButton>
          <h3 className="my-auto">{headerText ? headerText : buttonText}</h3>
          <div>{headerButton && headerButton}</div>
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
