import { Outlet, Link } from "react-router-dom";
import useAuth from "../hooks/useAuth";
import NotificationNavbarItem from "../components/NotificationNavbarItem";

const Navbar = () => {
  const { auth } = useAuth();

  return (
    <>
      <nav>
        <ul>
          {!auth && (
            <>
              <li>
                <Link className="text-link" to="/login">
                  Sign In
                </Link>
              </li>
              <li>
                <Link className="text-link" to="/signup">
                  Sign Up
                </Link>
              </li>
            </>
          )}
          {auth && <NotificationNavbarItem />}
          <li>
            <Link className="text-link" to="/profile">
              Profile
            </Link>
          </li>
          <li>
            <Link className="text-link" to="/posts">
              Posts
            </Link>
          </li>
          <li>
            <Link className="text-link" id="logout" to="/logout">
              <img className="text-link" src={logoutIcon} alt="Logout" />
            </Link>
          </li>
        </ul>
      </nav>
      <Outlet />
    </>
  );
};

export default Navbar;
