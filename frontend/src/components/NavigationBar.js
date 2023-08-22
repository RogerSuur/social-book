import { Outlet } from "react-router-dom";
import useAuth from "../hooks/useAuth";
import { Nav, Navbar } from "react-bootstrap";
import NotificationNavbarItem from "./NotificationNavbarItem";
import { LinkContainer } from "react-router-bootstrap";

const NavigationBar = () => {
  const { auth } = useAuth();

  return (
    <>
      <Navbar
        as={Nav}
        className="bg-warning flex-sm-row justify-content-evenly"
        fixed="top"
      >
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
        <LinkContainer to="/profile">
          <Nav.Link>Profile</Nav.Link>
        </LinkContainer>
        <LinkContainer to="/posts">
          <Nav.Link>Posts</Nav.Link>
        </LinkContainer>
        <LinkContainer to="/logout">
          <Nav.Link>
            <img src={`${process.env.PUBLIC_URL}/logout.png`} alt="Logout" />
          </Nav.Link>
        </LinkContainer>
      </Navbar>
      <div className="content-wrapper">
        <Outlet />
      </div>
    </>
  );
};

export default NavigationBar;
