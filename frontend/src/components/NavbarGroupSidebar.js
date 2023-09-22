import { useState } from "react";
import GroupSidebar from "../components/GroupSidebar";
import { Offcanvas, Nav } from "react-bootstrap";

const NavbarGroupSidebar = () => {
  const [show, setShow] = useState(false);

  const handleShow = () => setShow(true);
  const handleClose = () => setShow(false);

  return (
    <>
      <Nav.Link onClick={handleShow}>Groups and events</Nav.Link>
      {show && (
        <>
          <Offcanvas show={show} onHide={handleClose}>
            <Offcanvas.Header className="ms-auto" closeButton />
            <Offcanvas.Body onClick={handleClose}>
              <GroupSidebar />
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

export default NavbarGroupSidebar;
