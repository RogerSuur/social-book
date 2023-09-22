import { Outlet } from "react-router-dom";
import useAuth from "../hooks/useAuth";
import { Nav, Navbar, Container, Row, Col } from "react-bootstrap";
import NotificationNavbarItem from "./NotificationNavbarItem";
import { LinkContainer } from "react-router-bootstrap";
import SearchUtility from "../components/SearchUtility";
import { BoxArrowRight } from "react-bootstrap-icons";
import SmallNotificationNavbarItem from "../components/SmallNotificationNavbarItem";
import SearchSmallUtility from "../components/SearchSmallUtility";

const NavigationBar = () => {
  const { auth } = useAuth();

  return (
    <>
      <Navbar
        collapseOnSelect
        expand="md"
        className="bg-body-tertiary"
        fixed="top"
      >
        <Container>
          <LinkContainer to="/profile">
            <Navbar.Brand>React-Bootstrap</Navbar.Brand>
          </LinkContainer>
          {auth && <NotificationNavbarItem />}
          <Navbar.Toggle aria-controls="responsive-navbar-nav" />

          <Navbar.Collapse id="responsive-navbar-nav">
            <Nav className="me-auto">
              <LinkContainer to="/profile">
                <Nav.Link>Profile</Nav.Link>
              </LinkContainer>
              <LinkContainer to="/posts">
                <Nav.Link>Posts</Nav.Link>
              </LinkContainer>
            </Nav>
            <Nav>
              {auth && <SearchUtility />}
              {!auth && (
                <>
                  <LinkContainer to="/login">
                    <Nav.Link>Sign In</Nav.Link>
                  </LinkContainer>
                  <LinkContainer to="/signup">
                    <Nav.Link>Sign Up</Nav.Link>
                  </LinkContainer>
                </>
              )}

              <LinkContainer to="/logout">
                <Nav.Link>
                  <BoxArrowRight />
                </Nav.Link>
              </LinkContainer>
            </Nav>
          </Navbar.Collapse>
        </Container>
      </Navbar>
      {/* <Navbar
        className="bg-secondary-subtle border-bottom"
        fixed="top"
        expand="md"
        collapseOnSelect
      >
        <Navbar.Toggle aria-controls="responsive-navbar-nav" />
        <Navbar.Collapse
          className="justify-content-around"
          id="responsive-navbar-nav"
        >
          <Nav>
            {!auth && (
              <>
                <LinkContainer to="/login">
                  <Nav.Link>Sign In</Nav.Link>
                </LinkContainer>
                <LinkContainer to="/signup">
                  <Nav.Link>Sign Up</Nav.Link>
                </LinkContainer>
              </>
            )}
            {auth && <NotificationNavbarItem />}
            {auth && <SearchUtility />}
            <LinkContainer to="/profile">
              <Nav.Link>Profile</Nav.Link>
            </LinkContainer>
            <LinkContainer to="/posts">
              <Nav.Link>Posts</Nav.Link>
            </LinkContainer>
            <LinkContainer to="/logout">
              <Nav.Link>
                <BoxArrowRight />
              </Nav.Link>
            </LinkContainer>
          </Nav>
        </Navbar.Collapse>
      </Navbar>
      <div className="content-wrapper">
        <Outlet />
      </div> */}
    </>
  );
};

export default NavigationBar;
