import { useState } from "react";
import { Button } from "react-bootstrap";
import { Messenger, XLg } from "react-bootstrap-icons";
import Chat from "../components/Chat";
import { Offcanvas } from "react-bootstrap";

const NavbarChat = () => {
  const [show, setShow] = useState(false);

  const handleShow = () => setShow(true);
  const handleClose = () => setShow(false);

  return (
    <>
      <Messenger as={Button} onClick={handleShow} />
      {show && (
        <>
          <Offcanvas show={show} onHide={handleClose}>
            <Offcanvas.Header className="ms-auto" closeButton />
            <Offcanvas.Body>
              <Chat />
            </Offcanvas.Body>
          </Offcanvas>
        </>
        // <div className="d-md-none position-fixed top-0 start-0 vw-100 vh-100 bg-light">
        //   <XLg
        //     className="justify-content-end"
        //     as={Button}
        //     onClick={handleClose}
        //   />
        //   <Chat />
        // </div>
      )}
    </>
  );
};

export default NavbarChat;