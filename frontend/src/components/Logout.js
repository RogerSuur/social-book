import { useEffect } from "react";
import useAuth from "../hooks/useAuth";
import axios from "axios";
import { useNavigate } from "react-router-dom";

const LOGOUT_URL = "http://localhost:8000/logout";

const Logout = () => {
  const { setAuth } = useAuth();
  const navigate = useNavigate();

  useEffect(() => {
    const authorisation = async () => {
      await axios.get(LOGOUT_URL, {
        withCredentials: true,
      });

      setAuth(false);
      navigate("/login", { replace: true });
    };

    authorisation();
  }, []);
};

export default Logout;
