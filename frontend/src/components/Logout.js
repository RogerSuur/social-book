import { useEffect } from "react";
import useAuth from "../hooks/useAuth";
import axios from "axios";

const LOGOUT_URL = "http://localhost:8000/logout";

const Logout = () => {
  const { setAuth } = useAuth();

  useEffect(() => {
    const authorisation = async () => {
      await axios.get(LOGOUT_URL, {
        withCredentials: true,
      });

      setAuth(false);
    };

    authorisation();
  }, []);
};

export default Logout;
